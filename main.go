package main

import (
	"fmt"
	
	"modulo/rotas"
	"net/http"
)

func main() {
	rotas.CarregaRotas()
	fmt.Println("Servidor rodando na porta 8081")
	http.ListenAndServe(":8081", nil)
}