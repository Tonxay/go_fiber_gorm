package models

type User struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func (User) TableName() string {
	return "users"
}
