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

type ProjectController struct {
	Db *gorm.DB
}

func (t *ProjectController) GetProject(c *gin.Context) {
	proj_id, err := strconv.Atoi(c.Query("proj_id"))
	if err != nil {
		log.Fatal(err)
	}
	var project []entity.Project
	t.Db.Where("proj_id = ?", proj_id).
		Find(&project)
	c.JSON(http.StatusOK, project)

}

type CustomerController struct {
	Db *gorm.DB
}

func (t *CustomerController) GetCustomer(c *gin.Context) {
	customer_id, err := strconv.Atoi(c.Query("customer_id"))
	if err != nil {
		log.Fatal(err)
	}
	var customer []entity.Customer
	t.Db.Where("customer_id = ?", customer_id).
		Find(&customer)
	c.JSON(http.StatusOK, customer)

}
