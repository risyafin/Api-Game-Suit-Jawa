package entity


type User struct {
	IdGame int `json:"idGame"`
	Username string `json:"username"`
	Password string `json:"password"`
	Point string `json:"point"`
}