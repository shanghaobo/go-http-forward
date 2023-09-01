package client

import (
	"encoding/json"
	"github.com/shanghaobo/go-http-forward/utils"
)

func makeRegisterMsg(token string) ([]byte, error) {
	registerReqData := utils.MessageRegisterReqType{
		Token: token,
	}
	tmpJsonData, err := json.Marshal(registerReqData)
	if err != nil {
		return nil, err
	}

	msg := utils.MessageType{
		Type: utils.MessageRegister,
		Data: tmpJsonData,
	}
	jsonData, err := json.Marshal(msg)
	return jsonData, err
}

func makeHeartMsg() ([]byte, error) {
	msg := utils.MessageType{
		Type: utils.MessageHeart,
	}
	jsonData, err := json.Marshal(msg)
	return jsonData, err
}

func makeForwardGetRespMsg(uuid_ string, success bool, tip string) ([]byte, error) {
	registerReqData := utils.MessageHttpForwardGetRespType{
		Uuid:    uuid_,
		Success: success,
		Tip:     tip,
	}
	tmpJsonData, err := json.Marshal(registerReqData)
	if err != nil {
		return nil, err
	}

	msg := utils.MessageType{
		Type: utils.MessageHttpForward,
		Data: tmpJsonData,
	}
	jsonData, err := json.Marshal(msg)
	return jsonData, err
}
