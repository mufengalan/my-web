package service

import (
	"github.com/mufengalan/my-web/config/autoload"
	"github.com/mufengalan/my-web/internal/model"
	"github.com/mufengalan/my-web/internal/repository"
	"gorm.io/gorm"
)

type IExampleService interface {
	GetExampleList() ([]*model.Example, error)
}

type ExampleService struct {
	db *gorm.DB
}

func NewExampleService() IExampleService {
	return &ExampleService{
		db: autoload.GormDb,
	}
}

// GetList 获取example列表
func (s *ExampleService) GetExampleList() ([]*model.Example, error) {
	repo := repository.NewExampleRepository(s.db)
	list, err := repo.GetList()
	if err != nil {
		return nil, err
	}
	return list, nil
}
