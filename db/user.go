package db

import (
	"database/sql"
	"github.com/aulyarahman/bucketeer/models"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	user := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM users ORDER BY ID_USER DESC")

	if err != nil {
		return user, err
	}

	for rows.Next() {
		var users models.User
		err := rows.Scan(&users.IdUser, &users.Name, &users.Address, &users.PhoneNumber, &users.CreatedAt)

		if err != nil {
			return user, err
		}

		user.Users = append(user.Users, users)
	}
	return user, nil
}

func (db Database) AddUser(user *models.User) error {
	var id int
	var createdAt string

	query := `INSERT INTO users (name, address, phone_number) VALUES ($1, $2, $3) RETURNING id_user, created_at`
	err := db.Conn.QueryRow(query, user.Name, user.Address, user.PhoneNumber).Scan(&id, &createdAt)

	if err != nil {
		return err
	}

	user.IdUser = id
	user.CreatedAt = createdAt
	return nil
}

func (db Database) GetUserById(userId int) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE id_user = $1`
	row := db.Conn.QueryRow(query, userId)
	switch err := row.Scan(&user.IdUser, &user.Name, &user.Address, &user.PhoneNumber, &user.CreatedAt); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}

func (db Database) DeleteUser(userId int) error {
	query := `DELETE * FROM users WHERE id_user = $1`
	_, err := db.Conn.Exec(query, userId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateUser(userId int, userData models.User) (models.User, error) {
	user := models.User{}
	query := `UPDATE users SET name=$1, address=$2 WHERE id_user=$3 RETURNING id_user, name, address, created_at`
	err := db.Conn.QueryRow(query, userData.Name, userData.Address, userId).Scan(&user.Name, &user.Address, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNoMatch
		}
		return user, err
	}

	return user, nil
}
