package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/smtp"
	"strings"
)
const (
	emlUser = "xcl@xxx.com"
	emlPwd  = "-------"
	emlSMTP = "smtp.xxx.com:25"
)

func push2mail(title  string ,attaFile string,mailtype string,to string,attaFileName string)bool{

	cc := ""
	sendTo := strings.Split(to, ";")
	subject := title
	boundary := "ds13difsknfsifuere134"
	mime := bytes.NewBuffer(nil)
	mime.WriteString(fmt.Sprintf("From: %s<%s>\r\nTo: %s\r\nCC: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\n", emlUser, emlUser, to, cc, subject))
	mime.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", boundary))
	mime.WriteString("Content-Description: 自动邮件\r\n")
	//邮件普通Text正文
	mime.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	mime.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	mime.WriteString("This is a multipart message in MIME format.")
	// 第一个附件


	mime.WriteString(fmt.Sprintf("\n--%s\r\n", boundary))
	mime.WriteString("Content-Type: application/octet-stream\r\n")
	mime.WriteString(fmt.Sprintf("Content-Description: %s\r\n" ,attaFileName))
	mime.WriteString("Content-Transfer-Encoding: base64\r\n")
	mime.WriteString("Content-Disposition: attachment; filename=\"" + attaFileName + "\"\r\n\r\n")
	//读取并编码文件内容
	attaData, err := ioutil.ReadFile(attaFile)
	if err != nil {
		return false
	}
	b := make([]byte, base64.StdEncoding.EncodedLen(len(attaData)))
	base64.StdEncoding.Encode(b, attaData)
	mime.Write(b)
	//邮件结束
	mime.WriteString("\r\n--" + boundary + "--\r\n\r\n")
	fmt.Println(mime.String())
	//发送相关
	smtpHost, _, err := net.SplitHostPort(emlSMTP)
	if err != nil {
		return false
	}
	auth := smtp.PlainAuth("", emlUser, emlPwd, smtpHost)
	smtp.SendMail(emlSMTP, auth, emlUser, sendTo, mime.Bytes())
}