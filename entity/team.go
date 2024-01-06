package entity

type Team struct {
	TeamId   int    `gorm:"primaryKey" json:"team_id"`
	TeamName string ` json:"team_name"`
	ReaderId int    ` json:"reader_id"`
}
