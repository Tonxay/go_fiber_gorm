package services

import (
	"fiber_gorm/db"
	"fiber_gorm/models"
)
func PostCreate() (models.Studen, error) {
	dbIns := db.GetDB()
	   user :=  models.Studen{Frist_name: "Jin"}
	if err := dbIns.Create(&user).Error; err != nil {
		return models.Studen{}, err
	}
	return user, nil
}