package modelo

import (
	"database/sql"
	"fmt"
	"log"
	"modulo/db"
)

type LoginPaciente struct {
	Id_paciente       int
	Nome_paciente     string
	Tipo_plano        string
	Telefone_paciente int
	Endereco_paciente string
	Cidade_paciente   string
	Bairro_paciente   string
	Cpf_paciente      int
	Email_paciente    string
	Senha_paciente    int
}

type LoginProfissional struct {
	Id_profissional       int
	Nome_profissional     string
	Telefone_profissional int
	Endereco_profissional string
	Cidade_profissional   string
	Bairro_profissional   string
	Cpf_profissional      int
	Email_profissional    string
	Senha_profissional    int
}

type Procedimento struct {
	Id_procedimento        int
	Nome_procedimento      string
	Data_procedimento      string
	Hora_procedimento      string
	Tipo_procedimento      string
	Valor_procedimento     int
	Profissional_Designado string
	Paciente_id            int // Novo campo adicionado
}

//PACIENTE

// CADASTRAR PACIENTE
func Inseri(nome_paciente string, tipo_plano string, telefone_paciente int, endereco_paciente string, cidade_paciente string, bairro_paciente string, cpf_paciente int, email_paciente string, senha_paciente int) {
	db := db.Acesse()

	inserir, err := db.Prepare("INSERT INTO loginpaciente (nome_paciente, tipo_plano, telefone_paciente, endereco_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		log.Println("Erro ao tentar salvar", err)
		return
	}

	// Executar a inserção
	_, execErr := inserir.Exec(nome_paciente, tipo_plano, telefone_paciente, endereco_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente)
	if execErr != nil {
		log.Println("Erro ao executar SQL:", execErr)
	} else {
		log.Println("Dados inseridos com sucesso em:", nome_paciente)
	}

	defer db.Close()
}

//AUTENTICAÇÃO DE USUÁRIO

// Função de autenticação
func AutenticaUsuario(Email_paciente string, Senha_paciente int) (bool, error) {
	db := db.Acesse()

	var senhaHash int
	query := "SELECT senha_paciente FROM loginpaciente WHERE email_paciente=$1"
	err := db.QueryRow(query, Email_paciente).Scan(&senhaHash)

	if err != nil {
		if err == sql.ErrNoRows {
			// Usuário não encontrado
			fmt.Println("Usuário não encontrado:", Email_paciente)
			return false, nil
		}
		// Outro erro (ex: erro na query)
		fmt.Println("Erro ao executar query:", err)
		return false, err
	}
	fmt.Println("Senha encontrada:", senhaHash)
	if senhaHash == Senha_paciente {
		// Autenticado com sucesso
		return true, nil
	}

	// Senha incorreta
	fmt.Println("Senha incorreta para o email:", Email_paciente)
	return false, nil
}

//AUTENTICAÇÃO DE USUÁRIO

// Função de autenticação
func AutenticaProfissional(email_profissional string, senha_profissional int) (bool, error) {
	db := db.Acesse()

	var senhaHash int
	query := "SELECT senha_profissional FROM loginprofissional WHERE email_profissional=$1"
	err := db.QueryRow(query, email_profissional).Scan(&senhaHash)

	if err != nil {
		if err == sql.ErrNoRows {
			// Usuário não encontrado
			fmt.Println("Usuário não encontrado:", email_profissional)
			return false, nil
		}
		// Outro erro (ex: erro na query)
		fmt.Println("Erro ao executar query:", err)
		return false, err
	}
	fmt.Println("Senha encontrada:", senhaHash)
	if senhaHash == senha_profissional {
		// Autenticado com sucesso
		return true, nil
	}

	// Senha incorreta
	fmt.Println("Senha incorreta para o email:", email_profissional)
	return false, nil
}

