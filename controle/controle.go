package controle

import (
	"fmt"
	"html/template"
	"log"
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

//FUNÇÕES DE ENTRADA

// FUNÇÃO GERAL DE AUTH
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

// FUNÇÕES DE AUTH - PACIENTE
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

// FUNÇÕES DE AUTH - PROFISSIONAL
func LoginHandler_profissional(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtém os dados do formulário
		Email_profissional := r.FormValue("email_profissional")
		Senha_profissional := r.FormValue("senha_profissional")

		// Converter a senha para um inteiro (se necessário)
		convertsenha, err := strconv.Atoi(Senha_profissional)
		if err != nil {
			fmt.Println("Erro ao converter a senha:", err)
			http.Error(w, "Erro ao processar os dados", http.StatusBadRequest)
			return
		}

		// Chama a função de autenticação
		autenticado, err := modelo.AutenticaProfissional(Email_profissional, convertsenha)
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
			http.Redirect(w, r, "/Dashboard_Profissional", http.StatusSeeOther)

			// Log para depuração (apenas para ver o valor gerado)
			fmt.Println("Cookie de sessão gerado:", sessionToken)
		} else {
			// Exibe uma mensagem de erro de login inválido
			http.Error(w, "Email ou senha incorretos", http.StatusUnauthorized)
		}
	} else {
		// Se não for POST, renderiza o formulário de login
		http.ServeFile(w, r, "login_profissional.html")
	}
}

//FUNÇÕES DE ENTRADA(URL) E RENDERIZAÇÃO DO PACIENTE

