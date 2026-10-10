package baum

// Knoten beschreibt einen einzelnen Knoten in einem binären Suchbaum.
type Knoten struct {
	Wert   int
	Links  *Knoten
	Rechts *Knoten
}

// Baum stellt einen binären Suchbaum dar.
type Baum struct {
	Wurzel *Knoten
}

func NewKnoten(wert int) *Knoten {
	k := &Knoten{Wert: wert}
	k.Links = nil
	k.Rechts = nil
	return k
}

func NewBaum() *Baum {
	return &Baum{}
}

func NewBaumMitWurzel(wert int) *Baum {
	k := NewKnoten(wert)
	return &Baum{Wurzel: k}
}

func (k *Knoten) SetzeLinks(neu *Knoten) {
	if k == nil {
		return
	}
	k.Links = neu
}

func (k *Knoten) SetzeRechts(neu *Knoten) {
	if k == nil {
		return
	}
	k.Rechts = neu
}

func (k Knoten) IstBlatt() bool {
	return k.Links == nil && k.Rechts == nil
}

func (k *Knoten) Einfuegen(wert int) {
	if k == nil {
		return
	}

	if wert <= k.Wert {
		if k.Links == nil {
			k.SetzeLinks(NewKnoten(wert))
			return
		}
		k.Links.Einfuegen(wert)
		return
	}

	if k.Rechts == nil {
		k.SetzeRechts(NewKnoten(wert))
		return
	}
	k.Rechts.Einfuegen(wert)
}

func (k *Knoten) Suchen(wert int) bool {
	if k == nil {
		return false
	}
	if wert == k.Wert {
		return true
	}
	if wert < k.Wert {
		if k.Links == nil {
			return false
		}
		return k.Links.Suchen(wert)
	}
	if k.Rechts == nil {
		return false
	}
	return k.Rechts.Suchen(wert)
}

func (b *Baum) Einfuegen(wert int) {
	if b == nil {
		return
	}
	if b.Wurzel == nil {
		b.Wurzel = NewKnoten(wert)
		return
	}
	b.Wurzel.Einfuegen(wert)
}

func (b *Baum) Suchen(wert int) bool {
	if b == nil || b.Wurzel == nil {
		return false
	}
	return b.Wurzel.Suchen(wert)
}

func (b *Baum) AuslesenSortiert() []int {
	if b == nil || b.Wurzel == nil {
		return nil
	}

	werte := []int{}
	var inorder func(*Knoten)
	inorder = func(k *Knoten) {
		if k == nil {
			return
		}
		inorder(k.Links)
		werte = append(werte, k.Wert)
		inorder(k.Rechts)
	}
	inorder(b.Wurzel)
	return werte
}

func baumAusSortierterListe(werte []int, links int, rechts int) *Knoten {
	if links > rechts {
		return nil
	}

	mitte := links + (rechts-links)/2
	knoten := NewKnoten(werte[mitte])
	knoten.Links = baumAusSortierterListe(werte, links, mitte-1)
	knoten.Rechts = baumAusSortierterListe(werte, mitte+1, rechts)
	return knoten
}

func (b *Baum) EinlesenSortierteListe(werte []int) {
	if b == nil {
		return
	}
	if len(werte) == 0 {
		b.Wurzel = nil
		return
	}
	b.Wurzel = baumAusSortierterListe(werte, 0, len(werte)-1)
}

func (b *Baum) Ausbalancieren() {
	if b == nil {
		return
	}
	b.EinlesenSortierteListe(b.AuslesenSortiert())
}

func (b *Baum) IstLeer() bool {
	return b == nil || b.Wurzel == nil
}
