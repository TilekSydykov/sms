package entity

type Message struct {
	HasId
	Number string `json:"number"`
	Text   string `json:"text" gorm:"type:text" form:"text" binding:"required"`
	Send   bool   `json:"send" gorm:"default:false"`
	Model
}

func (Message) TableName() string {
	return "messages"
}
