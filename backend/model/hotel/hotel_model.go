package hotel

import "ucc-arq-soft-I/model/user"

type Hotel struct {
	IDHotel     uint64     `gorm:"primaryKey;autoIncrement"`
	IDAdmin     uint64     `gorm:"not null;index"`
	Admin       *user.User `gorm:"foreignKey:IDAdmin;references:IDUser;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Name        string     `gorm:"type:varchar(120);not null"`
	Place       string     `gorm:"type:varchar(120);not null"`
	Description string     `gorm:"type:varchar(255);not null"`
}
