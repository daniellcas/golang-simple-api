package main

import (
	"fmt"
	"github.com/Bhimmo/golang-simple-api/adapter/routes"
	"github.com/Bhimmo/golang-simple-api/pkg/rabbitmq"
	"github.com/Bhimmo/golang-simple-api/pkg/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	//Init database conection
	sqlite.Init()

	//Init rabbimq
	rabbitmq.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	//Servico
	r.Post("/servico", routes.NovoServico)
	r.Get("/servico/{id}", routes.PegandoServicoPeloId)

	//Solicitacao
	r.Post("/solicitacao", routes.SalvarSolicitacao)
	r.Get("/solicitacao/{id}", routes.PegandoSolicitacaoPeloId)
	r.Get("/solicitacao/{id}/atualizar-status", routes.AtualizarStatusSolicitacao)

	fmt.Println("Start api")
	http.ListenAndServe(":3000", r)
}
