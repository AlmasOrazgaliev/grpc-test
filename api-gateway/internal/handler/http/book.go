package http

import (
	"api-gateway/internal/domain/book"
	"api-gateway/pkg/server/status"
	desc "api-gateway/proto"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type BookHandler struct {
	bookServiceClient desc.BookClient
}

func NewBookHandler(s desc.BookClient) *BookHandler {
	return &BookHandler{bookServiceClient: s}
}

func (h *BookHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.list)
	r.Post("/", h.create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})

	return r
}

func (h *BookHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.bookServiceClient.List(r.Context(), &desc.BookData{})
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *BookHandler) create(w http.ResponseWriter, r *http.Request) {
	req := book.Request{}
	if err := render.Bind(r, &req); err != nil {
		status.BadRequest(w, r, err, req)
		return
	}
	data := &desc.BookData{
		Name:  req.Name,
		Isbn:  req.ISBN,
		Genre: req.Genre,
	}
	res, err := h.bookServiceClient.Add(r.Context(), data)
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *BookHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.bookServiceClient.Get(r.Context(), &desc.BookData{Id: id})
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *BookHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := book.Request{}
	if err := render.Bind(r, &req); err != nil {
		status.BadRequest(w, r, err, req)
		return
	}
	data := &desc.BookData{
		Id:    id,
		Name:  req.Name,
		Isbn:  req.ISBN,
		Genre: req.Genre,
	}
	res, err := h.bookServiceClient.Update(r.Context(), data)
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *BookHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := h.bookServiceClient.Get(r.Context(), &desc.BookData{Id: id})
	if err != nil {
		status.InternalServerError(w, r, err)
	}
}