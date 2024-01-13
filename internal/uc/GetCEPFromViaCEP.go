package uc

import (
	"Multithreading/internal/dto"
	"Multithreading/internal/infra/webclient"
	"encoding/json"
	"fmt"
)

type GetCEPFromViaCEP struct{}

const ViacepUrl = "https://viacep.com.br/ws/%s/json/"

func (g *GetCEPFromViaCEP) FetchCEP(cep string) dto.AddressDTO {

	response := webclient.GET(requestViaCEPUrl(cep))
	var addressViaCep dto.AddressViaCep
	err := json.Unmarshal([]byte(response), &addressViaCep)
	if err != nil {
		panic(err)
	}

	addressDTO := dto.AddressDTO{
		Cep:          addressViaCep.Cep,
		State:        addressViaCep.Uf,
		City:         addressViaCep.Localidade,
		Neighborhood: addressViaCep.Bairro,
		Street:       addressViaCep.Logradouro,
	}

	return addressDTO
}

func requestViaCEPUrl(cep string) string {
	return fmt.Sprintf(ViacepUrl, cep)
}
