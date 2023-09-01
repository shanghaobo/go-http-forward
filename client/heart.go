package client

import (
	"fmt"
	"github.com/shanghaobo/go-http-forward/utils"
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
			fmt.Println("停止心跳程序")
			return
		case <-ticker.C:
			fmt.Println("time.Now().Sub(lastHeartRespTime)=", time.Now().Sub(lastHeartRespTime))
			if time.Now().Sub(lastHeartRespTime) > HeartTimeout {
				//经过三次心跳连接未收到回复，重连
				fmt.Println("心跳超时，准备重连")
				restartConnChan <- true
				fmt.Println("haha")
				return
			}
			msgHandle.WriterChan <- heartMsg

		}
	}
}
