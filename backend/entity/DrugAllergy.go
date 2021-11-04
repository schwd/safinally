package entity

import (
	"time"

	"gorm.io/gorm"
)

//DrugAllergy
type Drug struct {
	gorm.Model
	Drug_Name       string
	Drug_properties string
	Drug_group      string
	Stock           uint
	DrugAllergies   []DrugAllergy `gorm:"foreignKey:DrugID"`
}

type DrugAllergy struct {
	gorm.Model

	MedicalRecordID *uint
	MedicalRecord   MedicalRecord

	DrugID *uint
	Drug   Drug

	DrugAllergy string

	NurseID *uint
	Nurse   Nurse

	AddedTime time.Time
}
