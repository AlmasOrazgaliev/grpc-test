package http

import (
	"api-gateway/internal/domain/author"
	"api-gateway/pkg/server/status"
	desc "api-gateway/proto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type AuthorHandler struct {
	authorServiceClient desc.AuthorClient
}

func NewAuthorHandler(s desc.AuthorClient) *AuthorHandler {
	return &AuthorHandler{authorServiceClient: s}
}

func (h *AuthorHandler) Routes() chi.Router {
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

func (h *AuthorHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.authorServiceClient.List(r.Context(), &desc.AuthorData{})
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *AuthorHandler) create(w http.ResponseWriter, r *http.Request) {
	req := author.Request{}
	if err := render.Bind(r, &req); err != nil {
		status.BadRequest(w, r, err, req)
		return
	}
	data := &desc.AuthorData{
		FullName:  req.FullName,
		Pseudonym: req.Pseudonym,
		Specialty: req.Specialty,
	}
	res, err := h.authorServiceClient.Add(r.Context(), data)
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *AuthorHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.authorServiceClient.Get(r.Context(), &desc.AuthorData{Id: id})
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *AuthorHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := author.Request{}
	if err := render.Bind(r, &req); err != nil {
		status.BadRequest(w, r, err, req)
		return
	}
	data := &desc.AuthorData{
		Id:        id,
		FullName:  req.FullName,
		Pseudonym: req.Pseudonym,
		Specialty: req.Specialty,
	}
	res, err := h.authorServiceClient.Update(r.Context(), data)
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)

}

func (h *AuthorHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := h.authorServiceClient.Get(r.Context(), &desc.AuthorData{Id: id})
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}
}
