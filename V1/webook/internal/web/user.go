package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandle struct {
	EmailRegexExp    *regexp.Regexp
	PasswordRegexExp *regexp.Regexp
}

func NewUserHandle() *UserHandle {
	const (
		emailRegexPattern    = "^[a-zA-Z0-9\\_-]+@[a-zA-Z0-9\\_-]+(\\.[a-zA-Z0-9\\_-]+)+$"
		passwordRegexPattern = "^(?=.*[A-Za-z])(?=.*\\d)(?=.*[$@$!%*#?&])[A-Za-z\\d$@$!%*#?&]{8,}$"
	)
	return &UserHandle{
		EmailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		PasswordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (u *UserHandle) RegisterUser(server *gin.Engine) {
	ug := server.Group("/user")
	ug.POST("/singup", u.SingUp)
	ug.POST("/longin", u.Longin)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandle) SingUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {

		return
	}
	ctx.String(200, "注册成功")

	//开始校验
	isEmail, err := u.EmailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱不正确")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入的密码不正确")
		return
	}
	isPassword, err := u.PasswordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码错误，密码不少于8位包含字母下划线")
		return
	}
}
func (u *UserHandle) Longin(ctx *gin.Context) {

}
func (u *UserHandle) Edit(ctx *gin.Context) {

}
func (u *UserHandle) Profile(ctx *gin.Context) {

}
