package queries

import (
	"database/sql"
	"demo/model"
)

func CreateUser(db *sql.DB, user model.User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func GetAllUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
