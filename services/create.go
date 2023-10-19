package services

import (
	"fiber_gorm/db"
	"fiber_gorm/models"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PostCreate(c *fiber.Ctx) (models.User, error) {
	dbIns := db.GetDB()
	intVer, err1 := strconv.Atoi(c.Params("id"))
	user := models.User{Name: c.Params("name"), ID: int32(intVer)}
	if err := dbIns.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	if err1 != nil {
		return models.User{}, err1
	}
	return user, nil
}
func GetRead(c *fiber.Ctx) ([]models.User, error) {
	dbIns := db.GetDB()
	var users []models.User
	if err := dbIns.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
func UpDate(c *fiber.Ctx) (models.User, error) {
	intVer, err1 := strconv.Atoi(c.Params("id"))

	dbIns := db.GetDB()
	var user models.User
	if err := dbIns.Model(&user).Where("id = ?", intVer).Update("name", c.Params("name")).Error; err != nil {
		return models.User{}, err
	}
	if err1 != nil {
		return models.User{}, err1
	}

	return user, nil
}
func Delete(c *fiber.Ctx) (int, error) {
	intVer, err1 := strconv.Atoi(c.Params("id"))
	dbIns := db.GetDB()

	if err := dbIns.Delete(&models.User{}, intVer).Error; err != nil {
		return intVer, err
	}
	if err1 != nil {
		return intVer, err1
	}
	return intVer, nil
	//test
}

func InsertOnetoOne(c *fiber.Ctx) error {
	dbIns := db.GetDB()
	var user *models.Author

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	result := dbIns.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetQueryOnetoOne(c *fiber.Ctx) ([]*models.Author, error) {
	dbIns := db.GetDB()
	id, err := c.ParamsInt("id")

	var authors []*models.Author
	if err != nil {
		return authors, err
	}
	result := dbIns.Preload("Article", "id != ?", id).Find(&authors)
	if result.Error != nil {
		return authors, result.Error
	}
	var filteredAuthors []*models.Author
	for _, author := range authors {
		if author.Article.AuthorID != 0 {
			filteredAuthors = append(filteredAuthors, author)
		}
	}
	return filteredAuthors, nil
}
