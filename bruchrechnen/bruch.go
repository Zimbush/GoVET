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

func KGV(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a * b / GGT(a, b)
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

func (bruch Bruch) Erweitere(faktor int) Bruch {
	if faktor == 0 {
		return Bruch{0, 1}
	}
	return Bruch{bruch.Zaehler * faktor, bruch.Nenner * faktor}
}

func (bruch Bruch) Inverses() Bruch {
	return Bruch{-bruch.Zaehler, bruch.Nenner}
}

func (bruch Bruch) Addiere(summand Bruch) Bruch {
	kgv := KGV(bruch.Nenner, summand.Nenner)
	return Bruch{bruch.Zaehler*kgv/bruch.Nenner + summand.Zaehler*kgv/summand.Nenner, kgv}
}

func (bruch Bruch) Subtrahiere(summand Bruch) Bruch {
	return bruch.Addiere(summand.Inverses())
}

func (bruch Bruch) Kehrwert() (Bruch, error) {
	kw, err := NewBruch(bruch.Nenner, bruch.Zaehler)
	return kw, err
}

func (bruch Bruch) Multipliziere(faktor Bruch) Bruch {
	return Bruch{bruch.Zaehler * faktor.Zaehler, bruch.Nenner * faktor.Nenner}
}

func (bruch Bruch) Dividiere(faktor Bruch) (Bruch, error) {
	kw, err := faktor.Kehrwert()
	return bruch.Multipliziere(kw), err
}
