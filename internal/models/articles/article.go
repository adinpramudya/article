package articles

import (
	"gorm.io/gorm"
)

type(
	Article struct{
		gorm.Model
		Title			string			`gorm:"type:varchar(200);not null"`
		Content			string			`gorm:"type:text;not null"`
		Category		string			`gorm:"type:varchar(100);not null"`
		Status 			string 			`gorm:"type:varchar(100);not null;default:'Draft'"`
	}
)
type ArticleRequest struct {
	Title    			string 			`json:"title" binding:"required,min=20"`
	Content  			string 			`json:"content" binding:"required,min=200"`
	Category 			string 			`json:"category" binding:"required,min=3"`
	Status   			string 			`json:"status" binding:"required,oneof=Publish Draft Thrash"`
}
