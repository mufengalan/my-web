package model

type BaseModel struct {
	ID uint `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
}
