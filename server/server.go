/**
 * @Desc    The server is part of http-api
 * @Author  wuchuheng<root@wuchuheng.com>
 * @Blog    https://wuchuheng.com
 * @wechat  wc20030318
 * @DATE    2021/6/23
 * @Listen  MIT
 */
package server

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/urfave/cli"
	gomail "gopkg.in/mail.v2"
	"log"
	"os/exec"
	"time"
)

/**
 * 启动监测服务
 */
func RunServer(c *cli.Context) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <- ticker.C:
			getCPUSample(c)
		}
	}
}

type Process struct {
	pid int
	cpu float64
}

/**
 * 获取cpu 样本
 */
func getCPUSample(c *cli.Context) {
	touchCpu := c.String("touchCpu")
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("ps aux  | awk '$3 > %s {print $0}'" , touchCpu))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	res := out.String()
	toEmail := c.String("email")
	SendMail(res, toEmail)
}

func SendMail(content string, toEmail string)  {
	from := "helloworlddev@163.com"
	password := "YSPHALPWNDBNLBII"

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", toEmail)

	// Set E-Mail subject
	m.SetHeader("Subject", "cpu 监测预警")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", content)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.163.com", 25, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return

}