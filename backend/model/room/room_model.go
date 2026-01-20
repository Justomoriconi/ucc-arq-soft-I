package room

import "ucc-arq-soft-I/model/hotel"

type Room struct {
	IDRoom   uint64       `gorm:"primaryKey;autoIncrement"`
	IDHotel  uint64       `gorm:"not null;index"`
	Hotel    *hotel.Hotel `gorm:"foreignKey:IDHotel;references:IDHotel;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Name     string       `gorm:"type:varchar(120);not null"`
	RoomType string       `gorm:"type:varchar(120);not null"`
	Capacity uint64       `gorm:"not null"`
	Price    float64      `gorm:"type:decimal(10,2);not null"`
}
