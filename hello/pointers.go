package main

import "fmt"

func main() {
	/* a := 10
	fmt.Println(a)
	fmt.Println(&a)

	var ponteiro *int = &a
	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)

	*ponteiro = 50
	fmt.Println("\n",*ponteiro)
	fmt.Println(a) */

	variavel := 10

	abc(&variavel)

	fmt.Println(variavel)
}

func abc(a *int) {
	*a = 200
}