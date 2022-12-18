package repo

import (
	"go-jwt/gorm/initializers"
	"go-jwt/gorm/models"
)

func init() {
	initializers.ConnectDb()
}
func GetAll() []models.TownDistrict {
	var output []models.TownDistrict
	//initializers.DB.Find(&output)
	return output
}
