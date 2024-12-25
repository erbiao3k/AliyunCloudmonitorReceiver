# AliyunCloudmonitorReceiver 说明文档
基于阿里云监控的webhook，消息推送到企业微信机器人

## 运行要求
- 完善`config/config.go`中的微信机器人地址集合
- 完善`config/config.go`中的`PlatformInfo`账号集合
- 修改`controller/webhook.go`账号id的匹配前缀`strings.HasPrefix(requestURI, config.AlertWebhook+"xod_aly")`
- `alert-template.json`为可用的一份告警模板，包含阿里云常用资源
