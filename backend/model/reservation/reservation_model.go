package reservation

import (
	"time"
	"ucc-arq-soft-I/model/room"
	"ucc-arq-soft-I/model/user"
)

type Reservation struct {
	IDReservation uint64     `gorm:"primaryKey;autoIncrement"`
	IDUser        uint64     `gorm:"not null;index"`
	IDRoom        uint64     `gorm:"not null;index"`
	User          *user.User `gorm:"foreignKey:IDUser;references:IDUser;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Room          *room.Room `gorm:"foreignKey:IDRoom;references:IDRoom;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CheckIn       time.Time  `gorm:"type:date;not null"`
	CheckOut      time.Time  `gorm:"type:date;not null"`
	Status        string     `gorm:"type:enum('CONFIRMED','CANCELLED');default:'CONFIRMED';not null"`
}
