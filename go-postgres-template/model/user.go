package model

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Gid   string `json:"gid" gorm:"primaryKey"`
}
