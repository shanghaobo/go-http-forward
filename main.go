package main

import (
	"github.com/shanghaobo/go-http-forward/client"
	"github.com/shanghaobo/go-http-forward/server"
	"sync"
)

// example
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		client.Start("127.0.0.1", "9919", "12333", "http://127.0.0.1:19000/api/toast")
	}()
	go func() {
		defer wg.Done()
		server.Start("9919", "12333", "111", "19009")
	}()
	wg.Wait()
}
