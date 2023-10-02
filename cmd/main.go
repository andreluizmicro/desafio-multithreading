package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreluizmicro/desafio-multithreading/internal/aplication"
	repository "github.com/andreluizmicro/desafio-multithreading/internal/infrastructure/repository"
)

func main() {
	http.HandleFunc("/", SearchCEP)
	http.ListenAndServe(":8080", nil)
}

func SearchCEP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	cepRepository := repository.NewCepRepository()
	service := aplication.NewCepService(cepRepository)

	input := aplication.InputDTO{
		Cep: r.URL.Query().Get("cep"),
	}

	cep, err := service.SearchCEP(&ctx, input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Retorno da api: %s\n", ctx.Value("api"))
	fmt.Printf("cep: %s\nstate: %s\ncity: %s\ndistrict: %s\nAddress: %s\n", cep.Cep, cep.State, cep.City, cep.Districit, cep.Address)
}
