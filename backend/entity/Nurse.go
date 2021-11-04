package entity

import (
	

	"gorm.io/gorm"
)

type Nurse struct {
	gorm.Model

	Name  string
	Email string `gorm:"uniqueIndex"`
	Pass  string

	// เชื่อมกับ Screenings
	Screenings []Screening `gorm:"foreignKey:NurseID"`

	// เชื่อมกับ DrugAllergies
	DrugAllergies []DrugAllergy `gorm:"foreignKey:NurseID"`
}