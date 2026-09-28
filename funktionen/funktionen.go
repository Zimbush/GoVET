package funktionen

import "strings"

func Halbiere(zahl int) float64 {
	return float64(zahl) / 2.0
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func NormalisiereWinkel(winkel int) int {
	winkel = winkel % 360
	if winkel < 0 {
		winkel += 360
	}
	return winkel
}

func IstWinkelNormalisiert(winkel int) bool {
	return winkel >= 0 && winkel < 360
}

func AddiereWinkel(winkel1 int, winkel2 int) int {
	return NormalisiereWinkel(winkel1 + winkel2)
}

func Grossschreibung(text string) string {
	text = strings.Trim(text, " \t\n\r")
	if text == "" {
		return ""
	}

	return strings.ToUpper(text[:1]) + strings.ToLower(text[1:])
}
