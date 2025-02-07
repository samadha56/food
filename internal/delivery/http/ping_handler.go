package http

import (
	"encoding/json"
	"food/internal/usecase"
	"net/http"
)

type PingHandler struct {
	useCase usecase.PingUseCase
}

func NewPingHandler(uc usecase.PingUseCase) *PingHandler {
	return &PingHandler{useCase: uc}
}

func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	response := h.useCase.GetPingResponse()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": response})
}