package main

import (
	"github.com/gin-gonic/gin"
	"src/lianxi/V1/webook/internal/web"
)

func main() {
	server := gin.Default()
	u := web.UserHandle{}
	u.RegisterUser(server)
	server.Run(":8080")
}
