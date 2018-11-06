package rec

import (
	"net/http"
	"strconv"
)

func GetUserName(req *http.Request) string {
	req.ParseForm()
	paramUsername, found1:= req.Form["username"]
	if !(found1){
	} else {
		username := paramUsername[0]
		return username
	}
	return ""
}

func GetPassWord(req *http.Request) string {
	req.ParseForm()
	paramPassword, found2:= req.Form["password"]
	if !(found2){
	} else {
		password := paramPassword[0]
		return password
	}
	return ""
}

func GetUserId(req *http.Request) int {
	req.ParseForm()
	paramUserid, found3:= req.Form["userid"]
	if !(found3){
	} else {
		userid, _ := strconv.Atoi(paramUserid[0])
		return userid
	}
	return 0
}

func GetToken(req *http.Request) string {
	paramToken, found4:= req.Form["token"]
	if !(found4){
		return "token is null!"
	} else {
		token := paramToken[0]
		return token
	}
	return ""
}

func GetSessionKey(req *http.Request) string {
	req.ParseForm()
	paramSessionkey, found5:= req.Form["sessionkey"]
	if !(found5){
		return "sessionkey is null!"
	} else {
		sessionkey := paramSessionkey[0]
		return sessionkey
	}
	return ""
}

func GetData(req *http.Request) string {
	req.ParseForm()
	paramData, found6:= req.Form["data"]
	if !(found6){
		return "data is null!"
	} else {
		data := paramData[0]
		return data
	}
	return ""
}

func GetEvent(req *http.Request) string {
	req.ParseForm()
	paramEvent, found7:= req.Form["event"]
	if !(found7){
		return "event is null!"
	} else {
		event := paramEvent[0]
		return event
	}
	return ""
}

func GetBankType(req *http.Request) string {
	req.ParseForm()
	paramBank, found8:= req.Form["bank"]
	if !(found8){
		return "bank is null!"
	}else {
		bank := paramBank[0]
		return bank
	}
	return ""
}

func GetCardNumber(req *http.Request) string {
	req.ParseForm()
	paramCardNumber, _:= req.Form["card_num"]

	if len(paramCardNumber) > 0{
		return paramCardNumber[0]
	}
	return ""
}

func GetAddr(req *http.Request) string {
	req.ParseForm()
	paramAddr, found10:= req.Form["addr"]
	if !(found10){
		return "addr is null!"
	}else {
		bank := paramAddr[0]
		return bank
	}
	return ""
}

func GetRealName(req *http.Request) string {
	req.ParseForm()
	paramRealName, found10:= req.Form["name"]
	if !(found10){
		return "name is null!"
	}else {
		bank := paramRealName[0]
		return bank
	}
	return ""
}

func GetBankId(req *http.Request)string {
	req.ParseForm()
	paramBankId, found11:= req.Form["bank_id"]
	if !(found11){

	}else {
		bank_id := paramBankId[0]
		return bank_id
	}
	return ""
}

func GetNewPassword(req *http.Request) string {
	req.ParseForm()
	paramNewPW, found12:= req.Form["new_password"]
	if !(found12){
	}else {
		bank_id := paramNewPW[0]
		return bank_id
	}
	return ""
}

