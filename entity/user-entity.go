package entity

import (
	"time"
)

type User struct {
	HasId
	UserName    string    `json:"user_name" gorm:"type:varchar(30)" form:"user_name"`
	Email       string    `json:"email" gorm:"type:varchar(100)" form:"email" binding:"required"`
	FirstName   string    `json:"first_name" gorm:"type:varchar(30)" form:"first_name"`
	LastName    string    `json:"last_name" gorm:"type:varchar(30)" form:"last_name"`
	Image       string    `json:"img" gorm:"type:varchar(1024)" form:"-"`
	Password    string    `json:"-" form:"password" binding:"required"`
	Level       uint      `json:"level" binding:"-"`
	LastLogin   time.Time `json:"-" binding:"-"`
	DateJoined  time.Time `json:"-" binding:"-"`
	IsActive    bool      `json:"-" binding:"-"`
	IsStaff     bool      `json:"-" binding:"-"`
	IsSuperuser bool      `json:"-" binding:"-"`

	ValidationCount uint64 `json:"-"`

	Model
}

func (User) TableName() string {
	return "users"
}
