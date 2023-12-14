package user

import "time"

type Role int

const (
	User  Role = 0
	Admin Role = 1
)

type UserDomain struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	Role      Role
	CreatedAt time.Time
	UpdatedAt *time.Time
}
