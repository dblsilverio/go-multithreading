package uc

import "Multithreading/internal/dto"

type GetCEP interface {
	FetchCEP(cep string) dto.AddressDTO
}

func FetchCepFromProvider(cepUC GetCEP, cep string, cAddress chan dto.AddressDTO) {
	address := cepUC.FetchCEP(cep)
	cAddress <- address
}
