package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint      `json:"-"`
	FirstName string    `json:first_name`
	LastName  string    `json:last_name`
	Email     string    `json:email`
	Password  []byte    `json:"-"`
	Phone     string    `json:phone`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) SetCreatedAt() {
	// Lấy thời gian hiện tại và chuyển đổi sang múi giờ Asia/Ho_Chi_Minh (UTC+7)
	location, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	user.CreatedAt = time.Now().In(location)
}

func (user *User) SetUpdatedAt() {
	// Lấy thời gian hiện tại và chuyển đổi sang múi giờ Asia/Ho_Chi_Minh (UTC+7)
	location, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	user.UpdatedAt = time.Now().In(location)
}
