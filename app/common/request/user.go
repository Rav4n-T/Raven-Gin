package request

type Register struct {
	Name     string `json:"name" binding:"required"`
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required"`
}

func (r *Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名必填",
		"mobile.required":   "手机号必填",
		"password.required": "密码必填",
	}
}

type Login struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required"`
}

func (l *Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required":   "手机号必填",
		"mobile.mobile":     "手机号格式错误",
		"password.required": "密码必填",
	}
}
