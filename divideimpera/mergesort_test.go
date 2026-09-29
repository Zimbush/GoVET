package divideimpera_test

import (
	"testing"

	"divideimpera"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "leer", input: []int{}, want: []int{}},
		{name: "ein_element", input: []int{42}, want: []int{42}},
		{name: "bereits_sortiert", input: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
		{name: "ungeordnet", input: []int{9, 4, 7, 1, 3, 8, 2, 5}, want: []int{1, 2, 3, 4, 5, 7, 8, 9}},
		{name: "negative_werte", input: []int{3, -1, 2, 0, -5}, want: []int{-5, -1, 0, 2, 3}},
		{name: "mehrfache_werte", input: []int{4, 2, 4, 1, 2, 3, 1}, want: []int{1, 1, 2, 2, 3, 4, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divideimpera.MergeSort(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
