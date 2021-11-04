package controller

import (
	"net/http"
	"github.com/Project/entity"
	"github.com/gin-gonic/gin"
)

func CreateDrugAllergy(c *gin.Context) {

	var Nurse entity.Nurse
	var MedicalRecord entity.MedicalRecord
	var Drug entity.Drug
	var DrugAllergy entity.DrugAllergy

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร DrugAllergy
	if err := c.ShouldBindJSON(&DrugAllergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา Nurse ด้วย id
	if tx := entity.DB().Where("id = ?", DrugAllergy.NurseID).First(&Nurse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// 11: ค้นหา MedicalRecord ด้วย id
	if tx := entity.DB().Where("id = ?", DrugAllergy.MedicalRecordID).First(&MedicalRecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	// 12: ค้นหา Drug ด้วย id
	if tx := entity.DB().Where("id = ?", DrugAllergy.DrugID).First(&Drug); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	// 14: สร้าง DrugAllergy
	lr := entity.DrugAllergy{
		Nurse:         Nurse,
		MedicalRecord: MedicalRecord,
		Drug:          Drug,
		DrugAllergy:   DrugAllergy.DrugAllergy,
		AddedTime:     DrugAllergy.AddedTime,
	}

	// 15: บันทึก
	if err := entity.DB().Create(&lr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lr})
}

func ListDrugAllergy(c *gin.Context) {
	var DrugAllergy []entity.DrugAllergy
	if err := entity.DB().Preload("Nurse").Preload("MedicalRecord").Preload("Drug").Table("drug_allergies").Find(&DrugAllergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": DrugAllergy})
}
