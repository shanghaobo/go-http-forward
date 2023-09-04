package server

import (
	"encoding/json"
	"fmt"
	"github.com/shanghaobo/go-http-forward/utils"
	"sync"
)

type ForwardRes int

const (
	ForwardResPending ForwardRes = 0
	ForwardResSuccess ForwardRes = 1
	ForwardResFail    ForwardRes = 9
)

var forwardRes map[string]ForwardRes
var forwardResLock sync.Mutex

func init() {
	forwardRes = make(map[string]ForwardRes)
	forwardResLock = sync.Mutex{}
}

func HandleMessage(handleChan chan utils.MsgHandleType) {
	for {
		select {
		case msgHandle := <-handleChan:
			fmt.Println("jsonData=", msgHandle.JsonData)
			switch msgHandle.JsonData.Type {
			case utils.MessageRegister:
				handleRegister(msgHandle)
			case utils.MessageHeart:
				handleHeart(msgHandle)
			case utils.MessageHttpForward:
				handleForwardGet(msgHandle)

			}
		}
	}

}

func handleRegister(msgHandle utils.MsgHandleType) {
	var msgData utils.MessageRegisterReqType
	err := json.Unmarshal(msgHandle.JsonData.Data, &msgData)
	if err != nil {
		fmt.Println("data解析失败")
		return
	}
	fmt.Println("msgData.Token=", msgData.Token)
	if msgData.Token != Token {
		msgHandle.WriterChan <- []byte(`{"type":"register", "data":{"success": false}}`)
		msgHandle.CancelFunc()
		return
	}
	msgHandle.WriterChan <- []byte(`{"type":"register", "data":{"success": true}}`)
}

func handleHeart(msgHandle utils.MsgHandleType) {
	heartRespMsg, _ := makeHeartMsg()
	msgHandle.WriterChan <- heartRespMsg
}

func handleForwardGet(msgHandle utils.MsgHandleType) {
	var msgData utils.MessageHttpForwardGetRespType
	err := json.Unmarshal(msgHandle.JsonData.Data, &msgData)
	if err != nil {
		fmt.Println("data解析失败")
		return
	}
	forwardResLock.Lock()
	if msgData.Success {
		forwardRes[msgData.Uuid] = ForwardResSuccess
	} else {
		forwardRes[msgData.Uuid] = ForwardResFail
	}
	forwardResLock.Unlock()

}
