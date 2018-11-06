package sw

import "strconv"


// int to string
func IntStr(integer int) string {
	str :=strconv.Itoa(integer)
	return str
}

// string to int64
func StrInt64(str string) int64 {
	Int64,_ := strconv.ParseInt(str,10,64)
	return Int64
}

// int64 to string
func Int64Str(Int64 int64) string {
	str:=strconv.FormatInt(Int64,10)
	return str
}

// string to int
func StrInt(str string) int {
	Int, _:=strconv.Atoi(str)
	return Int
}

// string to float64
func StrFt64(str string) float64 {
	float,_ := strconv.ParseFloat(str,64)
	return float
}

// string to float32
func StrFt32(str string) float64 {
	float,_ := strconv.ParseFloat(str,32)
	return float
}

// float64 to string
func Ft64Str(float float64) string {
	str := strconv.FormatFloat(float,'E',-1,64)
	return str
}