/*
func Buscar() []LoginPaciente {
	db := db.Acesse()

	selectReg, err := db.Query("select * from loginpaciente")
	if err != nil {
		panic(err.Error())
	}

	l := LoginPaciente{}
	logar := []LoginPaciente{}

	for selectReg.Next() {
		var id_paciente int
		var email_paciente string
		var senha_paciente int

		err = selectReg.Scan(&id_paciente, &email_paciente, &senha_paciente)
		if err != nil {
			panic(err.Error())
		}

		l.Id_paciente = id_paciente
		l.Email_paciente = email_paciente
		l.Senha_paciente = senha_paciente

		logar = append(logar, l)
	}

	defer db.Close()
	return logar

}
*/
//DELETAR USUÁRIO
func Deleta(id_paciente int) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from loginpaciente where id_paciente=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id_paciente)
	defer db.Close()

}

// EDITAR USUÁRIO
func Editar(id_paciente int) LoginPaciente {
	db := db.Acesse()

	atualizar, err := db.Query("select * from loginpaciente where id_paciente = $1", id_paciente)
	if err != nil {
		panic(err.Error())
	}

	Atualize := LoginPaciente{}

	//CRIA AS VARIAVEIS PARA INTERPRETAR VALORES.
	for atualizar.Next() {
		var id_paciente int
		var nome_paciente string
		var tipo_plano string
		var telefone_paciente int
		var endereco_paciente string
		var cidade_paciente string
		var bairro_paciente string
		var cpf_paciente int
		var email_paciente string
		var senha_paciente int
		//INTERPRETA OS DADOS DO BANCO, COLETANDO DADOS ATUAIS
		err = atualizar.Scan(&id_paciente, &nome_paciente, &tipo_plano, &telefone_paciente, &endereco_paciente, &cidade_paciente, &bairro_paciente, &cpf_paciente, &email_paciente, &senha_paciente)
		if err != nil {
			panic(err.Error())
		}
		//ATUALIZA O VALOR DAS VARIAVEIS
		Atualize.Id_paciente = id_paciente
		Atualize.Nome_paciente = nome_paciente
		Atualize.Tipo_plano = tipo_plano
		Atualize.Telefone_paciente = telefone_paciente
		Atualize.Endereco_paciente = endereco_paciente
		Atualize.Cidade_paciente = cidade_paciente
		Atualize.Bairro_paciente = bairro_paciente
		Atualize.Cpf_paciente = cpf_paciente
		Atualize.Email_paciente = email_paciente
		Atualize.Senha_paciente = senha_paciente
	}
	//SALVA AS ALTERAÇÕES
	defer db.Close()
	return Atualize
}

// LANÇA AS ATUALIZAÇÕES DAS VARIAVEIS NO BANCO DE DADOS
func Atualizar(id_paciente int, nome_paciente string, tipo_plano string, telefone_paciente int, endereco_paciente string, cidade_paciente string, bairro_paciente string, cpf_paciente int, email_paciente string, senha_paciente int) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update loginpaciente set nome_paciente=$1, tipo_plano=$2, telefone_paciente=$3, endereco_paciente=$4, cidade_paciente=$5, bairro_paciente=$6, cpf_paciente=$7, email_paciente=$8, senha_paciente=$9 where id_paciente=$10")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(id_paciente, nome_paciente, tipo_plano, telefone_paciente, endereco_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente)
	defer db.Close()
}

//PROFISSIONAL

// CADASTRAR PROFISSIONAL
func Inseriprofissional(nome_profissional string, telefone_profissional int, endereco_profissional string, cidade_profissional string, bairro_profissional string, cpf_profissional int, email_profissional string, senha_profissional int) {
	db := db.Acesse()

	inserir, err := db.Prepare("INSERT INTO loginprofissional (nome_profissional, telefone_profissional, endereco_profissional, cidade_profissional, bairro_profissional, cpf_profissional, email_profissional, senha_profissional) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		log.Println("Erro ao tentar salvar", err)
		return
	}

	// Executar a inserção
	_, execErr := inserir.Exec(nome_profissional, telefone_profissional, endereco_profissional, cidade_profissional, bairro_profissional, cpf_profissional, email_profissional, senha_profissional)
	if execErr != nil {
		log.Println("Erro ao executar SQL:", execErr)
	} else {
		log.Println("Dados inseridos com sucesso em:", nome_profissional)
	}

	defer db.Close()
}

