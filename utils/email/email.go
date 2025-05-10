package email

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
)

// SMTP configuration
const (
	SmtpHost = "smtp.qq.com"
	SmtpPort = 587
	SmtpUser = "1975611740@qq.com"
	SmtpPass = "xjerlkpwgahacjhf"
)
const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Subject string

type EmailApi struct {
	Subject Subject
}

func NewCode() EmailApi {
	return EmailApi{Subject: Code}
}
func NewNote() EmailApi {
	return EmailApi{Subject: Note}
}
func NewAlarm() EmailApi {
	return EmailApi{Subject: Alarm}
}

// EmailTokenMap 用于存储邮箱和令牌的映射
var EmailTokenMap = make(map[string]string)

// GenerateToken 生成唯一令牌
func GenerateToken() (string, error) {
	token := make([]byte, 16)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

// SendEmail 发送邮件
func (e EmailApi) SendEmail(email, Body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(SmtpUser, "htt-web"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", string(e.Subject)) //邮件主题
	m.SetBody("text/plain", fmt.Sprintf("登录验证码：%s", Body))

	d := gomail.NewDialer(SmtpHost, SmtpPort, SmtpUser, SmtpPass)

	return d.DialAndSend(m)
}
