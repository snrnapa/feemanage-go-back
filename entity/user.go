package entity

type User struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	LastName  string ` json:"last_name"`
	FirstName string ` json:"first_name"`
	DeptId    int    ` json:"dept_id"`
	TeamId    int    ` json:"team_id"`
}
