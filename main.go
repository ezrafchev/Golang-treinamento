package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	arquivoFlag := flag.String("arquivo", "pessoas.txt", "Arquivo contendo informações das pessoas (nome, idade)")
	flag.Parse()

	pessoas, err := lerDadosDoArquivo(*arquivoFlag)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	mediaIdades := calcularMediaIdades(pessoas)
	fmt.Printf("\nA média das idades é %.2f.\n", mediaIdades)
}

func lerDadosDoArquivo(nomeArquivo string) ([]Pessoa, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	var pessoas []Pessoa
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		parts := strings.Split(linha, ",")
		if len(parts) != 2 {
			continue
		}

		nome := strings.TrimSpace(parts[0])
		idade, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}

		pessoas = append(pessoas, Pessoa{Nome: nome, Idade: idade})
	}

	return pessoas, scanner.Err()
}

func calcularMediaIdades(pessoas []Pessoa) float64 {
	total := 0.0
	for _, pessoa := range pessoas {
		total += float64(pessoa.Idade)
	}

	return total / float64(len(pessoas))
}
