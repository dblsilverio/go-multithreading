package uc

import "Multithreading/internal/dto"

type GetCEP interface {
	FetchCEP(cep string) dto.AddressDTO
}
