package controller

import (
	"feemanage-go-back/entity"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func (u *UserController) ShowAllUser(c *gin.Context) {
	var Users []entity.User
	u.Db.Find(&Users)
	c.JSON(http.StatusOK, Users)

}

type DeptController struct {
	Db *gorm.DB
}

func (d *DeptController) GetDept(c *gin.Context) {
	dept_id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dept_id出力します")
	fmt.Println(dept_id)
	var dept []entity.Dept
	d.Db.Where("dept_id = ?", dept_id).
		Find(&dept)
	c.JSON(http.StatusOK, dept)
}

type TeamController struct {
	Db *gorm.DB
}

func (t *TeamController) GetTeam(c *gin.Context) {
	team_id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("team_id出力します")
	fmt.Println(team_id)
	var team []entity.Team
	t.Db.Where("team_id = ?", team_id).
		Find(&team)
	c.JSON(http.StatusOK, team)

}

type FeeController struct {
	Db *gorm.DB
}

func (t *FeeController) GetFee(c *gin.Context) {
	UserId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Fatal(err)
	}
	var Fee []entity.Fee
	t.Db.Where("User_id = ?", UserId).
		Find(&Fee)
	c.JSON(http.StatusOK, Fee)

}

type ReportController struct {
	Db *gorm.DB
}

func (t *ReportController) GetReport(c *gin.Context) {
	UserId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Fatal(err)
	}
	var Report []entity.Report
	t.Db.Where("User_id = ?", UserId).
		Find(&Report)
	c.JSON(http.StatusOK, Report)

}
