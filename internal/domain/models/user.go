package models

type User struct {
	ID       int64
	Email    string
	PassHash []byte //Password will be hashed for security reasons
}
