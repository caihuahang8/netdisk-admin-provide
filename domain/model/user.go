package model

type User struct {
	ID       int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Username int64  `gorm:"not_null" json:"username"`
	Password string `gorm:"not_null" json:"password"`
}
