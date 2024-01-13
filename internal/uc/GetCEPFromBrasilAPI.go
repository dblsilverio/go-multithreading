package uc

import (
	"Multithreading/internal/dto"
	"Multithreading/internal/infra/webclient"
	"encoding/json"
	"fmt"
)

type GetCEPFromBrasilAPI struct{}

const BrasilapiUrl = "https://brasilapi.com.br/api/cep/v1/%s"

func (g *GetCEPFromBrasilAPI) FetchCEP(cep string) dto.AddressDTO {

	response := webclient.GET(requestBrasilAPIUrl(cep))
	var addressViaCep dto.AddressBrasilAPI
	err := json.Unmarshal([]byte(response), &addressViaCep)
	if err != nil {
		panic(err)
	}
	addressDTO := dto.AddressDTO{
		Cep:          addressViaCep.Cep,
		State:        addressViaCep.State,
		City:         addressViaCep.City,
		Neighborhood: addressViaCep.Neighborhood,
		Street:       addressViaCep.Street,
	}
	return addressDTO

}

func requestBrasilAPIUrl(cep string) string {
	return fmt.Sprintf(BrasilapiUrl, cep)
}
