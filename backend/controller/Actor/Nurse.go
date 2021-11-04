package controller

import (
	"net/http"

	"github.com/Project/entity"
	"github.com/Project/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginNursePayload login body
type LoginNursePayload struct {
	Email    string `json:"Email"`
	Pass string `json:"Pass"`
}

// LoginNurseResponse token response
type LoginNurseResponse struct {
	Token string       `json:"token"`
	Nurse entity.Nurse `json:"nurses"`
}

// POST /LoginNurse
func LoginNurse(c *gin.Context) {
	var payload LoginNursePayload
	var Nurse entity.Nurse

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา Nurse ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM nurses WHERE email = ?", payload.Email).Scan(&Nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(Nurse.Pass), []byte(payload.Pass))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid nurses credentials"})
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

	signedToken, err := jwtWrapper.GenerateToken(Nurse.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginNurseResponse{
		Token: signedToken,
		Nurse: Nurse,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}