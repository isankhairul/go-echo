package request

import (
	validation "github.com/itgelo/ozzo-validation/v4"
	"go-echo/util"
)

type LoginRQ struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (req LoginRQ) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Phone, validation.Required.Error("is required"),
			validation.Length(10, 14).Error("minimum 10 characters and maximum 14 characters"),
			validation.By(util.CheckPhoneCountry())),
		validation.Field(&req.Password, validation.Required.Error("is required"), validation.By(util.CheckStrongPassword())),
	)
}
