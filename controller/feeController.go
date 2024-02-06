package controller

import (
	"feemanage-go-back/entity"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FeeController struct {
	Db *gorm.DB
}

func (t *FeeController) GetFee(c *gin.Context) {
	UserId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Fatal(err)
	}
	var fee []entity.Fee
	t.Db.Where("User_id = ?", UserId).
		Find(&fee)
	c.JSON(http.StatusOK, fee)

}

type ReportController struct {
	Db *gorm.DB
}

func (t *ReportController) GetReport(c *gin.Context) {
	UserId, err := strconv.Atoi(c.Query("user_id"))
	Date, err := strconv.Atoi(c.Query("date"))
	Seq, err := strconv.Atoi(c.Query("seq"))
	if err != nil {
		log.Fatal(err)
	}
	var report []entity.Report
	t.Db.Where("User_id = ?", UserId).
		Where("date = ?", Date).
		Where("seq = ?", Seq).
		Find(&report)
	c.JSON(http.StatusOK, report)

}
