package entity

import "time"

type Customer struct {
	CustomerId   int       `gorm:"primaryKey" json:"customer_id"`
	CustomerName time.Time ` json:"customer_name"`
	CustomerInfo int       ` json:"customer_info"`
}
