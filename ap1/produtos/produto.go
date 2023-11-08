package produtos

import "fmt"

var TotalProdutosJaCadastrados = 0

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
}

/*
Define um id de um produto, considerando todos os produtos já cadastrados.
*/
func (p *Produto) definirId() {
	TotalProdutosJaCadastrados++
	p.Id = TotalProdutosJaCadastrados
}

/*
Exibe as informações de um produto no terminal.
*/
func (p *Produto) Exibir() {
	fmt.Println("\nProduto", p.Id)
	fmt.Println(p.Nome)
	fmt.Println(p.Descricao)
	fmt.Printf("Preço: R$ %.2f\n", p.Preco)
}

/*
Retorna um elemento do tipo Produto, com um id a ser definido ou com um id
pré-definido.
*/
func criar(nome, descricao string, preco float64, id int) Produto {
	p := Produto { Nome: nome, Descricao: descricao, Preco: preco }
	if id == -1 {
		p.definirId()
	} else {
		p.Id = id
	}

	return p
}
