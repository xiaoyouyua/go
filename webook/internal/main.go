package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"src/webook/internal/repository"
	"src/webook/internal/repository/dao"
	"src/webook/internal/service"
	"src/webook/internal/web"
	"src/webook/internal/web/middleware"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	u := initUser(db)
	u.RegisterRoutes(server)
	server.Run(":8080")
}

func initUser(db *gorm.DB) *web.UserHandle {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandle(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(locaLhost:13316)/webook"))
	if err != nil {
		// 我只会在初始化过程中 panic
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"https://localhost:3000"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders:    []string{},
		//是否允许你带cookie之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//你的开发的环境
				return true
			}
			//你公司的域名
			return strings.Contains(origin, "youcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//步骤1
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	//步骤3
	server.Use(middleware.NewLoginMiddlewareBuilder().Build())
	return server
}
