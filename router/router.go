package router

import (
	"feemanage-go-back/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(dbConn *gorm.DB) {

	userHandler := controller.UserController{
		Db: dbConn,
	}
	deptHandler := controller.DeptController{
		Db: dbConn,
	}
	teamHandler := controller.TeamController{
		Db: dbConn,
	}

	r := gin.Default()
	// crosOriginの設定を適用する
	setCors(r)

	r.GET("/user", userHandler.ShowAllUser)
	r.GET("/dept/", deptHandler.GetDept)
	r.GET("/team/", teamHandler.GetTeam)
	r.Run(":9001")

}
