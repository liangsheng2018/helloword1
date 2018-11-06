package main

import (
	"hello/api"
	"hello/rec"
	"log"
	"net/http"
)

// 开启web服务
func startHttpServer(){

	http.HandleFunc("/request", request)

	err := http.ListenAndServe(":8083", nil)

	if err != nil{

		log.Fatal("ListenAndServe", err)
	}
}

func request(w http.ResponseWriter,req *http.Request){

	params := rec.GetEvent(req)

	switch  value := params ; value {

	case "login":
		api.UserLogin(w,req)

	case "register":
		api.UserRegister(w, req)

	case "getUserInfo":
		api.GetUserInfo(w, req)

	case "addUserBankCard":
		api.AddUserBankCard(w, req)

	case "deleteUserBankCard":
		api.DeleteUserBankCard(w, req)

	case "userLoginOut":
		api.UserLoginOut(w, req)

	case "changeUserPassword":
		api.ChangeUserPassWord(w, req)

	case "userSignIn":
		api.SignIn(w, req)
	}
}

func main() {
	startHttpServer()
}