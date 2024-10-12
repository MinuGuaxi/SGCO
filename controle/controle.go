package controle

import (
	"html/template"
	modelo "modulo/modelos"
	"net/http"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Login", nil)
}

func Tela_Dashboard(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "Dashboard", http.StatusFound)
}

// PACIENTES

// Listar pacientes registrados
func Listar(w http.ResponseWriter, r *http.Request) {
	registros := modelo.Buscar()
	tmp.ExecuteTemplate(w, "Listar", registros)
}

/*
// Cadastrar novo paciente
func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Nome_paciente := r.FormValue("Nome_paciente")
		Tipo_plano := r.FormValue("Tipo_plano")
		Telefone_paciente := r.FormValue("Telefone_paciente")
		Endereco_paciente := r.FormValue("Endereco_paciente")
		Cidade_paciente := r.FormValue("Cidade_paciente")
		Bairro_paciente := r.FormValue("Bairro_paciente")
		Cpf_paciente := r.FormValue("Cpf_paciente")
		Email_paciente := r.FormValue("Email_paciente")
		Senha_paciente := r.FormValue("Senha_paciente")

		convertsenha, err := strconv.Atoi(Senha_paciente)
		if err != nil {
			panic(err.Error())
		}

		convertelefone, err := strconv.Atoi(Telefone_paciente)
		if err != nil {
			panic(err.Error())
		}

		modelo.Inseri(Nome_paciente, Tipo_plano, convertelefone, Endereco_paciente, Cidade_paciente, Bairro_paciente, Email_paciente,
			convertcpf, convertsenha)
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
		Id_paciente := r.FormValue("Id_paciente")
		Nome_paciente := r.FormValue("Nome_paciente")
		Tipo_plano := r.FormValue("Tipo_plano")
		Telefone_paciente := r.FormValue("Telefone_paciente")
		Endereco_paciente := r.FormValue("Endereco_paciente")
		Cidade_paciente := r.FormValue("Cidade_paciente")
		Bairro_paciente := r.FormValue("Bairro_paciente")
		Cpf_paciente := r.FormValue("Cpf_paciente")
		Email_paciente := r.FormValue("Email_paciente")
		Senha_paciente := r.FormValue("Senha_paciente")

		idConvertidaParaInt, err := strconv.Atoi(Id_paciente)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		convertelefone, err := strconv.Atoi(Telefone_paciente)
		if err != nil {
			panic(err.Error())
		}

		convertcpf, err := strconv.Atoi(Cpf_paciente)
		if err != nil {
			panic(err.Error())
		}

		convertsenha, err := strconv.Atoi(Senha_paciente)
		if err != nil {
			panic(err.Error())
		}

		modelo.Atualizar(idConvertidaParaInt, Nome_paciente, Tipo_plano, Endereco_paciente, Cidade_paciente, Bairro_paciente,
			convertelefone, convertcpf, convertsenha, Email_paciente)
	}

	http.Redirect(w, r, "/listar.html", 301)

}
*/

/*
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

*/
