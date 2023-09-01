package client

var Token string
var ForwardToUrl string

func Start(token, forwardToUrl string) {
	Token = token
	ForwardToUrl = forwardToUrl
	clientMain()
}
