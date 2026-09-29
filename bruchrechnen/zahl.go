package bruchrechnen

type Zahl interface {
	BerechneWert() float64
	BerechneBruch() Bruch
	BerechneKommazahl() Kommazahl
	Mult(zahl Zahl) Zahl
}

func Mult(op1 Zahl, op2 Zahl) Zahl {
	return op1.Mult(op2)
}
