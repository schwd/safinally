package controller

import (
	"net/http"
	"github.com/Project/entity"
	"github.com/gin-gonic/gin"
)

func ListLabRoom(c *gin.Context) {
	var LabRoom []entity.LabRoom
	if err := entity.DB().Table("lab_rooms").Find(&LabRoom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": LabRoom})
}