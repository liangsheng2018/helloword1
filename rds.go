package main

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

func main(){
	bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379","key":"collectionName","dbNum":"0","password":""}`)
	if err != nil {
		fmt.Println("redis error:", err)
	}
	bm.Put("test", "hello", time.Second*100)
	v := bm.Get("test")
	fmt.Println("value:", string(v.([]byte)))

}
