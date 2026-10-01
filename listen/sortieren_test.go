package listen_test

import (
	"govet/listen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "leer", input: []int{}, want: []int{}},
		{name: "ein_element", input: []int{42}, want: []int{42}},
		{name: "sortiert", input: []int{3, 1, 2, 5, 4}, want: []int{1, 2, 3, 4, 5}},
		{name: "duplikate", input: []int{4, 2, 2, 1, 3}, want: []int{1, 2, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listen.InsertionSort(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "leer", input: []int{}, want: []int{}},
		{name: "ein_element", input: []int{9}, want: []int{9}},
		{name: "unsortiert", input: []int{4, 1, 3, 2, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "duplikate", input: []int{5, 2, 2, 1, 3}, want: []int{1, 2, 2, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listen.SelectionSort(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSelectionSortDelete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		want  []int
	}{
		{name: "leer", input: []int{}, want: []int{}},
		{name: "ein_element", input: []int{9}, want: []int{9}},
		{name: "unsortiert", input: []int{4, 1, 3, 2, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "duplikate", input: []int{5, 2, 2, 1, 3}, want: []int{1, 2, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listen.SelectionSortDelete(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
