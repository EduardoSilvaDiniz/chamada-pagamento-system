package handler

import (
	"encoding/json"
	"net/http"

	"projeto-integrador-mdm/internal/service"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		object, err := h.service.Create(r.Context(), r.Body)
		if err != nil {
			http.Error(w, "erro de execução Create: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(object); err != nil {
			http.Error(w, "Erro de execução JSON encode: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *PaymentHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := h.service.List(r.Context())
		if err != nil {
			http.Error(w, "erro de execução List: "+err.Error(), http.StatusBadRequest)
			return
		}

		if len(list) == 0 {
			http.Error(w, "nenhum registro encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(list); err != nil {
			http.Error(w, "erro ao serializar JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *PaymentHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.PathValue("payment_id")
		rows, err := h.service.Delete(ctx, id)
		if err != nil {
			http.Error(w, "Error de execução Delete: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if rows == 0 {
			http.Error(w, "Registro não encontrado", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(id); err != nil {
			http.Error(w, "Erro de execução JSON encode: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
