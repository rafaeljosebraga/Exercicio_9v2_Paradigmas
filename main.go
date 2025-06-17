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

	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&primes, StrategyFuncs.OrdenarCrescente)
	fmt.Println("Lista Ordenada: ", primes)

	executar_estrategia(&primes, StrategyFuncs.RemoverDuplicatas)
	fmt.Println("Lista de Remoção de duplicatas: ",primes)
  
	executar_estrategia(&primes, StrategyFuncs.FiltraPares)
	fmt.Println("Lista com Pares Filtrados: ", primes)

}
