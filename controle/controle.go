package controle

import (
	"fmt"
	"html/template"
	modelo "modulo/modelos"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Login_Option", nil)
}

func LoginPaciente(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Login_Paciente", nil)
}

func LoginProfissional(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Login_Profissional", nil)
}
func SetSessionToken(w http.ResponseWriter, sessionToken string) {
	// Definir o cookie com as opções corretas
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",   // O cookie estará disponível em todas as rotas
		HttpOnly: true,  // Impede acesso via JavaScript
		Secure:   false, // Use true se estiver usando HTTPS
		MaxAge:   3600,  // 1 hora de duração
	}

	// Adicionar o cookie à resposta HTTP
	http.SetCookie(w, cookie)
	fmt.Println("Cookie criado com sucesso:", cookie)
}

// Função que renderiza o dashboard usando o template base
func Tela_Dashboard(w http.ResponseWriter, r *http.Request) {
	// Tentar recuperar o cookie de sessão
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Se o cookie não for encontrado, redirecionar para a página de login
		fmt.Println("Cookie não encontrado:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Verificar o valor do cookie para depuração
	fmt.Println("Valor do cookie recebido:", cookie.Value)

	// Carregar os templates
	tmpl, err := template.ParseFiles("templates/index.html", "templates/dashboard.html") // Template base e específico
	if err != nil {
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página usando o template base
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}
}

// Função que renderiza as configurações usando o template base
func Tela_Configuracoes(w http.ResponseWriter, r *http.Request) {
	// Recuperar o cookie de sessão
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Se o cookie não for encontrado, redirecionar para a página de login
		fmt.Println("Cookie não encontrado:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redireciona para a página inicial de login
		return
	}

	// Verificar o valor do cookie (para depuração)
	fmt.Println("Valor do cookie recebido:", cookie.Value)

	// Se o cookie for válido, carregue os templates
	tmpl, err := template.ParseFiles("templates/index.html", "templates/configuracoes.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}

// Função que renderiza o atendimento usando o template base
func Tela_Atendimento(w http.ResponseWriter, r *http.Request) {
	// Recuperar o cookie de sessão
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Se o cookie não for encontrado, redirecionar para a página de login
		fmt.Println("Cookie não encontrado:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redireciona para a página inicial de login
		return
	}

	// Verificar o valor do cookie (para depuração)
	fmt.Println("Valor do cookie recebido:", cookie.Value)

	// Se o cookie for válido, carregue os templates
	tmpl, err := template.ParseFiles("templates/index.html", "templates/atendimento.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de atendimento renderizada com sucesso.")
}

//Função para renderizar a pagina calendario com template base
func Tela_Calendario(w http.ResponseWriter, r *http.Request) {
	// Recuperar o cookie de sessão
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Se o cookie não for encontrado, redirecionar para a página de login
		fmt.Println("Cookie não encontrado:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redireciona para a página inicial de login
		return
	}

	// Verificar o valor do cookie (para depuração)
	fmt.Println("Valor do cookie recebido:", cookie.Value)

	// Se o cookie for válido, carregue os templates
	tmpl, err := template.ParseFiles("templates/index.html", "templates/calendario_paciente.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}

//Função para renderizar a tela de historico de tratamentos

func Tela_Tratamentos(w http.ResponseWriter, r *http.Request) {
	// Recuperar o cookie de sessão
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Se o cookie não for encontrado, redirecionar para a página de login
		fmt.Println("Cookie não encontrado:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redireciona para a página inicial de login
		return
	}

	// Verificar o valor do cookie (para depuração)
	fmt.Println("Valor do cookie recebido:", cookie.Value)

	// Se o cookie for válido, carregue os templates
	tmpl, err := template.ParseFiles("templates/index.html", "templates/tratamentos_paciente.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}
//Tela para renderizar a pagina do Financeiro com o template padrão

func Tela_Financeiro(w http.ResponseWriter, r *http.Request) {
	// Recuperar o cookie de sessão
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Se o cookie não for encontrado, redirecionar para a página de login
		fmt.Println("Cookie não encontrado:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redireciona para a página inicial de login
		return
	}

	// Verificar o valor do cookie (para depuração)
	fmt.Println("Valor do cookie recebido:", cookie.Value)

	// Se o cookie for válido, carregue os templates
	tmpl, err := template.ParseFiles("templates/index.html", "templates/financeiro_paciente.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}

// PACIENTES

// Listar pacientes registrados
/*
func Listar(w http.ResponseWriter, r *http.Request) {
	registros := modelo.Buscar()
	tmp.ExecuteTemplate(w, "Listar", registros)
}
*/

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtém os dados do formulário
		Email_paciente := r.FormValue("Email_paciente")
		Senha_paciente := r.FormValue("Senha_paciente")

		// Converter a senha para um inteiro (se necessário)
		convertsenha, err := strconv.Atoi(Senha_paciente)
		if err != nil {
			fmt.Println("Erro ao converter a senha:", err)
			http.Error(w, "Erro ao processar os dados", http.StatusBadRequest)
			return
		}

		// Chama a função de autenticação
		autenticado, err := modelo.AutenticaUsuario(Email_paciente, convertsenha)
		if err != nil {
			fmt.Println("Erro ao autenticar o usuário:", err)
			http.Error(w, "Erro no servidor, tente novamente mais tarde.", http.StatusInternalServerError)
			return
		}

		if autenticado {
			// Gerar um UUID único para o cookie de sessão
			sessionToken := uuid.New().String()

			// Definir o cookie com o valor único gerado
			expiration := time.Now().Add(24 * time.Hour) // Expiração do cookie em 24 horas
			cookie := http.Cookie{
				Name:     "session_token",
				Value:    sessionToken,         // UUID gerado
				Expires:  expiration,           // Definir tempo de expiração
				Path:     "/",                  // Disponível em todas as rotas
				HttpOnly: true,                 // Impede acesso via JavaScript
				Secure:   false,                // Usar true se estiver em ambiente HTTPS
				SameSite: http.SameSiteLaxMode, // Reduz risco de CSRF
			}

			// Armazenar o token de sessão no cookie
			http.SetCookie(w, &cookie)

			// Redirecionar para o Dashboard
			http.Redirect(w, r, "/Dashboard", http.StatusSeeOther)

			// Log para depuração (apenas para ver o valor gerado)
			fmt.Println("Cookie de sessão gerado:", sessionToken)
		} else {
			// Exibe uma mensagem de erro de login inválido
			http.Error(w, "Email ou senha incorretos", http.StatusUnauthorized)
		}
	} else {
		// Se não for POST, renderiza o formulário de login
		http.ServeFile(w, r, "login.html")
	}
}

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
		convertcpf, err := strconv.Atoi(Cpf_paciente)
		if err != nil {
			panic(err.Error())
		}
		modelo.Inseri(
			Nome_paciente,
			Tipo_plano,
			convertelefone,
			Endereco_paciente,
			Cidade_paciente,
			Bairro_paciente,
			convertcpf,
			Email_paciente,
			convertsenha,
		)
	}

	http.Redirect(w, r, "/login_paciente.html", 301)

}

/*

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
,
			/*
// Historico do paciente
func Historico_Paciente() {

}

// PROFISSIONAIS

// Listar os profissionais registrados
func Listar_Profissional() {

}

// Cadastra um novo profissional
func Inserir_Profissional(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Nome_profissional := r.FormValue("Nome_profissional")
		Telefone_profissional := r.FormValue("Telefone_profissional")
		Endereco_profissional := r.FormValue("Endereco_profissional")
		Cidade_profissional := r.FormValue("Cidade_profissional")
		Bairro_profissional := r.FormValue("Bairro_profissional")
		Cpf_profissional := r.FormValue("Cpf_profissional")
		Email_profissional := r.FormValue("Email_profissional")
		Senha_profissional := r.FormValue("Senha_profissional")

		convertsenha, err := strconv.Atoi(Senha_profissional)
		if err != nil {
			panic(err.Error())
		}

		convertelefone, err := strconv.Atoi(Telefone_profissional)
		if err != nil {
			panic(err.Error())
		}
		convertcpf, err := strconv.Atoi(Cpf_profissional)
		if err != nil {
			panic(err.Error())
		}
		modelo.Inseriprofissional(
			Nome_paciente,
			Tipo_plano,
			convertelefone,
			Endereco_paciente,
			Cidade_paciente,
			Bairro_paciente,
			convertcpf,
			Email_paciente,
			convertsenha,
		)
	}

	http.Redirect(w, r, "/login_profissional.html", 301)
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
func Inserir_Procedimentos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Nome_procedimento := r.FormValue("Nome_paciente")
		Data_procedimento := r.FormValue("Data_procedimento")
		Hora_procedimento := r.FormValue("Hora_procedimento")
		Tipo_procedimento := r.FormValue("Tipo_procedimento")
		Valor_procedimento := r.FormValue("Valor_procedimento")
		Profissional_designado := r.FormValue("Profissional_designado")


		modelo.Inseriprocedimento(
			Nome_procedimento,
			Data_procedimento,
			Hora_procedimento,
			Tipo_procedimento,
			Valor_procedimento,
			Profissional_designado,
		)
	}

	http.Redirect(w, r, "/cadastro_procedimento.html", 301)
}

// Deleta um profissional
func Deletar_Procedimentos() {

}

// Localiza as informações do profissional
func Editar_Procedimentos() {

}

// Atualiza as informações ,

*/
