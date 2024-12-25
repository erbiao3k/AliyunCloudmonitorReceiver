package main

import (
	"AliyunCloudmonitorReceiver/router"
	"log"
)

func init() {
	log.Println("程序已启动")
	log.Println("webhook接口地址：http://domainName/alert/webhook/${id},id为各平台告警的URI")
}
func main() {
	r := router.SetupRouter()
	if err := r.Run(":8888"); err != nil {
		log.Printf("server startup failed, err:%v\n", err)
	}
}
