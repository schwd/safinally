package controller

import (
	"github.com/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)


// GET /hospitals
func ListHospitals(c *gin.Context) {
	var hospitals []entity.Hospital
	if err := entity.DB().Raw("SELECT * FROM hospitals").Scan(&hospitals).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hospitals})
}

