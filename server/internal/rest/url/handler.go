package url

import "net/http"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	h.service.Get(w, r)
}
