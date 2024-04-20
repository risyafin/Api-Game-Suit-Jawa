package entity

type User struct {
	Id_Game  int    `gorm:"not null, primaryKey, autoIncrement" json:"id_game"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Point    string `gorm:"not null" json:"point"`
}
