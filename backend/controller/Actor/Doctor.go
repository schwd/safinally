package controller

import (
	"net/http"

	"github.com/Project/entity"
	"github.com/Project/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginDoctorPayload login body
type LoginDoctorPayload struct {
	Email string `json:"email"`
	Password  string `json:"password"`
}

// LoginDoctorResponse token response
type LoginDoctorResponse struct {
	Token       string		`json:"token"`
	Doctor entity.Doctor	`json:"doctor"`
}

// POST /LoginDoctor
func LoginDoctor(c *gin.Context) {
	var payload LoginDoctorPayload
	var Doctor entity.Doctor

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา Doctor ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM doctors WHERE email = ?", payload.Email).Scan(&Doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(Doctor.Password), []byte(payload.Password))
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

	signedToken, err := jwtWrapper.GenerateToken(Doctor.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginDoctorResponse{
		Token:       signedToken,
		Doctor: Doctor,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}
