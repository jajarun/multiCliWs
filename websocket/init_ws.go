package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	clientManager = newClientManager()
	upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// InitWebSocket 在Gin中初始化WebSocket路由
func InitWebSocket(r *gin.Engine) {
	go clientManager.start()
	// 添加WebSocket路由，支持路径参数
	r.Use(AuthWs).GET("/:client/:sign", func(c *gin.Context) {
		// 检查是否是WebSocket升级请求
		if c.GetHeader("Upgrade") == "websocket" {
			handleWebSocket(c)
			return
		}
		// 如果不是WebSocket请求，返回404或其他处理
		c.JSON(404, gin.H{"error": "Not found"})
	})
}

func handleWebSocket(c *gin.Context) {
	accountId, _ := c.GetQuery("account_id")
	userId, _ := c.GetQuery("user_id")

	fmt.Println("accountId:", accountId, "userId:", userId)

	// 升级HTTP连接为WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 这里可以添加你的WebSocket处理逻辑
	// 比如将连接添加到客户端管理器等
}

func AuthWs(c *gin.Context) {
	client := c.Param("client")
	sign := c.Param("sign")
	fmt.Println("client:", client, "sign:", sign)
	// 验证客户端类型
	validClients := map[string]bool{
		"web":      true,
		"master":   true,
		"whatsapp": true,
	}
	if !validClients[client] {
		fmt.Println("invalid client")
		c.JSON(404, gin.H{"error": "Invalid client"})
		return
	}
}
