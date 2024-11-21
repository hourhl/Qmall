package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHash string
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Model(&User{}).Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Model(&User{}).Where("email = ?", email).First(&user).Error
	return &user, err
}
