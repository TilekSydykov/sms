package repository

import (
	"gorm.io/gorm"
	"solar-faza/entity"
	"solar-faza/repository/database"
)

type MessageRepository interface {
	Create(message *entity.Message)
	Update(message *entity.Message)
	Delete(message *entity.Message)
	All() []*entity.Message
	GetById(id uint64) *entity.Message
	AllFromPost(vid uint64) []*entity.Message
}

type messageRepository struct {
	connection *gorm.DB
}

func NewMessageRepository() *messageRepository {
	return &messageRepository{
		connection: database.DB,
	}
}

func (db *messageRepository) GetById(id uint64) *entity.Message {
	model := &entity.Message{}
	db.connection.Where("id = ?", id).First(&model)
	return model
}

func (db *messageRepository) All() []*entity.Message {
	var model []*entity.Message
	db.connection.Find(&model)
	return model
}

func (db *messageRepository) Delete(model *entity.Message) {
	db.connection.Delete(model)
}

func (db *messageRepository) Update(model *entity.Message) {
	db.connection.Save(&model)
}

func (db *messageRepository) Create(model *entity.Message) *entity.Message {
	db.connection.Save(model)
	return model
}
