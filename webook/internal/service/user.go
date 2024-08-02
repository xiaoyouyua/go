package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"src/webook/internal/domain"
	"src/webook/internal/repository"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvaLidUserOrPassword = errors.New("账号/邮箱或密码不正确")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc UserService) SignUp(ctx context.Context, u domain.User) error {
	//你要考虑加密放在哪里的问题了
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	//然后就是，存起来
	return svc.repo.Create(ctx, u)
}
func (svc UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	//先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrUserDuplicateEmail
	}
	if err != nil {
		return domain.User{}, err
	}
	//比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		//DEBUG
		return domain.User{}, ErrInvaLidUserOrPassword
	}
	return u, nil
}