// AUTENTICAÇÃO DE USUÁRIO
func Buscarprofissional() []LoginProfissional {
	db := db.Acesse()

	selectReg, err := db.Query("select * from loginprofissional")
	if err != nil {
		panic(err.Error())
	}

	l := LoginProfissional{}
	logar := []LoginProfissional{}

	for selectReg.Next() {
		var id_profissional int
		var nome_profissional string
		var telefone_profissional int
		var endereco_profissional string
		var cidade_profissional string
		var bairro_profissional string
		var cpf_profissional int
		var email_profissional string
		var senha_profissional int

		err = selectReg.Scan(&id_profissional, &nome_profissional, &telefone_profissional, &endereco_profissional, &cidade_profissional, &bairro_profissional, &cpf_profissional, &email_profissional, &senha_profissional)
		if err != nil {
			panic(err.Error())
		}

		l.Id_profissional = id_profissional
		l.Nome_profissional = nome_profissional
		l.Telefone_profissional = telefone_profissional
		l.Endereco_profissional = endereco_profissional
		l.Cidade_profissional = cidade_profissional
		l.Bairro_profissional = bairro_profissional
		l.Cpf_profissional = cpf_profissional
		l.Email_profissional = email_profissional
		l.Senha_profissional = senha_profissional

		logar = append(logar, l)
	}

	defer db.Close()
	return logar

}

// DELETAR USUÁRIO
func Deletaprofissional(id_profissional string) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from loginprofissional where id_profissional=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id_profissional)
	defer db.Close()

}

// EDITAR USUÁRIO
func Editarprofissional(id_profissional string) LoginProfissional {
	db := db.Acesse()

	atualizar, err := db.Query("select * from loginprofissional where id_profissional = $1", id_profissional)
	if err != nil {
		panic(err.Error())
	}

	Atualize := LoginProfissional{}

	//CRIA AS VARIAVEIS PARA INTERPRETAR VALORES
	for atualizar.Next() {
		var id_profissional int
		var nome_profissional string
		var telefone_profissional int
		var endereco_profissional string
		var cidade_profissional string
		var bairro_profissional string
		var cpf_profissional int
		var email_profissional string
		var senha_profissional int
		//INTERPRETA OS DADOS DO BANCO, COLETANDO DADOS ATUAIS
		err = atualizar.Scan(&id_profissional, &nome_profissional, &telefone_profissional, &endereco_profissional, &cidade_profissional, &bairro_profissional, &cpf_profissional, &email_profissional, &senha_profissional)
		if err != nil {
			panic(err.Error())
		}
		//ATUALIZA O VALOR DAS VARIAVEIS
		Atualize.Id_profissional = id_profissional
		Atualize.Nome_profissional = nome_profissional
		Atualize.Telefone_profissional = telefone_profissional
		Atualize.Endereco_profissional = endereco_profissional
		Atualize.Cidade_profissional = cidade_profissional
		Atualize.Bairro_profissional = bairro_profissional
		Atualize.Cpf_profissional = cpf_profissional
		Atualize.Email_profissional = email_profissional
		Atualize.Senha_profissional = senha_profissional
	}
	//SALVA AS ALTERAÇÕES
	defer db.Close()
	return Atualize
}

// LANÇA AS ATUALIZAÇÕES DAS VARIAVEIS NO BANCO DE DADOS
func Atualizarprofissional(id_profissional int, nome_profissional string, telefone_profissional int, endereco_profissional string, cidade_profissional string, bairro_profissional string, cpf_profissional int, email_profissional string, senha_profissional int) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update loginprofissional set nome_profissional=$1, telefone_profissional=$2, endereco_profissional=$3, cidade_profissional=$4, bairro_profissional=$5, cpf_profissionale=$6, email_profissional=$7, senha_profissional=$8 where id_profissional=$9")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(nome_profissional, telefone_profissional, endereco_profissional, cidade_profissional, bairro_profissional, cpf_profissional, email_profissional, senha_profissional, id_profissional)
	defer db.Close()
}

//PROCEDIMENTO

