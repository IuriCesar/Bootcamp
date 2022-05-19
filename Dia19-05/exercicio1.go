package main

import "fmt"

func main() {
	fmt.Println("Exercício 1:")
	exercicio1Manha()
	fmt.Println()
	fmt.Println("Exercício 2:")
	exercicio2Manha()
}

func exercicio1Manha() {
	var nome = "Iuri"
	var idade = 25

	fmt.Println(nome, idade)
}
func exercicio2Manha() {
	var temp int32 = 36
	var umid int32 = 25
	var atm int32 = 79

	fmt.Printf("A sua cidade está:\n  Com %d ºC\n  Umidade em %d\n  Com %d ATM\n", temp, umid, atm)
}
