package main

type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint
}

type Estudante struct {
	Pessoa   // pega todos os campos de Pessoa, reaproveitando os campos
	curso    string
	semestre uint
}

func main() {

}
