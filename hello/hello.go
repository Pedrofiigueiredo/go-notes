package main

import (
	"fmt"
	/* "hello/math" */
)

var a string /* Declaração de variável */

func hello() {
	a = "Pedro"
	b := "Figueiredo"

	e := 10
	f := "Mundo"
	g := 3.144
	h := true

	fmt.Println("Pedro")
	fmt.Printf("%v", b)

	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", h)

	resultado := soma(1, 1)
	fmt.Printf("%v", resultado)
}

// não exportavel
func soma(a int, b int) int {
	return a + b
}