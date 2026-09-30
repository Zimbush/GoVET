package rechenbaum

import (
	"github.com/orsinium-labs/enum"
)

type Grundrechenart struct {
	Zeichen  string
	Operator func(op1 float64, op2 float64) float64
}

type Operator enum.Member[*Grundrechenart]

func Addf(a, b float64) float64  { return a + b }
func Subf(a, b float64) float64  { return a - b }
func Multf(a, b float64) float64 { return a * b }
func Divf(a, b float64) float64  { return a / b }

var (
	Add           = Operator{&Grundrechenart{"+", Addf}}
	Sub           = Operator{&Grundrechenart{"-", Subf}}
	Mult          = Operator{&Grundrechenart{"*", Multf}}
	Div           = Operator{&Grundrechenart{"/", Divf}}
	InfixOperator = enum.New(Add, Sub, Mult, Div)
)
