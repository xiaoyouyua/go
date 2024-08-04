package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&user).Error
	return user, err
}

func (dao *UserDao) Inseter(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		const uniqueCinflictsErrNo uint16 = 1062
		if mysqlError.Number == uniqueCinflictsErrNo {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

type User struct {
	Id       int64  `gorm:"primary_key;AUTO_INCREMENT"`
	Email    string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255);not null"`

	Ctime int64 `gorm:"not null"`
	Utime int64 `gorm:"not null"`
}
