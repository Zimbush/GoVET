package funktionen_test

import (
	"govet/funktionen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHalbiere(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		zahl int
		want float64
	}{
		{name: "null", zahl: 0, want: 0.0},
		{name: "positiv_ganzzahlig", zahl: 8, want: 4.0},
		{name: "positiv_ungerade", zahl: 5, want: 2.5},
		{name: "negativ", zahl: -6, want: -3.0},
		{name: "negativ_ungerade", zahl: -7, want: -3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funktionen.Halbiere(tt.zahl)
			assert.Equal(t, got, tt.want, tt.name)
		})
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name    string
		celsius float64
		want    float64
	}{
		{name: "nullpunkt", celsius: 0, want: 32.0},
		{name: "ein_grad", celsius: 1, want: 33.8},
		{name: "wasser_siedepunkt", celsius: 100, want: 212.0},
		{name: "minus_10", celsius: -10, want: 14.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funktionen.CelsiusToFahrenheit(tt.celsius)
			assert.InDelta(t, tt.want, got, 0.0001)
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name       string
		fahrenheit float64
		want       float64
	}{
		{name: "frostpunkt", fahrenheit: 32, want: 0.0},
		{name: "dreiunddreissig_komma_acht", fahrenheit: 33.8, want: 1.0},
		{name: "wasser_siedepunkt", fahrenheit: 212, want: 100.0},
		{name: "minus_4", fahrenheit: 14, want: -10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funktionen.FahrenheitToCelsius(tt.fahrenheit)
			assert.InDelta(t, tt.want, got, 0.0001)
		})
	}
}

func TestNormalisiereWinkel(t *testing.T) {
	tests := []struct {
		name   string
		winkel int
		want   int
	}{
		{name: "null", winkel: 0, want: 0},
		{name: "gleich_360", winkel: 360, want: 0},
		{name: "positiv_innerhalb", winkel: 135, want: 135},
		{name: "positiv_ueberschritten", winkel: 450, want: 90},
		{name: "negativ", winkel: -45, want: 315},
		{name: "negativ_stark", winkel: -480, want: 240},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funktionen.NormalisiereWinkel(tt.winkel)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIstWinkelNormalisiert(t *testing.T) {
	tests := []struct {
		name   string
		winkel int
		want   bool
	}{
		{name: "normalisiert_min", winkel: 0, want: true},
		{name: "normalisiert_mitte", winkel: 180, want: true},
		{name: "normalisiert_max_ohne_360", winkel: 359, want: true},
		{name: "nicht_normalisiert_360", winkel: 360, want: false},
		{name: "nicht_normalisiert_negativ", winkel: -1, want: false},
		{name: "nicht_normalisiert_zu_gross", winkel: 450, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funktionen.IstWinkelNormalisiert(tt.winkel)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAddiereWinkel(t *testing.T) {
	tests := []struct {
		name string
		w1   int
		w2   int
		want int
	}{
		{name: "ohne_ueberlauf", w1: 30, w2: 45, want: 75},
		{name: "mit_ueberlauf", w1: 300, w2: 120, want: 60},
		{name: "negativ_plus_positiv", w1: -30, w2: 45, want: 15},
		{name: "beide_negativ", w1: -90, w2: -180, want: 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funktionen.AddiereWinkel(tt.w1, tt.w2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestErstesZeichenGrossRestKlein(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "leer", text: "", want: ""},
		{name: "ein_buchstabe", text: "a", want: "A"},
		{name: "gemischt", text: "hELLO", want: "Hello"},
		{name: "mit_leerzeichen", text: " welt", want: "Welt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funktionen.Grossschreibung(tt.text)
			assert.Equal(t, tt.want, got)
		})
	}
}
