package repository

import (
	"github.com/mufengalan/my-web/internal/model"
	"gorm.io/gorm"
)

type ExampleRepository interface {
	GetList() (list []*model.Example, err error)
}

type Example struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return &Example{
		db: db,
	}
}

func (r *Example) GetList() (list []*model.Example, err error) {
	err = r.db.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}
