package controller

import (
	"AliyunCloudmonitorReceiver/config"
	"AliyunCloudmonitorReceiver/public"
	"net/url"
	"strings"
)

var replaceMap = map[string]string{
	"alertName": "告警事件：", "alertState": "告警状态：", "curValue": "触发阈值：", "dimensions": "实例信息：", "instanceName": "实例标签：",
	"regionName": "所在区域：", "triggerLevel": "告警级别：", "lastTime": "持续时间：", "userId": "账号ID",
}

func buildAliAlterMsg(str, path string) (data string) {

	// 表单转map
	formMap, _ := public.Form2Map(str)

	// 拿出要用到的key，并转中文
	cnMap := make(map[string]string)

	for key, value := range formMap {

		cnValue, ok := replaceMap[key]

		if !ok {
			continue
		}

		// 将值urldecode
		str, _ = url.QueryUnescape(value)
		cnMap[cnValue] = str
	}

	if cnMap["告警状态："] == "OK" {
		data += "[坏笑][坏笑][坏笑]\n"
	} else {
		data += "[流泪][流泪][流泪]\n"
	}

	// 拼接平台信息
	for info, uri := range config.PlatformInfo {
		if strings.Contains(path, uri) {
			data += info + "\n"
		}
	}

	// 拼接data
	delete(cnMap, "账号ID")

	data += "告警事件：" + cnMap["告警事件："] + "\n"
	data += "实例信息：" + cnMap["实例信息："] + "\n"
	data += "实例标签：" + cnMap["实例标签："] + "\n"
	data += "触发阈值：" + cnMap["触发阈值："] + "\n"
	data += "触发时间：" + public.NowTime() + "\n"
	data += "持续时间：" + cnMap["持续时间："] + "\n"
	data += "所在区域：" + cnMap["所在区域："] + "\n"
	data += "告警级别：" + cnMap["告警级别："] + "\n"
	data += "告警状态：" + cnMap["告警状态："] + "\n"

	data = strings.ReplaceAll(data, "{instanceId=", "")
	// data = strings.ReplaceAll(data, ", userId=1774428049377092}", "")
	return
}
