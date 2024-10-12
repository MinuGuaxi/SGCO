package rotas

import (
	"modulo/controle"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controle.Login)
	http.HandleFunc("/Dashboard", controle.Tela_Dashboard)

	//Paciente
	http.HandleFunc("/listar.html", controle.Listar)
	/*
		http.HandleFunc("/inserir", controle.Inserir)
		http.HandleFunc("/delete", controle.Delete)
		http.HandleFunc("/edit.html", controle.Edit)
		http.HandleFunc("/atualizar", controle.Atualize)
	*/
	//Profissional

	//http.HandleFunc("/listar_profissionais", controle.Listar_Profissional)
	//http.HandleFunc("/cadastrar_profissional", controle.Inserir_Profissional)
	//http.HandleFunc("/deletar_profissional", controle.Editar_Profissional)
	//http.HandleFunc("/atualizar_profissional", controle.Atualizar_Profissional)

	//Procedimentos
	//http.HandleFunc("/listar_procedimentos", controle.Listar_Procedimentos)
	//http.HandleFunc("/cadastrar_procedimento", controle.Inserir_Procedimentos)
	//http.HandleFunc("/deletar_procedimento", controle.Editar_Procedimentos)
	//http.HandleFunc("/atualizar_procedimento", controle.Atualizar_Procedimentos)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}