func LoginPaciente(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Login_Paciente", nil)
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
	tmpl, err := template.ParseFiles("templates/index.html", "templates/pacientes/dashboard.html") // Template base e específico
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
	tmpl, err := template.ParseFiles("templates/index.html", "templates/pacientes/configuracoes.html") // Carrega o template base e o específico
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
	tmpl, err := template.ParseFiles("templates/index.html", "templates/pacientes/atendimento.html") // Carrega o template base e o específico
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

// Função para renderizar a pagina calendario com template base
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
	tmpl, err := template.ParseFiles("templates/index.html", "templates/pacientes/calendario_paciente.html") // Carrega o template base e o específico
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

// Função para renderizar a tela de historico de tratamentos
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
	tmpl, err := template.ParseFiles("templates/index.html", "templates/pacientes/tratamentos_paciente.html") // Carrega o template base e o específico
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

// Tela para renderizar a pagina do Financeiro com o template padrão
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
	tmpl, err := template.ParseFiles("templates/index.html", "templates/pacientes/financeiro_paciente.html") // Carrega o template base e o específico
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

//FUNÇÕES DE ENTRADA(URL) E RENDERIZAÇÃO DO PROFISSIONAL

func LoginProfissional(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Login_Profissional", nil)
}

// Função que renderiza o dashboard do profissional usando o template base
func Tela_Dashboard_Profissional(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.ParseFiles("templates/index_profissional.html", "templates/profissionais/dashboard_profissional.html") // Template base e específico
	if err != nil {
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página usando o template base
	err = tmpl.ExecuteTemplate(w, "index_profissional.html", nil)
	if err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}
}

// Função que renderiza as configurações do profissional usando o template base
func Tela_Configuracoes_Profissional(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.ParseFiles("templates/index_profissional.html", "templates/profissionais/configuracoes_profissional.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index_profissional.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}

// Função que renderiza o atendimento do profissional usando o template base
func Tela_Atendimento_Profissional(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.ParseFiles("templates/index_profissional.html", "templates/profissionais/atendimento_profissional.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index_profissional.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de atendimento renderizada com sucesso.")
}

// Função para renderizar a pagina calendario com template base
func Tela_Calendario_Profissional(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.ParseFiles("templates/index_profissional.html", "templates/profissionais/calendario_profissional.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index_profissional.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}

// Função para renderizar a tela para o profissional gerenciar o historico de tratamento do paciente
func Tela_Tratamentos_Profissional(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.ParseFiles("templates/index_profissional.html", "templates/profissionais/tratamentos_profissional.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index_profissional.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}

// Tela para renderizar a pagina para o profissional gerenciar o Financeiro do paciente
func Tela_Financeiro_Profissional(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.ParseFiles("templates/index_profissional.html", "templates/profissionais/financeiro_profissional.html") // Carrega o template base e o específico
	if err != nil {
		// Se houver erro ao carregar os templates
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		fmt.Println("Erro ao carregar o template:", err)
		return
	}

	// Renderizar a página de configurações usando o template base (index.html)
	err = tmpl.ExecuteTemplate(w, "index_profissional.html", nil)
	if err != nil {
		// Caso haja um erro ao renderizar o template
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		fmt.Println("Erro ao renderizar o template:", err)
		return
	}

	// Para depuração: Verificar se a página foi renderizada corretamente
	fmt.Println("Página de configurações renderizada com sucesso.")
}

// CENTRAL DE REGISTROS - URL E RENDER
func Central_Registros(w http.ResponseWriter, r *http.Request) {
	// Adicionando log para verificar se a função foi chamada
	fmt.Println("Acessando a página Central_Registros.")

	procedimentos := modelo.BuscarProcedimento()

	// Recuperar o cookie de sessão
	cookie, err := r.Cookie("session_token")
	if err != nil {
		fmt.Println("Cookie não encontrado:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("Valor do cookie recebido:", cookie.Value)

	// Carregando os templates
	tmpl, err := template.ParseFiles("templates/index_profissional.html", "templates/central_registros.html")
	if err != nil {
		fmt.Println("Erro ao carregar o template:", err)
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		return
	}

	// Tentando renderizar a página com os dados de procedimentos
	err = tmpl.ExecuteTemplate(w, "index_profissional.html", procedimentos)
	if err != nil {
		fmt.Println("Erro ao renderizar o template:", err)
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}

	fmt.Println("Página Central_Registros renderizada com sucesso.")
}

// FUNÇÕES CRUD - PACIENTE

// Listar pacientes registrados
/*
func Listar(w http.ResponseWriter, r *http.Request) {
	registros := modelo.Buscar()
	tmp.ExecuteTemplate(w, "Listar", registros)
}
*/

// Cadastrar novo paciente
func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome_paciente := r.FormValue("nome_paciente")
		tipo_plano := r.FormValue("tipo_plano")
		telefone_paciente := r.FormValue("telefone_paciente")
		endereco_paciente := r.FormValue("endereco_paciente")
		cidade_paciente := r.FormValue("cidade_paciente")
		bairro_paciente := r.FormValue("bairro_paciente")
		cpf_paciente := r.FormValue("cpf_paciente")
		email_paciente := r.FormValue("email_paciente")
		senha_paciente := r.FormValue("senha_paciente")

		convertsenha, err := strconv.Atoi(senha_paciente)
		if err != nil {
			panic(err.Error())
		}

		convertelefone, err := strconv.Atoi(telefone_paciente)
		if err != nil {
			panic(err.Error())
		}
		convertcpf, err := strconv.Atoi(cpf_paciente)
		if err != nil {
			panic(err.Error())
		}
		modelo.Inseri(
			nome_paciente,
			tipo_plano,
			convertelefone,
			endereco_paciente,
			cidade_paciente,
			bairro_paciente,
			convertcpf,
			email_paciente,
			convertsenha,
		)
	}

	http.Redirect(w, r, "/login_paciente.html", 301)

}

// Atualiza as informações do paciente
func Atualize(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id_paciente := r.FormValue("id_paciente")
		nome_paciente := r.FormValue("nome_paciente")
		tipo_plano := r.FormValue("tipo_plano")
		telefone_paciente := r.FormValue("telefone_paciente")
		endereco_paciente := r.FormValue("endereco_paciente")
		cidade_paciente := r.FormValue("cidade_paciente")
		bairro_paciente := r.FormValue("bairro_paciente")
		cpf_paciente := r.FormValue("cpf_paciente")
		email_paciente := r.FormValue("email_paciente")
		senha_paciente := r.FormValue("senha_paciente")

		idConvertidaParaInt, err := strconv.Atoi(id_paciente)

		convertelefone, err := strconv.Atoi(telefone_paciente)
		if err != nil {
			panic(err.Error())
		}

		convertcpf, err := strconv.Atoi(cpf_paciente)
		if err != nil {
			panic(err.Error())
		}
		convertsenha, err := strconv.Atoi(senha_paciente)
		if err != nil {
			panic(err.Error())
		}

		modelo.Atualizar(
			idConvertidaParaInt,
			nome_paciente,
			tipo_plano,
			convertelefone,
			endereco_paciente,
			cidade_paciente,
			bairro_paciente,
			convertcpf,
			email_paciente,
			convertsenha,
		)
	}
}

// FUNÇÕES CRUD - PROFISSIONAL

// Cadastra um novo profissional
func Inserir_Profissional(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome_profissional := r.FormValue("nome_profissional")
		telefone_profissional := r.FormValue("telefone_profissional")
		endereco_profissional := r.FormValue("endereco_profissional")
		cidade_profissional := r.FormValue("cidade_profissional")
		bairro_profissional := r.FormValue("bairro_profissional")
		cpf_profissional := r.FormValue("cpf_profissional")
		email_profissional := r.FormValue("email_profissional")
		senha_profissional := r.FormValue("senha_profissional")

		convertsenha, err := strconv.Atoi(senha_profissional)
		if err != nil {
			panic(err.Error())
		}

		convertelefone, err := strconv.Atoi(telefone_profissional)
		if err != nil {
			panic(err.Error())
		}
		convertcpf, err := strconv.Atoi(cpf_profissional)
		if err != nil {
			panic(err.Error())
		}
		modelo.Inseriprofissional(
			nome_profissional,
			convertelefone,
			endereco_profissional,
			cidade_profissional,
			bairro_profissional,
			convertcpf,
			email_profissional,
			convertsenha,
		)
	}

	http.Redirect(w, r, "/login_profissional.html", 301)
}

// Atualiza as informações do profissional
func Atualizar_Profissional(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id_profissional := r.FormValue("id_profissional")
		nome_profissional := r.FormValue("nome_profissional")
		telefone_profissional := r.FormValue("telefone_profissional")
		endereco_profissional := r.FormValue("endereco_profissional")
		cidade_profissional := r.FormValue("cidade_profissional")
		bairro_profissional := r.FormValue("bairro_profissional")
		cpf_profissional := r.FormValue("cpf_profissional")
		email_profissional := r.FormValue("email_profissional")
		senha_profissional := r.FormValue("senha_profissional")

		idConvertidaParaInt, err := strconv.Atoi(id_profissional)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		convertelefone, err := strconv.Atoi(telefone_profissional)
		if err != nil {
			panic(err.Error())
		}

		convertcpf, err := strconv.Atoi(cpf_profissional)
		if err != nil {
			panic(err.Error())
		}

		convertSenha, err := strconv.Atoi(senha_profissional)
		if err != nil {
			panic(err.Error())
		}

		modelo.Atualizarprofissional(
			idConvertidaParaInt,
			nome_profissional,
			convertelefone,
			endereco_profissional,
			cidade_profissional,
			bairro_profissional,
			convertcpf,
			email_profissional,
			convertSenha,
		)
	}
	http.Redirect(w, r, "/listar_profissional.html", 301)

}

//FUNÇÕES CRUD - PROCEDIMENTOS

func Inserir_Procedimentos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome_procedimento := r.FormValue("nome_procedimento")
		data_procedimento := r.FormValue("data_procedimento")
		hora_procedimento := r.FormValue("hora_procedimento")
		tipo_procedimento := r.FormValue("tipo_procedimento")
		valor_procedimento := r.FormValue("valor_procedimento")
		profissional_designado := r.FormValue("profissional_designado")
		paciente_id := r.FormValue("paciente_id")

		convertnumero, err := strconv.Atoi(valor_procedimento)
		if err != nil {
			log.Println("Erro ao converter valor_procedimento para inteiro:", err)
			http.Error(w, "Valor do procedimento inválido", http.StatusBadRequest)
			return
		}
		// Convertendo `paciente_id` para inteiro
		pacienteIdInt, err := strconv.Atoi(paciente_id)
		if err != nil {
			log.Println("Erro ao converter paciente_id para inteiro:", err)
			http.Error(w, "ID do paciente inválido", http.StatusBadRequest)
			return
		}

		// Inserindo procedimento
		modelo.InseriProcedimento(
			nome_procedimento,
			data_procedimento,
			hora_procedimento,
			tipo_procedimento,
			convertnumero,
			profissional_designado,
			pacienteIdInt,
		)

		http.Redirect(w, r, "/Central_Registros", http.StatusSeeOther)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// Deletar procedimento
func Deletar_Procedimentos(w http.ResponseWriter, r *http.Request) {
	idProcedimento := r.URL.Query().Get("id_procedimento")
	modelo.DeletaProcedimento(idProcedimento)
	http.Redirect(w, r, "/", 301)

}

// Listar os procedimentos registrados
func Listar_Procedimentos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listando procedimentos registrados.")
	registros := modelo.BuscarProcedimento()

	// Tentando renderizar os dados no template
	err := tmp.ExecuteTemplate(w, "central_registros", registros)
	if err != nil {
		fmt.Println("Erro ao renderizar procedimentos:", err)
		http.Error(w, "Erro ao renderizar procedimentos", http.StatusInternalServerError)
		return
	}

	fmt.Println("Procedimentos listados e renderizados com sucesso.")
}

/*

// Deletar paciente
func Delete(w http.ResponseWriter, r *http.Request) {
	idUsuario := r.URL.Query().Get("id_paciente")
	modelo.Deleta(idUsuario)
	http.Redirect(w, r, "/listar.html", 301)
}

// Localiza as informações do paciente
func Edit(w http.ResponseWriter, r *http.Request) {
	idatualizar := r.URL.Query().Get("id_paciente")
	atualize := modelo.Editar(idatualizar)
	tmp.ExecuteTemplate(w, "Edit", atualize)
}


,
			/*
// Historico do paciente
func Historico_Paciente() {

}


// PROFISSIONAIS

// Listar os profissionais registrados
func Listar_Profissional() {

}


// Deleta um profissional
func Deletar_Profissional() {
	idProfissional := r.URL.Query().Get("id_profissional")
    modelo.Deletaprofissional(idProfissional)
    http.Redirect(w, r, "/listar_profissional.html", 301)

}

// Localiza as informações do profissional
func Editar_Profissional() {
	idProfissional := r.URL.Query().Get("id_profissional")
    atualize := modelo.Editarprofissional(idProfissional)
    tmp.ExecuteTemplate(w, "Edit_profissional", atualize)

}


// PROCEDIMENTOS

// Listar os procedimentos registrados
func Listar_Procedimentos() {

}

// Cadastra um novo profissional
func Inserir_Procedimentos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id_procedimento := r.FormValue("id_procedimento")
		nome_procedimento := r.FormValue("Nome_paciente")
		data_procedimento := r.FormValue("Data_procedimento")
		hora_procedimento := r.FormValue("Hora_procedimento")
		tipo_procedimento := r.FormValue("Tipo_procedimento")
		valor_procedimento := r.FormValue("Valor_procedimento")
		profissional_designado := r.FormValue("Profissional_designado")

		idConvertidaParaInt, err := strconv.Atoi(Id_procedimento)
		if err != nil {
            log.Println("Erro na convesão do ID para int:", err)
        }


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


// Localiza as informações do profissional
func Editar_Procedimentos() {

    idProcedimento := r.URL.Query().Get("id_procedimento")
	atualize := modelo.Editarprocedimento(idProcedimento)
	tmp.ExecuteTemplate(w, "Edit_procedimento", atualize)
}

// Atualiza as informações


func Atualizar_Procedimentos() {
	if r.Method == "POST" {
        Id_procedimento := r.FormValue("Id_procedimento")
        Nome_procedimento := r.FormValue("Nome_procedimento")
        Data_procedimento := r.FormValue("Data_procedimento")
        Hora_procedimento := r.FormValue("Hora_procedimento")
        Tipo_procedimento := r.FormValue("Tipo_procedimento")
        Valor_procedimento := r.FormValue("Valor_procedimento")
        Profissional_designado := r.FormValue("Profissional_designado")

        idConvertidaParaInt, err := strconv.Atoi(Id_procedimento)
        if err != nil {
            log.Println("Erro na convesão do ID para int:", err)
        }

		modelo.Atualizarprocedimento(
            idConvertidaParaInt,
            Nome_procedimento,
            Data_procedimento,
            Hora_procedimento,
			Tipo_procedimento,
            Valor_procedimento,
            Profissional_designado,
        )
	}
		http.Redirect(w, r, "/listar_procedimentos.html", 301)

}

*/
