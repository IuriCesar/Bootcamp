package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type produto struct {
	Id    int32
	Preco float32
	Qtd   int
}

func main() {

	var saida []produto

	var id int32 = 0
	var p float32 = 1
	qtd := 1
	file := "teste.txt"

	// linhas referentes ao exercicio 1
	for i := 0; i < 4; i++ {
		item := produto{id, p, qtd}
		saida = append(saida, item)
		id++
		p++
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	for _, c := range saida {
		dados := fmt.Sprintf("%d;%.2f;%d", c.Id, c.Preco, c.Qtd)
		f.WriteString(dados + "\n")
	}
	//

	// linhas referente ao exercicio 2
	linha, err := os.Open(file)
	if err != nil {
		log.Fatalf("Não consegui abrir o arquivo")
	}
	scanner := bufio.NewScanner(linha)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	linha.Close()
	var soma float32 = 0
	for _, each_ln := range text {
		dados := strings.Split(each_ln, ";")
		ident, _ := strconv.Atoi(dados[0])
		prec, _ := strconv.ParseFloat(dados[1], 32)
		qt, _ := strconv.Atoi(dados[2])
		fmt.Printf("%d\t %.2f\t %d\n", ident, prec, qt)
		soma += float32(prec) * float32(qt)
	}
	fmt.Printf("\t%.2f\n", soma)
	//
}
