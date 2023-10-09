package services

import (
	"fiber_gorm/db"
	"fiber_gorm/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)
func PostCreate(c *fiber.Ctx) (models.User, error) {
	dbIns := db.GetDB()
intVer,err1 :=	strconv.Atoi(c.Params("id"))
	   user :=  models.User{Name:c.Params("name"),Id:intVer}
	if err := dbIns.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	if err1 != nil {
			return models.User{}, err1
		}
		return user, nil
	}
func GetRead(c *fiber.Ctx) ( []models.User, error) {
	dbIns := db.GetDB()
	  var users []models.User
	if err := dbIns.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
func UpDate(c *fiber.Ctx) ( models.User, error) {
	intVer,err1 :=	strconv.Atoi(c.Params("id"))
	 
	dbIns := db.GetDB()
	  var user models.User
	  dbIns.Model(&user).Where("id = ?", intVer).Update("name",c.Params("name"))
	if err := dbIns.Model(&user).Error; err != nil {
		return user, err
	}
		if err1 !=  nil {
		return user, err1
	}
	return user, nil
}
func Delete(c *fiber.Ctx) ( int, error) {
	intVer,err1 :=	strconv.Atoi(c.Params("id"))
	dbIns := db.GetDB()
	

	if err := dbIns.Delete(&models.User{}, intVer).Error; err != nil {
		return intVer, err
	}
		if err1 !=  nil {
		return intVer, err1
	}
	return intVer, nil
}