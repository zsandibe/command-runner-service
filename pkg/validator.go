package pkg

import "errors"

func ValidateCommandRequest(script, desc string) error {
	if len(script) == 0 {
		return errors.New("script must not be empty")
	}

	if len(desc) == 0 {
		return errors.New("description for script must not be empty")
	}
	return nil
}
