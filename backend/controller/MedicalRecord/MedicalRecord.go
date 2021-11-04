package controller

import (
	"github.com/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateMedicalRecord(c *gin.Context) {

	var MedicalRecord entity.MedicalRecord
	var MedRecOfficer entity.MedicalRecordOfficer
	var NameTitle entity.NameTitle
	var HealthInsurance entity.HealthInsurance

	if err := c.ShouldBindJSON(&MedicalRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา medrecofficer ด้วย id
	if tx := entity.DB().Where("id = ?", MedicalRecord.MedRecOfficerID).First(&MedRecOfficer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medical record officer not found"})
		return
	}

	// 10: ค้นหา nametitle ด้วย id
	if tx := entity.DB().Where("id = ?", MedicalRecord.NameTitleID).First(&NameTitle); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name title not found"})
		return
	}

	// 11: ค้นหา healthinsurance ด้วย id
	if tx := entity.DB().Where("id = ?", MedicalRecord.HealthInsuranceID).First(&HealthInsurance); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "health insurance not found"})
		return
	}
	// 12: สร้าง MedicalRecord
	mr := entity.MedicalRecord{
		Hospital_Number: MedicalRecord.Hospital_Number,
		Personal_ID:     MedicalRecord.Personal_ID,

		NameTitle: NameTitle,

		Patient_Name: MedicalRecord.Patient_Name,
		Patient_Age:  MedicalRecord.Patient_Age,
		Patient_dob:  MedicalRecord.Patient_dob,
		Patient_Tel:  MedicalRecord.Patient_Tel,

		HealthInsurance: HealthInsurance,

		MedRecOfficer: MedRecOfficer,

		Register_Date: MedicalRecord.Register_Date,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&mr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mr})
}

// GET /medical_records
func ListMedicalRecord(c *gin.Context) {
	var MedicalRecord []*entity.MedicalRecord
	if err :=
		entity.DB().Preload("MedRecOfficer").Preload("NameTitle").Preload("HealthInsurance").Table("medical_records").Find(&MedicalRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": MedicalRecord})
}
