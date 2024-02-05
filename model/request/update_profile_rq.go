package request

import (
	validation "github.com/itgelo/ozzo-validation/v4"
	"go-echo/generated"
	"go-echo/util"
)

func ValidateUpdateProfileBodyRequest(req generated.UpdateProfileBodyRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Phone, validation.Required.Error("is required"),
			validation.Length(10, 14).Error("minimum 10 characters and maximum 14 characters"),
			validation.By(util.CheckPhoneCountry())),
		validation.Field(&req.FullName, validation.Required.Error("is required"),
			validation.Length(3, 60).Error("minimum 3 characters and maximum 60 characters")),
	)
}
