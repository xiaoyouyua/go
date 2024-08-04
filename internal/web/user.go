package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/lianxi/V1/webook/internal/domain"
	"src/lianxi/V1/webook/internal/service"
)

type UserHandle struct {
	svc         *service.UserService
	EmailExp    *regexp.Regexp
	PasswordExp *regexp.Regexp
}

func NewUserHandle(svc *service.UserService) *UserHandle {
	const (
		emailRegexPattern    = "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
		passwordRegexPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$"
	)
	EmailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	PasswordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandle{
		svc:         svc,
		EmailExp:    EmailExp,
		PasswordExp: PasswordExp,
	}
}

func (u *UserHandle) RegisterUser(server *gin.Engine) {
	ug := server.Group("/users")
	//注册
	ug.POST("/signup", u.SignUp)
	//登录
	ug.POST("/login", u.Login)

	ug.POST("/edit", u.Edit)
	//用户信息
	ug.GET("/profile", u.Profile)
}

func (u *UserHandle) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirm_password"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {

		return
	}

	//开始校验
	isEmail, err := u.EmailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱不正确")
		return
	}
	//if req.ConfirmPassword != req.Password {
	//	ctx.String(http.StatusOK, "两次输入的密码不正确")
	//	return
	//}
	//isPassword, err := u.PasswordExp.MatchString(req.Password)
	//if err != nil {
	//	ctx.String(http.StatusOK, "系统错误")
	//	return
	//}
	//if !isPassword {
	//	ctx.String(http.StatusOK, "密码错误，密码不少于8位包含字母下划线")
	//	return
	//}
	//调用一下svc的方法
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicateEmail {
		ctx.String(200, "重复邮箱，请重新注册")
		return
	}
	if err != nil {
		ctx.String(200, "系统错误")
		return
	}
	ctx.String(200, "注册成功")

}

func (u *UserHandle) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvaLidUserOrPassword {
		ctx.String(200, "用户密码不对")
		return
	}
	if err != nil {
		ctx.String(200, "系统错误")
		return
	}
	//在这里登录成功了
	//设置session
	sess := sessions.Default(ctx)
	sess.Set("user_id", user.Id)
	sess.Save()
	ctx.String(200, "登陆成功")
	return

}
func (u *UserHandle) Edit(ctx *gin.Context) {

}
func (u *UserHandle) Profile(ctx *gin.Context) {

}
