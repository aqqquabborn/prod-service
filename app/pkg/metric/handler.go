package metric

import (
	"log"
	"net/http"

	_ "prod-service/app/docs"

	"github.com/julienschmidt/httprouter"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

// Heartbeat
// @Summary Heartbeat metric
// @Tags Metrics
// @Success 204
// @Failure 400
// @Router /api/heartbeat/ [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	log.Print("hehfehe")
	w.WriteHeader(204)
}
