package bruchrechnen

import "fmt"

func GGT(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type Zahl interface {
	BerechneWert() float64
}

type Bruch struct {
	Zaehler int
	Nenner  int
}

func NewBruch(zaehler int, nenner int) (Bruch, error) {
	if nenner == 0 {
		return Bruch{1, 1}, fmt.Errorf("Nenner gleich 0")
	}
	return Bruch{zaehler, nenner}, nil
}

func (bruch Bruch) BerechneWert() float64 {
	return float64(bruch.Zaehler) / float64(bruch.Nenner)
}

func (bruch Bruch) Kuerze() Bruch {
	ggt := GGT(bruch.Zaehler, bruch.Nenner)
	return Bruch{bruch.Zaehler / ggt, bruch.Nenner / ggt}
}
