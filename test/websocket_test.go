package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
	"time"
)

// 服务端
func TestWEBs(t *testing.T) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
	}

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		tryTime := 1
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			fmt.Printf("%s TestWEBs func receive msg: %s,the time is %d \n", conn.RemoteAddr(), string(msg), tryTime)
			tryTime += 1
			time.Sleep(time.Second * 2)
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":12345", nil)
}

// 客户端
func TestWebSocket(t *testing.T) {

	url := "ws://localhost:12345/websocket"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	log.Printf("响应:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	tryTime := 1
	for {

		err = c.WriteMessage(websocket.TextMessage, []byte("你好,我是FunTester"))
		if err != nil {
			fmt.Println(err)
		}

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("TestWebSocket func 收到消息 第 %d: %s", tryTime, message)
		tryTime += 1
	}
	<-done
}
