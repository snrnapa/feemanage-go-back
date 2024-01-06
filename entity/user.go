package entity

type User struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Title   string ` json:"title"`
	Content string ` json:"content"`
}
