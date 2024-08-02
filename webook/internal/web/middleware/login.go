package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct {
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//不需要登录校验
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}
		session := sessions.Default(ctx)
		id := session.Get("id")
		if id == nil {
			//没有登录
			ctx.AbortWithStatus(401)
			return
		}
	}
}
