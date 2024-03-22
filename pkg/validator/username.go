package validator

import (
	"errors"
	"kms/pkg/helpers"
	"regexp"
)

func Username(username string) error {
	if len(username) == 0 {
		msg := "Username is required"
		return errors.New(msg)
	}

	rule, err := regexp.Compile(defaultUsernameRules)
	if err != nil {
		return err
	}

	if !rule.MatchString(username) {
		msg := "use 4 to 15 characters, mix letters and numbers"
		return errors.New(msg)
	}

	if helpers.IsArrayContains(usernameBlackList, username) {
		return errors.New("username is not allow")

	}

	return nil
}
