package schleifen

func Produkt(von int, bis int) int {
	produkt := 1
	for faktor := von; faktor <= bis; faktor++ {
		produkt *= faktor
	}
	return produkt
}

func ProduktRekursiv(von int, bis int) int {
	if von > bis {
		return 1
	}
	return von * ProduktRekursiv(von+1, bis)
}

func Fakultaet(n int) int {
	return Produkt(1, n)
}

func FakultaetRekursiv(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FakultaetRekursiv(n-1)
}

func SechsAusNeunundvierzig() int {
	return Produkt(44, 49) / Fakultaet(6)
}
