package domain

import "time"

// User 领域对象
type User struct {
	Id       int
	Email    string
	Password string
	Ctime    time.Time
}
