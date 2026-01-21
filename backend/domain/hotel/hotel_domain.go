package hotel

type Hotel struct {
	IDHotel     uint64 `json:"id_hotel"`
	IDAdmin     uint64 `json:"id_admin"`
	Name        string `json:"name"`
	Place       string `json:"place"`
	Description string `json:"description"`
}

type Hotels []Hotel
