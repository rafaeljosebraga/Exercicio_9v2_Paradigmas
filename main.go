package main

import (
	"fmt"
	"github.com/rafaeljosebraga/Exercicio_9v2_Paradigmas/StrategyFuncs"
)

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func main() {

	lista := []int{1, 2, 25, 7, 5, 4, 3, 12}
	executar_estrategia(&lista, StrategyFuncs.OrdenarDecrescente)
	fmt.Println("Lista Decrescente: ", lista)
}
