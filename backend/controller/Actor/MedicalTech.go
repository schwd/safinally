package controller

import (
	"net/http"

	"github.com/Project/entity"
	"github.com/Project/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginMedicalTechPayload login body
type LoginMedicalTechPayload struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// LoginMedicalTechResponse token response
type LoginMedicalTechResponse struct {
	Token       string             `json:"token"`
	MedicalTech entity.MedicalTech `json:"medicaltech"`
}

// POST /loginMedicalTech
func LoginMedicalTech(c *gin.Context) {
	var payload LoginMedicalTechPayload
	var MedicalTech entity.MedicalTech

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา MedicalTech ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM medical_teches WHERE email = ?", payload.Email).Scan(&MedicalTech).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(MedicalTech.Pass), []byte(payload.Pass))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid medical_teches credentials"})
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

	signedToken, err := jwtWrapper.GenerateToken(MedicalTech.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginMedicalTechResponse{
		Token:       signedToken,
		MedicalTech: MedicalTech,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}
