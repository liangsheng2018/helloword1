package data

import (
	"encoding/json"
	"fmt"
	"hello/create"
	"hello/db"
	"hello/rec"
	"net/http"
)

type UserInfoSuccessReturn struct {
	Code    	 int 		`json:"code"`
	Message 	 string 		`json:"message"`
	Token 		 string 		`json:"token"`
	Data		 UserInfo		`json:"data"`
}

type UserInfo struct {
	UserName    string 		`json:"username"`
	UserId		int	   		`json:"userid"`
	Userage 	int64  		`json:"userage"`
	Usersex 	string 		`json:"usersex"`
	Name		string		`json:"name"`
	Remain  	int64 		`json:"remain"`
	Addr		string 		`json:"addr"`
	CardCount	int64  		`json:"card_count"`
	UserBankCard []db.Data	`json:"bank_info"`
}

func InfoSuccessReturn(w http.ResponseWriter, req *http.Request)(map[string]*UserInfoSuccessReturn){

	userid := rec.GetUserId(req)

	queryByUserId := fmt.Sprintf("select * from user where userid='%v'", userid)

	sv,_:= db.Query(queryByUserId)

	queryUserBankCard := fmt.Sprintf("select * from user_bank_card where userid='%v'", userid)

	_,bankInfo,_ := db.QueryUserBankCard(queryUserBankCard)

	infoSuccessReturn := map[string]*UserInfoSuccessReturn{}

	userInfoReturn := map[string]*UserInfo{}
	userInfo := UserInfo{
		UserName: 		sv.UserName,
		UserId:			sv.UserId,
		Userage: 		sv.UserAge,
		Usersex: 		sv.UserSex,
		Name:			sv.Name,
		Remain: 		sv.Remain,
		Addr:			sv.Addr,
		CardCount:		sv.CardCount,
		UserBankCard:	bankInfo,
	}
	userInfoReturn["userInfo"] = &userInfo

	result := UserInfoSuccessReturn {
		Code: 			200,
		Message: 		"查询用户信息成功!",
		Token:			sv.Token,
		Data:			userInfo,
	}
	infoSuccessReturn["result"] = &result
	return infoSuccessReturn
}

func LoginSucceedReturn(w http.ResponseWriter, req *http.Request)(map[string]*UserInfoSuccessReturn){

	username := rec.GetUserName(req)

	queryByUserId := fmt.Sprintf("select * from user where username='%v'", username)

	sv,_:= db.Query(queryByUserId)

	queryUserBankCard := fmt.Sprintf("select * from user_bank_card where userid='%v'", sv.UserId)

	_,BankInfo,_ := db.QueryUserBankCard(queryUserBankCard)

	infoSuccessReturn := map[string]*UserInfoSuccessReturn{}

	userInfoReturn := map[string]*UserInfo{}
	userInfo := UserInfo{
		UserName: 		sv.UserName,
		UserId:			sv.UserId,
		Userage: 		sv.UserAge,
		Usersex: 		sv.UserSex,
		Name:			sv.Name,
		Remain: 		sv.Remain,
		Addr:			sv.Addr,
		CardCount:		sv.CardCount,
		UserBankCard:	BankInfo,
	}
	userInfoReturn["userInfo"] = &userInfo

	result := UserInfoSuccessReturn {
		Code: 			200,
		Message: 		"ok",
		Token:			create.CreateToken(),
		Data:			userInfo,
	}
	infoSuccessReturn["result"] = &result
	return infoSuccessReturn
}

type UserRequestRightReturn struct {
	Code    	string 		`json:"code"`
	Message 	string 		`json:"message"`
}

func RequestRightReturn()map[string]*UserRequestRightReturn{
	infoFailReturn := map[string]*UserRequestRightReturn{}
	result := UserRequestRightReturn{Code: "200", Message: "ok"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

func ServerResponseRight(w http.ResponseWriter){
	returnData := RequestRightReturn()["result"]
	js, err := json.Marshal(returnData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
}

type UserRequestErrReturn struct {
	Code    	int 		`json:"code"`
	Message 	string 		`json:"message"`
}


func ParamErrReturn(Err string)map[string]*UserRequestErrReturn{
	infoFailReturn := map[string]*UserRequestErrReturn{}
	result := UserRequestErrReturn{Code: 400, Message: Err}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type ServerResponseErrReturn struct {
	M1 	string 		`json:"message"`
	M2 	string 		`json:"message"`
	M3 	string 		`json:"message"`
	M4 	string 		`json:"message"`
	M5 	string 		`json:"message"`
	M6 	string 		`json:"message"`
	M7 	string 		`json:"message"`
	M8 	string 		`json:"message"`
	M9 	string 		`json:"message"`
	M10 string 		`json:"message"`
	M11 string 		`json:"message"`
	M12 string 		`json:"message"`
	M13 string 		`json:"message"`
	M14 string 		`json:"message"`
	M15 string 		`json:"message"`
	M16 string 		`json:"message"`
	M17 string 		`json:"message"`
	M18 string 		`json:"message"`
	M19 string 		`json:"message"`
	M20 string 		`json:"message"`
	M21 string 		`json:"message"`
	M22 string 		`json:"message"`
	M23 string 		`json:"message"`
}

func ResponseErrReturn()*ServerResponseErrReturn{

	var msg ServerResponseErrReturn
	msg.M1 = "用户名不合法！"
	msg.M2 = "登录密码不合法！"
	msg.M3 = "请输入正确的验证码！"
	msg.M4 = "用户名已存在，请重新输入！"
	msg.M5 = "用户名必须由6-14位字符组成（支持数字/字母）！"
	msg.M6 = "登录密码必须由6-14位字符组成（支持数字/字母）！"
	msg.M7 = "新密码与旧密码不能相同！"
	msg.M8 = "银行卡账号必须由16-20位数字组成！"
	msg.M9 = "姓名不能为空！"
	msg.M10 = "地址不能为空！"
	msg.M11 = "请输入正确的用户名和密码！"
	msg.M12 = "您已离线，请重新登录！"
	msg.M13 = "参数错误！"
	msg.M14 = "银行名称不能为空！"
	msg.M15 = "银行账号已存在，请重新输入！"
	msg.M16 = "请输入正确的姓名！"
	msg.M17 = "最多支持绑定5张银行卡！"
	msg.M18 = "默认银行卡不支持解除绑定！"
	msg.M19 = "用户名或密码不能为空！"
	msg.M20 = "原密码校验失败！"
	msg.M21 = "活动未开始！"
	msg.M22 = "签到成功！"
	msg.M23 = "您今天已签到！"

	return &msg
}

func ServerErrResponse(w http.ResponseWriter, msg string){
	returnData := ParamErrReturn(msg)["result"]
	js, err := json.Marshal(returnData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
}