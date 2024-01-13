package handler

import (
	"Multithreading/internal/dto"
	"Multithreading/internal/uc"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

var ErrApiTimeout = errors.New("CEP Api Timeout")

type CepHandler struct {
	apiTimeout int
}

func NewCepHandler(apiTimeout int) *CepHandler {
	return &CepHandler{apiTimeout: apiTimeout}
}

// GetCEP godoc
// @Summary Get CEP/Address from provider
// @Description Fetch CEP/Address from two different providers - ViaCEP and BrasilAPI.
// @Tags cep
// @Accept  json
// @Produce  json
// @Param cep path string true "cep" pattern(\d{8})
// @Success 200 {object} dto.AddressDTO "CEP found"
// @Success 500 "Internal server error"
// @Router /cep/{cep} [get]
func (h *CepHandler) GetCEP(w http.ResponseWriter, request *http.Request) {
	cep := chi.URLParam(request, "cep")

	w.Header().Set("Content-Type", "application/json")

	addressDTO, err := getCep(cep, h.apiTimeout)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(addressDTO)
	if err != nil {
		panic(err)
	}

}

func getCep(cep string, timeout int) (dto.AddressDTO, error) {
	cViaCep := make(chan dto.AddressDTO)
	cBrasilApi := make(chan dto.AddressDTO)

	go func() {
		uc.FetchCepFromProvider(&uc.GetCEPFromViaCEP{}, cep, cViaCep)
	}()

	go func() {
		uc.FetchCepFromProvider(&uc.GetCEPFromBrasilAPI{}, cep, cBrasilApi)
	}()

	select {
	case m := <-cViaCep:
		fmt.Printf("VIA Cep response received: %v\n", m)
		return m, nil
	case m := <-cBrasilApi:
		fmt.Printf("BrasilAPI response received: %v\n", m)
		return m, nil
	case <-time.After(time.Duration(timeout) * time.Second):
		println(ErrApiTimeout.Error())
		return dto.AddressDTO{}, ErrApiTimeout
	}

}
