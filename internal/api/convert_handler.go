package api

import (
	"encoding/json"
	"net/http"

	"github.com/johnnyFR26/GoMicroservice/internal/service"
	"github.com/johnnyFR26/GoMicroservice/pkg/model"
)

type ConvertHandler struct {
	Service *service.ConverterService
}

func NewConvertHandler(s *service.ConverterService) *ConvertHandler {
	return &ConvertHandler{Service: s}
}

func (h *ConvertHandler) HandleConvert(w http.ResponseWriter, r *http.Request) {
	var req model.ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	converted, rate, err := h.Service.Convert(req.From, req.To, req.Amount)
	if err != nil {
		http.Error(w, "Erro na conversão: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.ConversionResponse{
		ConvertedAmount: converted,
		Rate:            rate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
