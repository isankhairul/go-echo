package util

import (
	"errors"
	"fmt"
	validation "github.com/itgelo/ozzo-validation/v4"
	"github.com/matthewhartstonge/argon2"
	"unicode"
)

func HashPassword(password []byte) (string, error) {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded(password)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func VerifyPassword(password []byte, encoded []byte) error {
	b, err := argon2.VerifyEncoded(password, encoded)
	if err != nil {
		return err
	}

	if !b {
		return errors.New("mismatch")
	}

	return nil
}

func CheckStrongPassword() validation.RuleFunc {
	return func(value interface{}) error {
		password := value.(string)
		var uppercasePresent bool
		var lowercasePresent bool
		var numberPresent bool
		var specialCharPresent bool
		const minPassLength = 8
		const maxPassLength = 64
		var passLen int

		for _, ch := range password {
			switch {
			case unicode.IsNumber(ch):
				numberPresent = true
				passLen++
			case unicode.IsUpper(ch):
				uppercasePresent = true
				passLen++
			case unicode.IsLower(ch):
				lowercasePresent = true
				passLen++
			case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
				specialCharPresent = true
				passLen++
			case ch == ' ':
				passLen++
			}
		}

		if !lowercasePresent {
			return errors.New("lowercase letter missing")
		}
		if !uppercasePresent {
			return errors.New("uppercase letter missing")
		}
		if !numberPresent {
			return errors.New("atleast one numeric character required")
		}
		if !specialCharPresent {
			return errors.New("special character missing")
		}
		if !(minPassLength <= passLen && passLen <= maxPassLength) {
			return errors.New(fmt.Sprintf("length must be between %d to %d characters long", minPassLength, maxPassLength))
		}

		return nil
	}
}

func CheckPhoneCountry() validation.RuleFunc {
	return func(value interface{}) error {
		phone := value.(string)
		if len(phone) < 10 {
			return errors.New("minimum characters ")
		}

		if phone[0:3] != "+62" {
			return errors.New("must start with the indonesia country code +62")
		}

		return nil
	}
}
