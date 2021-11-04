package entity

import (
	"time"

	"gorm.io/gorm"
)

//Refer
type Hospital struct {
	gorm.Model
	Name string
	Tel  string
	// 1 Hospital เป็นเจ้าของได้หลาย Refer
	Refers []Refer `gorm:"foreignKey:HospitalID"`
}

type Refer struct {
	gorm.Model
	Date  time.Time
	Cause string
	// DoctorID ทำหน้าที่เป็น FK
	DoctorID *uint
	// เป็นข้อมูล Doctor ใช้เพื่อให้ join ตาราง
	Doctor Doctor
	// MedRecID ทำหน้าที่เป็น FK
	MedicalRecordID *uint
	// เป็นข้อมูล MedicalRecord ใช้เพื่อให้ join ตาราง
	MedicalRecord MedicalRecord

	// HospitalID ทำหน้าที่เป็น FK
	HospitalID *uint
	// เป็นข้อมูล Hospital ใช้เพื่อให้ join ตาราง
	Hospital Hospital

	DiseaseID *uint
	Disease   Disease
}
