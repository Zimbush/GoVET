package bruchrechnen_test

import (
	"govet/bruchrechnen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErmittleNachkommastellen(t *testing.T) {
	tests := []struct {
		name string
		zahl float64
		want int
	}{
		{name: "ganze_zahl", zahl: 42.0, want: 0},
		{name: "einfacher_bruch", zahl: 3.14, want: 2},
		{name: "negative_wert", zahl: -2.75, want: 2},
		{name: "viele_nachkommastellen", zahl: 0.000001, want: 6},
		{name: "null", zahl: 0.0, want: 0},
		{name: "decimal_edge_case", zahl: 2.5000001, want: 7},
		{name: "trailing_zero_ignored", zahl: 2.500, want: 1},
		{name: "kleiner_als_1", zahl: 0.125, want: 3},
		{name: "kleiner_als_1_negativ", zahl: -0.125, want: 3},
		{name: "nahe_bei_ganze_zahl", zahl: 10.0000001, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruchrechnen.ErmittleNachkommastellen(tt.zahl)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestKommmazahl_BerechneWert(t *testing.T) {
	tests := []struct {
		name string
		zahl bruchrechnen.Kommazahl
		want float64
	}{
		{name: "ganze_zahl", zahl: bruchrechnen.Kommazahl{Zahl: 42.0}, want: 42.0},
		{name: "kommazahl", zahl: bruchrechnen.Kommazahl{Zahl: 3.14}, want: 3.14},
		{name: "negative", zahl: bruchrechnen.Kommazahl{Zahl: -2.75}, want: -2.75},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.zahl.BerechneWert()
			assert.InDelta(t, tt.want, got, 0.0001)
		})
	}
}

func TestKommmazahl_BerechneBruch(t *testing.T) {
	tests := []struct {
		name string
		zahl bruchrechnen.Kommazahl
		want bruchrechnen.Bruch
	}{
		{name: "ganze_zahl", zahl: bruchrechnen.Kommazahl{Zahl: 42.0}, want: bruchrechnen.Bruch{Zaehler: 42, Nenner: 1}},
		{name: "einfacher_bruch", zahl: bruchrechnen.Kommazahl{Zahl: 3.14}, want: bruchrechnen.Bruch{Zaehler: 314, Nenner: 100}},
		{name: "negative", zahl: bruchrechnen.Kommazahl{Zahl: -2.75}, want: bruchrechnen.Bruch{Zaehler: -275, Nenner: 100}},
		{name: "null", zahl: bruchrechnen.Kommazahl{Zahl: 0.0}, want: bruchrechnen.Bruch{Zaehler: 0, Nenner: 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.zahl.BerechneBruch()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestKommmazahl_Arithmetik(t *testing.T) {
	base := bruchrechnen.Kommazahl{Zahl: 2.5}

	t.Run("subtrahiere", func(t *testing.T) {
		got := base.Subtrahiere(bruchrechnen.Kommazahl{Zahl: 0.5})
		assert.InDelta(t, 2.0, got.BerechneWert(), 0.0001)
	})

	t.Run("multipliziere", func(t *testing.T) {
		got := base.Multipliziere(bruchrechnen.Kommazahl{Zahl: 2.0})
		assert.InDelta(t, 5.0, got.BerechneWert(), 0.0001)
	})

	t.Run("dividiere", func(t *testing.T) {
		got, err := base.Dividiere(bruchrechnen.Kommazahl{Zahl: 0.5})
		assert.NoError(t, err)
		assert.InDelta(t, 5.0, got.BerechneWert(), 0.0001)
	})

	t.Run("dividiere_durch_null", func(t *testing.T) {
		_, err := base.Dividiere(bruchrechnen.Kommazahl{Zahl: 0})
		assert.Error(t, err)
	})
}
