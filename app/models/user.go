package models

import "strconv"

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名"`
	Mobile   string `json:"mobile" gorm:"not null;comment:手机号"`
	Password string `json:"-" gorm:"not null;comment:密码"`
	Timestamps
	SoftDelete
}

func (user User) GetUid() string {
	return strconv.Itoa(user.ID.ID)
}
