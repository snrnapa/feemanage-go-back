package entity

import "time"

type Report struct {
	UserId    int       `gorm:"primaryKey" json:"id"`
	Date      time.Time `gorm:"primaryKey" json:"date"`
	Seq       int       `gorm:"primaryKey" json:"seq"`
	Location  string    ` json:"location"`
	StartTime time.Time ` json:"start_time"`
	EndTime   time.Time ` json:"end_time"`
	RestFlg   int       ` json:"rest_flg"`
	ProjId    int       ` json:"proj_id"`
}
