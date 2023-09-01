package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shanghaobo/go-http-forward/utils"
)

func makeHeartMsg() ([]byte, error) {
	msg := utils.MessageType{
		Type: utils.MessageHeart,
	}
	jsonData, err := json.Marshal(msg)
	return jsonData, err
}

func makeHttpGetForwardMsg(params string) (string, []byte, error) {
	uuid_ := uuid.New().String()
	data := utils.MessageHttpForwardGetReqType{
		Uuid:   uuid_,
		Params: params,
	}
	tmpJsonData, err := json.Marshal(data)
	if err != nil {
		return "", nil, err
	}

	msg := utils.MessageType{
		Type: utils.MessageHttpForward,
		Data: tmpJsonData,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return "", nil, err
	}
	return uuid_, jsonData, nil
}
