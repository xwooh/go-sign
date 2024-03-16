package main

import (
	"fmt"
	"time"

	env "github.com/yuanblq/go-sign/pkg/env"
	notify "github.com/yuanblq/go-sign/pkg/notify"
	sign "github.com/yuanblq/go-sign/pkg/sign"
)

func main() {
	t := time.Now()
	subject := fmt.Sprintf("京东每日签到通知 - %s", t.Format("1/2/2006"))
	content := fmt.Sprint("京东签到成功，获得京豆 2 颗")
	notify.SendEmail(
		subject, content, notify.EmailOpt{AdminEmail: env.GetEnv("ADMIN_EMAIL"), AdminName: "京东每日签到", Password: env.GetEnv("EMAIL_PASSWORD")},
	)

	sign.JD(sign.JdKey{})
}
