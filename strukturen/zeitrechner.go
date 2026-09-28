package strukturen

func StundenZuMinuten(stunden int) int {
	return stunden * 60
}

func Gesamtminuten_alt(stunden int, minuten int) int {
	return StundenZuMinuten(stunden) + minuten
}

func MinutenZuStundenMinuten(minuten int) (int, int) {
	return minuten / 60, minuten % 60
}

func Differenz_alt(von_h int, von_min int, bis_h int, bis_min int) (int, int) {
	minuten_von := Gesamtminuten_alt(von_h, von_min)
	minuten_bis := Gesamtminuten_alt(bis_h, bis_min)
	differenz := minuten_bis - minuten_von
	return MinutenZuStundenMinuten(differenz)
}

type Uhrzeit struct {
	Stunden int
	Minuten int
}

func Gesamtminuten(zeit Uhrzeit) int {
	return zeit.Stunden*60 + zeit.Minuten
}

func MinutenZuUhrzeit(minuten int) Uhrzeit {
	return Uhrzeit{minuten / 60, minuten % 60}
}

func Differenz(von Uhrzeit, bis Uhrzeit) Uhrzeit {
	minuten_von := Gesamtminuten(von)
	minuten_bis := Gesamtminuten(bis)
	differenz := minuten_bis - minuten_von
	return MinutenZuUhrzeit(differenz)
}
