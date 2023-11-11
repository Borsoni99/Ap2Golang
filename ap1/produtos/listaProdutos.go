package produtos

import (
	m "mcronalds/metricas"
	"strings"
)


type No struct {
	chave Produto
	prox  *No
}

type Lista struct {
	cab   *No
}


func (l *Lista) insere(produto Produto) {
	novoNo := &No{chave: produto}

	no := l.cab
	if no == nil {
		l.cab = novoNo
	} else {
		for no.prox != nil { no = no.prox }
		no.prox = novoNo
	}
}

func (l *Lista) exibe() {
	no := l.cab

	for no != nil {
		no.chave.Exibir()
		no = no.prox
	}
}

func (l *Lista) buscaPorProduto(produto Produto) *No {
	no := l.cab

	for no != nil {
		if no.chave == produto { return no }
		no = no.prox
	}

	return nil
}

func (l *Lista) buscaPorId(id int) *No {
	no := l.cab

	for no != nil {
		if no.chave.Id == id { return no }
		no = no.prox
	}

	return nil

}

func (l *Lista) buscaPorNome(nome string) *No {
	no := l.cab

	for no != nil {
		if no.chave.Nome == nome { return no }
		no = no.prox
	}

	return nil
}

func (l *Lista) remove(produto Produto) {
	if l.cab == nil {
		return
	}

	if l.cab.chave == produto {
		l.cab = l.cab.prox
		return
	}

	noAnterior := l.cab
	for noAtual := l.cab.prox; noAtual != nil; noAtual = noAtual.prox {
		if noAtual.chave == produto {
			noAnterior.prox = noAtual.prox
			return
		}
		noAnterior = noAtual
	}
}

const maxProdutos = 50

var lista = Lista{cab: nil}

func tentarCriar(nome, descricao string, preco float64, id int) Produto {
	if id != -1 {
		_, idProcurado := BuscarId(id)
		if idProcurado != -1 { return Produto{} }
	}

	return criar(nome, descricao, preco, id)
}

/*
Adiciona um produto com nome, descrição e preço à lista de produtos.
Adiciona o produto no primeiro espaço vazio da lista.
Caso já exista um produto com o mesmo id, não adiciona e retorna -3.
Caso já exista um produto com o mesmo nome, não adiciona e retorna erro -2.
Retorna -1 caso a lista esteja cheia, ou o número de produtos cadastrados em caso de sucesso.
*/
func AdicionarUnico(nome, descricao string, preco float64, id int) int {
	if listaTamanho() == maxProdutos {
		return -1 // Overflow
	}

	if lista.buscaPorNome(nome) != nil {
		return -2 // Já existe um produto com o mesmo nome
	}

	if lista.buscaPorProduto(Produto{Id: id}) != nil {
		return -3 // Já existe um produto com o mesmo id
	}

	produtoCriado := tentarCriar(nome, descricao, preco, id)
	if produtoCriado == (Produto{}) {
		return -3 // Falha ao criar o produto
	}

	lista.insere(produtoCriado)
	m.M.SomaProdutosCadastrados(1)
	return listaTamanho()
}

/*
Localiza um produto a partir do seu id.
Retorna o produto encontrado e a sua posição na lista, em caso de sucesso.
Retorna um produto vazio e -1 em caso de erro.
*/
func BuscarId(id int) (*Produto, int) {
	no := lista.buscaPorId(id)
	if no != nil {
		return &no.chave, 0
	}

	return nil, -1
}

/*
Localiza produtos que iniciem com a string passada.
Retorna um slice com todos os produtos encontrados, e o tamanho do slice.
*/
func BuscarNome(comecaCom string) ([]Produto, int) {
	var produtosEncontrados []Produto

	no := lista.buscaPorNome(comecaCom)
	for no != nil {
		produtosEncontrados = append(produtosEncontrados, no.chave)
		no = no.prox
	}

	return produtosEncontrados, len(produtosEncontrados)
}

/*
Exibe todos os produtos cadastrados.
*/
func Exibir() {
	lista.exibe()
}

/*
Remove um produto da lista a partir do seu id.
Retorna -2 caso não haja produtos na lista.
Retorna -1 caso não haja um produto com o id passado, ou 0 em caso de sucesso.
*/
func Excluir(id int) int {
	if listaTamanho() == 0 {
		return -2 // Lista vazia
	}

	no := lista.buscaPorId(id)
	if no == nil {
		return -1 // Produto com o id não encontrado
	}

	lista.remove(no.chave)
	m.M.SomaProdutosCadastrados(-1)
	return 0
}

/*
Atualiza um produto da lista a partir do seu id.
Retorna -2 caso não haja produtos na lista.
Retorna -1 caso não haja um produto com o id passado, ou 0 em caso de sucesso.
*/
func Atualizar(id int, valor float64) int {
	if listaTamanho() == 0 {
		return -2 // Lista vazia
	}

	no := lista.buscaPorId(id)
	if no == nil {
		return -1 // Produto com o id não encontrado
	}

	no.chave.Preco = valor
	return 0
}

// Função auxiliar para obter o tamanho da lista
func listaTamanho() int {
	no := lista.cab
	count := 0

	for no != nil {
		count++
		no = no.prox
	}

	return count
}

func OrdenarPorNome() {
    trocou := true
    limite := listaTamanho()

    for trocou && limite > 1 {
        trocou = false
        noAtual := lista.cab
        noAnterior := lista.cab
        proximo := lista.cab.prox

        for i := 0; i < limite-1 && noAtual != nil && proximo != nil; i++ {
            if strings.Compare(noAtual.chave.Nome, proximo.chave.Nome) > 0 {
                if noAtual == lista.cab {
                    lista.cab = proximo
                } else {
                    noAnterior.prox = proximo
                }
                noAtual.prox, proximo.prox = proximo.prox, noAtual
                noAtual, trocou = proximo, true
            }
            noAtual, noAnterior, proximo = proximo, noAtual, proximo.prox
        }
        limite--
    }
}



