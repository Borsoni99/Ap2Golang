package pedidos

import (
	"fmt"
	"time"
	i "mcronalds/itens"
)

const maxItensPedido = 10
const taxaDelivery = 10.0
var TotalPedidosJaCadastrados = 0

type Pedido struct {
	Id          int
	Delivery    bool
	ItensPedido [maxItensPedido]i.Item
	DataHora    time.Time
	TotalItens  int
	PrecoTotal  float64
}

/*
Define um id de um pedido, considerando todos os pedidos já cadastrados.
*/
func (p *Pedido) definirId() {
	TotalPedidosJaCadastrados++
	p.Id = TotalPedidosJaCadastrados
}

/*
Exibe as informações de um pedido no terminal.
*/
func (p *Pedido) exibir() {
	fmt.Println("\nPedido", p.Id)
	fmt.Println("Hora do pedido:", p.DataHora.Format("02/01/2006 15:04"))

	if p.Delivery {
		fmt.Println("É delivery? Sim")
	} else {
		fmt.Println("É delivery? Não")
	}

	fmt.Println("Itens no pedido:")

	for _, item := range p.ItensPedido {
		if (item == i.Item{}) { break }

		fmt.Println(item.Quant, "x", item.Prod.Nome)
	}

	fmt.Printf("Preço total: R$ %.2f\n", p.PrecoTotal)
}

/*
Adiciona um produto ao pedido.
Retorna -1 caso não seja mais possível adicionar itens.
Retorna -2 caso o id buscado não corresponda a um produto.
Retorna 0 em caso de sucesso.
*/
func (p *Pedido) AdicionarItem(id, quant int) int {
	if p.TotalItens == maxItensPedido { return -1 }

	item := i.Criar(id, quant)
	if (item == i.Item{}) { return -2 }

	p.ItensPedido[p.TotalItens] = item
	p.TotalItens++
	p.atualizarPreco(item)
	return 0
}

/*
Atualiza o preço total do pedido com o item fornecido.
*/
func (p *Pedido) atualizarPreco(item i.Item) {
	p.PrecoTotal += item.CalcularPrecoParcial()
}

/*
Retorna um elemento do tipo Pedido, com um id já definido.
*/
func criar(delivery bool) Pedido {
	p := Pedido{Delivery: delivery}
	p.TotalItens = 0
	p.DataHora = time.Now()
	p.definirId()

	p.PrecoTotal = 0.0
	if p.Delivery {
		p.PrecoTotal += taxaDelivery
	}

	return p
}
