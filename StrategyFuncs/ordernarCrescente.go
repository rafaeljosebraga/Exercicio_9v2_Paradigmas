package strategy

func swap(pos1, pos2 *int) {
	aux := *pos1
	*pos1 = *pos2
	*pos2 = aux
}

func ordenarCrescente(array []int) []int { // // Ordena uma cópia de array de números recebidos por Cocktail Sort e o retorna
	var swapped bool = true
	var start int = 0
	var end int = len(array) - 1

	for swapped {
		swapped = false

		for i := start; i < end; i++ { // Empurra os maiores pro final
			if array[i] > array[i+1] {
				swap(&array[i], &array[i+1])
				swapped = true
			}
		} 
		if !swapped { // Termina se não empurrou ninguém
			break
		}
		swapped = false
		end--

		for i := end - 1; i >= start; i-- { // Empurra os menores pro começo
			if array[i] > array[i+1] {
				swap(&array[i], &array[i+1])
				swapped = true
			}
		}
		start++
	}
	return array
}