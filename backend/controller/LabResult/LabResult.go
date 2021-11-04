package controller

import (
	"net/http"
	"github.com/Project/entity"
	"github.com/gin-gonic/gin"
) 

// POST /LabResult
func CreateLabResult(c *gin.Context) {
	var MedicalTech entity.MedicalTech
	var MedicalRecord entity.MedicalRecord
	var LabType entity.LabType
	var LabRoom entity.LabRoom
	var LabResult entity.LabResult

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร LabResult
	if err := c.ShouldBindJSON(&LabResult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา MedicalTech ด้วย id
	if tx := entity.DB().Where("id = ?", LabResult.MedicalTechID).First(&MedicalTech); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalTech not found"})
		return
	}

	// 11: ค้นหา MedicalRecord ด้วย id
	if tx := entity.DB().Where("id = ?", LabResult.MedicalRecordID).First(&MedicalRecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalRecord not found"})
		return
	}

	// 12: ค้นหา LabType ด้วย id
	if tx := entity.DB().Where("id = ?", LabResult.LabTypeID).First(&LabType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "LabType not found"})
		return
	}
	//13: ค้นหา LabRoom ด้วย id
	if tx := entity.DB().Where("id = ?", LabResult.LabRoomID).First(&LabRoom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "LabRoom not found"})
		return
	}
	// 14: สร้าง LabResult
	lr := entity.LabResult{
		MedicalTech:  MedicalTech,             
		MedicalRecord:       MedicalRecord,                  
		LabType:    LabType,
		Lab_Result: LabResult.Lab_Result,    
		Lab_Detail: LabResult.Lab_Detail,    
		LabRoom: LabRoom,       
		AddedTime: LabResult.AddedTime,
	}

	// 15: บันทึก
	if err := entity.DB().Create(&lr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lr})
}

// GET: /api/ListLabResult
func ListLabResult(c *gin.Context) {
	var LabResult []*entity.LabResult
	if err := entity.DB().Preload("MedicalTech").Preload("MedicalRecord").Preload("LabType").Preload("LabRoom").Table("lab_results").Find(&LabResult).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": LabResult})
}





