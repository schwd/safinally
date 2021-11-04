package entity

import (
	"time"

	"gorm.io/gorm"
)

//Screening
type Screening struct {
	gorm.Model
	SaveTime time.Time

	Symptoms        string
	Weight          float32
	Height          float32
	Temperature     float32
	PulseRate       int
	RespirationRate int

	MedRecID *uint
	MedRec   MedicalRecord

	DiseaseID *uint
	Disease   Disease

	NurseID *uint
	Nurse   Nurse
}
