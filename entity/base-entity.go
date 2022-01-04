package entity

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type HasIdInterface interface {
	GetId() uint64
	SetId(uint64)
}

type HasOwnerInterface interface {
	GetOwner() *User
	GetOwnerId() uint64
}

type HasOwner struct {
	UserId uint64 `json:"user_id"`
	User   *User  `json:"user" gorm:"foreignKey:UserId;references:Id"`
}

func (h *HasOwner) GetOwner() *User {
	return h.User
}

func (h *HasOwner) GetOwnerId() uint64 {
	return h.UserId
}

type HasId struct {
	HasIdInterface `gorm:"-"`
	Id             uint64 `json:"id" gorm:"primaryKey;autoIncrement" form:"-"`
}

func (h *HasId) GetId() uint64 {
	return h.Id
}

func (h *HasId) SetId(id uint64) {
	h.Id = id
}
