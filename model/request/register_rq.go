package request

import (
	validation "github.com/itgelo/ozzo-validation/v4"
	"go-echo/util"
)

type RegisterRQ struct {
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

func (req RegisterRQ) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Phone, validation.Required.Error("is required"),
			validation.Length(10, 14).Error("minimum 10 characters and maximum 14 characters"),
			validation.By(util.CheckPhoneCountry())),
		validation.Field(&req.FullName, validation.Required.Error("is required"),
			validation.Length(3, 60).Error("minimum 3 characters and maximum 60 characters")),
		validation.Field(&req.Password, validation.Required.Error("is required"), validation.By(util.CheckStrongPassword())),
	)
}
