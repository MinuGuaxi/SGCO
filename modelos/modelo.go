package modelo

import "modulo/db"

type LoginPaciente struct {
	Id_paciente       int
	Nome_paciente     string
	Tipo_plano        string
	Telefone_paciente int
	Cidade_paciente   string
	Bairro_paciente   string
	Cpf_paciente      int
	Email_paciente    string
	Senha_paciente    int
}

func Inseri(nome_paciente, tipo_plano string, telefone_paciente int, cidade_paciente string, bairro_paciente string, cpf_paciente int, email_paciente string, senha_paciente int) {
	db := db.Acesse()

	inserir, err := db.Prepare("insert into loginpaciente (nome_paciente, tipo_plano, telefone_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}

	inserir.Exec(nome_paciente, tipo_plano, telefone_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente)
	defer db.Close()
}

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
		var cidade_paciente string
		var bairro_paciente string
		var cpf_paciente int
		var email_paciente string
		var senha_paciente int

		err = selectReg.Scan(&id_paciente, &nome_paciente, &tipo_plano, &telefone_paciente, &cidade_paciente, &bairro_paciente, &cpf_paciente, &email_paciente, &senha_paciente)
		if err != nil {
			panic(err.Error())
		}

		l.Id_paciente = id_paciente
		l.Nome_paciente = nome_paciente
		l.Tipo_plano = tipo_plano
		l.Telefone_paciente = telefone_paciente
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

func Deleta(id_paciente string) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from loginpaciente where id_paciente=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id_paciente)
	defer db.Close()

}

func Editar(id_paciente string) LoginPaciente {
	db := db.Acesse()

	atualizar, err := db.Query("select * from loginpaciente where id_paciente = $1", id_paciente)
	if err != nil {
		panic(err.Error())
	}

	Atualize := LoginPaciente{}

	for atualizar.Next() {
		var id_paciente int
		var nome_paciente, tipo_plano string
		var telefone_paciente int
		var cidade_paciente string
		var bairro_paciente string
		var cpf_paciente int
		var email_paciente string
		var senha_paciente int
		err = atualizar.Scan(&id_paciente, &nome_paciente, &telefone_paciente, &tipo_plano, &cidade_paciente, &bairro_paciente, &cpf_paciente, &email_paciente, &senha_paciente)
		if err != nil {
			panic(err.Error())
		}
		Atualize.Id_paciente = id_paciente
		Atualize.Nome_paciente = nome_paciente
		Atualize.Tipo_plano = tipo_plano
		Atualize.Telefone_paciente = telefone_paciente
		Atualize.Cidade_paciente = cidade_paciente
		Atualize.Bairro_paciente = bairro_paciente
		Atualize.Cpf_paciente = cpf_paciente
		Atualize.Email_paciente = email_paciente
		Atualize.Senha_paciente = senha_paciente
	}

	defer db.Close()
	return Atualize
}

func Atualizar(id_paciente int, nome_paciente, tipo_plano string, telefone_paciente int, cidade_paciente string, bairro_paciente string, cpf_paciente int, email_paciente string, senha_paciente int) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update loginpaciente set nome_paciente=$1, tipo_plano=$2, telefone_paciente=$3, cidade_paciente=$4, bairro_paciente=$5, cpf_paciente=$6, email_paciente=$7, senha_paciente=$8 where id_paciente=$9")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(nome_paciente, tipo_plano, telefone_paciente, cidade_paciente, bairro_paciente, cpf_paciente, email_paciente, senha_paciente, id_paciente)
	defer db.Close()
}
