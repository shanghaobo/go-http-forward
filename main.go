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
		client.Start("12333", "http://127.0.0.1:19000/api/toast")
	}()
	go func() {
		defer wg.Done()
		server.Start("12333", "111", "9919")
	}()
	wg.Wait()
}
