package entity

type PostComment struct {
	HasId
	Title     string `json:"title"`
	Text      string `json:"text" gorm:"type:text" form:"text" binding:"required"`
	PostId uint64 `json:"post_id"`
	ParentId  uint64 `json:"parent_id" form:"parent_id"`

	Parent   *PostComment   `json:"parent" gorm:"foreignKey:ParentId;references:Id;" form:"-" `
	Children []*PostComment `json:"children" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;ForeignKey:ParentId;references:Id" form:"-"`
	HasOwner
	Model
}

func (PostComment) TableName() string {
	return "post_comments"
}
