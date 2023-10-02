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

	fmt.Printf("cep: %s - api: %s\n", cep.Cep, ctx.Value("api"))
}
