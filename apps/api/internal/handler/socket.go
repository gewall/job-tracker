package handler

import (
	"job-tracker/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type SocketHandler struct {
	upgrader websocket.Upgrader
	socket *utils.Socket
}

func NewSocketHandler(upgrader websocket.Upgrader, socket *utils.Socket) *SocketHandler{
	return &SocketHandler{
		upgrader: upgrader,
		socket: socket,
	}
}

func RegisterSocketHandler(r *gin.RouterGroup, h *SocketHandler) {
	r.GET("/notif", h.SendNotification)
}


func (s *SocketHandler) SendNotification(c *gin.Context) {
	conn, err := s.upgrader.Upgrade(c.Writer,c.Request,nil)
		if err != nil {
			log.Print("err upgrade", err)
			return
		}
		
		client := &utils.Client{Conn: conn, Send: make(chan []byte,256)}
		s.socket.Register <- client
		log.Println("Oke")
	
		go client.WritePump()
}