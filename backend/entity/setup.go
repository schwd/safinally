package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64-project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&MedicalRecord{}, &MedicalRecordOfficer{}, &NameTitle{}, &HealthInsurance{},
		&Nurse{}, &Disease{}, &Screening{},
		&Drug{}, &DrugAllergy{},
		&Doctor{}, &Department{}, &MedicalHistory{},
		&Hospital{}, &Refer{},
		&MedicalTech{}, &LabType{}, &LabRoom{}, &LabResult{},
	)

	db = database

	//setup MedicalRecord
	PasswordMedicalRecordOfficer1, err := bcrypt.GenerateFromPassword([]byte("111111a"), 14)
	MedicalRecordOfficer1 := MedicalRecordOfficer{
		MedRecOfficer_Name:  "Rosie",
		MedRecOfficer_Email: "rosie@gmail.com",
		MedRecOfficer_Pass:  string(PasswordMedicalRecordOfficer1),
	}
	db.Model(&MedicalRecordOfficer{}).Create(&MedicalRecordOfficer1)

	PasswordMedicalRecordOfficer2, err := bcrypt.GenerateFromPassword([]byte("2222222a"), 14)
	MedicalRecordOfficer2 := MedicalRecordOfficer{
		MedRecOfficer_Name:  "Carla",
		MedRecOfficer_Email: "carla@gmail.com",
		MedRecOfficer_Pass:  string(PasswordMedicalRecordOfficer2),
	}
	db.Model(&MedicalRecordOfficer{}).Create(&MedicalRecordOfficer2)

	// setup nametitle
	NameTitle1 := NameTitle{
		Title: "นาง",
	}
	db.Model(&NameTitle{}).Create(&NameTitle1)

	NameTitle2 := NameTitle{
		Title: "นางสาว",
	}
	db.Model(&NameTitle{}).Create(&NameTitle2)

	NameTitle3 := NameTitle{
		Title: "นาย",
	}
	db.Model(&NameTitle{}).Create(&NameTitle3)

	NameTitle4 := NameTitle{
		Title: "เด็กชาย",
	}
	db.Model(&NameTitle{}).Create(&NameTitle4)

	NameTitle5 := NameTitle{
		Title: "เด็กหญิง",
	}
	db.Model(&NameTitle{}).Create(&NameTitle5)


	//setup HealthInsurance
	HealthInsurance1 := HealthInsurance{
		HealthInsurance_Name: "นักศึกษา",
		Detail:               "นักศึกษามหาวิทยาลัยเทคโนโลยีสุรนารีรักษาฟรี",
	}
	db.Model(&HealthInsurance{}).Create(&HealthInsurance1)

	HealthInsurance2 := HealthInsurance{
		HealthInsurance_Name: "บัตรทอง",
		Detail:               "สวัสดิการแห่งรัฐ 30 บาทรักษาทุกโรค",
	}
	db.Model(&HealthInsurance{}).Create(&HealthInsurance2)

	//setup MedicalRecord1
	MedicalRecord1 := MedicalRecord{
		Hospital_Number: "2001",
		Personal_ID:     "1234567891234",
		Patient_Name:    "Saifon",
		Patient_Age:     21,
		Patient_dob:     time.Now(),
		Patient_Tel:     "0823642199",
		Register_Date:   time.Now(),
		HealthInsurance: HealthInsurance2,
		MedRecOfficer:   MedicalRecordOfficer2,
		NameTitle:       NameTitle2,
	}
	db.Model(MedicalRecord{}).Create(&MedicalRecord1)

	MedicalRecord2 := MedicalRecord{
		Hospital_Number: "2002",
		Personal_ID:     "9876543210123",
		Patient_Name:    "Sainam",
		Patient_Age:     26,
		Patient_dob:     time.Now(),
		Patient_Tel:     "0987475566",
		Register_Date:   time.Now(),
		HealthInsurance: HealthInsurance1,
		MedRecOfficer:   MedicalRecordOfficer1,
		NameTitle:       NameTitle3,
	}
	db.Model(MedicalRecord{}).Create(&MedicalRecord2)

	//setup Disease
	Disease1 := Disease{
		Name:        "Dengue",
		Description: "illnesses that cause fever, aches and pains, or a rash. The most common symptom of dengue is fever",
	}
	db.Model(&Disease{}).Create(&Disease1)

	Disease2 := Disease{
		Name:        "Heart attack",
		Description: "the chest can feel like it's being pressed or squeezed by a heavy object, and pain can radiate from the chest to the jaw, neck, arms and back",
	}
	db.Model(&Disease{}).Create(&Disease2)

	Disease3 := Disease{
		Name:        "Gastritis",
		Description: "Gnawing or burning ache or pain (indigestion) in your upper abdomen that may become either worse or better with eating",
	}
	db.Model(&Disease{}).Create(&Disease3)

	//setup Nurse
	PasswordNurse1, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	Nurse1 := Nurse{
		Name:  "pakapon seepakdee",
		Email: "pakapon@gmail.com",
		Pass:  string(PasswordNurse1),
	}
	db.Model(&Nurse{}).Create(&Nurse1)

	PasswordNurse2, err := bcrypt.GenerateFromPassword([]byte("123"), 14)
	Nurse2 := Nurse{
		Name:  "kritsada papakdee",
		Email: "big16635@gmail.com",
		Pass:  string(PasswordNurse2),
	}
	db.Model(&Nurse{}).Create(&Nurse2)

	PasswordNurse3, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	Nurse3 := Nurse{
		Name:  "somying kondee",
		Email: "somying@gmail.com",
		Pass:  string(PasswordNurse3),
	}
	db.Model(&Nurse{}).Create(&Nurse3)

	
	//setup Drug
	Drug1 := Drug{
		Drug_Name:       "Amantadine",
		Drug_properties: "รักษาไข้หวัดใหญ่สายพันธุ์เอ และโรคพาร์กินสัน",
		Drug_group:      "ยาต้านไวรัส",
		Stock:           20,
	}
	db.Model(&Drug{}).Create(&Drug1)

	Drug2 := Drug{
		Drug_Name:       "Lithium",
		Drug_properties: "บรรเทาหรือป้องกันการเกิดซ้ำของอาการจากโรคอารมณ์สองขั้ว ภาวะอารมณ์ดีตื่นตัวผิดปกติ",
		Drug_group:      "ยารักษาโรคจิต (Antipsychotics)",
		Stock:           500,
	}
	db.Model(&Drug{}).Create(&Drug2)

	//setup MedicalTech
	PassMedicalTech1, err := bcrypt.GenerateFromPassword([]byte("123"), 14)
	PassMedicalTech2, err := bcrypt.GenerateFromPassword([]byte("12345"), 14)

	MedicalTech1 := MedicalTech{
		Name:  "Chalermkiet kongkapan",
		Email: "chalermkiet@gmail.com",
		Pass:  string(PassMedicalTech1),
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech1)

	MedicalTech2 := MedicalTech{
		Name:  "Somkiat Kongkapan",
		Email: "Somkiat@gmail.com",
		Pass:  string(PassMedicalTech2),
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech2)

	//setup LabType
	LabType1 := LabType{
		Name: "Blood test",
	}
	db.Model(&LabType{}).Create(&LabType1)

	LabType2 := LabType{
		Name: "Urinalysis test",
	}
	db.Model(&LabType{}).Create(&LabType2)

	//setup LabRoom
	LabRoom1 := LabRoom{
		Name:     "Hematology",
		Building: "Rattanavejjapat",
		floor:    1,
	}
	db.Model(&LabRoom{}).Create(&LabRoom1)

	LabRoom2 := LabRoom{
		Name:     "Microbiology",
		Building: "Research Center",
		floor:    7,
	}
	db.Model(&LabRoom{}).Create(&LabRoom2)

	//setup Doctor
	PasswordDoctor1, err := bcrypt.GenerateFromPassword([]byte("yhyh555"), 14)
	Doctor1 := Doctor{
		Name:     "Yohan Song",
		Tel:      "0885556699",
		Email:    "yh.s@hp.ac.th",
		Password: string(PasswordDoctor1),
	}
	db.Model(&Doctor{}).Create(&Doctor1)

	PasswordDoctor2, err := bcrypt.GenerateFromPassword([]byte("password1234"), 14)
	Doctor2 := Doctor{
		Name:     "Supot Jamsai",
		Tel:      "0877774412",
		Email:    "supot.j@hp.ac.th",
		Password: string(PasswordDoctor2),
	}
	db.Model(&Doctor{}).Create(&Doctor2)

	PasswordDoctor3, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	Doctor3 := Doctor{
		Name:     "Suchawadee Teangtrong",
		Tel:      "0644416289",
		Email:    "Suchawadee@gmail.com",
		Password: string(PasswordDoctor3),
	}
	db.Model(&Doctor{}).Create(&Doctor3)

	//setup Department
	Department1 := Department{
		Name:     "Emergency Room",
		Building: "Surapiphat",
		Floor:    1,
	}
	db.Model(&Department{}).Create(&Department1)

	Department2 := Department{
		Name:     "Radiology Department",
		Building: "Thepnipa",
		Floor:    5,
	}
	db.Model(&Department{}).Create(&Department2)

	Department3 := Department{
		Name:     "Pediatrics Department",
		Building: "Pataranavee",
		Floor:    5,
	}
	db.Model(&Department{}).Create(&Department3)

	//setup Hospital
	Hospital1 := Hospital{
		Name: "Srithanya Hospital",
		Tel:  "025287800",
	}
	db.Model(&Hospital{}).Create(&Hospital1)

	Hospital2 := Hospital{
		Name: "Suranaree Hospital",
		Tel:  "020000000",
	}
	db.Model(&Hospital{}).Create(&Hospital2)

	

}
