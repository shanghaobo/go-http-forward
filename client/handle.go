package client

import (
	"context"
	"encoding/json"
	"github.com/shanghaobo/go-http-forward/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func handleMessage(handleChan chan utils.MsgHandleType, ctx context.Context) {
	for {
		select {
		case msgHandle := <-handleChan:
			log.Println("handle jsonData type=", msgHandle.JsonData.Type)
			log.Println("handle jsonData=", msgHandle.JsonData)
			log.Println("handle jsonData.data=", string(msgHandle.JsonData.Data))
			switch msgHandle.JsonData.Type {
			case utils.MessageRegister:
				handleRegister(msgHandle)
			case utils.MessageHeart:
				handleHeart(msgHandle)
			case utils.MessageHttpForward:
				handleHttpForward(msgHandle)
			}
		case <-ctx.Done():
			return
		}
	}

}

func handleRegister(msgHandle utils.MsgHandleType) {
	var msgData utils.MessageRegisterRespType
	err := json.Unmarshal(msgHandle.JsonData.Data, &msgData)
	if err != nil {
		log.Println("msgData解析失败")
	}
	if msgData.Success {
		log.Println("注册成功")
		go startHeart(msgHandle)
	} else {
		log.Println("注册失败")
	}

}

func handleHeart(msgHandle utils.MsgHandleType) {
	lastHeartRespTime = time.Now()
}

func handleHttpForward(msgHandle utils.MsgHandleType) {
	var msgData utils.MessageHttpForwardGetReqType
	err := json.Unmarshal(msgHandle.JsonData.Data, &msgData)
	if err != nil {
		log.Println("msgData解析失败")
	}
	forwardParams := msgData.Params
	Url, err := url.Parse(ForwardToUrl)
	Url.RawQuery = forwardParams
	urlPath := Url.String()
	log.Println("urlPath=", urlPath)
	resp, err := http.Get(urlPath)
	//time.Sleep(3 * time.Second)
	if err != nil {
		log.Println("err=", err)
		respMsg, _ := makeForwardGetRespMsg(msgData.Uuid, false, "本机服务请求失败")
		msgHandle.WriterChan <- respMsg
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Println("body=", string(body))
	respMsg, _ := makeForwardGetRespMsg(msgData.Uuid, true, "本机服务请求成功")
	msgHandle.WriterChan <- respMsg
}
