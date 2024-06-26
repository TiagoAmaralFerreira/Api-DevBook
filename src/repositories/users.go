package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewRepositoriesUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (repositoryUser Users) Create(user models.User) (uint64, error) {
	statement, erro := repositoryUser.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if erro != nil {
		return 0, erro
	}

	lastInsertId, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertId), nil

}
