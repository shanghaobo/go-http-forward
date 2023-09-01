package client

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
