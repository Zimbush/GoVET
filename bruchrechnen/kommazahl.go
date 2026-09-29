package bruchrechnen

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ErmittleNachkommastellen(zahl float64) int {
	if zahl == 0 {
		return 0
	}

	abs := math.Abs(zahl)
	if abs == math.Trunc(abs) {
		return 0
	}

	text := strconv.FormatFloat(abs, 'f', -1, 64)
	if idx := strings.Index(text, "."); idx >= 0 {
		frac := strings.TrimRight(text[idx+1:], "0")
		return len(frac)
	}

	return 0
}

type Kommazahl struct {
	Zahl float64
}

func (zahl Kommazahl) BerechneWert() float64 {
	return zahl.Zahl
}

func (zahl Kommazahl) BerechneBruch() Bruch {
	nks := ErmittleNachkommastellen(zahl.Zahl)
	nenner := int(math.Pow10(nks))
	return Bruch{int(zahl.Zahl * float64(nenner)), nenner}
}

func (zahl Kommazahl) BerechneKommazahl() Kommazahl {
	return zahl
}

func (zahl Kommazahl) Addiere(summand Kommazahl) Kommazahl {
	return Kommazahl{zahl.Zahl + summand.Zahl}
}

func (zahl Kommazahl) Subtrahiere(summand Kommazahl) Kommazahl {
	return Kommazahl{zahl.Zahl - summand.Zahl}
}

func (zahl Kommazahl) Multipliziere(faktor Kommazahl) Kommazahl {
	return Kommazahl{zahl.Zahl * faktor.Zahl}
}

func (zahl Kommazahl) Dividiere(faktor Kommazahl) (Kommazahl, error) {
	divisor := faktor.Zahl
	if divisor == 0 {
		return Kommazahl{}, fmt.Errorf("Division durch 0")
	}
	return Kommazahl{zahl.Zahl / divisor}, nil
}

func (zahl Kommazahl) Mult(faktor Zahl) Zahl {
	return zahl.Multipliziere(faktor.BerechneKommazahl())
}
