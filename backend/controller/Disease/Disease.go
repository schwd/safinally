package controller

import (
	"net/http"

	"github.com/Project/entity"
	"github.com/gin-gonic/gin"
)

// GET ListDiseases
func ListDiseases(c *gin.Context) {
	var diseases []entity.Disease
	if err := entity.DB().Raw("SELECT * FROM diseases").Scan(&diseases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diseases})
}
