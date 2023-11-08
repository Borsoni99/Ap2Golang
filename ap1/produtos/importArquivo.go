package produtos

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const nomeArquivo = "dados.csv"

func LerArquivo() {
	fmt.Println("Importando dados de produtos de arquivo .csv...")

	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer arquivo.Close()

	reader := csv.NewReader(arquivo)

	linhas, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo CSV:", err)
		return
	}

	for _, linha := range linhas[1:] {
		id, _ := strconv.Atoi(linha[0])
		nome := linha[1]
		descricao := linha[2]
		preco, _ := strconv.ParseFloat(linha[3], 64)

		ret := AdicionarUnico(nome, descricao, preco, id)
		if ret < 0 {
			fmt.Println("Ocorreu um erro ao adicionar o produto:", id, nome)
		}
	}

	fmt.Printf("Leitura de arquivo concluída! Iniciando programa...\n\n\n")
}
