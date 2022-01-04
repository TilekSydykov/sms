package repository

import (
	"gorm.io/gorm"
	"solar-faza/entity"
	"solar-faza/repository/database"
)

type PostCommentRepository interface {
	Create(user *entity.PostComment)
	Update(user *entity.PostComment)
	Delete(user *entity.PostComment)
	All() []*entity.PostComment
	GetById(id uint64) *entity.PostComment
	AllFromPost(vid uint64) []*entity.PostComment
}

type postCommentRepository struct {
	connection *gorm.DB
}

func NewPostCommentRepository() *postCommentRepository {
	return &postCommentRepository{
		connection: database.DB,
	}
}

func (db *postCommentRepository) GetById(id uint64) *entity.PostComment {
	model := &entity.PostComment{}
	db.connection.Where("id = ?", id).First(&model)
	return model
}

func (db *postCommentRepository) AllFromPost(vid uint64) []*entity.PostComment {
	var model []*entity.PostComment
	db.connection.Preload("Children").Preload("User").Where("post_id = ?", vid).Find(&model)
	return model
}

func (db *postCommentRepository) All() []*entity.PostComment {
	var model []*entity.PostComment
	db.connection.Find(&model)
	return model
}

func (db *postCommentRepository) Delete(model *entity.PostComment) {
	db.connection.Delete(model)
}

func (db *postCommentRepository) Update(model *entity.PostComment) {
	db.connection.Save(&model)
}

func (db *postCommentRepository) Create(model *entity.PostComment) {
	db.connection.Save(model)
}
