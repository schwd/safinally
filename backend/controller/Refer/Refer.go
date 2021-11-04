package controller

import (
	"github.com/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Refer
func CreateRefer(c *gin.Context) {

	var medicalrecord entity.MedicalRecord
	var refer entity.Refer
	var doctor entity.Doctor
	var hospital entity.Hospital
	var disease entity.Disease

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร refer
	if err := c.ShouldBindJSON(&refer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Doctor ด้วย DoctorID
	if tx := entity.DB().Where("id = ?", refer.DoctorID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
		return
	}

	// 10: ค้นหา MedicalRecord ด้วย MedicalRecordID
	if tx := entity.DB().Where("id = ?", refer.MedicalRecordID).First(&medicalrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalRecord not found"})
		return
	}

	// 11: ค้นหา Hospital ด้วย HospitalID
	if tx := entity.DB().Where("id = ?", refer.HospitalID).First(&hospital); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hospital not found"})
		return
	}

	// 11: ค้นหา Disease ด้วย DiseaseID
	if tx := entity.DB().Where("id = ?", refer.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Disease not found"})
		return
	}

	// 12: สร้าง Refer
	wv := entity.Refer{
		MedicalRecord: medicalrecord, // โยงความสัมพันธ์กับ Entity medrec
		Doctor:        doctor,        // โยงความสัมพันธ์กับ Entity Doctor
		Hospital:      hospital,	// โยงความสัมพันธ์กับ Entity Hospital
		Cause:         refer.Cause, // ตั้งค่าฟิลด์ diag
		Date:          refer.Date,  // ตั้งค่าฟิลด์ date
		Disease:       disease,		// โยงความสัมพันธ์กับ Entity Disease
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}


// GET /Refer
func ListRefer(c *gin.Context) {
	var refers []entity.Refer
	if err := entity.DB().Preload("MedicalRecord").Preload("Doctor").Preload("Hospital").Preload("Disease").Raw("SELECT * FROM refers").Find(&refers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": refers})
}
