package server

import (
	"log"
	"os"
)

var Port string
var Token string
var ApiToken string
var BindPort string

func Start(port string, token string, apiToken string, bindPort string) {
	Port = port
	Token = token
	ApiToken = apiToken
	BindPort = bindPort
	serverMain()
}

func SetLogPath(logPath string) {
	logFile, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(logFile)
}
