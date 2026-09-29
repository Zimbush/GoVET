package divideimpera

func MergeSort(eingabe []int) []int {
	if len(eingabe) <= 1 {
		result := make([]int, len(eingabe))
		copy(result, eingabe)
		return result
	}

	mitte := len(eingabe) / 2
	links := MergeSort(eingabe[:mitte])
	rechts := MergeSort(eingabe[mitte:])
	return merge(links, rechts)
}

func merge(links, rechts []int) []int {
	result := make([]int, 0, len(links)+len(rechts))
	l, r := 0, 0

	for l < len(links) && r < len(rechts) {
		if links[l] <= rechts[r] {
			result = append(result, links[l])
			l++
		} else {
			result = append(result, rechts[r])
			r++
		}
	}

	result = append(result, links[l:]...)
	result = append(result, rechts[r:]...)
	return result
}
