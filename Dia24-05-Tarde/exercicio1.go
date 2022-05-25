package main

import "fmt"

// Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
// que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
// que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
// para as funções:
// - A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
// senha
// E devem implementar as seguintes funcionalidades:
// - mudar o nome: me permite mudar o nome e sobrenome
// - mudar a idade: me permite mudar a idade
// - mudar e-mail: me permite mudar o e-mail
// - mudar a senha: me permite mudar a senha
type Produto struct {
	Id        int
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     string
}

type produtos []*Produto

func (p *produtos) addProduto(produto *Produto) {
	*p = append(*p, produto)
}

func (p *produtos) editNomeProduto(id int, nome string, sobrenome string) {
	for _, item := range *p {
		if item.Id == id {
			item.Nome = nome
			item.Sobrenome = sobrenome
		}
	}
}

func (p *produtos) editIdadeProduto(id int, idade int) {
	for _, item := range *p {
		if item.Id == id {
			item.Idade = idade
		}
	}
}

func (p *produtos) editEmailProduto(id int, email string) {
	for _, item := range *p {
		if item.Id == id {
			item.Email = email
		}
	}
}

func (p *produtos) editSenhaProduto(id int, senha string) {
	for _, item := range *p {
		if item.Id == id {
			item.Senha = senha
		}
	}
}

func main() {
	pessoa := Produto{1, "Iuri", "Cesar Caliman", 26, "lol@teste.com", "123456"}
	pessoas := produtos{}

	pessoas.addProduto(&pessoa)

	for _, item := range pessoas {
		fmt.Println(*item)
	}

	pessoas.editEmailProduto(1, "pucca@geor.com")
	pessoas.editIdadeProduto(1, 25)
	pessoas.editNomeProduto(1, "Irui", "Rasec Namilac")
	pessoas.editSenhaProduto(1, "654321")

	for _, item := range pessoas {
		fmt.Println(*item)
	}
}
