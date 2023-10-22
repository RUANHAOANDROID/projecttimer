package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"projecttimer/utils"
	"sync"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}} // use default options
	connections = make([]*websocket.Conn, 0)
	messageChan = make(chan []byte) // 用于接收客户端消息的通道
	mutex       sync.Mutex
)

func SendMsg(msg Msg[interface{}]) {
	if connections == nil {
		println("ws server is nil")
		return
	}
	mutex.Lock()
	for _, conn := range connections {
		utils.Log.Println("WebSocket msg: ", msg)
		err := conn.WriteJSON(msg)
		if err != nil {
			utils.Log.Println("Failed to send message:", err)
			conn.Close()
		}
	}
	mutex.Unlock()
}
func handlerHoldWS(r *gin.RouterGroup) {
	r.GET("/flow", func(c *gin.Context) {
		wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			utils.Log.Print("upgrade:", err)
			return
		}
		defer func() {
			// 关闭连接并从连接列表中移除
			wsConn.Close()
			mutex.Lock()
			for i := range connections {
				if connections[i] == wsConn {
					connections = append(connections[:i], connections[i+1:]...)
					break
				}
			}
			mutex.Unlock()
		}()
		// 将连接添加到连接列表中
		mutex.Lock()
		connections = append(connections, wsConn)
		mutex.Unlock()

		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					utils.Log.Println("read:", err)
				}
				break
			}

			utils.Log.Printf("recv: %s", message)

			// 将消息发送到处理通道
			messageChan <- message
		}
	})
}
