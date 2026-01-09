package usecase_dto

type User struct {
	ID         string `json:"id"`
	Fullname   string `json:"fullname"`
	Password   string `jsono:"password"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DateCreate string `json:"dateCreated"`
}
