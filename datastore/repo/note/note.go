package note

import (
	"github.com/goava/di"
	"gorm.io/gorm"

	"blog/models"
)

type (
	db struct {
		orm *gorm.DB
	}
)

func NewNoteStorage(dic *di.Container) (Storage, error) {
	var orm *gorm.DB
	if err := dic.Resolve(&orm); err != nil {
		return nil, err
	}

	return &db{
		orm: orm,
	}, nil
}

func (d *db) Create(note *models.Note) error {
	return d.orm.Create(note).Error
}

func (d *db) GetList() ([]models.Note, error) {
	var s []models.Note
	err := d.orm.Model(models.Note{}).Find(&s).Error
	return s, err
}
