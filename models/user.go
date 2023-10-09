package models

type User struct{
	Id   int `gorm:"primaryKey"`
	Name string
}