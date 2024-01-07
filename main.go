package main

import (
	"feemanage-go-back/entity"
	"feemanage-go-back/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// dbを作成します
	db := dbInit()

	// dbをmigrateします
	db.AutoMigrate(&entity.User{}, &entity.Dept{}, &entity.Team{}, &entity.Fee{}, &entity.Customer{}, &entity.Project{}, &entity.Report{})
	router.GetRouter(db)

}

func dbInit() *gorm.DB {
	// dsn := fmt.Sprintf("%s%s", s1, s2)

	dsn := "host=localhost user=feemanageuser password=feemanageuser dbname=feemanagedb port=15432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
