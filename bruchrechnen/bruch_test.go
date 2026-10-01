package bruchrechnen_test

import (
	"govet/bruchrechnen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ggt(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "beide_null", a: 0, b: 0, want: 0},
		{name: "null_mit_wert", a: 0, b: 15, want: 15},
		{name: "gleiche_zahlen", a: 12, b: 12, want: 12},
		{name: "gemeinsamer_teiler", a: 24, b: 18, want: 6},
		{name: "teilerfremd", a: 21, b: 16, want: 1},
		{name: "groessere_zahl_erst", a: 81, b: 27, want: 27},
		{name: "kleinere_zahl_erst", a: 462, b: 1071, want: 21},
		{name: "neg_1", a: -462, b: 1071, want: 21},
		{name: "neg_2", a: 462, b: -1071, want: 21},
		{name: "neg_3", a: -462, b: -1071, want: 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruchrechnen.GGT(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewBruch(t *testing.T) {
	tests := []struct {
		name      string
		zaehler   int
		nenner    int
		wantValue bruchrechnen.Bruch
		wantErr   bool
	}{
		{name: "positiv", zaehler: 3, nenner: 4, wantValue: bruchrechnen.Bruch{3, 4}},
		{name: "negativer_zaehler", zaehler: -3, nenner: 4, wantValue: bruchrechnen.Bruch{-3, 4}},
		{name: "negativer_nenner", zaehler: 3, nenner: -4, wantValue: bruchrechnen.Bruch{3, -4}},
		{name: "null_zaehler", zaehler: 0, nenner: 5, wantValue: bruchrechnen.Bruch{0, 5}},
		{name: "nenner_null", zaehler: 7, nenner: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bruchrechnen.NewBruch(tt.zaehler, tt.nenner)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantValue, got)
		})
	}
}

func TestKGV(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "beide_null", a: 0, b: 0, want: 0},
		{name: "ein_zerowert", a: 0, b: 9, want: 0},
		{name: "teilerfremd", a: 4, b: 9, want: 36},
		{name: "gemeinsame_teiler", a: 6, b: 9, want: 18},
		{name: "vielfaches", a: 12, b: 6, want: 12},
		{name: "vielfaches_negativ", a: -15, b: 5, want: 15},
		{name: "negative_werte", a: -6, b: 9, want: 18},
		{name: "gleiche_zahlen", a: 7, b: 7, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruchrechnen.KGV(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_BerechneWert(t *testing.T) {
	tests := []struct {
		name  string
		bruch bruchrechnen.Bruch
		want  float64
	}{
		{"einfacher_bruch", bruchrechnen.Bruch{1, 2}, 0.5},
		{"dreiviertel", bruchrechnen.Bruch{3, 4}, 0.75},
		{"negative_werte", bruchrechnen.Bruch{-3, 4}, -0.75},
		{"null_zaehler", bruchrechnen.Bruch{0, 7}, 0.0},
		{"ganze_zahl", bruchrechnen.Bruch{9, 4}, 2.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bruch.BerechneWert()
			assert.InDelta(t, tt.want, got, 0.0001)
		})
	}
}

func TestBruch_Kuerze(t *testing.T) {
	tests := []struct {
		name  string
		bruch bruchrechnen.Bruch
		want  bruchrechnen.Bruch
	}{
		{name: "bereits_gekürzt", bruch: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, want: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}},
		{name: "gemeinsamer_teiler", bruch: bruchrechnen.Bruch{Zaehler: 8, Nenner: 12}, want: bruchrechnen.Bruch{Zaehler: 2, Nenner: 3}},
		{name: "negative_werte", bruch: bruchrechnen.Bruch{Zaehler: -12, Nenner: 18}, want: bruchrechnen.Bruch{Zaehler: -2, Nenner: 3}},
		{name: "null_zaehler", bruch: bruchrechnen.Bruch{Zaehler: 0, Nenner: 7}, want: bruchrechnen.Bruch{Zaehler: 0, Nenner: 1}},
		{name: "ganzer_bruch", bruch: bruchrechnen.Bruch{Zaehler: 9, Nenner: 3}, want: bruchrechnen.Bruch{Zaehler: 3, Nenner: 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bruch.Kuerze()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_Erweitere(t *testing.T) {
	tests := []struct {
		name   string
		bruch  bruchrechnen.Bruch
		faktor int
		want   bruchrechnen.Bruch
	}{
		{name: "normal", bruch: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, faktor: 2, want: bruchrechnen.Bruch{Zaehler: 6, Nenner: 8}},
		{name: "negativer_faktor", bruch: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, faktor: -2, want: bruchrechnen.Bruch{Zaehler: -6, Nenner: -8}},
		{name: "null_faktor", bruch: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, faktor: 0, want: bruchrechnen.Bruch{Zaehler: 0, Nenner: 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bruch.Erweitere(tt.faktor)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_Inverses(t *testing.T) {
	tests := []struct {
		name  string
		bruch bruchrechnen.Bruch
		want  bruchrechnen.Bruch
	}{
		{name: "positiv", bruch: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, want: bruchrechnen.Bruch{Zaehler: -3, Nenner: 4}},
		{name: "negativ", bruch: bruchrechnen.Bruch{Zaehler: -3, Nenner: 4}, want: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}},
		{name: "null", bruch: bruchrechnen.Bruch{Zaehler: 0, Nenner: 7}, want: bruchrechnen.Bruch{Zaehler: 0, Nenner: 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bruch.Inverses()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_Addiere(t *testing.T) {
	tests := []struct {
		name    string
		bruch   bruchrechnen.Bruch
		summand bruchrechnen.Bruch
		want    bruchrechnen.Bruch
	}{
		{name: "verschiedene_nenner", bruch: bruchrechnen.Bruch{Zaehler: 1, Nenner: 2}, summand: bruchrechnen.Bruch{Zaehler: 1, Nenner: 3}, want: bruchrechnen.Bruch{Zaehler: 5, Nenner: 6}},
		{name: "gleicher_nenner", bruch: bruchrechnen.Bruch{Zaehler: 2, Nenner: 5}, summand: bruchrechnen.Bruch{Zaehler: 1, Nenner: 5}, want: bruchrechnen.Bruch{Zaehler: 3, Nenner: 5}},
		{name: "negative_werte", bruch: bruchrechnen.Bruch{Zaehler: -1, Nenner: 3}, summand: bruchrechnen.Bruch{Zaehler: 1, Nenner: 6}, want: bruchrechnen.Bruch{Zaehler: -1, Nenner: 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bruch.Addiere(tt.summand)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_Subtrahiere(t *testing.T) {
	tests := []struct {
		name    string
		bruch   bruchrechnen.Bruch
		summand bruchrechnen.Bruch
		want    bruchrechnen.Bruch
	}{
		{name: "normal", bruch: bruchrechnen.Bruch{Zaehler: 1, Nenner: 2}, summand: bruchrechnen.Bruch{Zaehler: 1, Nenner: 3}, want: bruchrechnen.Bruch{Zaehler: 1, Nenner: 6}},
		{name: "negative_ergebnis", bruch: bruchrechnen.Bruch{Zaehler: 1, Nenner: 3}, summand: bruchrechnen.Bruch{Zaehler: 1, Nenner: 2}, want: bruchrechnen.Bruch{Zaehler: -1, Nenner: 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bruch.Subtrahiere(tt.summand)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_Kehrwert(t *testing.T) {
	tests := []struct {
		name    string
		bruch   bruchrechnen.Bruch
		want    bruchrechnen.Bruch
		wantErr bool
	}{
		{name: "normal", bruch: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, want: bruchrechnen.Bruch{Zaehler: 4, Nenner: 3}},
		{name: "null_zaehler", bruch: bruchrechnen.Bruch{Zaehler: 0, Nenner: 5}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bruch.Kehrwert()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_Multipliziere(t *testing.T) {
	tests := []struct {
		name   string
		bruch  bruchrechnen.Bruch
		faktor bruchrechnen.Bruch
		want   bruchrechnen.Bruch
	}{
		{name: "normal", bruch: bruchrechnen.Bruch{Zaehler: 2, Nenner: 3}, faktor: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, want: bruchrechnen.Bruch{Zaehler: 6, Nenner: 12}},
		{name: "negative", bruch: bruchrechnen.Bruch{Zaehler: -2, Nenner: 3}, faktor: bruchrechnen.Bruch{Zaehler: 4, Nenner: -5}, want: bruchrechnen.Bruch{Zaehler: -8, Nenner: -15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bruch.Multipliziere(tt.faktor)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBruch_Dividiere(t *testing.T) {
	tests := []struct {
		name    string
		bruch   bruchrechnen.Bruch
		faktor  bruchrechnen.Bruch
		want    bruchrechnen.Bruch
		wantErr bool
	}{
		{name: "normal", bruch: bruchrechnen.Bruch{Zaehler: 1, Nenner: 2}, faktor: bruchrechnen.Bruch{Zaehler: 3, Nenner: 4}, want: bruchrechnen.Bruch{Zaehler: 4, Nenner: 6}},
		{name: "divisor_null", bruch: bruchrechnen.Bruch{Zaehler: 1, Nenner: 2}, faktor: bruchrechnen.Bruch{Zaehler: 0, Nenner: 4}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bruch.Dividiere(tt.faktor)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
