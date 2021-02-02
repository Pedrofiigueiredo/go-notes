package main

import (
	"fmt"
	/* "net/http" */
	"log"
	"errors"
)

func erros() {
	/* res, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(res.Header) */

	/* 
		blank identifier = _
		sempre tem que receber as variaveis,
		se - por qualquer motivo - eu não quiser tratar os erros,
		so usar _ e deixar vazio
	*/

	res2, err2 := somaa(10, 10)
	if err2 != nil {
		/* log.Fatal(err2.Error()) */
		fmt.Println("Total maior que zero")
	} else {
		fmt.Println(res2)
	}

	res3, err3 := somaa(1, 8)
	if err3 != nil {
		log.Fatal(err3.Error())
	}
	fmt.Println(res3)
}

func somaa(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("Total maior que zero")
	}
	
	return res, nil
}