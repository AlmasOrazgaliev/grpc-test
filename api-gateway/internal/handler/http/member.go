package http

import (
	"api-gateway/internal/domain/member"
	"api-gateway/pkg/server/status"
	desc "api-gateway/proto"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type MemberHandler struct {
	memberServiceClient desc.MemberClient
}

func NewMemberHandler(s desc.MemberClient) *MemberHandler {
	return &MemberHandler{memberServiceClient: s}
}

func (h *MemberHandler) Routes() chi.Router {
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

func (h *MemberHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.memberServiceClient.List(r.Context(), &desc.MemberData{})
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
	data := &desc.MemberData{
		FullName: req.FullName,
	}
	res, err := h.memberServiceClient.Add(r.Context(), data)
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *MemberHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.memberServiceClient.Get(r.Context(), &desc.MemberData{Id: id})
	if err != nil {
		status.InternalServerError(w, r, err)
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
	data := &desc.MemberData{
		Id:       id,
		FullName: req.FullName,
	}
	res, err := h.memberServiceClient.Update(r.Context(), data)
	if err != nil {
		status.InternalServerError(w, r, err)
		return
	}

	status.OK(w, r, res)
}

func (h *MemberHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := h.memberServiceClient.Get(r.Context(), &desc.MemberData{Id: id})
	if err != nil {
		status.InternalServerError(w, r, err)
	}
}
