package bruchrechnen_test

import (
	"govet/bruchrechnen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMult(t *testing.T) {
	tests := []struct {
		name string
		op1  bruchrechnen.Zahl
		op2  bruchrechnen.Zahl
		want float64
	}{
		{name: "bruch_mal_bruch", op1: bruchrechnen.Bruch{Zaehler: 2, Nenner: 3}, op2: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, want: 0.5},
		{name: "bruch_mal_kommazahl", op1: bruchrechnen.Bruch{Zaehler: 1, Nenner: 2}, op2: bruchrechnen.Kommazahl{Zahl: 0.5}, want: 0.25},
		{name: "kommazahl_mal_bruch", op1: bruchrechnen.Kommazahl{Zahl: 2.4}, op2: bruchrechnen.Bruch{Zaehler: 2, Nenner: 3}, want: 1.6},
		{name: "kommazahl_mal_kommazahl", op1: bruchrechnen.Kommazahl{Zahl: 2.5}, op2: bruchrechnen.Kommazahl{Zahl: 2.0}, want: 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruchrechnen.Mult(tt.op1, tt.op2)
			assert.InDelta(t, tt.want, got.BerechneWert(), 0.0001)
		})
	}
}
