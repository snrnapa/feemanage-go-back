package entity

import "time"

type Project struct {
	ProjId     int       `gorm:"primaryKey" json:"proj_id"`
	ProjName   time.Time ` json:"proj_name"`
	CustomerId int       ` json:"customer_id"`
	LeaderId   string    ` json:"leader_id"`
}
