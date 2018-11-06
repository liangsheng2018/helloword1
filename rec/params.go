package rec

import "net/http"

type Params struct {
	UserName 	string  	// 用户名称
	PassWord 	string		// 登录密码
	UserId 	 	int			// 用户ID
	Token		string		// 令牌
	Event		string		// 接口名称
	Bank 		string		// 银行名称
	Addr		string		// 地址
	Name   		string		// 真是姓名
	BankId		string		// 银行卡ID
	CardNum		string		// 银行卡号
	NewPassWord	string		// 新密码
	SessionKey	string		// 会话秘钥
	Data 		string		// 用户信息
}

// 获取请求参数
func GetParams(req *http.Request)*Params{
	var params Params

	params.UserName = GetUserName(req)				// 获取用户名称
	params.PassWord = GetPassWord(req)				// 获取登录密码
	params.UserId = GetUserId(req)					// 获取用户ID
	params.Token = GetToken(req)					// 获取令牌
	params.Event = GetEvent(req)
	params.Bank = GetBankType(req)
	params.Addr = GetAddr(req)
	params.Name = GetRealName(req)
	params.BankId = GetBankId(req)
	params.CardNum = GetCardNumber(req)
	params.NewPassWord	= GetNewPassword(req)
	params.SessionKey = GetSessionKey(req)
	params.Data = GetData(req)

	return &params
}


//func GetParams(req *http.Request)map[string]*Params{
//	params := map[string]*Params{}
//	data := Params{
//		UserName : GetUserName(req),
//		PassWord : GetPassWord(req),
//		UserId : GetUserId(req),
//		Token : GetToken(req),
//		Event : GetEvent(req),
//		Bank : GetBankType(req),
//		Addr : GetAddr(req),
//		Name : GetRealName(req),
//		BankId : GetBankId(req),
//		CardNum : GetCardNumber(req),
//		NewPassWord	: GetNewPassword(req),
//		SessionKey : GetSessionKey(req),
//		Data:GetData(req),
//	}
//	params["dt"] = &data
//
//	return params
//}

