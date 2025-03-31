package repository

import (
	"github.com/mufengalan/my-web/internal/model"
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) Create(instance *model.BaseModel) error {
	return r.db.Create(instance).Error
}

func (r *BaseRepository) Update(instance *model.BaseModel) error {
	return r.db.Save(instance).Error
}

func (r *BaseRepository) Delete(instance *model.BaseModel) error {
	if instance.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	return r.db.Delete(instance).Error
}

func (r *BaseRepository) FindByID(instance *model.BaseModel, id uint) error {
	return r.db.First(instance, id).Error
}

func (r *BaseRepository) FindAll(instances interface{}) error {
	return r.db.Find(instances).Error
}

func (r *BaseRepository) GetDB() *gorm.DB {
	return r.db
}
