package server

import (
	"context"
	"github.com/shanghaobo/go-http-forward/utils"
	"log"
	"net"
)

var clients = make(map[net.Conn]utils.MsgHandleType) // 存储客户端连接

func createMessage(message string) {
	for _, client := range clients {
		client.WriterChan <- []byte(message)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("handleConnection:", conn.RemoteAddr())

	tmpBuffer := make([]byte, 0)
	readerChan := make(chan []byte, 16)
	writerChan := make(chan []byte)
	handleChan := make(chan utils.MsgHandleType)

	ctx, cancel := context.WithCancel(context.Background())

	msgHandle := utils.MsgHandleType{
		Conn:       conn,
		ReaderChan: readerChan,
		WriterChan: writerChan,
		CancelFunc: cancel,
	}

	clients[conn] = msgHandle

	go utils.ReaderMessage(handleChan, msgHandle, ctx)
	go utils.WriteMessage(writerChan, conn, ctx)
	go HandleMessage(handleChan)
	defer cancel()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Client disconnected:", conn.RemoteAddr())
			delete(clients, conn)
			return
		}
		tmpBuffer = utils.Depack(append(tmpBuffer, buffer[:n]...), readerChan)
	}
}

func startServer() {
	listen, err := net.Listen("tcp", "0.0.0.0:"+Port)
	if err != nil {
		log.Println("Error listening:", err)
		return
	}
	defer listen.Close()
	log.Println("Server started, waiting for connections")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func serverMain() {
	go startHttp()
	startServer()
}
