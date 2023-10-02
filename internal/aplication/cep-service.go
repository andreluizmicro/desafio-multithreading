package aplication

import (
	"context"

	repository "github.com/andreluizmicro/desafio-multithreading/internal/infrastructure/repository"
)

type InputDTO struct {
	Cep string
}

type OutputDTO struct {
	Cep       string `json:"cep"`
	State     string `json:"state"`
	City      string `json:"city"`
	Districit string `json:"district"`
	Address   string `json:"address"`
}

type CepService struct {
	repository *repository.CepRepository
}

func NewCepService(repository *repository.CepRepository) *CepService {
	return &CepService{
		repository: repository,
	}
}

func (service *CepService) SearchCEP(ctx *context.Context, input InputDTO) (*OutputDTO, error) {
	cep, err := service.repository.SearchCEP(ctx, input.Cep)
	if err != nil {
		return nil, err
	}

	return &OutputDTO{
		Cep:       cep.Cep,
		State:     cep.State,
		City:      cep.City,
		Districit: cep.Districit,
		Address:   cep.Address,
	}, nil
}
