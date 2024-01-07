package entity

import "time"

type Fee struct {
	UserId    int       `gorm:"primaryKey" json:"user_id"`
	Seq       int       `gorm:"primaryKey" json:"seq"`
	Date      time.Time `gorm:"primaryKey" json:"date"`
	Departure int       ` json:"departure"`
	Arraival  int       ` json:"arraival"`
	Price     int       ` json:"price"`
}
