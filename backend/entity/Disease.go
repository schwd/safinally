package entity

import (
	"gorm.io/gorm"
)



type Disease struct {
	gorm.Model

	Name        string
	Description string

	// เชื่อมกับ Screening
	Screenings []Screening `gorm:"foreignKey:DiseaseID"`

	//เชื่อมกับ MedicalHistory
	MedicalHistories []MedicalHistory `gorm:"foreignKey:DiseaseID"`

	//เชื่อมกับ refers
	Refers []Refer `gorm:"foreignKey:DiseaseID"`
}
