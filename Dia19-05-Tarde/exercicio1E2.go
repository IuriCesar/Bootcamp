package main

import "fmt"

func main() {
	fmt.Printf("Exercicio 1:\n")
	exercicio1()
	fmt.Println()
	fmt.Println()
	fmt.Printf("\nExercicio 2:\n")
	exercicio2()
}

//A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
//das letras separadamente para soletrá-la. Para isso terão que:
//1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
//letras que ela contém.
//2. Em seguida, imprimi cada uma das letras.
func exercicio1() {
	palavra := "Ornitorrinco"

	fmt.Println(palavra)
	for _, c := range palavra {
		fmt.Printf("%s\n", string(c))
	}
	fmt.Printf("Essa palavra tem: %d letras.\n", len(palavra))
}

/*Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.*/
func exercicio2() {

	var idade = []int{20, 25, 30}
	var emprego = []bool{true, true, true}
	var anoAtividade = []int{2, 1, 2}
	var salario = []float32{1000.5, 50000, 200000}

	for i := 0; i < 3; i++ {
		fmt.Printf("\nIdade: %d, Trabalha: %t, Anos atividade: %d, Salário: %f\n", idade[i], emprego[i], anoAtividade[i], salario[i])
		if idade[i] > 22 && emprego[i] && anoAtividade[i] > 1 {
			if salario[i] > 100000 {
				fmt.Println("  È possível conceder um empréstimo a você, no qual não cobrarei juros.")
			} else {
				fmt.Println("  È possível conceder um empréstimo a você, no qual cobrarei juros.")
			}
		} else {
			fmt.Println("  Não é possível conceder um empréstimo a você. Tem que ter mais de 22 anos, estar empregado e com mais de um ano serviço.")
		}
	}

}
