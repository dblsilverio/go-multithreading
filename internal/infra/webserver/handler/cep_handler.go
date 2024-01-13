package handler

import (
	"Multithreading/internal/dto"
	"Multithreading/internal/uc"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type CepHandler struct {
	apiTimeout int
}

func NewCepHandler(apiTimeout int) *CepHandler {
	return &CepHandler{apiTimeout: apiTimeout}
}

func (h *CepHandler) GetCEP(w http.ResponseWriter, request *http.Request) {
	cep := chi.URLParam(request, "cep")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	viaCep := uc.GetCEPFromViaCEP{}
	brasilApi := uc.GetCEPFromBrasilAPI{}

	addressViaCep := viaCep.FetchCEP(cep)
	addressBrasilApi := brasilApi.FetchCEP(cep)
	addressList := []dto.AddressDTO{addressViaCep, addressBrasilApi}
	json.NewEncoder(w).Encode(addressList)

}
