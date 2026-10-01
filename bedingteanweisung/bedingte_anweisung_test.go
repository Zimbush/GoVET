package bedingteanweisung_test

import (
	"govet/bedingteanweisung"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersandkosten(t *testing.T) {
	tests := []struct {
		name        string
		bestellwert float64
		want        float64
		wantErr     bool
	}{
		{name: "unter_grenze", bestellwert: 49.99, want: 4.90},
		{name: "genau_50", bestellwert: 50.0, want: 0.0},
		{name: "ueber_grenze", bestellwert: 75.5, want: 0.0},
		{name: "null", bestellwert: 0.0, want: 4.90},
		{name: "negativ", bestellwert: -1.0, want: 0.0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bedingteanweisung.Versandkosten(tt.bestellwert)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIstGerade(t *testing.T) {
	tests := []struct {
		name string
		zahl int
		want bool
	}{
		{name: "null", zahl: 0, want: true},
		{name: "positiv_gerade", zahl: 8, want: true},
		{name: "positiv_ungerade", zahl: 5, want: false},
		{name: "negativ_gerade", zahl: -6, want: true},
		{name: "negativ_ungerade", zahl: -7, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bedingteanweisung.IstGerade(tt.zahl)
			assert.Equal(t, tt.want, got)
		})
	}
}
