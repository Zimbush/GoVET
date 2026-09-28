package listen

import "slices"

func InsertionSort(eingabe []int) []int {
	ergebnis := make([]int, 0)

	for _, wert := range eingabe {
		position := len(ergebnis)
		for position > 0 && ergebnis[position-1] > wert {
			position--
		}

		ergebnis = append(ergebnis, 0)
		copy(ergebnis[position+1:], ergebnis[position:])
		ergebnis[position] = wert
	}

	return ergebnis
}

func SelectionSort(eingabe []int) []int {
	ergebnis := make([]int, len(eingabe))
	copy(ergebnis, eingabe)

	for i := 0; i < len(ergebnis)-1; i++ {
		kleinstes := i
		for j := i + 1; j < len(ergebnis); j++ {
			if ergebnis[j] < ergebnis[kleinstes] {
				kleinstes = j
			}
		}

		if kleinstes != i {
			ergebnis[i], ergebnis[kleinstes] = ergebnis[kleinstes], ergebnis[i]
		}
	}

	return ergebnis
}

func SelectionSortDelete(eingabe []int) []int {
	ergebnis := make([]int, 0)
	for len(eingabe) > 0 {
		min_index := 0
		kandidat := eingabe[min_index]
		for index := 1; index < len(eingabe); index++ {
			if eingabe[index] < kandidat {
				min_index = index
				kandidat = eingabe[index]
			}
		}
		ergebnis = append(ergebnis, kandidat)
		eingabe = slices.Delete(eingabe, min_index, min_index+1)
	}

	return ergebnis

}
