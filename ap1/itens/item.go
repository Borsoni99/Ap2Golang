package itens

import (
	p "mcronalds/produtos"
)

type Item struct {
	Prod *p.Produto
	Quant int
}

func (p *Item) CalcularPrecoParcial() float64 {
	return p.Prod.Preco * float64(p.Quant)
}

/*
Retorna um Item com as informações solicitadas.
Se o id não existir para um produto, retorna um Item vazio.
*/
func Criar(id int, quant int) Item {
	produto, _ := p.BuscarId(id)
	if (produto == p.Produto{}) { return Item{} }

	return Item{Prod: &produto, Quant: quant}
}
