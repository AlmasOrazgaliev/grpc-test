package http

import (
	"api-gateway/internal/domain/member"
	"api-gateway/pkg/server/status"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type MemberHandler struct {
	subscriptionService *subscription.Service
}

func NewMemberHandler(s *subscription.Service) *MemberHandler {
	return &MemberHandler{subscriptionService: s}
}

func (h *MemberHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.list)
	r.Post("/", h.create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
		r.Get("/books", h.listBooks)
	})

	return r
}

func (h *MemberHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.subscriptionService.ListMembers(r.Context())
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *MemberHandler) create(w http.ResponseWriter, r *http.Request) {
	req := member.Request{}
	if err := render.Bind(r, &req); err != nil {
		status.BadRequest(w, r, err, req)
		return
	}

	res, err := h.subscriptionService.CreateMember(r.Context(), req)
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *MemberHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.subscriptionService.GetMember(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			status.NotFound(w, r, err)
		default:
			status.InternalServerError(w, r, err)
		}
		return
	}

	status.OK(w, r, res)
}

func (h *MemberHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := member.Request{}
	if err := render.Bind(r, &req); err != nil {
		status.BadRequest(w, r, err, req)
		return
	}

	if err := h.subscriptionService.UpdateMember(r.Context(), id, req); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			status.NotFound(w, r, err)
		default:
			status.InternalServerError(w, r, err)
		}
		return
	}
}

func (h *MemberHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.subscriptionService.DeleteMember(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			status.NotFound(w, r, err)
		default:
			status.InternalServerError(w, r, err)
		}
		return
	}
}

func (h *MemberHandler) listBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.subscriptionService.ListMemberBooks(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			status.NotFound(w, r, err)
		default:
			status.InternalServerError(w, r, err)
		}
		return
	}

	status.OK(w, r, res)
}