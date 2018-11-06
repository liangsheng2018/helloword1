package enc

import "encoding/base64"


// 加密
func Encrypt(str string) string {
	strBytes := []byte(str)
	encoded := base64.StdEncoding.EncodeToString(strBytes)
	return encoded
}

// 解密
func Decode(encoded string) string {
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	decodeStr := string(decoded)
	return decodeStr
}