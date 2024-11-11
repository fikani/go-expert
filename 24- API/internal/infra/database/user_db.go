package database

import (
	"app-example/internal/entity"
	pkg_entity "app-example/pkg/entity"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db}
}

func (u *User) Create(user *entity.User) error {
	_, err := u.db.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID.String(), user.Name, user.Email, user.Password)
	return err
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) FindByID(id pkg_entity.ID) (*entity.User, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserTables(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE users (id TEXT PRIMARY KEY, name TEXT, email TEXT, password TEXT)")
	return err
}
