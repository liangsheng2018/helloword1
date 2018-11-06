package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hello/create"
	"hello/rec"
	_ "hello/routers" //自动注册路由
	"io"
	"log"
	"net/http"
)

var db *sql.DB
func init(){
	db, _ = sql.Open("mysql","root:000@tcp(192.168.100.103:3306)/sky001?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

// 新增数据
func Insert(sqlInsert string) error {
	stmt, err := db.Prepare(sqlInsert)
	checkErr(err)
	_, err = stmt.Exec()
	return err
}

// 删除数据
func Remove(sqlUpdate string){
	stmt, err := db.Prepare(sqlUpdate)
	checkErr(err)
	res, err := stmt.Exec()
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("影响的行数为：",num)
}

// 更新数据
func Update(sqlUpdate string){
	stmt, err := db.Prepare(sqlUpdate)
	checkErr(err)
	res, err := stmt.Exec()
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("影响的行数为：", num)
}

type UserInfo struct {
	UserId			int			`json:"userid"`				// 用户ID
	UserName		string		`json:"username"`			// 用户名称
	PassWord		string		`json:"password"`			// 登录密码
	Token			string 		`json:"token"`				// 令牌
	SessionKey		string		`json:"session_key"`		// 会话秘钥
	UserAge			int64		`json:"userage"`			// 年龄
	UserSex			string		`json:"usersex"`			// 性别
	Name			string		`json:"name"`				// 真实姓名
	Remain			int64		`json:"remain"`				// 余额
	Addr			string		`json:"addr"`				// 地址
	CardCount		int64		`json:"card_count"`			// 银行卡总数
	Status 			int										// 在线状态
	CreateTime		string									// 注册时间
	UpdateTime		string									// 更新时间
}

// 查询数据
func Query(sqlQuery string)(*UserInfo, error){
	rows, err := db.Query(sqlQuery)
	checkErr(err)

	var userid 	 	int
	var username 	string
	var password 	string
	var token 		sql.NullString
	var sessionkey  sql.NullString
	var userage 	sql.NullInt64
	var usersex 	sql.NullString
	var name 		sql.NullString
	var remain 		sql.NullInt64
	var addr		sql.NullString
	var card_count	sql.NullInt64
	var status		int
	var create_time	string
	var update_time	sql.NullString
	var	st			UserInfo

	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&userid, &username, &password, &token, &sessionkey, &userage, &usersex, &remain, &name, &addr, &card_count, &status, &create_time, &update_time)
		checkErr(err)

		if err != nil{
			fmt.Println("Query error.", err)
		}
		defer rows.Close()

		st.UserId = userid
		st.UserName = username
		st.PassWord = password
		st.Token = token.String
		st.SessionKey = sessionkey.String
		st.UserAge = userage.Int64
		st.UserSex = usersex.String
		st.Name	= name.String
		st.Remain = remain.Int64
		st.Addr = addr.String
		st.CardCount = card_count.Int64
		st.Status = status
		st.CreateTime = create_time
		st.UpdateTime = update_time.String
	}
	return &st, nil
}

type Data struct{

	UserId        		int			`json:"userid"`
	Bank          		string		`json:"bank"`
	BandId        		int64		`json:"bank_id"`
	Card_Num     		string		`json:"card_num"`
	Bank_Default    	int64		`json:"default"`
}

// 查询用户的银行卡表
func QueryUserBankCard(sqlQuery string)(*Data, []Data, error){
	rows, err := db.Query(sqlQuery)
	checkErr(err)

	var userid 	 	 	int						// 用户ID
	var bank 		 	sql.NullString			// 银行名称
	var bank_id 	 	sql.NullInt64			// 银行卡ID
	var card_num 	 	sql.NullString			// 银行卡账号
	var bank_default 	sql.NullInt64			// 默认银行卡
	var st 				Data
	var sm				[]Data

	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&userid, &bank, &bank_id, &card_num, &bank_default)
		checkErr(err)

		if err != nil{
			fmt.Println("Query error.", err)
		}

		defer rows.Close()

		st.UserId = userid
		st.Bank = bank.String
		st.BandId = bank_id.Int64
		st.Card_Num = card_num.String
		st.Bank_Default = bank_default.Int64

		sd := Data {
			UserId:userid,
			Bank:bank.String,
			BandId:bank_id.Int64,
			Card_Num:card_num.String,
			Bank_Default:bank_default.Int64,
		}
		sm = append(sm, sd)
	}
	return &st, sm, err
}

// 查询用户已绑定的银行卡数
func QueryCardCount(sqlQuery string)(int64){
	rows, err := db.Query(sqlQuery)
	checkErr(err)

	var card_count	 sql.NullInt64
	for rows.Next(){
		rows.Columns()
		err := rows.Scan(&card_count)
		checkErr(err)

		if err != nil{
			fmt.Println("Query error.", err)
		}

		defer rows.Close()
		return card_count.Int64
	}
	return 0
}

type UserSignIn struct {
	UserId 		int  	`json:"userid"`
	Id 			int		`json:"id"`
	Integral	int		`json:"integral"`
	SignTime	string	`json:"sign_time"`
	Types 		int		`json:"type"`
	Configs		int		`json:"configs"`
	OpenLog		int		`json:"open_log"`
	IsOpen		int		`json:"is_open"`
	SaveDays	int		`json:"save_days"`
}

func UserSignInQuery(sqlQuery string) *UserSignIn {
	rows, err := db.Query(sqlQuery)
	checkErr(err)

	var userid	 	int
	var id 			int
	var integral	int
	var types 		int
	var configs		int
	var open_log	int
	var is_open		int
	var sign_time	string
	var save_days	int
	var st			UserSignIn

	for rows.Next() {
		rows.Columns()
		err := rows.Scan(&id, &userid, &integral, &types, &configs, &open_log, &is_open, &sign_time, &save_days)
		checkErr(err)

		if err != nil {
			fmt.Println("Query error.", err)
		}
		defer rows.Close()
		st.UserId = userid
		st.Id = id
		st.Integral = integral
		st.Types = types
		st.Configs = configs
		st.OpenLog = open_log
		st.SignTime = sign_time
		st.IsOpen = is_open
		st.SaveDays = save_days
	}
	return &st
}

func httpServer(){
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func pool(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM user WHERE username = 102")
	io.WriteString(w, "Hello Golang")
	defer rows.Close()
	checkErr(err)
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values{
		scanArgs[j] = &values[j]
	}

	record := make(map[string] string)
	for rows.Next(){
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values{
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}
	fmt.Println(record)
	fmt.Println(w, "finish")
}

// 写入token
func InsertToken(w http.ResponseWriter, req *http.Request)string{
	token := create.CreateToken()
	sqlUpdate := fmt.Sprintf("UPDATE user SET token='%v' where username='%v'", token, rec.GetUserName(req))
	Update(sqlUpdate)
	return token
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
