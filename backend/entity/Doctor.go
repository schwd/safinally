package entity

import (

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Name             string
	Tel              string
	Email            string `gorm:"uniqueIndex"`
	Password         string

	//เชื่อมกับ MedicalHistories
	MedicalHistories []MedicalHistory `gorm:"foreignKey:DoctorID"`

	//เชื่อมกับ refers
	Refers []Refer `gorm:"foreignKey:DoctorID"`
}