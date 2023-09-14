package client

import (
	"log"
	"os"
)

var Host string
var Port string
var Token string
var ForwardToUrl string

func Start(host, port, token, forwardToUrl string) {
	Host = host
	Port = port
	Token = token
	ForwardToUrl = forwardToUrl
	clientMain()
}

func SetLogPath(logPath string) {
	logFile, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(logFile)
}
