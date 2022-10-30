package auth

type ID uint
type Email string
type Password string
type HashedPassword string

type User struct {
	ID             ID
	Email          Email
	HashedPassword HashedPassword
}
