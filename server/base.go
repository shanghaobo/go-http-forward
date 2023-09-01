package server

var Addr string
var Token string
var ApiToken string
var BindPort string

func Start(addr string, token string, apiToken string, bindPort string) {
	Addr = addr
	Token = token
	ApiToken = apiToken
	BindPort = bindPort
	serverMain()
}
