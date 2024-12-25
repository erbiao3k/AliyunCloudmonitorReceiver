package config

const (
	AlertWebhook = "/alert/webhook/"
)

var (
	WecomRobotAddr = []string{
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=111",
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=222",
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=333"}

	PlatformInfo = map[string]string{
		"阿里云账号：新欧达":         "xod_aly_1755048621309203",
		"阿里云账号：newouda.com": "xod_aly_1114624318124043",
	}
)
