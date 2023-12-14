package model

import "time"

type UserUpdateHistory struct {
	Id        int64     `db:"id"`
	UserId    int64     `db:"user_id"`
	EmailOld  string    `db:"email_old"`
	EmailNew  string    `db:"email_new"`
	NameOld   string    `db:"name_old"`
	NameNew   string    `db:"name_new"`
	RoleOld   int       `db:"name_old"`
	RoleNew   int       `db:"name_new"`
	CreatedAt time.Time `db:"created_at"`
}