// CADASTRAR PROCEDIMENTO
func InseriProcedimento(nome_procedimento string, data_procedimento string, hora_procedimento string, tipo_procedimento string, valor_procedimento int, profissional_designado string, paciente_id int) {
	db := db.Acesse()

	inserir, err := db.Prepare("INSERT INTO procedimento (nome_procedimento, data_procedimento, hora_procedimento, tipo_procedimento, valor_procedimento, profissional_designado, paciente_id) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Println("Erro ao preparar a inserção:", err)
	}

	_, execErr := inserir.Exec(nome_procedimento, data_procedimento, hora_procedimento, tipo_procedimento, valor_procedimento, profissional_designado, paciente_id)
	if execErr != nil {
		log.Println("Erro ao executar a inserção:", execErr)
	}

	defer db.Close()
}

// REALIZA UMA BUSCA PELO BANCO DE DADOS
func BuscarProcedimento() []Procedimento {
	db := db.Acesse()

	selectReg, err := db.Query("select * from procedimento")
	if err != nil {
		panic(err.Error())
	}
	//DEFINE A FORMA DE AUTENTICAÇÃO
	l := Procedimento{}
	logar := []Procedimento{}

	for selectReg.Next() {
		var id_procedimento int
		var nome_procedimento string
		var data_procedimento string
		var hora_procedimento string
		var tipo_procedimento string
		var valor_procedimento int
		var profissional_designado string
		var paciente_id int

		err = selectReg.Scan(&id_procedimento, &nome_procedimento, &data_procedimento, &hora_procedimento, &tipo_procedimento, &valor_procedimento, &profissional_designado, &paciente_id)
		if err != nil {
			panic(err.Error())
		}

		l.Id_procedimento = id_procedimento
		l.Nome_procedimento = nome_procedimento
		l.Data_procedimento = data_procedimento
		l.Hora_procedimento = hora_procedimento
		l.Tipo_procedimento = tipo_procedimento
		l.Valor_procedimento = valor_procedimento
		l.Profissional_Designado = profissional_designado
		l.Paciente_id = paciente_id

		logar = append(logar, l)
	}

	defer db.Close()
	return logar

}

// DELETAR PROCEDIMENTO
func DeletaProcedimento(id_procedimento string) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from procedimento where id_procedimento=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id_procedimento)
	defer db.Close()

}

// EDITAR PROCEDIMENTO
func EditarProcedimento(id_procedimento string) Procedimento {
	db := db.Acesse()

	atualizar, err := db.Query("select * from procedimento where id_procedimento = $1", id_procedimento)
	if err != nil {
		panic(err.Error())
	}

	Atualize := Procedimento{}

	//CRIA AS VARIAVEIS PARA INTERPRETAR VALORES.
	for atualizar.Next() {
		var id_procedimento int
		var nome_procedimento string
		var data_procedimento string
		var hora_procedimento string
		var tipo_procedimento string
		var valor_procedimento int
		var profissional_designado string
		//INTERPRETA OS DADOS DO BANCO, COLETANDO DADOS ATUAIS
		err = atualizar.Scan(&id_procedimento, &nome_procedimento, &data_procedimento, &hora_procedimento, &tipo_procedimento, &valor_procedimento, &profissional_designado)
		if err != nil {
			panic(err.Error())
		}
		//ATUALIZA O VALOR DAS VARIAVEIS
		Atualize.Id_procedimento = id_procedimento
		Atualize.Nome_procedimento = nome_procedimento
		Atualize.Data_procedimento = data_procedimento
		Atualize.Hora_procedimento = hora_procedimento
		Atualize.Tipo_procedimento = tipo_procedimento
		Atualize.Valor_procedimento = valor_procedimento
		Atualize.Profissional_Designado = profissional_designado

	}
	//SALVA AS ALTERAÇÕES
	defer db.Close()
	return Atualize
}

// LANÇA AS ATUALIZAÇÕES DAS VARIAVEIS NO BANCO DE DADOS
func AtualizarProcedimento(id_procedimento int, nome_procedimento string, data_procedimento string, hora_procedimento string, tipo_procedimento string, valor_procedimento string, profissional_designado string) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update procedimento set nome_paciente=$1, data_procedimento=$2, hora_procedimento=$3, tipo_procedimento=$4, valor_procedimento=$5, profissional_designado=$6 where id_paciente=$7")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(nome_procedimento, data_procedimento, hora_procedimento, tipo_procedimento, valor_procedimento, profissional_designado, id_procedimento)
	defer db.Close()
}
