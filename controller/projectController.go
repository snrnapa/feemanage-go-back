package controller

import (
	"feemanage-go-back/entity"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
