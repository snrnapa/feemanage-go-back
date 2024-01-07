package entity

import "time"

type Report struct {
	UserId   int       `gorm:"primaryKey" json:"id"`
	Date     time.Time `gorm:"primaryKey" json:"date"`
	Seq      int       `gorm:"primaryKey" json:"seq"`
	Location string    ` json:"location"`
	Start    time.Time ` json:"start"`
	End      time.Time ` json:"end"`
	RestFlg  int       ` json:"rest_flg"`
	ProjId   int       ` json:"proj_id"`
}
