package entity

type Dept struct {
	DeptId   int    `gorm:"primaryKey" json:"dept_id"`
	DeptName string ` json:"dept_name"`
	ReaderId int    ` json:"reader_id"`
}
