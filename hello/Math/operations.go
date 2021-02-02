package math /* "namespace" */

// A para ser exportado tem que ter 1a letra maiuscula
var A string = "Isso é visível quando math é importado"

// Sum é uma função para somar dois números
func Sum(a int, b int) int {
	return a + b
}

/* 
	Sum != sum
	A primeira letra maiuscula deixa a func exportavel

	Documentação das funções -> // {NOME DA FUNÇÂO} ...
	Util com Go Docs + organização dentro do codigo
*/