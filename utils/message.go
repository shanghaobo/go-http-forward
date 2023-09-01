package utils

import "encoding/json"

const (
	MessageRegister    = "register"
	MessageHeart       = "heart"
	MessageHttpForward = "http-forward"
)

type MessageType struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type MessageRegisterReqType struct {
	Token string `json:"token"`
}

type MessageRegisterRespType struct {
	Success bool `json:"success"`
}

type MessageHttpForwardGetReqType struct {
	Uuid   string `json:"uuid"`
	Params string `json:"params"`
}

type MessageHttpForwardGetRespType struct {
	Uuid    string `json:"uuid"`
	Success bool   `json:"success"`
	Tip     string `json:"tip"`
}
