package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goava/di"
	json "github.com/json-iterator/go"

	"blog/api/request"
	"blog/models"
	"blog/service"
)

const (
	notesURL = "/notes"
)

func Register(r chi.Router, dic *di.Container) {
	router := chi.NewRouter()
	router.Get(notesURL, func(w http.ResponseWriter, r *http.Request) {
		getListNotes(dic, w, r)
	})
	router.Post(notesURL, func(w http.ResponseWriter, r *http.Request) {
		addNote(dic, w, r)
	})

	r.Mount("/", router)
}

func addNote(dic *di.Container, w http.ResponseWriter, r *http.Request) {
	var note service.Note
	if err := dic.Resolve(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	n := &models.NoteDTO{}
	v, err := request.ParseHTTPRequest(n, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	n = v.(*models.NoteDTO)

	if err = note.Create(n); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getListNotes(dic *di.Container, w http.ResponseWriter, _ *http.Request) {
	var note service.Note
	if err := dic.Resolve(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	notes, err := note.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	blob, err := json.Marshal(notes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(blob)
}
