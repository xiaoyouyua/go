package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/webook/internal/domain"
	"src/webook/internal/service"
)

var ErrUserDuplicateEmail = service.ErrUserDuplicateEmail

// UserHandLer 我准备在它上面定义跟用户有关的路由
type UserHandle struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandle(svc *service.UserService) *UserHandle {
	const (
		emaiLRegexPattern    = "ada"
		passwordRegexPattern = "xxx"
	)
	return &UserHandle{
		svc:         svc,
		emailExp:    regexp.MustCompile(emaiLRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (u *UserHandle) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	//注册
	ug.POST("/signup", u.SingUp)
	//登录
	ug.POST("/login", u.Login)

	ug.POST("/edit", u.Edit)
	//用户信息
	ug.GET("/profile", u.Profile)
}

func (u *UserHandle) SingUp(c *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	// Bind 方法会根据 Content-Type 来解析你的数据到 reg
	//里面解析错了，就会直接写回一个400的错误
	if err := c.Bind(&req); err != nil {
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		c.String(500, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "邮箱错误格式不对")
		return
	}

	if req.ConfirmPassword != req.Password {
		c.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		//记录日志
		c.String(500, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	//数据库操作
	err = u.svc.SignUp(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	//errors.Is(err, service.ErrUserDuplicateEmail)
	if err == ErrUserDuplicateEmail {
		c.String(http.StatusOK, "邮箱冲突")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统异常")
		return
	}
	c.String(http.StatusOK, "注册成功")
	//fmt.Printf("req:%v\n", req)
}

func (u *UserHandle) Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(c, req.Email, req.Password)
	if err == service.ErrInvaLidUserOrPassword {
		c.String(http.StatusOK, "用户名密码不对")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}

	//步骤2
	//在这里登录成功了
	//设置session
	session := sessions.Default(c)
	//我可以随便设置值了
	//你要放在sessionn里面的值
	session.Set("user_id", user.Id)
	session.Save()
	c.String(http.StatusOK, "登录成功")
	return

}
func (u *UserHandle) Edit(c *gin.Context) {

}
func (u *UserHandle) Profile(c *gin.Context) {

}
