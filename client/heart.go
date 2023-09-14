package client

import (
	"github.com/shanghaobo/go-http-forward/utils"
	"log"
	"time"
)

var heartMsg []byte
var lastHeartRespTime time.Time

const HeartInterval = 30 * time.Second
const HeartTimeout = 60 * time.Second

func init() {
	heartMsg, _ = makeHeartMsg()
}

func startHeart(msgHandle utils.MsgHandleType) {
	ticker := time.NewTicker(HeartInterval)
	defer ticker.Stop()
	lastHeartRespTime = time.Now()

	for {
		select {
		case <-clientCtx.Done():
			log.Println("停止心跳程序")
			return
		case <-ticker.C:
			log.Println("time.Now().Sub(lastHeartRespTime)=", time.Now().Sub(lastHeartRespTime))
			if time.Now().Sub(lastHeartRespTime) > HeartTimeout {
				//经过三次心跳连接未收到回复，重连
				log.Println("心跳超时，准备重连")
				restartConnChan <- true
				log.Println("haha")
				return
			}
			msgHandle.WriterChan <- heartMsg

		}
	}
}
