package entity

import (
	"time"

	"gorm.io/gorm"
)

//MedicalHistory

type Department struct {
	gorm.Model
	Name             string
	Building         string
	Floor            int
	MedicalHistories []MedicalHistory `gorm:"foreignKey:DepartmentID"`
}

type MedicalHistory struct {
	gorm.Model

	MedicalRecordID *uint
	MedicalRecord   MedicalRecord

	DiseaseID *uint
	Disease   Disease

	Diagnosis string
	Treatment string

	DepartmentID *uint
	Department   Department

	DoctorID *uint
	Doctor   Doctor

	Date time.Time
}
