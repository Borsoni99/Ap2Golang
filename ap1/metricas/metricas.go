package metricas

import "fmt"

type Metricas struct {
	tempoMedioExpedicao float64
	faturamentoTotal    float64
	produtosCadastrados int
	pedidosEncerrados   int
	pedidosEmAndamento  int
}

var M = Metricas{
	tempoMedioExpedicao: 0.0,
	faturamentoTotal:    0.0,
	produtosCadastrados: 0,
	pedidosEncerrados:   0,
	pedidosEmAndamento:  0,
}

func (m *Metricas) SomaProdutosCadastrados(valor int) {
	m.produtosCadastrados += valor
}

func (m *Metricas) SomaPedidosEmAndamento(valor int) {
	m.pedidosEmAndamento += valor
}

func (m *Metricas) AtualizaExpedicao(novoTempo int, valorVenda float64) {
	tempoTotalExpedicao := m.tempoMedioExpedicao * float64(m.pedidosEncerrados) + float64(novoTempo)
	m.pedidosEncerrados++
	m.tempoMedioExpedicao = tempoTotalExpedicao / float64(m.pedidosEncerrados)
	m.faturamentoTotal += valorVenda
}

func (m *Metricas) ExibirMetricas() {
	fmt.Println("\nMétricas de resultados do sistema:")
	fmt.Println("Número de produtos cadastrados:", m.produtosCadastrados)
	fmt.Println("Número de pedidos em andamento:", m.pedidosEmAndamento)
	fmt.Println("Número de pedidos encerrados:", m.pedidosEncerrados)
	fmt.Printf("Tempo médio para expedição de pedidos (em min): %.2f\n", m.tempoMedioExpedicao)
	fmt.Printf("Faturamento total: R$ %.2f\n", m.faturamentoTotal)
}
