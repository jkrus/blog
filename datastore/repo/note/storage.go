package note

import (
	"blog/models"
)

type (
	Storage interface {
		Create(note *models.Note) error
		GetList() ([]models.Note, error)
	}
)
