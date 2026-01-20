package user

type User struct {
	IDUser   uint64 `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(120);not null"`
	LastName string `gorm:"type:varchar(120);not null"`
	Email    string `gorm:"type:varchar(120);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:enum('USER','ADMIN');default:'USER';not null"`
}
