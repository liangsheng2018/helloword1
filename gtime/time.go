package gtime

import (
	"strconv"
	"time"
)

// 获取时间字符串
func TimeStr(timestamp int64) string {
	timeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return timeStr
}


// 获取当前时间字符串
func CurrentTimeStr() string {
	timestamp := time.Now().Unix()
	currentTime := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return currentTime
}

// 获取时间字符串
func TimeString(timestamp string) string {
	t,_ := strconv.ParseInt(timestamp,10,64)
	timeStr := time.Unix(t,0).Format("2006-01-02 15:04:05")
	return timeStr
}
