package entity

type Dept struct {
	DeptId   int    `gorm:"primaryKey" json:"dept_id"`
	DeptName string ` json:"dept_name"`
	LeaderId int    ` json:"leader_id"`
}
