package http

import (
	"encoding/json"
	"food/internal/usecase"
	"net/http"
)

// PingHandler struct holds a reference to the PingUseCase.
type PingHandler struct {
	useCase usecase.PingUseCase
}

// NewPingHandler function returns a new instance of PingHandler.
func NewPingHandler(uc usecase.PingUseCase) *PingHandler {
	return &PingHandler{useCase: uc}
}

// Ping method handles HTTP requests for the "ping" endpoint.
func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	response := h.useCase.GetPingResponse()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": response})
}
