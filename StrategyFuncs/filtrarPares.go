package StrategyFuncs

//Essa função filtraPares recebe um array de inteiros e retorna um novo array contendo apenas os números pares.
func FiltraPares(array []int) []int {
	rem := make([]int, len(array))
	k := 0

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			rem[k] = array[i]
			k++
		}
	}
	return rem[0:k]
}


