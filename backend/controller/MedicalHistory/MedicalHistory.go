package controller

import (
	"net/http"

	"github.com/Project/entity"
	"github.com/gin-gonic/gin"
)

// POST /medical_histories
func CreateMedicalHistory(c *gin.Context) {

	var medicalhistory entity.MedicalHistory
	var medicalrecord entity.MedicalRecord
	var department entity.Department
	var doctor entity.Doctor
	var disease entity.Disease

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร medicalhistory
	if err := c.ShouldBindJSON(&medicalhistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา medicalrecord ด้วย id
	if tx := entity.DB().Where("id = ?", medicalhistory.MedicalRecordID).First(&medicalrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medicalrecord not found"})
		return
	}

	// 10: ค้นหา doctor ด้วย id
	if tx := entity.DB().Where("id = ?", medicalhistory.DoctorID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
		return
	}

	// 11: ค้นหา disease ด้วย id
	if tx := entity.DB().Where("id = ?", medicalhistory.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Disease not found"})
		return
	}

	// 12: ค้นหา department ด้วย id
	if tx := entity.DB().Where("id = ?", medicalhistory.DepartmentID).First(&department); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Department not found"})
		return
	}

	// 13: สร้าง WatchVideo
	mh := entity.MedicalHistory{
		MedicalRecord: medicalrecord,            // โยงความสัมพันธ์กับ Entity medrec
		Doctor:        doctor,                   // โยงความสัมพันธ์กับ Entity Doctor
		Disease:       disease,                  // โยงความสัมพันธ์กับ Entity Disease
		Diagnosis:     medicalhistory.Diagnosis, // ตั้งค่าฟิลด์ diag
		Treatment:     medicalhistory.Treatment, // ตั้งค่าฟิลด์ treat
		Department:    department,               // โยงความสัมพันธ์กับ Entity Department
		Date:          medicalhistory.Date,      // ตั้งค่าฟิลด์ date
	}

	// 14: บันทึก
	if err := entity.DB().Create(&mh).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mh})
}

// GET /medical_histories
func ListMedicalHistories(c *gin.Context) {
	var medicalhistories []entity.MedicalHistory
	if err := entity.DB().Preload("MedicalRecord").Preload("Doctor").Preload("Disease").Preload("Department").Raw("SELECT * FROM medical_histories").Find(&medicalhistories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicalhistories})
}
