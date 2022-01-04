package repository

import (
	"gorm.io/gorm"
	"solar-faza/entity"
	"solar-faza/repository/database"
)

type ParticipantRepository interface {
	Create(user *entity.Participant)
	Update(user *entity.Participant)
	Delete(user *entity.Participant)
	All() []*entity.Participant
	Paginate(offset uint64, limit int, active bool, isDesc bool, orderByViews bool, deleted bool) ([]*entity.Participant, int64)
	PaginateByUserID(offset uint64, limit int, active bool, isDesc bool, orderByViews bool, deleted bool) ([]*entity.Participant, int64)
	GetById(id uint64) *entity.Participant
	GetByUserId(id uint64) []*entity.Participant
	SoftDelete(model *entity.Participant)
}

type participantRepository struct {
	connection *gorm.DB
}

func NewParticipantRepository() *participantRepository {
	return &participantRepository{
		connection: database.DB,
	}
}

func (db *participantRepository) GetById(id uint64) *entity.Participant {
	model := &entity.Participant{}
	db.connection.
		Where("id = ?", id).
		First(&model)
	return model
}

func (db *participantRepository) GetByIdWithJoins(id uint64) *entity.Participant {
	model := &entity.Participant{}
	db.connection.
		Joins("User").
		Where("participants.id = ?", id).
		First(&model)
	return model
}

func (db *participantRepository) GetByUserId(id uint64) []*entity.Participant {
	var model []*entity.Participant
	db.connection.
		Where("user_id = ?", id, false).
		Find(&model)
	return model
}

func (db *participantRepository) SoftDelete(model *entity.Participant) {
	db.connection.
		Where("id = ?", model.Id).
		Update("deleted", true)
}

func (db *participantRepository) All() []*entity.Participant {
	var model []*entity.Participant
	db.connection.
		Order("views desc").
		Where("deleted = ? AND is_active = ?", false, true).
		Find(&model)
	return model
}

func (db *participantRepository) Paginate(
	offset int,
	limit int,
	active bool,
	isDesc bool,
	orderByViews bool,
	deleted bool,
) ([]*entity.Participant, int64) {
	var model []*entity.Participant
	var order string
	if isDesc {
		order = "desc"
	} else {
		order = "asc"
	}
	query := db.connection.
		Where("deleted = ? AND is_active = ?", deleted, active)

	if orderByViews {
		query.Order("views " + order)
	} else {
		query.Order(order)
	}
	var total int64
	db.connection.Model(entity.Participant{}).
		Where("deleted = ? AND is_active = ?", deleted, active).Count(&total)
	query.Offset(offset)
	query.Limit(limit)
	query.Find(&model)
	return model, total
}

func (db *participantRepository) PaginateByUserID(
	offset int,
	limit int,
	isDesc bool,
	orderByViews bool,
	deleted bool,
	userId uint64,
) ([]*entity.Participant, int64) {
	var model []*entity.Participant
	var order string
	if isDesc {
		order = "desc"
	} else {
		order = "asc"
	}
	query := db.connection.
		Where("deleted = ? AND user_id = ?", deleted, userId)
	if orderByViews {
		query.Order("views " + order)
	} else {
		query.Order(order)
	}
	var total int64
	db.connection.
		Model(entity.Participant{}).
		Where("deleted = ? AND user_id = ?", deleted, userId).Count(&total)
	query.
		Offset(offset).
		Limit(limit).
		Find(&model)
	return model, total
}

func (db *participantRepository) Delete(model *entity.Participant) {
	db.connection.Delete(model)
}

func (db *participantRepository) Update(model *entity.Participant) {
	db.connection.Save(&model)
}

func (db *participantRepository) Create(model *entity.Participant) {
	db.connection.Save(model)
}
