package main

import (
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
	cepRepository := repository.NewCepRepository()
	service := aplication.NewCepService(cepRepository)

	input := aplication.InputDTO{
		Cep: r.URL.Query().Get("cep"),
	}

	cep, err := service.SearchCEP(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(cep.Cep)
}
