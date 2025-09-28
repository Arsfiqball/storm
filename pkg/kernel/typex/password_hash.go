package typex

import "fmt"

type PasswordHash string

func (p PasswordHash) Validate() error {
	if string(p) == "" {
		return fmt.Errorf("can not be empty")
	}

	return nil
}
