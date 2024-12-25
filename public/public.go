package public

import (
	"AliyunCloudmonitorReceiver/config"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Form2Map form表单转map
func Form2Map(form string) (map[string]string, error) {
	m := make(map[string]string)

	// 首先分解&（sessionid=22222&token=3333 )变成sessionid=222222 和 token=3333
	pars := strings.Split(form, "&")
	for _, par := range pars {
		// 然后分解 sessionid=222222 和 token=3333
		parkv := strings.Split(par, "=")
		m[parkv[0]] = parkv[1] // 等号前面是key,后面是value
	}
	return m, nil
}

// NowTime 获取当前时间
func NowTime() string {
	t := time.Unix(time.Now().Unix(), 0)
	return t.Format("2006.01.02 15:04:05")
}

func MsgSender(data string) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 产生一个 [0, len(wecomRobotAddr)-1) 的随机整数
	randNum := rand.Intn(len(config.WecomRobotAddr))
	postData := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s","mentioned_mobile_list":["%s"]}}`, data, "")

	if _, err := http.Post(config.WecomRobotAddr[randNum], "application/json", strings.NewReader(postData)); err != nil {
		log.Println("群机器人消息推送失败，err：", err)
	}

}
