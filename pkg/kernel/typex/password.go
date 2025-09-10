package typex

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Password string

func (p Password) Validate() error {
	if string(p) == "" {
		return fmt.Errorf("can not be empty")
	}

	return nil
}

func (p Password) Hash() (string, error) {
	return argon2idHash(string(p))
}

// argon2idParams contains the parameters needed for Argon2id hashing
type argon2idParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// defaultParams returns reasonable default parameters for Argon2id
func defaultParams() *argon2idParams {
	return &argon2idParams{
		Memory:      64 * 1024, // 64MB
		Iterations:  1,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}
}

// generateSalt generates a random salt of the specified length
func generateSalt(saltLength uint32) ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// hash generates an Argon2id hash with the given parameters and password
func hash(password string, params *argon2idParams) (string, error) {
	salt, err := generateSalt(params.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	// Format: $argon2id$v=19$m=65536,t=1,p=4$salt$hash
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		params.Memory,
		params.Iterations,
		params.Parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encodedHash, nil
}

func argon2idHash(password string) (string, error) {
	params := defaultParams()
	return hash(password, params)
}
