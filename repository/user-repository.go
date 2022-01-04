package repository

import (
	"gorm.io/gorm"
	"solar-faza/entity"
	"solar-faza/repository/database"
	"time"
)

type UserRepository interface {
	Create(user *entity.User)
	Update(user *entity.User)
	Delete(user *entity.User)
	All() []*entity.User
	GetUserByEmail(phone string) *entity.User
	GetUserById(id uint64) *entity.User
	GetUserByIdAndCountValidation(id uint64) *entity.User
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: database.DB,
	}
}

func (db *userRepository) GetUserByEmail(email string) *entity.User {
	user := &entity.User{}
	db.connection.Where("email = ?", email).First(&user)
	if user.Id == 0 {
		return nil
	}
	return user
}

func (db *userRepository) GetUserById(id uint64) *entity.User {
	user := &entity.User{}
	db.connection.Where("id = ?", id).First(&user)
	return user
}

func (db *userRepository) GetUserByIdAndCountValidation(id uint64) *entity.User {
	user := &entity.User{}
	db.connection.Where("id = ?", id).First(&user)
	res := 0
	db.connection.Raw("UPDATE users SET validation_count = validation_count + 1 WHERE id = ?", user.Id).Scan(&res)
	return user
}

func (db *userRepository) All() []*entity.User {
	var users []*entity.User
	db.connection.Find(&users)
	return users
}

func (db *userRepository) Delete(user *entity.User) {
	db.connection.Delete(user)
}

func (db *userRepository) Update(user *entity.User) {
	db.connection.Save(&user)
}

func (db *userRepository) Create(user *entity.User) {
	user.IsActive = true
	user.IsStaff = true
	user.IsSuperuser = false
	user.DateJoined = time.Now()
	user.LastLogin = time.Now()
	db.connection.Save(user)
}
