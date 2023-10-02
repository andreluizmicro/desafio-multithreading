package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/andreluizmicro/desafio-multithreading/internal/domain/entity"
)

const (
	API_CEP     = "https://cdn.apicep.com/file/apicep/"
	VIA_CEP_API = "http://viacep.com.br/ws"
)

type JsonResponseViaCepApi struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type JsonResponseApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type CepRepository struct {
}

func NewCepRepository() *CepRepository {
	return &CepRepository{}
}

func (repository *CepRepository) SearchCEP(cep string) (*entity.Cep, error) {

	chViaCepApi := make(chan entity.Cep)
	chApiCep := make(chan entity.Cep)

	go buscaViaCep(cep, chViaCepApi)
	go buscaApiCep(cep, chApiCep)

	select {
	case cepFound := <-chViaCepApi:
		return &cepFound, nil
	case cepFound := <-chApiCep:
		return &cepFound, nil
	case <-time.After(time.Second * 1):
		return nil, errors.New("timeout")
	}
}

func buscaViaCep(cep string, ch chan<- entity.Cep) error {
	req, err := http.Get(VIA_CEP_API + "/" + cep + "/json")
	if err != nil {
		return err
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var response JsonResponseViaCepApi
	json.Unmarshal(data, &response)

	cepFound := entity.Cep{
		Cep:       response.Cep,
		State:     response.Uf,
		City:      response.Localidade,
		Districit: response.Bairro,
		Address:   response.Logradouro,
	}

	ch <- cepFound

	return nil
}

func buscaApiCep(cep string, ch chan<- entity.Cep) error {
	req, err := http.Get(API_CEP + "/" + cep + "/json")
	if err != nil {
		return err
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var response JsonResponseApiCep
	json.Unmarshal(data, &response)

	cepFound := entity.Cep{
		Cep:       response.Code,
		State:     response.State,
		City:      response.City,
		Districit: response.District,
		Address:   response.Address,
	}

	ch <- cepFound

	return nil
}
