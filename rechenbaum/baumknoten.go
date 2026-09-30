package rechenbaum

type Baumknoten interface {
	Berechne() float64
}

type Wertknoten struct {
	Zahl float64
}

func (knoten Wertknoten) Berechne() float64 {
	return knoten.Zahl
}

type Operatorknoten struct {
	Op     Operator
	Links  Baumknoten
	Rechts Baumknoten
}

func NewSingleOperatorknoten(op Operator) Operatorknoten {
	links := Wertknoten{0.0}
	rechts := Wertknoten{1.0}
	knoten := Operatorknoten{op, links, rechts}
	return knoten
}

func NewOperatorknoten(op Operator, links Baumknoten, rechts Baumknoten) Operatorknoten {
	return Operatorknoten{op, links, rechts}
}

func (knoten *Operatorknoten) SetzeLinks(kind Baumknoten) {
	knoten.Links = kind
}

func (knoten *Operatorknoten) SetzeRechts(kind Baumknoten) {
	knoten.Rechts = kind
}

func (knoten Operatorknoten) Berechne() float64 {
	a := knoten.Links.Berechne()
	b := knoten.Rechts.Berechne()
	return knoten.Op.Value.Operator(a, b)
}
