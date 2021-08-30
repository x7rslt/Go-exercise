
package main

import (
"crypto/tls"
	"fmt"

	"gopkg.in/mail.v2"
)

func main() {
	d := mail.NewDialer("smtp.qq.com", 465, "example", "key")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
	m := mail.NewMessage()
	m.SetHeader("From", "example@qq.com")
	m.SetHeader("To", "example@qq.com")
	//抄送：
	//m.SetAddressHeader("Cc", "example@qq.com", "Dan")
	m.SetHeader("Subject", "Mail send test")
	m.SetBody("text/html", "<h1>This is a email test.<br/></h><b>No message.</b>")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}else{
		fmt.Println("Mail send success.")
	}

}