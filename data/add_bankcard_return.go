package data


type UserBankCardParamErrReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardParamErrReturn()map[string]*UserBankCardParamErrReturn{
	infoFailReturn := map[string]*UserBankCardParamErrReturn{}
	result := UserBankCardParamErrReturn{Data: "null", Code: "400", Message: "请求参数错误！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardNameLengthOutReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardNameLengthOutReturn()map[string]*UserBankCardNameLengthOutReturn{
	infoFailReturn := map[string]*UserBankCardNameLengthOutReturn{}
	result := UserBankCardNameLengthOutReturn{Data: "null", Code: "400", Message: "姓名不合法"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardNameSpecialReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardNameSpecialReturn()map[string]*UserBankCardNameSpecialReturn{
	infoFailReturn := map[string]*UserBankCardNameSpecialReturn{}
	result := UserBankCardNameSpecialReturn{Data: "null", Code: "400", Message: "姓名不能含特殊符号！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardBankNullReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardBankNullReturn()map[string]*UserBankCardBankNullReturn{
	infoFailReturn := map[string]*UserBankCardBankNullReturn{}
	result := UserBankCardBankNullReturn{Data: "null", Code: "400", Message: "银行账号不能为空！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardAddrLengthOutReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardAddrLengthOutReturn()map[string]*UserBankCardAddrLengthOutReturn{
	infoFailReturn := map[string]*UserBankCardAddrLengthOutReturn{}
	result := UserBankCardAddrLengthOutReturn{Data: "null", Code: "400", Message: "地址不合法"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardNumSameReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardNumSameReturn()map[string]*UserBankCardNumSameReturn{
	infoFailReturn := map[string]*UserBankCardNumSameReturn{}
	result := UserBankCardNumSameReturn{Data: "null", Code: "400", Message: "银行账户已存在！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardNumNullReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardNumNullReturn()map[string]*UserBankCardNumNullReturn{
	infoFailReturn := map[string]*UserBankCardNumNullReturn{}
	result := UserBankCardNumNullReturn{Data: "null", Code: "400", Message: "请重新输入银行卡账号！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardCountOverReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardCountOverReturn()map[string]*UserBankCardCountOverReturn{
	infoFailReturn := map[string]*UserBankCardCountOverReturn{}
	result := UserBankCardCountOverReturn{Data: "null", Code: "200", Message: "最多支持绑定5张银行卡！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserBankCardSuccessReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func BankCardSuccessReturn()map[string]*UserBankCardSuccessReturn{
	infoFailReturn := map[string]*UserBankCardSuccessReturn{}
	result := UserBankCardSuccessReturn{Data: "null", Code: "200", Message: "添加银行卡信息成功！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserDeleteBankCardErrReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func DeleteBankCardErrReturn()map[string]*UserDeleteBankCardErrReturn{
	infoFailReturn := map[string]*UserDeleteBankCardErrReturn{}
	result := UserDeleteBankCardErrReturn{Data: "null", Code: "400", Message: "默认银行卡不支持解除绑定！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserDeleteBankCardSuccessReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func DeleteBankCardSuccessReturn()map[string]*UserDeleteBankCardSuccessReturn{
	infoFailReturn := map[string]*UserDeleteBankCardSuccessReturn{}
	result := UserDeleteBankCardSuccessReturn{Data: "null", Code: "200", Message: "该卡已解除绑定！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserChangePasswordSuccessReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func ChangePasswordSuccessReturn()map[string]*UserChangePasswordSuccessReturn{
	infoFailReturn := map[string]*UserChangePasswordSuccessReturn{}
	result := UserChangePasswordSuccessReturn{Data: "null", Code: "200", Message: "ok"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}


type UserNewPasswordSameReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func NewPasswordSameReturn()map[string]*UserNewPasswordSameReturn{
	infoFailReturn := map[string]*UserNewPasswordSameReturn{}
	result := UserNewPasswordSameReturn{Data: "null", Code: "400", Message: "新密码和旧密码不能相同！"}
	infoFailReturn["result"] = &result
	return infoFailReturn
}

type UserNewPasswordNullReturn struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func NewPasswordNullReturn()map[string]*UserNewPasswordNullReturn{
	infoFailReturn := map[string]*UserNewPasswordNullReturn{}
	result := UserNewPasswordNullReturn{Data: "null", Code: "400", Message: ""}
	infoFailReturn["result"] = &result
	return infoFailReturn
}
