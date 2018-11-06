package redis

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"hello/data"
	"hello/db"
	"time"
)

var redis, err = cache.NewCache("redis", `{"conn":"127.0.0.1:6379","key":"collectionName","dbNum":"0","password":""}`)

func init(){
	if err != nil {
		fmt.Println("redis error:", err)
	}
	return
}

type Cache interface {
	Get(key string) interface{}
	GetMulti(keys []string) []interface{}
	Put(key string, val interface{}, timeout time.Duration) error
	Delete(key string) error
	Incr(key string) error
	Decr(key string) error
	IsExist(key string) bool
	ClearAll() error
	StartAndGC(config string) error
}

func RedisSetUserInfo(userid string, data interface{}){
	err := redis.Put(userid, data, time.Second*10000)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func RedisGetUserInfo(userid string)interface{}{
	values := redis.Get(userid)
	return values
}

func CheckRedisKey(userid string)bool{
	if redis.IsExist(userid) == true {
		return true
	}else {
		return false
	}
}

func RedisUserInfo(userid string)*db.UserInfo{
	value := redis.Get(userid).([]byte)
	var v data.UserInfoSuccessReturn
	err = json.Unmarshal(value, &v)
	if err != nil{
		fmt.Println("err:", err)
	}
	var rds db.UserInfo

	rds.Token = v.Token
	rds.UserId = v.Data.UserId
	rds.UserName = v.Data.UserName
	rds.UserSex = v.Data.Usersex
	rds.UserAge = v.Data.Userage
	rds.Addr = v.Data.Addr
	rds.Remain = v.Data.Remain
	rds.CardCount = v.Data.CardCount
	rds.Name = v.Data.Name

	return &rds
}

func UpdateTime(username string){

}

func RemoveRedis(userid string)bool{
	redis.Delete(userid)
	if redis.IsExist(userid) == true {
		fmt.Println("删除redis失败！")
		return false
	}
	fmt.Println("删除redis成功！")
	return true
}