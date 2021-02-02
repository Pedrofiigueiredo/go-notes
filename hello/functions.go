package main

import "fmt"

// Carro struct, +- POO...
type Carro struct {
	Nome string
}


// Metodo
func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func functions() {
	carro := Carro {
		Nome: "BMW",
	}

	carro.andar()

	res, str := mult(3, 2)
	fmt.Println(res, str)

	res2 := somaTudo(3, 2, 5, 10, 1)
	fmt.Println(res2) 
}

// Retorno multiplo
func mult(a int, b int) (int, string) {
	return a * b, "multiplicou"
}

// Retorno nomeado
func mult2(a int, b int) (resultado int) {
	resultado = a * b
	return 
}

// Funcoes variadicas
func somaTudo(x ...int) int { /* Recebe como se fosse um array */
	/* Voltado para quando eu nao sei o num de entradas  */
	resultado := 0

	for _, v := range x { /* Faz um loop em todos os valores de x */
		resultado += v
	}

	return resultado
}