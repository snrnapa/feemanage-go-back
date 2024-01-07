package entity

type Customer struct {
	CustomerId   int    `gorm:"primaryKey" json:"customer_id"`
	CustomerName string ` json:"customer_name"`
	CustomerInfo string ` json:"customer_info"`
}
