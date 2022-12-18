package service

import (
	"github.com/gin-gonic/gin"
	"go-jwt/gorm/initializers"
	"go-jwt/gorm/models"
	"net/http"
	"strings"
)

func GetAll() gin.HandlerFunc {
	db := initializers.ConnectDb()
	return func(c *gin.Context) {
		var output []models.TownDistrict
		if err := db.Find(&output).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": output})
	}
}

func GetById() gin.HandlerFunc {
	db := initializers.ConnectDb()
	return func(c *gin.Context) {
		var ouput models.TownDistrict
		id := c.Param("id")
		if err := db.First(&ouput, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": ouput})
	}
}

func GetByNameAndId() gin.HandlerFunc {
	db := initializers.ConnectDb()
	return func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")
		var output models.TownDistrict
		if err := db.Raw("select * from CL_MB_TOWN_DISTRICT where LOC_NAME = ? AND ID = ?", name, id).Scan(&output).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": output})
	}
}

func Create() gin.HandlerFunc {
	db := initializers.ConnectDb()
	return func(c *gin.Context) {
		var dataItem models.TownDistrict
		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Id
		dataItem.Id = strings.TrimSpace(dataItem.Id)

		if dataItem.Id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id cannot be blank"})
			return
		}

		if err := db.Create(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}
