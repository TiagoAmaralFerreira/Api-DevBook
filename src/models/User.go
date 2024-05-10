package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json: "id,omitempty"`
	Name     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Password string    `json: "password,omitempty"`
	Created  time.Time `json: "created,omitempty"`
}

func (user *User) ExecuteValidation() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pose estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O userName é obrigatório e não pose estar em branco")
	}

	if user.Email == "" {
		return errors.New("O E-mail é obrigatório e não pose estar em branco")
	}

	if user.Password == "" {
		return errors.New("A senha é obrigatória e não pose estar em branco")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
