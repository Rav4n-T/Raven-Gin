package models

import "strconv"

type User struct {
	ID
	Username string `json:"username" gorm:"not null;comment:用户名"`
	Mobile   string `json:"mobile" gorm:"not null;default:18888888888;comment:手机号"`
	Password string `json:"-" gorm:"not null;comment:密码"`
	Role     uint   `json:"role" gorm:"not null;default:6;default:customer;comment:角色:1-超管、2-管理员、3-客服、4-客户、5-商家、6-普通用户"`
	Timestamps
	SoftDelete
}

func (user User) GetUid() string {
	return strconv.Itoa(user.ID.ID)
}
