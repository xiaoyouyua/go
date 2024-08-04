package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct{}

func NewLoginMiddleware() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

// 步骤2
func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/users/login" ||
			c.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(c)
		id := sess.Get("userid")
		if id == nil {
			c.AbortWithStatus(401)
			return
		}
	}
}
