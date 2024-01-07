package entity

type Project struct {
	ProjId     int    `gorm:"primaryKey" json:"proj_id"`
	ProjName   string ` json:"proj_name"`
	CustomerId int    ` json:"customer_id"`
	LeaderId   string ` json:"leader_id"`
}
