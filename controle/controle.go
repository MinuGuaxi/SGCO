package controle

import (
	"html/template"
	"log"
	"modulo/modelo"
	"net/http"
	"strconv"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Index", nil)
}

// PACIENTES

// Listar pacientes registrados
func Listar(w http.ResponseWriter, r *http.Request) {
	registros := modelo.Buscar()
	tmp.ExecuteTemplate(w, "Listar", registros)
}

// Cadastrar novo paciente
func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		nomedamae := r.FormValue("nomedamae")
		nomedopai := r.FormValue("nomedopai")
		telefone := r.FormValue("telefone")
		cpf := r.FormValue("cpf")
		email := r.FormValue("email")
		senha := r.FormValue("senha")

		convertsenha, err := strconv.Atoi(senha)
		if err != nil {
			panic(err.Error())
		}

		convertelefone, err := strconv.Atoi(telefone)
		if err != nil {
			panic(err.Error())
		}

		convertcpf, err := strconv.Atoi(cpf)
		if err != nil {
			panic(err.Error())
		}

		modelo.Inseri(nome, nomedamae, nomedopai, convertelefone,
			convertcpf, email, convertsenha)
	}

	http.Redirect(w, r, "/listar.html", 301)

}

// Deletar paciente
func Delete(w http.ResponseWriter, r *http.Request) {
	idUsuario := r.URL.Query().Get("id")
	modelo.Deleta(idUsuario)
	http.Redirect(w, r, "/listar.html", 301)
}

// Localiza as informações do paciente
func Edit(w http.ResponseWriter, r *http.Request) {
	idatualizar := r.URL.Query().Get("id")
	atualize := modelo.Editar(idatualizar)
	tmp.ExecuteTemplate(w, "Edit", atualize)
}

// Atualiza as informações do paciente
func Atualize(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		nomedamae := r.FormValue("nomedamae")
		nomedopai := r.FormValue("nomedopai")
		telefone := r.FormValue("telefone")
		cpf := r.FormValue("cpf")
		email := r.FormValue("email")
		senha := r.FormValue("senha")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		convertelefone, err := strconv.Atoi(telefone)
		if err != nil {
			panic(err.Error())
		}

		convertcpf, err := strconv.Atoi(cpf)
		if err != nil {
			panic(err.Error())
		}

		convertsenha, err := strconv.Atoi(senha)
		if err != nil {
			panic(err.Error())
		}

		modelo.Atualizar(idConvertidaParaInt, nome, nomedamae, nomedopai,
			convertelefone, convertcpf, email, convertsenha)
	}

	http.Redirect(w, r, "/listar.html", 301)

}

// Historico do paciente
func Historico_Paciente() {

}

// PROFISSIONAIS

// Listar os profissionais registrados
func Listar_Profissional() {

}

// Cadastra um novo profissional
func Inserir_Profissional() {

}

// Deleta um profissional
func Deletar_Profissional() {

}

// Localiza as informações do profissional
func Editar_Profissional() {

}

// Atualiza as informações do profissional
func Atualizar_Profissional() {

}

// PROCEDIMENTOS

// Listar os procedimentos registrados
func Listar_Procedimentos() {

}

// Cadastra um novo profissional
func Inserir_Procedimentos() {

}

// Deleta um profissional
func Deletar_Procedimentos() {

}

// Localiza as informações do profissional
func Editar_Procedimentos() {

}

// Atualiza as informações do profissional
func Atualizar_Procedimentos() {

}
