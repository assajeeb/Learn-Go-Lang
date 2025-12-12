package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName  string
	SecondName string
	Birthdrate string
	CreatedAt  time.Time
}
type Admin struct {
	Email    string
	Password string
	User     // annoonomus embending
}

func NewAdmin(email, passowrd string) *Admin {

	return &Admin{
		Email:    email,
		Password: passowrd,
		User: User{
			FirstName:  "Admin",
			SecondName: "Admin",
			Birthdrate: "---",
		},
	}

}

func (u *User) OutputUserDetails() {

	fmt.Println(u.FirstName, u.SecondName, u.Birthdrate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.SecondName = ""
}

func New(firstName, secondName, birthDate string) (*User, error) {

	if firstName == "" || secondName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		FirstName:  firstName,
		SecondName: secondName,
		Birthdrate: birthDate,
		CreatedAt:  time.Now(),
	}, nil

}
