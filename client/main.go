package client

import (
	"context"
	"github.com/shanghaobo/go-http-forward/utils"
	"log"
	"net"
	"sync"
	"time"
)

var restartConnChan chan bool

var clientCtx context.Context
var clientCancel context.CancelFunc

func readConnMain(conn net.Conn, readerChan chan []byte) {
	tmpBuffer := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Client disconnected:", conn.RemoteAddr())
			//stopChan <- true
			return
		}
		tmpBuffer = utils.Depack(append(tmpBuffer, buffer[:n]...), readerChan)
	}
}

func startClient(stopChan chan bool) {
	clientCtx, clientCancel = context.WithCancel(context.Background())
	defer clientCancel()

	conn, err := net.Dial("tcp", Host+":"+Port)
	if err != nil {
		log.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	readerChan := make(chan []byte, 16)
	writerChan := make(chan []byte, 16)
	handleChan := make(chan utils.MsgHandleType)

	msgHandle := utils.MsgHandleType{
		Conn:       conn,
		ReaderChan: readerChan,
		WriterChan: writerChan,
		CancelFunc: clientCancel,
	}

	readMainWg := sync.WaitGroup{}
	readMainWg.Add(1)

	go func() {
		defer log.Println("ReadConnMain end")
		defer readMainWg.Done()
		readConnMain(conn, readerChan)
	}()
	go func() {
		defer log.Println("ReaderMessage end")
		utils.ReaderMessage(handleChan, msgHandle, clientCtx)
	}()
	go func() {
		defer log.Println("WriteMessage end")
		utils.WriteMessage(writerChan, conn, clientCtx)
	}()
	go func() {
		defer log.Println("HandleMessage end")
		handleMessage(handleChan, clientCtx)
	}()

	registerMsg, err := makeRegisterMsg(Token)
	if err != nil {
		log.Println("注册报文创建失败")
		return
	}
	writerChan <- registerMsg

	go func() {
		select {
		case <-stopChan:
			conn.Close()
		}
	}()

	readMainWg.Wait()
	clientCancel()
	log.Println("client close")
}

func clientMain() {
	restartConnChan = make(chan bool)

	wg := sync.WaitGroup{}
	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			startClient(restartConnChan)
		}()
		wg.Wait()
		log.Println("5秒后重连")
		time.Sleep(5 * time.Second)
	}

}
