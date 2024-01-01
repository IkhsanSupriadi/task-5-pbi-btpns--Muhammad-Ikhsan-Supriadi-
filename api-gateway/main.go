package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/auth"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/config"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/logistic"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/reverse_proxy"
)

// @title API Service
// @version 1.0
// @description API in go using Gin framework
// @host localhost:3000
// @BasePath /auth
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.GET("/docs/*any", reverse_proxy.ReverseProxy)

	authSvc := *auth.RegisterRoutes(r, &c)
	logistic.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
