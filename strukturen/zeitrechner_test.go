package strukturen_test

import (
	"strukturen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStundenZuMinuten(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		stunden int
		want    int
	}{
		{"standard", 2, 120},
		{"null", 0, 0},
		{"eins", 1, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strukturen.StundenZuMinuten(tt.stunden)
			assert.Equal(t, got, tt.want, tt.name)
		})
	}
}

func TestDifferenz_alt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		von_h   int
		von_min int
		bis_h   int
		bis_min int
		want    int
		want2   int
	}{
		{"gleiche_uhrzeit", 10, 30, 10, 30, 0, 0},
		{"keine_stunde_ueberschritten", 9, 15, 10, 40, 1, 25},
		{"stunde_ueberschritten", 8, 45, 10, 15, 1, 30},
		{"nur_minuten_unterschied", 9, 50, 10, 10, 0, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := strukturen.Differenz_alt(tt.von_h, tt.von_min, tt.bis_h, tt.bis_min)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.want2, got2, tt.name)
		})
	}
}

func TestDifferenz(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		von  strukturen.Uhrzeit
		bis  strukturen.Uhrzeit
		want strukturen.Uhrzeit
	}{
		{"gleiche_uhrzeit", strukturen.Uhrzeit{10, 30}, strukturen.Uhrzeit{10, 30}, strukturen.Uhrzeit{0, 0}},
		{"keine_stunde_ueberschritten", strukturen.Uhrzeit{9, 15}, strukturen.Uhrzeit{10, 40}, strukturen.Uhrzeit{1, 25}},
		{"stunde_ueberschritten", strukturen.Uhrzeit{8, 45}, strukturen.Uhrzeit{10, 15}, strukturen.Uhrzeit{1, 30}},
		{"nur_minuten_unterschied", strukturen.Uhrzeit{9, 50}, strukturen.Uhrzeit{10, 10}, strukturen.Uhrzeit{0, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strukturen.Differenz(tt.von, tt.bis)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}
