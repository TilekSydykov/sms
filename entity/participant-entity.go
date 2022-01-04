package entity

type Participant struct {
	HasId
	Title      string `json:"title" gorm:"type:varchar(1024)" form:"title" binding:"required"`
	Desc       string `json:"desc" binding:"required"`
	IsActive   bool   `json:"is_active"`
	Views      uint64 `json:"views" sql:"DEFAULT:0"`
	Likes      uint64 `json:"likes" sql:"DEFAULT:0"`
	Stars      uint64 `json:"stars" sql:"DEFAULT:0"`
	LastVolume uint64 `json:"last_volume" sql:"DEFAULT:0"`

	Deleted bool `json:"deleted"`

	User []User `json:"users" gorm:"many2many:owners;"`
	Model
}

func (Participant) Participant() string {
	return "participants"
}
