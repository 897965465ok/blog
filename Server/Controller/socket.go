package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Child struct {
	Hub        *Hub
	Message    chan []byte
	UUID       string
	Conn       *websocket.Conn
	IsRegister bool
}

type Hub struct {
	Children   []*Child
	Broadcast  chan []byte
	Register   chan *Child
	Unregister chan *Child
}

var hub *Hub = &Hub{
	Children:   []*Child{},
	Broadcast:  make(chan []byte),
	Register:   make(chan *Child),
	Unregister: make(chan *Child),
}

func ReadLoop(c *Child) {
	defer func() {
		c.Conn.Close()
		for i, item := range c.Hub.Children {
			if item.UUID == c.UUID {
				c.Hub.Children = append(c.Hub.Children[0:i], c.Hub.Children[(i+1):]...)
				fmt.Println(c.Hub.Children)
			}
		}

	}()
	var (
		message []byte
		err     error
	)
	for {
		if _, message, err = c.Conn.ReadMessage(); err != nil {
			// 关闭通道
			return
		}
		c.Hub.Broadcast <- message
	}

}

func WriteLoop(c *Child) {
	defer func() {
		c.Conn.Close()
		for i, item := range c.Hub.Children {
			if item.UUID == c.UUID {
				c.Hub.Children = append(c.Hub.Children[0:i], c.Hub.Children[(i+1):]...)
				fmt.Println(c.Hub.Children)
			}
		}
	}()
	for {
		select {
		case message, err := <-c.Message:
			if err {
				c.Conn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}

}
func WsCnection(ctx *gin.Context) {
	var (
		conn *websocket.Conn
		err  error
	)
	go func() {
		for {
			select {
			case child := <-hub.Register:
				//标记为已注册
				child.IsRegister = true
				// 存入孩子列表
				hub.Children = append(hub.Children, child)
			case message := <-hub.Broadcast:
				for _, item := range hub.Children {
					if item.IsRegister {
						item.Message <- message
					}
				}
			}
		}

	}()
	// 升级为ws协议
	if conn, err = upGrader.Upgrade(ctx.Writer, ctx.Request, nil); err != nil {
		return
	}
	c := &Child{
		Hub:        hub,
		Conn:       conn,
		UUID:       uuid.NewV4().String(),
		IsRegister: false,
		Message:    make(chan []byte, 256),
	}
	c.Hub.Register <- c
	go ReadLoop(c)
	go WriteLoop(c)

}
