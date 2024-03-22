package validator

import (
	"errors"
	"regexp"
)

func Password(password string) error {
	if len(password) == 0 {
		return errors.New("password is required")
	}

	rule, err := regexp.Compile(defaultPasswordRules)
	if err != nil {
		return err
	}

	if !rule.MatchString(password) {
		return errors.New("use 6 to 15 characters, mix letters and numbers")
	}

	return nil
}
