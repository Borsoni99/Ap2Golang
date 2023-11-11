package pedidos

import (
	"fmt"
	m "mcronalds/metricas"
	"time"
)

const maxPedidos = 1000

var Pedidos [maxPedidos]Pedido
var totalPedidos = 0
var inicioFila, fimFila = -1, -1

/*
Adiciona um pedido com a informação se é ou não delivery.
Retorna o ponteiro para o pedido, no caso de sucesso, ou nil em caso de erro
*/
func Adicionar(delivery bool) *Pedido {
	if totalPedidos == maxPedidos {
		return nil
	} // Overflow

	fimFila++
	totalPedidos++
	if fimFila == maxPedidos {
		fimFila = 0
	}
	Pedidos[fimFila] = criar(delivery)

	if inicioFila == -1 {
		inicioFila = 0
	}

	m.M.SomaPedidosEmAndamento(1)
	return &Pedidos[fimFila]
}

/*
Remove um primeiro pedido criado.
Retorna o pedido removido, ou um pedido vazio em caso de erro.
*/
func excluir() Pedido {
	if totalPedidos == 0 {
		return Pedido{}
	} // Underflow

	pedidoRemovido := Pedidos[inicioFila]
	Pedidos[inicioFila] = Pedido{}

	if inicioFila == fimFila {
		inicioFila, fimFila = -1, -1
	} else {
		inicioFila++
		if inicioFila == maxPedidos {
			inicioFila = 0
		}
	}

	totalPedidos--

	m.M.SomaPedidosEmAndamento(-1)
	return pedidoRemovido
}

/*
Realiza a expedição de um pedido.
Remove o pedido da fila, exibe as informações na tela e o horário da expedição.
Atualiza as métricas do sistema.
Retorna 0 em caso de sucesso, ou -1 caso não haja pedidos na fila.
*/
func Expedir() int {
	pedidoExpedido := excluir()
	if (pedidoExpedido == Pedido{}) {
		return -1
	}

	timeStampExpedicao := time.Now()
	fmt.Println("\nPedido entregue!")
	fmt.Println(timeStampExpedicao.Format("02/01/2006 15:04"))
	pedidoExpedido.exibir()

	tempoExpedicao := timeStampExpedicao.Sub(pedidoExpedido.DataHora).Minutes()
	m.M.AtualizaExpedicao(int(tempoExpedicao), pedidoExpedido.PrecoTotal)
	return 0
}

/*
Exibe todos os pedidos ativos.
*/
func Exibir() {
	if inicioFila == -1 {
		fmt.Println("Fila de pedidos vazia!")
		return
	}

	i := inicioFila

	for {
		if (Pedidos[i] == Pedido{}) { break }
		Pedidos[i].exibir()

		i++
		if i == maxPedidos { i = 0 }
	}
}
