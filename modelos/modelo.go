package modelo

import "modulo/db"

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
	Funcao_profissional   string
	Telefone_profissional int
	Endereco_profissional string
	Cidade_profissional   string
	Bairro_profissional   string
	Cpf_profissional      int
	Email_profissional    string
	Senha_profissional    int
}

//PACIENTE

//CADASTRAR PACIENTE
func Inseri(nome_paciente, tipo_plano string, telefone_paciente int, endereco_paciente string, cidade_paciente string, bairro_paciente string, cpf_paciente int, email_paciente string, senha_paciente int) {
	db := db.Acesse()

	inserir, err := db.Prepare("insert into loginpaciente (nome_paciente, tipo_plano, telefone_paciente, endereco_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		panic(err.Error())
	}

	inserir.Exec(nome_paciente, tipo_plano, telefone_paciente, endereco_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente)
	defer db.Close()
}

//AUTENTICAÇÃO DE USUÁRIO
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
		var nome_paciente string
		var tipo_plano string
		var telefone_paciente int
		var endereco_paciente string
		var cidade_paciente string
		var bairro_paciente string
		var cpf_paciente int
		var email_paciente string
		var senha_paciente int

		err = selectReg.Scan(&id_paciente, &nome_paciente, &tipo_plano, &telefone_paciente, &endereco_paciente, &cidade_paciente, &bairro_paciente, &cpf_paciente, &email_paciente, &senha_paciente)
		if err != nil {
			panic(err.Error())
		}

		l.Id_paciente = id_paciente
		l.Nome_paciente = nome_paciente
		l.Tipo_plano = tipo_plano
		l.Telefone_paciente = telefone_paciente
		l.Endereco_paciente = endereco_paciente
		l.Cidade_paciente = cidade_paciente
		l.Bairro_paciente = bairro_paciente
		l.Cpf_paciente = cpf_paciente
		l.Email_paciente = email_paciente
		l.Senha_paciente = senha_paciente

		logar = append(logar, l)
	}

	defer db.Close()
	return logar

}

//DELETAR USUÁRIO
func Deleta(id_paciente string) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from loginpaciente where id_paciente=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id_paciente)
	defer db.Close()

}

//EDITAR USUÁRIO
func Editar(id_paciente string) LoginPaciente {
	db := db.Acesse()

	atualizar, err := db.Query("select * from loginpaciente where id_paciente = $1", id_paciente)
	if err != nil {
		panic(err.Error())
	}

	Atualize := LoginPaciente{}

	//CRIA AS VARIAVEIS PARA INTERPRETAR VALORES.
	for atualizar.Next() {
		var id_paciente int
		var nome_paciente, tipo_plano string
		var telefone_paciente int
		var endereco_paciente string
		var cidade_paciente string
		var bairro_paciente string
		var cpf_paciente int
		var email_paciente string
		var senha_paciente int
		//INTERPRETA OS DADOS DO BANCO, COLETANDO DADOS ATUAIS
		err = atualizar.Scan(&id_paciente, &nome_paciente, &telefone_paciente, &tipo_plano, &endereco_paciente, &cidade_paciente, &bairro_paciente, &cpf_paciente, &email_paciente, &senha_paciente)
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

//LANÇA AS ATUALIZAÇÕES DAS VARIAVEIS NO BANCO DE DADOS
func Atualizar(id_paciente int, nome_paciente, tipo_plano string, telefone_paciente int, endereco_paciente string, cidade_paciente string, bairro_paciente string, cpf_paciente int, email_paciente string, senha_paciente int) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update loginpaciente set nome_paciente=$1, tipo_plano=$2, telefone_paciente=$3, endereco_paciente=$4, cidade_paciente=$5, bairro_paciente=$6, cpf_paciente=$7, email_paciente=$8, senha_paciente=$9 where id_paciente=$10")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(nome_paciente, tipo_plano, telefone_paciente, endereco_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente, id_paciente)
	defer db.Close()
}

//PROFISSIONAL

//CADASTRAR PROFISSIONAL
func Inseriprofissional(nome_profissional, funcao_profissional string, telefone_profissional int, endereco_profissional string, cidade_profissional string, bairro_profissional string, cpf_profissional int, email_profissional string, senha_profissional int) {
	db := db.Acesse()

	inserir, err := db.Prepare("insert into loginprofissional (nome_profissional, funcao_profissional, telefone_profissional, endereco_profissional, cidade_profissional, bairro_profissional, cpf_profissional, email_profissional, senha_profissional) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		panic(err.Error())
	}

	inserir.Exec(nome_profissional, funcao_profissional, telefone_profissional, endereco_profissional, cidade_profissional, bairro_profissional, cpf_profissional, email_profissional, senha_profissional)
	defer db.Close()
}

//AUTENTICAÇÃO DE USUÁRIO
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
		var funcao_profissional string
		var telefone_profissional int
		var endereco_profissional string
		var cidade_profissional string
		var bairro_profissional string
		var cpf_profissional int
		var email_profissional string
		var senha_profissional int

		err = selectReg.Scan(&id_profissional, &nome_profissional, &funcao_profissional, &telefone_profissional, &endereco_profissional, &cidade_profissional, &bairro_profissional, &cpf_profissional, &email_profissional, &senha_profissional)
		if err != nil {
			panic(err.Error())
		}

		l.Id_profissional = id_profissional
		l.Nome_profissional = nome_profissional
		l.Funcao_profissional = funcao_profissional
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

//DELETAR USUÁRIO
func Deletaprofissional(id_profissional string) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from loginprofissional where id_profissional=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id_profissional)
	defer db.Close()

}

//EDITAR USUÁRIO
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
		var nome_profissional, funcao_profissional string
		var telefone_profissional int
		var endereco_profissional string
		var cidade_profissional string
		var bairro_profissional string
		var cpf_profissional int
		var email_profissional string
		var senha_profissional int
		//INTERPRETA OS DADOS DO BANCO, COLETANDO DADOS ATUAIS
		err = atualizar.Scan(&id_profissional, &nome_profissional, &telefone_profissional, &funcao_profissional, &endereco_profissional, &cidade_profissional, &bairro_profissional, &cpf_profissional, &email_profissional, &senha_profissional)
		if err != nil {
			panic(err.Error())
		}
		//ATUALIZA O VALOR DAS VARIAVEIS
		Atualize.Id_profissional = id_profissional
		Atualize.Nome_profissional = nome_profissional
		Atualize.Funcao_profissional = funcao_profissional
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

//LANÇA AS ATUALIZAÇÕES DAS VARIAVEIS NO BANCO DE DADOS
func Atualizarprofissional(id_profissional int, nome_profissional, funcao_profissional string, telefone_profissional int, endereco_profissional string, cidade_profissional string, bairro_profissional string, cpf_profissional int, email_profissional string, senha_profissional int) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update loginprofissional set nome_profissional=$1, funcao_profissional=$2, telefone_profissional=$3, endereco_profissional=$4, cidade_profissional=$5, bairro_profissional=$6, cpf_profissionale=$7, email_profissional=$8, senha_profissional=$9 where id_profissional=$10")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(nome_profissional, funcao_profissional, telefone_profissional, endereco_profissional, cidade_profissional, bairro_profissional, cpf_profissional, email_profissional, senha_profissional, id_profissional)
	defer db.Close()
}
