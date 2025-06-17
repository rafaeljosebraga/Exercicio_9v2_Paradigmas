package main

import (
	"fmt"
	"StrategyFuncs/strategy"
)

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func main() {
	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	// executar_estrategia(&primes, removeDuplicates)
	fmt.Println("Remoção de duplicatas: ",primes)
	primes2 := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&primes, StrategyFuncs.FiltraPares)
	fmt.Println("Pares Filtrados: ",primes2)
}
