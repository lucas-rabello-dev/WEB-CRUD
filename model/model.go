package model

type Livro struct {
	Titulo           string
	Autor            string
	Idioma           string
	NumPages         int
	Editora          string
	Genero           string
	AnoPublicacao    string   // data de publicação do livro
	LinkCompra       []string // pode ser um ou mais
	Sinopse          string
	ValorMultaDiaria float64
}

func (l *Livro) CalcularMulta(valorOriginal float64) {
	l.ValorMultaDiaria = valorOriginal * 30 / 100
}

type User struct {
	NomeCompleto string
	Idade        int
	CPF          string
	Email        string
	Email_Rec    string // email de recuperação se precisar
	senha        string
}

// escrever o getter e setter para o atributo senha
