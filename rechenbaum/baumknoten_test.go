package rechenbaum_test

import (
	"govet/rechenbaum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorknoten_Berechne(t *testing.T) {
	tests := []struct {
		name   string
		op     rechenbaum.Operator
		links  float64
		rechts float64
		want   float64
	}{
		{
			name:   "Addition",
			op:     rechenbaum.Add,
			links:  2,
			rechts: 3,
			want:   5,
		},
		{
			name:   "Subtraktion",
			op:     rechenbaum.Sub,
			links:  9,
			rechts: 4,
			want:   5,
		},
		{
			name:   "Multiplikation",
			op:     rechenbaum.Mult,
			links:  6,
			rechts: 7,
			want:   42,
		},
		{
			name:   "Division",
			op:     rechenbaum.Div,
			links:  21,
			rechts: 3,
			want:   7,
		},
		{
			name:   "Negative Zahlen",
			op:     rechenbaum.Add,
			links:  -2.5,
			rechts: 4.5,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			links := rechenbaum.Wertknoten{Zahl: tt.links}
			rechts := rechenbaum.Wertknoten{Zahl: tt.rechts}
			knoten := rechenbaum.NewOperatorknoten(tt.op, links, rechts)

			got := knoten.Berechne()
			assert.Equal(t, tt.want, got)
		})
	}

	ll := rechenbaum.Wertknoten{1.0}
	lr := rechenbaum.Wertknoten{2.0}
	lo := rechenbaum.NewOperatorknoten(rechenbaum.Add, ll, lr)
	rr := rechenbaum.Wertknoten{3.0}
	root := rechenbaum.NewOperatorknoten(rechenbaum.Mult, lo, rr)
	assert.Equal(t, 9.0, root.Berechne())

	ro := rechenbaum.NewOperatorknoten(rechenbaum.Mult, lr, rr)
	root = rechenbaum.NewOperatorknoten(rechenbaum.Add, ll, ro)
	assert.Equal(t, 7.0, root.Berechne())
}
