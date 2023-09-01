package server

import (
	"context"
	"fmt"
	"github.com/shanghaobo/go-http-forward/utils"
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
	fmt.Println("handleConnection:", conn.RemoteAddr())

	tmpBuffer := make([]byte, 0)
	readerChan := make(chan []byte, 16)
	writerChan := make(chan []byte)
	handleChan := make(chan utils.MsgHandleType)

	msgHandle := utils.MsgHandleType{
		Conn:       conn,
		ReaderChan: readerChan,
		WriterChan: writerChan,
	}

	clients[conn] = msgHandle

	ctx, cancel := context.WithCancel(context.Background())

	go utils.ReaderMessage(handleChan, msgHandle, ctx)
	go utils.WriteMessage(writerChan, conn, ctx)
	go HandleMessage(handleChan)
	defer cancel()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			delete(clients, conn)
			return
		}
		tmpBuffer = utils.Depack(append(tmpBuffer, buffer[:n]...), readerChan)
	}
}

func startServer() {
	listen, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server started, waiting for connections")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func serverMain() {
	go startHttp()
	startServer()
}
