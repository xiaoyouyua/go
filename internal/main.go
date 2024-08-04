package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"src/lianxi/V1/webook/internal/repository"
	"src/lianxi/V1/webook/internal/repository/dao"
	"src/lianxi/V1/webook/internal/service"
	"src/lianxi/V1/webook/internal/web"
	"src/lianxi/V1/webook/internal/web/middleware"
	"strings"
	"time"
)

func main() {
	db := iniDB()
	server := initWebServer()
	u := initUser(db)
	u.RegisterUser(server)
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type", "authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//步骤1
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	//步骤2
	server.Use(middleware.NewLoginMiddleware().Build())
	return server
}

func initUser(db *gorm.DB) *web.UserHandle {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandle(svc)
	return u
}

func iniDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/newdb2"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
