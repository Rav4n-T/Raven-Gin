package request

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "用户名必填",
		"password.required": "密码必填",
	}
}

type RegisterWithMobile struct {
	Username string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required"`
}

func (r *RegisterWithMobile) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "用户名必填",
		"mobile.required":   "手机号必填",
		"password.required": "密码必填",
	}
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "用户名必填",
		"password.required": "密码必填",
	}
}

type LoginWithMobile struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginWithMobile) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required":   "手机号必填",
		"mobile.mobile":     "手机号格式错误",
		"password.required": "密码必填",
	}
}
