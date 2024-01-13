package uc

import (
	"Multithreading/internal/dto"
)

type GetCEP interface {
	FetchCEP(cep string) (*dto.AddressDTO, error)
}

func FetchCepFromProvider(cepUC GetCEP, cep string, cAddress chan dto.AddressDTO) {
	address, err := cepUC.FetchCEP(cep)
	if err != nil {
		return
	}
	cAddress <- *address
}
