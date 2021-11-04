package controller

import (
    "github.com/Project/entity"

    "github.com/gin-gonic/gin"

    "net/http"
)

func CreateScreening(c *gin.Context) {

    var Screening entity.Screening
    var Nurse entity.Nurse
    var MedicalRecord entity.MedicalRecord
    var Disease entity.Disease

    // ผลลัพธ์ที่ได้จากขั้นตอนที่ 7 จะถูก bind เข้าตัวแปร Screening
    if err := c.ShouldBindJSON(&Screening); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 8: ค้นหา MedicalRecord ด้วย id
    if tx := entity.DB().Where("id = ?", Screening.MedRecID).First(&MedicalRecord); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalRecord not found"})
        return
    }

    // 9: ค้นหา Nurse ด้วย id
    if tx := entity.DB().Where("id = ?", Screening.NurseID).First(&Nurse); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Nurse not found"})
        return
    }

    // 10: ค้นหา Disease ด้วย id
    if tx := entity.DB().Where("id = ?", Screening.DiseaseID).First(&Disease); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Disease not found"})
        return
    }
    // 11: สร้าง Screening
    Sr := entity.Screening{

        SaveTime: Screening.SaveTime,

        Symptoms:        Screening.Symptoms,
        Weight:          Screening.Weight,
        Height:          Screening.Height,
        Temperature:     Screening.Temperature,
        PulseRate:       Screening.PulseRate,
        RespirationRate: Screening.RespirationRate,

        Nurse:   Nurse,
        Disease: Disease,
        MedRec:  MedicalRecord,
    }

    // 12: บันทึก
    if err := entity.DB().Create(&Sr).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": Sr})

}

//Get ListScreenings
func ListScreenings(c *gin.Context) {
    var Screenings []entity.Screening
    if err := entity.DB().Preload("Nurse").Preload("MedRec").Preload("Disease").Raw("SELECT * FROM screenings").Find(&Screenings).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": Screenings})
}