package entity

import "github.com/ducthangng/geofleet/user-service/service/helper"

type User struct {
	ID          string
	Fullname    string
	Username    string
	Password    string
	Address     string
	Phone       string
	DateCreated string
}

func (u *User) Verify() bool {
	if len(u.Fullname) < 8 || len(u.Username) < 8 || len(u.Phone) != 10 || len(u.Address) < 10 {
		return false
	}

	// check if Fullname or Username contain special characters.
	fullNameChecker := helper.ContainsSpecialStrict(u.Fullname)
	if !fullNameChecker {
		return fullNameChecker
	}

	usernameChecker := helper.ContainsSpecialStrict(u.Username)
	if !usernameChecker {
		return usernameChecker
	}

	return true
}
