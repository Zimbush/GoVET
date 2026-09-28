package bedingteanweisung

import "fmt"

func Versandkosten(bestellwert float64) (float64, error) {
	if bestellwert < 0.0 {
		return 0.0, fmt.Errorf("negativer Bestellwert: %.2f", bestellwert)
	}
	if bestellwert < 50.0 {
		return 4.90, nil
	}
	return 0.0, nil
}

func IstGerade(zahl int) bool {
	return zahl%2 == 0
}
