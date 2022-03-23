package service

import (
	"context"
	"log"
	"sync"

	"github.com/goava/di"
	json "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"blog/datastore/repo/note"
	"blog/models"
)

type (
	Note interface {
		Service
		Create(note *models.NoteDTO) error
		GetAll() ([]models.Note, error)
	}

	noteService struct {
		ctx  context.Context
		repo note.Storage
		dic  *di.Container
		wg   *sync.WaitGroup
	}
)

func NewNoteService(ctx context.Context, dic *di.Container) Note {
	return &noteService{ctx: ctx, dic: dic}
}

func (ns *noteService) Create(note *models.NoteDTO) error {
	var n models.Note
	bytes, err := json.Marshal(note.Tags)
	if err != nil {
		return err
	}
	n.Tags = string(bytes)
	n.Header = note.Header
	n.Body = note.Body
	return ns.repo.Create(&n)
}

func (ns *noteService) GetAll() ([]models.Note, error) {
	return ns.repo.GetList()
}

func (ns *noteService) Start() error {
	log.Println("Start Note service...")
	if err := ns.dic.Resolve(&ns.wg); err != nil {
		return errors.Wrap(err, "resolve application wait group filed")
	}
	var err error
	ns.repo, err = note.NewNoteStorage(ns.dic)
	if err != nil {
		return errors.Wrap(err, "Create models storage filed")
	}

	ns.createHandlerContext()

	log.Println("Note service start success.")

	return nil
}

func (ns *noteService) Stop() error {
	log.Println("Stop Note Service...")

	log.Println("Note service stopped.")

	return nil
}

func (ns *noteService) createHandlerContext() {
	ns.wg.Add(1)
	go func() {
		for {
			<-ns.ctx.Done()
			_ = ns.Stop()
			ns.wg.Done()
			return
		}
	}()

}
