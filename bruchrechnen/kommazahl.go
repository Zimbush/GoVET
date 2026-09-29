package bruchrechnen

import (
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

type Kommmazahl struct {
	Zahl float64
}

func (zahl Kommmazahl) BerechneWert() float64 {
	return zahl.Zahl
}

func (zahl Kommmazahl) BerechneBruch() Bruch {
	nks := ErmittleNachkommastellen(zahl.Zahl)
	nenner := int(math.Pow10(nks))
	return Bruch{int(zahl.Zahl * float64(nenner)), nenner}
}
