package server

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
