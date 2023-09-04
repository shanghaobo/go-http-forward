# go-http-forward

> go实现的通过socket转发http的功能

### 使用示例

```go
package main

import (
	"github.com/shanghaobo/go-http-forward/client"
	"github.com/shanghaobo/go-http-forward/server"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	
	//启动客户端
	go func() {
		defer wg.Done()
		client.Start("127.0.0.1", "9919", "12333", "http://127.0.0.1:19000/api/toast")
	}()
	
	//启动服务端
	go func() {
		defer wg.Done()
		server.Start("9919", "12333", "111", "9919")
	}()
	
	wg.Wait()
}
```