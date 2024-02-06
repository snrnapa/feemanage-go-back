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
