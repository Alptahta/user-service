package postgresql

import "gorm.io/gorm"

type User struct {
	db gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: *db,
	}
}
