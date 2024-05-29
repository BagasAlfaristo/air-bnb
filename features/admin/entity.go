package admin

import (
	"time"
)

type Core struct {
	ID                 uint
	FullName           string
	Email              string
	NumberPhone        string
	Address            string
	Password           string
	KetikUlangPassword string
	ProfilePicture     string
	Token              string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type AdminModel interface {
	Register(admin Core) error
	AdminByEmail(email string) (*Core, error)
}

type AdminService interface {
	Register(admin Core) error
	Login(email string, password string) (data *Core, token string, err error)
}