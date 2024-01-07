package entity

type Team struct {
	TeamId   int    `gorm:"primaryKey" json:"team_id"`
	TeamName string ` json:"team_name"`
	LeaderId int    ` json:"leader_id"`
}
