/*A função ordenarDecrescente() recebe uma lista de inteiros, realiza uma cópia dessa lista
e utiliza o algoritmo Bubble Sort para ordená-la em ordem decrescente e retorna a cópia 
ordenada da lista.*/

package StrategyFuncs

func OrdenarDecrescente(lista []int) []int {
	tamanho := len(lista)
	copia := make([]int, tamanho)
	copy(copia, lista)

	for i := 0; i < tamanho-1; i++ {
		for j := 0; j < tamanho-1-i; j++ {
			if copia[j] < copia[j+1] {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}
	return copia
}