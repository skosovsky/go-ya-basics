package main

import (
	"errors"
	"fmt"
)

var ErrUserExists = errors.New("user exists")

type DBStorage interface {
	UserExists(email string) bool
}

func NewUser(db DBStorage, email string) error {
	if db.UserExists(email) {
		return fmt.Errorf("email: %s %w", email, ErrUserExists)
	}

	return nil
}

func main() {

}
