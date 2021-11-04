package controller

import (
	"github.com/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func ListHealthInsurance(c *gin.Context) {
	var HealthInsurance []entity.HealthInsurance
	if err := entity.DB().Table("health_insurances").Find(&HealthInsurance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": HealthInsurance})
}
