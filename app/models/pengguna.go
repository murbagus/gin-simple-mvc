package models

type Pengguna struct {
	Username string `gorm:"primaryKey"`
	Password string
}

func (Pengguna) TableName() string {
	return "pengguna"
}
