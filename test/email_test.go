package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Jordan Wright <3287776797@qq.com>"
	e.To = []string{"2020112834@ctgu.edu.cn"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<h1>您的验证码是：<b>123456<b></h1>")
	err := e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "3287776797@qq.com", "pbutenycpwcidbia", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
