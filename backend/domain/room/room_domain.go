package room

type Room struct {
	IDRoom   uint64  `json:"id_room"`
	IDHotel  uint64  `json:"id_hotel"`
	Name     string  `json:"name"`
	RoomType string  `json:"room_type"`
	Capacity uint64  `json:"capacity"`
	Price    float64 `json:"price"`
}

type Rooms []Room
