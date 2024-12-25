package controller

import (
	"AliyunCloudmonitorReceiver/config"
	"AliyunCloudmonitorReceiver/public"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func WebHook(c *gin.Context) {

	requestURI := c.Request.RequestURI

	uri := c.Param("id")

	var iTag bool
	for _, id := range config.PlatformInfo {
		if strings.Contains(uri, id) {
			iTag = true
		}

	}

	if !iTag {
		c.JSON(http.StatusBadRequest, gin.H{
			"errmsg":   "The access is to an unknown platform",
			"time":     public.NowTime(),
			"yourip":   c.ClientIP(),
			"resource": c.Request.Host + requestURI,
		})
		log.Println("The access is to an unknown platform")
		c.Abort()
		return
	}

	bodyByts, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errmsg":   "get GetRawData error: " + err.Error(),
			"time":     public.NowTime(),
			"yourip":   c.ClientIP(),
			"resource": c.Request.Host + requestURI,
		})
		log.Println("get GetRawData error: " + err.Error())
		c.Abort()
		return
	}

	str := string(bodyByts)

	if len(str) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errmsg":   "empty message",
			"time":     public.NowTime(),
			"yourip":   c.ClientIP(),
			"resource": c.Request.Host + requestURI,
		})
		log.Println("empty message")
		c.Abort()
		return
	}

	log.Println("原始数据：", str)

	// 暂不推送INFO级别告警消息
	if strings.Contains(str, "triggerLevel=INFO") {
		log.Println("INFO级别的告警消息暂不推送")
		c.Abort()
		return
	}

	msg := ""

	log.Println(requestURI, config.AlertWebhook+"xod_aly")
	if strings.HasPrefix(requestURI, config.AlertWebhook+"xod_aly") {
		log.Println("此次为阿里云云监控告警")
		msg = buildAliAlterMsg(str, requestURI)
	}

	if len(msg) == 0 {
		log.Println("最终数据为空")
		c.String(http.StatusBadRequest, "最终数据为空")
		c.Abort()
		return
	}

	log.Println("最终数据：", msg)

	public.MsgSender(msg)
}
