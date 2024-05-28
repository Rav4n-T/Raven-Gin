package request

import "github.com/go-playground/validator/v10"

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

func GetErrorMsg(req interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := req.(Validator)

		for _, v := range err.(validator.ValidationErrors) {
			if isValidator {
				if msg, exist := req.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return msg
				}
			}
			return v.Error()
		}
	}
	return "参数j错误"
}
