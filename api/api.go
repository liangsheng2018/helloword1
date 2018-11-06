package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"hello/business"
	"hello/data"
	"hello/db"
	"hello/enc"
	"hello/gtime"
	"hello/rec"
	"hello/redis"
	_ "hello/routers" //自动注册路由
	"hello/sw"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 账号登录
func UserLogin(w http.ResponseWriter, req *http.Request){

	p := rec.GetParams(req)

	timestamp := time.Now().Unix()

	queryByUserName := fmt.Sprintf("select * from user where username='%v'", p.UserName)

	UpdateInfo := fmt.Sprintf("update user set status=1, update_time='%v' where username='%v'", timestamp, p.UserName)

	DB,_:= db.Query(queryByUserName)

	decPassword := enc.Decode(DB.PassWord)

	msg := data.ResponseErrReturn()

	if DB.UserName == p.UserName && decPassword == p.PassWord {

		db.Update(UpdateInfo)
		returnData := data.LoginSucceedReturn(w, req)["result"]
		js, err := json.Marshal(returnData)
		redis.RedisSetUserInfo(sw.IntStr(DB.UserId), js)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	data.ServerErrResponse(w, msg.M11)
	return
}

// 账号注册
func UserRegister(w http.ResponseWriter, req *http.Request) {

	p := rec.GetParams(req)

	msg := data.ResponseErrReturn()

	encPassWord := enc.Encrypt(p.PassWord)

	if len(p.UserName) < 6 || len(p.UserName) > 14 {
		data.ServerErrResponse(w, msg.M5)
		return
	}
	if len(p.UserName) < 6 || len(p.PassWord) > 14 {
		data.ServerErrResponse(w, msg.M6)
		return
	}

	a, b, c, d := business.CharacterType(p.UserName)
	if a == 0 || b == 0 || c > 0 || d > 0 {
		data.ServerErrResponse(w, msg.M5)
		return
	}

	e, f, g, h := business.CharacterType(p.PassWord)
	if e == 0 || f == 0 || g != 0 || h != 0 {
		data.ServerErrResponse(w, msg.M6)
		return
	}

	sqlInsert := fmt.Sprintf("INSERT INTO user (username,password) values ('%v','%v')", p.UserName, encPassWord)
	err := db.Insert(sqlInsert)
	Err := fmt.Sprint("error", err)

	if strings.Contains(Err, "1062") {
		data.ServerErrResponse(w, msg.M4)
		return
	}
	data.ServerResponseRight(w)
	return
}

//获取用户信息
func GetUserInfo(w http.ResponseWriter, req *http.Request) {

	p := rec.GetParams(req)

	timestamp := time.Now().Unix()

	UpdateInfo := fmt.Sprintf("update user set status=1, update_time='%v' where userid='%v'", timestamp, p.UserId)

	queryById := fmt.Sprintf("select * from user where userid='%v'", p.UserId)

	DB, _ := db.Query(queryById)

	msg := data.ResponseErrReturn()

	if redis.CheckRedisKey(sw.IntStr(DB.UserId)) == true {

		rds := redis.RedisUserInfo(sw.IntStr(DB.UserId))

		if  p.UserId == DB.UserId &&  p.Token == rds.Token {

			db.Update(UpdateInfo)
			UserInfo := redis.RedisGetUserInfo(strconv.Itoa(DB.UserId))
			w.Header().Set("Content-Type", "application/json")
			w.Write(UserInfo.([]byte))
			return
		}
		data.ServerErrResponse(w, msg.M13)
		return
	}
	data.ServerErrResponse(w, msg.M12)
	return
}

// 添加银行卡
func AddUserBankCard(w http.ResponseWriter, req *http.Request){

	p := rec.GetParams(req)

	InsertNameAddr := fmt.Sprintf("UPDATE user SET name='%v',addr='%v' where userid=%v", p.Name, p.Addr, p.UserId)

	InsertAddr := fmt.Sprintf("UPDATE user SET addr='%v' where userid=%v", p.Addr, p.UserId)

	setBankDefault_1 := fmt.Sprintf("INSERT INTO user_bank_card (userid,bank,card_num,bank_default) VALUES (%v,'%v','%v',1)", p.UserId, p.Bank, p.CardNum)

	setBankDefault_0 := fmt.Sprintf("INSERT INTO user_bank_card (userid,bank,card_num,bank_default) VALUES (%v,'%v','%v',0)", p.UserId, p.Bank, p.CardNum)

	Query := fmt.Sprintf("select * from user where userid=%v", p.UserId)

	DB,_:= db.Query(Query)

	CheckCardNum := fmt.Sprintf("select * from user_bank_card where card_num='%v'", p.CardNum)

	checkCardCount := fmt.Sprintf("select count(*) from user_bank_card where userid ='%v'", p.UserId)

	card_count:= db.QueryCardCount(checkCardCount)

	st,_,_ := db.QueryUserBankCard(CheckCardNum)

	timestamp := time.Now().Unix()

	msg := data.ResponseErrReturn()

	_,_,_,s :=  business.CharacterType(p.Name)

	// 校验userid与token
	if redis.CheckRedisKey(sw.IntStr(DB.UserId)) == true {
		rds := redis.RedisUserInfo(sw.IntStr(DB.UserId))

		if p.UserId == DB.UserId &&  p.Token == rds.Token  {

			if len(p.Bank) == 0 {
				data.ServerErrResponse(w, msg.M14)
				return
			}
			// 检查银行卡号长度是否合法
			if len(p.CardNum) < 15 || len(p.CardNum) > 19 {
				data.ServerErrResponse(w, msg.M8)
				return
			}
			// 检查地址长度是否合法
			if len(p.Addr) < 2 || len(p.Addr) > 100 {
				data.ServerErrResponse(w, msg.M10)
				return
			}
			// 检查银行卡号是否已经存在
			if p.CardNum == st.Card_Num {
				data.ServerErrResponse(w, msg.M15)
				return
			}
			// 检查是否为首次绑定
			if card_count == 0 {

				// 检查该用户是否存已绑定真实姓名
				if len(DB.Name) == 0 {

					if len(p.Name) < 4 || len(p.Name) > 10 {
						data.ServerErrResponse(w, msg.M16)
						return
					}

					// 检查用户输入的姓名是否包含特殊符号
					if  s != 0 {
						data.ServerErrResponse(w, msg.M16)
						return
					}
					db.Update(InsertNameAddr)
				}
				db.Update(InsertAddr)
				db.Insert(setBankDefault_1)
				card_count = card_count +1
				UpdateUserTable := fmt.Sprintf("UPDATE user SET card_count='%v',update_time='%v' where userid='%v'", card_count, timestamp, p.UserId)
				db.Update(UpdateUserTable)
				data.ServerResponseRight(w)
				return
			}

			// 检查已绑定银行卡数量是否大于5张
			if 0 < card_count && card_count < 5 {
				db.Update(InsertAddr)
				err := db.Insert(setBankDefault_0)

				if err == nil {

					card_count = card_count +1
					UpdateUserTable := fmt.Sprintf("UPDATE user SET card_count='%v',update_time='%v' where userid='%v'", card_count, timestamp, p.UserId)
					db.Update(UpdateUserTable)
					data.ServerResponseRight(w)
					return
				}
				return
			}
			data.ServerErrResponse(w, msg.M16)
			return
		}
		data.ServerErrResponse(w, msg.M13)
		return
	}
	data.ServerErrResponse(w, msg.M12)
	return
}

// 解除绑定银行卡
func DeleteUserBankCard(w http.ResponseWriter, req *http.Request){

	p := rec.GetParams(req)

	QueryBankDefault := fmt.Sprintf("select * from user_bank_card where bank_id ='%v'", p.BankId)

	queryById := fmt.Sprintf("select * from user where userid=%v", p.UserId)

	DB,_:= db.Query(queryById)

	ss,_,_ := db.QueryUserBankCard(QueryBankDefault)

	QueryBankId := fmt.Sprintf("select * from user_bank_card where userid='%v' and bank_id='%v'", p.UserId, p.BankId)

	st,_,_:= db.QueryUserBankCard(QueryBankId)

	checkCardCount := fmt.Sprintf("select count(*) from user_bank_card where userid ='%v'", p.UserId)

	card_count:= db.QueryCardCount(checkCardCount)

	DeleteBankCard := fmt.Sprintf("delete from user_bank_card where bank_id=%v and userid=%v", p.BankId, p.UserId)

	msg := data.ResponseErrReturn()

	timestamp := time.Now().Unix()

	if redis.CheckRedisKey(sw.IntStr(DB.UserId)) == true {

		rds := redis.RedisUserInfo(sw.IntStr(DB.UserId))

		if p.UserId == DB.UserId && p.Token == rds.Token && p.BankId == strconv.FormatInt(st.BandId,10) {

			if len(p.BankId) == 0 {
				data.ServerErrResponse(w, msg.M13)
				return
			}
			if ss.Bank_Default == 1 {
				data.ServerErrResponse(w, msg.M18)
				return
			}
			card_count -= 1
			UpdateUserTable := fmt.Sprintf("UPDATE user SET card_count='%v',update_time='%v' where userid='%v'", card_count, timestamp, p.UserId)
			db.Remove(DeleteBankCard)
			db.Update(UpdateUserTable)
			data.ServerResponseRight(w)
			return
		}
		data.ServerErrResponse(w, msg.M13)
		return
	}
	data.ServerErrResponse(w, msg.M12)
	return
}

//用户退出登录
func UserLoginOut(w http.ResponseWriter, req *http.Request){
	p := rec.GetParams(req)

	queryById := fmt.Sprintf("select * from user where userid=%v", p.UserId)

	UpdateStatus := fmt.Sprintf("update user set status=0, token='' where userid=%v", p.UserId)

	DB,_:= db.Query(queryById)

	msg := data.ResponseErrReturn()

	if redis.CheckRedisKey(sw.IntStr(DB.UserId)) == true {

		rds := redis.RedisUserInfo(sw.IntStr(DB.UserId))

		if p.UserId == DB.UserId && p.Token == rds.Token{

			db.Update(UpdateStatus)

			if redis.RemoveRedis(strconv.Itoa(DB.UserId)) == true {
				data.ServerResponseRight(w)
				return
			}
			return
		}
		data.ServerErrResponse(w, msg.M13)
		return
	}
	data.ServerErrResponse(w, msg.M12)
	return
}

// 修改用户登录密码
func ChangeUserPassWord(w http.ResponseWriter, req *http.Request){

	p := rec.GetParams(req)

	encPassword := enc.Encrypt(p.NewPassWord)

	UpdatePW:=  fmt.Sprintf("update user set password='%v' where userid=%v", encPassword, p.UserId)

	queryById := fmt.Sprintf("select * from user where userid=%v", p.UserId)

	DB,_:= db.Query(queryById)

	msg := data.ResponseErrReturn()

	decPassWord := enc.Decode(DB.PassWord)

	if redis.CheckRedisKey(sw.IntStr(DB.UserId)) == true {  // 检查用户是否已离线

		rds := redis.RedisUserInfo(sw.IntStr(DB.UserId))

		if p.UserId == DB.UserId && p.Token == rds.Token {  // 检查请求参数是否正确

 			if p.PassWord == decPassWord {    // 校验原密码是否正确

				if len(p.NewPassWord) != 0 {  // 检查新密码是否为空

					if decPassWord != p.NewPassWord{
						db.Update(UpdatePW)
						data.ServerResponseRight(w)
						return
					}
					data.ServerErrResponse(w, msg.M7)
					return
				}
				data.ServerErrResponse(w, msg.M19)
				return
			}
			data.ServerErrResponse(w, msg.M20)
			return
		}
		data.ServerErrResponse(w, msg.M13)
		return
	}
	data.ServerErrResponse(w, msg.M12)
	return
}

// 签到功能
func SignIn(w http.ResponseWriter, req *http.Request){
	p := rec.GetParams(req)

	msg := data.ResponseErrReturn()

	queryByUserid := fmt.Sprintf("select * from user_signin where userid=%v", p.UserId)

	sign := db.UserSignInQuery(queryByUserid)

	timestamp := time.Now().Unix()

	updateSignTime := fmt.Sprintf("update user_signin set sign_time='%v' where userid=%v", timestamp, sign.UserId)

	updateSaveDays := fmt.Sprintf("update user_signin set save_days='0' where userid=%v", sign.UserId)

	currentTime := gtime.TimeStr(timestamp)

	latest_date := gtime.TimeString(sign.SignTime)[:10]

	latest_time := sw.StrInt64(sign.SignTime)

	if redis.CheckRedisKey(sw.IntStr(sign.UserId)) == true {

		rds := redis.RedisUserInfo(sw.IntStr(sign.UserId))

		if p.UserId == sign.UserId && p.Token == rds.Token {

			if sign.IsOpen == 1 {

				if sign.OpenLog == 1 {
					logF := flag.String("log", "test.log", "Log file name")
					flag.Parse()
					outfile, err := os.OpenFile(*logF, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

					if err != nil {
						fmt.Println(*outfile, "open failed")
						os.Exit(1)
					}
					//log.SetOutput(outfile)                              				      // 设置log的输出文件，不设置log输出默认为stdout
					//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) 					  // 设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名
					//log.Printf("test out:%v \n", "test log")                       		  // 向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了
				}
				if currentTime[:7] != latest_date[:7] || timestamp - latest_time > 172800 {   // 上月签到记录清零; 超过24小时未签到，签到记录清零
					sign.SaveDays = 0
					db.Update(updateSaveDays)
				}
				if currentTime[:10] == latest_date{                                           // 检查是当天否已签到
					data.ServerErrResponse(w, msg.M23)
					return
				}
				sign.SaveDays += 1

				if sign.Types == 1 {  														  // 判断活动模式，type为1时，表示每天签到奖励积分数
					sign.Integral += 3
					updateSignInfo := fmt.Sprintf("update user_signin set sign_time='%v',integral='%v',save_days='%v' where userid=%v", timestamp, sign.Integral, sign.SaveDays, sign.UserId)
					db.Update(updateSignInfo)
					data.ServerResponseRight(w)
					return
				}

				if sign.Types == 2 {  														  // type为2时，表示连续签到天数以及奖励阶梯

					if sign.SaveDays == 3 {
						sign.Integral += 5
						updateSignInfo := fmt.Sprintf("update user_signin set sign_time='%v',integral='%v',save_days='%v' where userid=%v", timestamp, sign.Integral, sign.SaveDays, sign.UserId)
						db.Update(updateSignInfo)
						data.ServerResponseRight(w)
						return
					}
					if sign.SaveDays == 5 {
						sign.Integral += 10
						updateSignInfo := fmt.Sprintf("update user_signin set sign_time='%v',integral='%v',save_days='%v' where userid=%v", timestamp, sign.Integral, sign.SaveDays, sign.UserId)
						db.Update(updateSignInfo)
						data.ServerResponseRight(w)
						return
					}
					if sign.SaveDays >= 7 {
						sign.Integral += 20
						updateSignInfo := fmt.Sprintf("update user_signin set sign_time='%v',integral='%v',save_days='%v' where userid=%v", timestamp, sign.Integral, sign.SaveDays, sign.UserId)
						db.Update(updateSignInfo)
						data.ServerResponseRight(w)
						return
					}

					updateSignInfo := fmt.Sprintf("update user_signin set sign_time='%v',integral='%v',save_days='%v' where userid=%v", timestamp, sign.Integral, sign.SaveDays, sign.UserId)
					db.Update(updateSignInfo)
					data.ServerResponseRight(w)
					return
				}
				sign.SaveDays += 1
				db.Update(updateSignTime)
				return
			}
			data.ServerErrResponse(w, msg.M21)
			return
		}
		data.ServerErrResponse(w, msg.M13)
		return
	}
	data.ServerErrResponse(w, msg.M12)
	return
}

