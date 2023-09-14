package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"time"
)

func forwardHandle(c *gin.Context) {
	forwardParams := url.Values{}
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		for _, value := range values {
			log.Println("c=", key, value)
			forwardParams.Set(key, value)
		}
	}
	apiToken := forwardParams.Get("token")
	log.Println("apiToken=", apiToken)
	if apiToken != ApiToken {
		c.JSON(http.StatusOK, gin.H{
			"success": "false",
			"tip":     "认证失败",
		})
		return
	}
	forwardParams.Del("token")

	if len(clients) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": "false",
			"tip":     "没有客户端连接",
		})
		return
	}

	uuid_, msg, err := makeHttpGetForwardMsg(forwardParams.Encode())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "false",
			"tip":     "组装报文失败",
		})
	}
	createMessage(string(msg))
	forwardResLock.Lock()
	forwardRes[uuid_] = ForwardResPending
	forwardResLock.Unlock()
	for i := 0; i < 10; i++ {
		log.Println("forwardRes[uuid_]=", uuid_, forwardRes[uuid_])
		if forwardRes[uuid_] != ForwardResPending {
			if forwardRes[uuid_] == ForwardResSuccess {
				c.JSON(http.StatusOK, gin.H{
					"success": "true",
					"tip":     "转发成功",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"success": "true",
					"tip":     "转发失败",
				})
				return
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"tip":     "转发失败",
	})

}
func startHttp() {
	r := gin.Default()
	r.GET("/", forwardHandle)
	r.Run("0.0.0.0:" + BindPort)
}
