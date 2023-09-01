package server

var Token string
var ApiToken string
var BindPort string

func Start(token string, apiToken string, bindPort string) {
	Token = token
	ApiToken = apiToken
	BindPort = bindPort
	serverMain()
}
