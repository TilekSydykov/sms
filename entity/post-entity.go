package entity

type Post struct {
	HasId
	Title     string `json:"title"`
	Text      string `json:"text" gorm:"type:text" form:"text" binding:"required"`

	ParticipantId uint64 `json:"participant_id"`
	Participant   *Participant `json:"participants" gorm:"foreignKey:ParticipantId;references:Id"`
	Model
}

func (Post) TableName() string {
	return "post_comments"
}
