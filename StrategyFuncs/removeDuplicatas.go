/* 
A função removeDuplicatas() recebe uma lista de inteiros com duplicatas e retorna uma nova lista contendo apenas os elementos únicos.
Ela preserva a ordem original dos itens, mantendo apenas a primeira vez que cada número aparece e descartando as repetições.
*/

package StrategyFuncs

func RemoverDuplicatas(lista []int) []int {
	chavesVistas := make(map[int]bool)
	resultado := []int{}

	for _, item := range lista {
		if _, visto := chavesVistas[item]; !visto {
			chavesVistas[item] = true
			resultado = append(resultado, item)
		}
	}
	return resultado
}

