package router

import (
	"feemanage-go-back/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(dbConn *gorm.DB) {

	userHandler := controller.UserController{
		Db: dbConn,
	}

	r := gin.Default()
	// crosOriginの設定を適用する
	setCors(r)

	r.GET("/users", userHandler.ShowAllUser)
	// r.GET("/show/:id", ShowOneBlog)
	r.Run(":9001")

}

// フロントエンドからのcrosOriginの許可設定
func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

}
