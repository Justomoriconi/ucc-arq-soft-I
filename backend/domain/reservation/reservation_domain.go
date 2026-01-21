package reservation

import "time"

type Reservation struct {
	IDReservation uint64    `json:"id_reservation"`
	IDUser        uint64    `json:"id_user"`
	IDRoom        uint64    `json:"id_room"`
	CheckIn       time.Time `json:"check_in"`
	CheckOut      time.Time `json:"check_out"`
	Status        string    `json:"status"`
}

type Reservations []Reservation
