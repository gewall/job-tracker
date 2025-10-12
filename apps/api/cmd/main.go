package main

import (
	"job-tracker/internal/config"
	"job-tracker/internal/handler"
	"job-tracker/internal/module"
	"job-tracker/internal/utils"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main () {
	cfg := config.InitConfig()

	wsUpgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request)bool {
			return true
		},
	}
	socket := utils.NewSocket()
	go socket.RunHub()

	mail := utils.NewEmail(cfg.Email.Email, cfg.Email.AppPassword, socket.Broadcast)

	db := config.NewConnect(cfg).Connect()
	
	


	r := gin.Default()
	

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		 AllowOrigins:     []string{"*"}, // jangan gunakan "*" jika pakai credentials
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // kalau ingin mengizinkan cookie/credentials
		MaxAge:           12 * time.Hour,
	}))
	
	userModule := module.InitUser(mail,cfg, db)
	jobModule := module.InitJob(db)

	public := r.Group("/")
	// publicHandler := handler.NewPublicHandler()
	// handler.RegisterPublicHandler(public, publicHandler)


	// auth := r.Group("/auth")
	// authR := handler.NewAuth(cfg)
	// handler.RegisterAuth(auth, authR)
	
	upgrader := handler.NewSocketHandler(wsUpgrader, socket)
	socketHandler := r.Group("/socket")
	handler.RegisterSocketHandler(socketHandler,upgrader)
	

	api := r.Group("/api")
	api.Use(jwt.Auth(cfg.JwtSecretKey))
	handler.RegisterUserRoutes(api,public,userModule)
	handler.RegisterJobHandler(api, jobModule)

	
	r.Run(":"+cfg.Port)
}