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
