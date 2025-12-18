package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {

	query := `
	
	INSERT INTO users (email, password)
	VALUES (?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	HashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, HashedPassword)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	return err

}

func (u *User) ValidateUser() error {

	query := `
 SELECT password, id FROM users WHERE email= ?
 `
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.Email)

	var HashedPassword string

	err = row.Scan(&HashedPassword, &u.ID)
	if err != nil {
		return err
	}

	err = utils.CheckPasswordHash(u.Password, HashedPassword)

	return err

}

func (u *User) GenerateToken() (string, error) {
	return utils.GenerateJWtToken(u.Email, int(u.ID))
}
