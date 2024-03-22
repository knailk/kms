package validator

import (
	"errors"
	"net/mail"

	normalizer "github.com/dimuska139/go-email-normalizer/v2"
)

func Email(email string) (string, error) {
	n := normalizer.NewNormalizer()
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", errors.New("invalid password")
	}

	return n.Normalize(addr.Address), nil
}
