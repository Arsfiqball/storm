package typex

import (
	"fmt"
	"regexp"
)

type Email string

func (e Email) Validate() error {
	if string(e) == "" {
		return fmt.Errorf("can not be empty")
	}

	re := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(re, string(e))
	if err != nil {
		return fmt.Errorf("failed to validate email format: %w", err)
	}

	if !matched {
		return fmt.Errorf("format is invalid")
	}

	return nil
}

func (e Email) String() string {
	return string(e)
}
