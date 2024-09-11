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

	inserir.Exec(nome, nomedamae, nomedopai, telefone, cpf, email, senha)
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
		var id int
		var nome string
		var nomedamae string
		var nomedopai string
		var telefone int
		var cpf int
		var email string
		var senha int

		err = selectReg.Scan(&id, &nome, &nomedamae, &nomedopai, &telefone, &cpf, &email, &senha)
		if err != nil {
			panic(err.Error())
		}

		l.Id = id
		l.Nome = nome
		l.Nomedamae = nomedamae
		l.Nomedopai = nomedopai
		l.Telefone = telefone
		l.CPF = cpf
		l.Email = email
		l.Senha = senha

		logar = append(logar, l)
	}

	defer db.Close()
	return logar

}

func Deleta(id string) {
	db := db.Acesse()

	deletar, err := db.Prepare("delete from login where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id)
	defer db.Close()

}

func Editar(id string) Login {
	db := db.Acesse()

	atualizar, err := db.Query("select * from login where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	Atualize := Login{}

	for atualizar.Next() {
		var id int
		var nome, nomedamae, nomedopai string
		var telefone, cpf int
		var email string
		var senha int
		err = atualizar.Scan(&id, &nome, &nomedamae, &nomedopai, &telefone, &cpf, &email, &senha)
		if err != nil {
			panic(err.Error())
		}
		Atualize.Id = id
		Atualize.Nome = nome
		Atualize.Nomedamae = nomedamae
		Atualize.Nomedopai = nomedopai
		Atualize.Telefone = telefone
		Atualize.CPF = cpf
		Atualize.Email = email
		Atualize.Senha = senha
	}

	defer db.Close()
	return Atualize
}

func Atualizar(id int, nome, nomedamae, nomedopai string, telefone, cpf int, email string, senha int) {
	db := db.Acesse()

	atualiza, err := db.Prepare("update login set nome=$1, nomedamae=$2, nomedopai=$3, telefone=$4, cpf=$5, email=$6, senha=$7 where id=$8")
	if err != nil {
		panic(err.Error())
	}
	atualiza.Exec(nome, nomedamae, nomedopai, telefone, cpf, email, senha, id)
	defer db.Close()
}
