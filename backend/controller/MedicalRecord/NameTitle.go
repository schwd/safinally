package controller

import (
	"github.com/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func ListNameTitle(c *gin.Context) {
	var NameTitle []entity.NameTitle
	if err := entity.DB().Table("name_titles").Find(&NameTitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": NameTitle})
}
