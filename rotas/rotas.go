package rotas

import (
	"modulo/controle"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controle.Index)

	http.HandleFunc("/listar.html", controle.Listar)

	http.HandleFunc("/inserir", controle.Inserir)
	http.HandleFunc("/delete", controle.Delete)

	http.HandleFunc("/edit.html", controle.Edit)
	http.HandleFunc("/atualizar", controle.Atualize)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}
