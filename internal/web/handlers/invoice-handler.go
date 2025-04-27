package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mauFade/go-payment-gateway/internal/dto"
	"github.com/mauFade/go-payment-gateway/internal/service"
)

type InvoiceHandler struct {
	service service.InvoiceService
}

func NewInvoiceHandler(s service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		service: s,
	}
}

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")

	if apiKey == "" {
		http.Error(w, "api key is required", http.StatusBadRequest)
		return
	}

	var input dto.CreateInvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.Create(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
