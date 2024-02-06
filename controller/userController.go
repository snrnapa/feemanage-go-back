package controller

import (
	"feemanage-go-back/entity"
	"net/http"

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
