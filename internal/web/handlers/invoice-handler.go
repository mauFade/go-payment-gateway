package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mauFade/go-payment-gateway/internal/domain"
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

func (h *InvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "api key is required", http.StatusBadRequest)
		return
	}

	output, err := h.service.FindByID(id, apiKey)

	if err != nil {
		switch err {
		case domain.ErrUnaithorizedAccess:
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *InvoiceHandler) ListByAccount(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "api key is required", http.StatusUnauthorized)
		return
	}

	output, err := h.service.ListByAPIKey(apiKey)

	if err != nil {
		switch err {
		case domain.ErrUnaithorizedAccess:
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
