package utils

import (
	"context"
	"encoding/json"
	"log"
	"net"
)

type MsgHandleType struct {
	Conn       net.Conn
	JsonData   MessageType
	ReaderChan chan []byte
	WriterChan chan []byte
	CancelFunc context.CancelFunc
}

func WriteMessage(writerChan chan []byte, conn net.Conn, ctx context.Context) {
	for {
		select {
		case msg := <-writerChan:
			log.Println("write msg=", string(msg))
			msgData := Enpack(msg)
			_, err := conn.Write(msgData)
			if err != nil {
				log.Println("write err, msg=", msg, "err=", err)
				return
			}
		case <-ctx.Done():
			conn.Close()
			return

		}
	}
}

func ReaderMessage(handleChan chan MsgHandleType, msgHandle MsgHandleType, ctx context.Context) {
	for {
		select {
		case msg := <-msgHandle.ReaderChan:
			log.Println("read msg=", string(msg))
			var jsonData MessageType
			err := json.Unmarshal(msg, &jsonData)
			if err != nil {
				log.Println("data解析失败")
			}
			msgHandle.JsonData = jsonData
			handleChan <- msgHandle
		case <-ctx.Done():
			return
		}
	}
}
