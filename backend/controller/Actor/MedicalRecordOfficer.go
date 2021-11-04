package controller

import (
	"net/http"

	"github.com/Project/entity"
	"github.com/Project/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginMedicalRecordOfficerPayload login body
type LoginMedicalRecordOfficerPayload struct {
	MedRecOfficer_Email string `json:"email"`
	MedRecOfficer_Pass  string `json:"pass"`
}

// LoginMedicalRecordOfficerResponse token response
type LoginMedicalRecordOfficerResponse struct {
	Token                string                      `json:"token"`
	MedicalRecordOfficer entity.MedicalRecordOfficer `json:"medicalrecordofficer"`
}

// POST /LoginMedicalRecordOfficer
func LoginMedicalRecordOfficer(c *gin.Context) {
	var payload LoginMedicalRecordOfficerPayload
	var MedicalRecordOfficer entity.MedicalRecordOfficer

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา MedicalRecordOfficer ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM medical_record_officers WHERE med_rec_officer_email = ?", payload.MedRecOfficer_Email).Scan(&MedicalRecordOfficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(MedicalRecordOfficer.MedRecOfficer_Email), []byte(payload.MedRecOfficer_Pass))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user credentials"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(MedicalRecordOfficer.MedRecOfficer_Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginMedicalRecordOfficerResponse{
		Token:       signedToken,
		MedicalRecordOfficer: MedicalRecordOfficer,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}
