package uc

import (
	"Multithreading/internal/dto"
	"Multithreading/internal/infra/webclient"
	"encoding/json"
	"fmt"
)

type GetCEPFromViaCEP struct{}

const ViacepUrl = "https://viacep.com.br/ws/%s/json/"

func (g *GetCEPFromViaCEP) FetchCEP(cep string) (*dto.AddressDTO, error) {

	response, err := webclient.GET(requestViaCEPUrl(cep))

	if err != nil {
		return nil, err
	}

	var addressViaCep dto.AddressViaCep
	err = json.Unmarshal([]byte(response), &addressViaCep)
	if err != nil {
		fmt.Printf("Via Cep query failed: %s\n", err.Error())
		return nil, err
	}
	addressDTO := dto.AddressDTO{
		Cep:          addressViaCep.Cep,
		State:        addressViaCep.Uf,
		City:         addressViaCep.Localidade,
		Neighborhood: addressViaCep.Bairro,
		Street:       addressViaCep.Logradouro,
	}

	return &addressDTO, nil
}

func requestViaCEPUrl(cep string) string {
	return fmt.Sprintf(ViacepUrl, cep)
}
