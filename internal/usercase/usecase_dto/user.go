package usecase_dto

type User struct {
	ID         string `json:"id"`
	Fullname   string `json:"fullname"`
	Username   string `json:"username"`
	Password   string `jsono:"password"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	DateCreate string `json:"dateCreated"`
}
