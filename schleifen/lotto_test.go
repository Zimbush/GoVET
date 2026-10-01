package schleifen_test

import (
	"govet/schleifen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProdukt(t *testing.T) {
	tests := []struct {
		name string
		von  int
		bis  int
		want int
	}{
		{name: "einfach", von: 1, bis: 3, want: 6},
		{name: "mit_fuenf", von: 3, bis: 5, want: 60},
		{name: "einzelwert", von: 5, bis: 5, want: 5},
		{name: "nullbereich", von: 1, bis: 0, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := schleifen.Produkt(tt.von, tt.bis)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFakultaet(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "null", n: 0, want: 1},
		{name: "eins", n: 1, want: 1},
		{name: "fuenf", n: 5, want: 120},
		{name: "sechs", n: 6, want: 720},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := schleifen.Fakultaet(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProduktRekursiv(t *testing.T) {
	tests := []struct {
		name string
		von  int
		bis  int
		want int
	}{
		{name: "einfach", von: 1, bis: 3, want: 6},
		{name: "mit_fuenf", von: 3, bis: 5, want: 60},
		{name: "einzelwert", von: 5, bis: 5, want: 5},
		{name: "nullbereich", von: 1, bis: 0, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := schleifen.ProduktRekursiv(tt.von, tt.bis)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFakultaetRekursiv(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "null", n: 0, want: 1},
		{name: "eins", n: 1, want: 1},
		{name: "fuenf", n: 5, want: 120},
		{name: "sechs", n: 6, want: 720},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := schleifen.FakultaetRekursiv(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSechsAusNeunundvierzig(t *testing.T) {
	got := schleifen.SechsAusNeunundvierzig()
	assert.Equal(t, 13983816, got)
}
