package model

type Livro struct {
	Titulo string
	Autor string
	Idioma string
	NumPages int
	Editora string
	Genero string
	AnoPublicacao string // data de publicação do livro
	LinkCompra []string // pode ser um ou mais
	Sinopse string
}

type User struct {
	NomeCompleto string
	Idade int
	CPF string
	Email string
	Email_Rec string // email de recuperação se precisar
	senha string
	IsAdmin bool
}


// escrever o getter e setter para o atributo senha