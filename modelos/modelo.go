package modelo

import "modulo/db"

type LoginPaciente struct {
	Id_paciente       int
	Nome_paciente     string
	Tipo_plano        string
	Telefone_paciente int
	Cidade_paciente   string
	bairro_paciente   string
	Cpf_paciente      int
	Email_paciente    string
	Senha_paciente    int
}

func Inseri(nome_paciente, tipo_plano string, telefone_paciente int, cpf int, email string, senha int) {
	db := db.Acesse()

	inserir, err := db.Prepare("insert into login (nome, nomedamae, nomedopai, telefone, cpf, email, senha) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}

	inserir.Exec(nome_paciente, tipo_plano, telefone_paciente, cpf, email, senha)
	defer db.Close()
}

func Buscar() []Login {
	db := db.Acesse()

	selectReg, err := db.Query("select * from login")
	if err != nil {
		panic(err.Error())
	}

	l := Login{}
	logar := []Login{}

	for selectReg.Next() {
		var id_paciente int
		var nome_paciente string
		var tipo_plano string
		var telefone_paciente int
		var cpf int
		var email string
		var senha int

		err = selectReg.Scan(&id_paciente, &nome_paciente, &tipo_plano, &telefone_paciente, &cpf, &email, &senha)
		if err != nil {
			panic(err.Error())
		}

		l.Id_paciente = id_paciente
		l.Nome_paciente = nome_paciente
		l.Tipo_plano = tipo_plano
		l.Telefone_paciente = telefone_paciente
		l.CPF = cpf
		l.Email = email
		l.Senha = senha

		logar = append(logar, l)
	}

	defer db.Close()
	return logar

}

func Deleta(id_paciente string) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from login where id_paciente=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id_paciente)
	defer db.Close()

}

func Editar(id_paciente string) Login {
	db := db.Acesse()

	atualizar, err := db.Query("select * from login where id_paciente = $1", id_paciente)
	if err != nil {
		panic(err.Error())
	}

	Atualize := Login{}

	for atualizar.Next() {
		var id_paciente int
		var nome_paciente, tipo_plano string
		var telefone_paciente, cpf int
		var email string
		var senha int
		err = atualizar.Scan(&id_paciente, &nome_paciente, &telefone_paciente, &tipo_plano, &cpf, &email, &senha)
		if err != nil {
			panic(err.Error())
		}
		Atualize.Id_paciente = id_paciente
		Atualize.Nome_paciente = nome_paciente
		Atualize.Tipo_plano = tipo_plano
		Atualize.Telefone_paciente = telefone_paciente
		Atualize.CPF = cpf
		Atualize.Email = email
		Atualize.Senha = senha
	}

	defer db.Close()
	return Atualize
}

func Atualizar(id_paciente int, nome_paciente, tipo_plano string, telefone_paciente, cpf int, email string, senha int) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update login set nome_paciente=$1, tipo_plano=$2, telefone_paciente=$3, cpf=$4, cpf=$5, email=$6, senha=$7 where id_paciente=$8")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(nome_paciente, tipo_plano, telefone_paciente, cpf, email, senha, id_paciente)
	defer db.Close()
}
