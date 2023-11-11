package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"mcronalds/metricas"
	pedidos "mcronalds/pedidos"
	produtos "mcronalds/produtos"
)

var scanner = bufio.NewReader(os.Stdin)

func opcoes() {
	fmt.Println("")
	fmt.Println("1 - Cadastrar produto;")
	fmt.Println("2 - Remover produto;")
	fmt.Println("3 - Buscar produto por id;")
	fmt.Println("4 - Buscar produto por nome;")
	fmt.Println("5 - Exibir todos os produtos;")
	fmt.Println("6 - Adicionar pedido;")
	fmt.Println("7 - Expedir pedido;")
	fmt.Println("8 - Exibir métricas do sistema;")
	fmt.Println("9 - Atualizar valor Produto;")
	fmt.Println("10 - Ordenar produtos por nome;")
	fmt.Println("20 - Exibir todos os pedidos em andamento;")
	fmt.Println("21 - Cadastrar produtos em lote;")
	fmt.Println("100 - Sair do programa;")
}

func Cli() {
	var opcao string
	fmt.Println("Bem-vindo ao McRonald's! Selecione uma das opções abaixo ou qualquer outra tecla para encerrar o programa:")
	for {
		opcoes()
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			cadastrarProduto()
		case "2":
			removerProduto()
		case "3":
			buscarProdutoId()
		case "4":
			buscarProdutoNome()
		case "5":
			produtos.Exibir()
		case "6":
			adicionarPedido()
		case "7":
			pedidos.Expedir()
		case "8":
			metricas.M.ExibirMetricas()
		case "9":
			atualizarProduto()
		case "10":
			produtos.OrdenarPorNome()
		case "20":
			pedidos.Exibir()
		case "21":
			cadastrarProdutosEmLote()
		case "100":
			fmt.Println("Volte sempre!")
			return
		}
	}
}

func atualizarProduto() {
	id := leInt("Informe o id do produto a ser atualizado: ")
	valor := leFloat("Informe o novo valor do produto: ")
	produtos.Atualizar(id,valor)
}

func leTexto(prompt string) string {
	fmt.Print(prompt)
	msg, _ := scanner.ReadString('\n')
	msg = strings.TrimRight(msg, "\n")
	msg = strings.TrimRight(msg, "\r\n")
	return msg
}

func leFloat(prompt string) float64 {
	var valor float64
	fmt.Print(prompt)
	fmt.Scanln(&valor)
	return valor
}

func leInt(prompt string) int {
	var valor int
	fmt.Print(prompt)
	fmt.Scanln(&valor)
	return valor
}

func cadastrarProduto() {
	nome := leTexto("Nome do produto: ")
	descricao := leTexto("Descrição: ")
	preco := leFloat("Preço do produto (em R$): ")

	ret := produtos.AdicionarUnico(nome, descricao, preco, -1)
	switch ret {
	case -2:
		fmt.Println("Erro! Produto já existe no cadastro.")
	case -1:
		fmt.Println("Erro! Lista de produtos está cheia.")
	default:
		fmt.Println("Produto cadastrado com sucesso!")
	}
}

func removerProduto() {
	id := leInt("Informe o id do produto a ser removido: ")

	ret := produtos.Excluir(id)
	switch ret {
	case -2:
		fmt.Println("Erro! Lista de produtos está vazia.")
	case -1:
		fmt.Println("Erro! Produto buscado não existe.")
	default:
		fmt.Println("Produto removido com sucesso!")
	}
}

func buscarProdutoId() {
	id := leInt("Informe o id do produto a ser buscado: ")

	produtoEncontrado, indice := produtos.BuscarId(id)
	if indice == -1 {
		fmt.Println("Erro! Produto buscado não existe.")
	} else {
		produtoEncontrado.Exibir()
	}
}

func buscarProdutoNome() {
	comecaCom := leTexto("Informe o nome do produto ou o início do nome: ")

	produtosEncontrados, totalProdutosEncontrados := produtos.BuscarNome(comecaCom)
	if totalProdutosEncontrados == 0 {
		fmt.Println("Erro! Não foi encontrado nenhum produto com esse nome.")
	} else {
		for _, produtoEncontrado := range produtosEncontrados {
			produtoEncontrado.Exibir()
		}
	}
}

func adicionarPedido() {
	var delivery = false
	var idProduto, quantidade int
	opDelivery := leTexto("O pedido é para delivery (s/n)? ")
	if opDelivery == "s" {
		delivery = true
	}

	pedido := pedidos.Adicionar(delivery)
	fmt.Println("Digite, em cada linha abaixo, o id do produto e a quantidade, separados por um espaço. Digite 0 0 para encerrar.")
	for {
		fmt.Scanln(&idProduto, &quantidade)
		if idProduto == 0 && quantidade == 0 { break }

		ret := pedido.AdicionarItem(idProduto, quantidade)
		switch ret {
		case -1:
			fmt.Println("Não é possível adicionar mais itens ao pedido.")
		case -2:
			fmt.Println("Id passado não corresponde a um produto.")
		default:
			fmt.Println("Item adicionado com sucesso!")
		}
	}
}

func cadastrarProdutosEmLote() {
	var opcao string
	for {
		cadastrarProduto()

		opcao = leTexto("\nDeseja continuar com as inserções (s/n)? ")
		if opcao != "s" {
			break
		}
	}
}
