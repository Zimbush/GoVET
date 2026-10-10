package baum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnoten_Einfuegen(t *testing.T) {
	root := NewKnoten(10)
	root.Einfuegen(5)
	root.Einfuegen(15)
	root.Einfuegen(12)

	assert.Equal(t, 10, root.Wert)
	assert.NotNil(t, root.Links)
	assert.NotNil(t, root.Rechts)
	assert.Equal(t, 5, root.Links.Wert)
	assert.Equal(t, 15, root.Rechts.Wert)
	assert.Equal(t, 12, root.Rechts.Links.Wert)
	assert.True(t, root.IstBlatt() == false)
}

func TestBaum_Einfuegen(t *testing.T) {
	baum := NewBaum()
	baum.Einfuegen(20)
	baum.Einfuegen(10)
	baum.Einfuegen(30)
	baum.Einfuegen(25)

	assert.NotNil(t, baum.Wurzel)
	assert.Equal(t, 20, baum.Wurzel.Wert)
	assert.Equal(t, 10, baum.Wurzel.Links.Wert)
	assert.Equal(t, 30, baum.Wurzel.Rechts.Wert)
	assert.Equal(t, 25, baum.Wurzel.Rechts.Links.Wert)
}

func TestKnoten_Einfuegen_DoppelteWerte(t *testing.T) {
	root := NewKnoten(10)
	root.Einfuegen(10)
	root.Einfuegen(10)
	root.Einfuegen(5)

	assert.Equal(t, 10, root.Wert)
	assert.NotNil(t, root.Links)
	assert.Equal(t, 10, root.Links.Wert)
	assert.NotNil(t, root.Links.Links)
	assert.Equal(t, 10, root.Links.Links.Wert)
	assert.NotNil(t, root.Links.Links.Links)
	assert.Equal(t, 5, root.Links.Links.Links.Wert)
	assert.Nil(t, root.Rechts)
}

func TestKnoten_Einfuegen_AufsteigendSortiert(t *testing.T) {
	root := NewKnoten(5)
	root.Einfuegen(7)
	root.Einfuegen(9)
	root.Einfuegen(11)

	assert.Equal(t, 5, root.Wert)
	assert.Equal(t, 7, root.Rechts.Wert)
	assert.Equal(t, 9, root.Rechts.Rechts.Wert)
	assert.Equal(t, 11, root.Rechts.Rechts.Rechts.Wert)
	assert.Nil(t, root.Links)
}

func TestBaum_AuslesenSortiert(t *testing.T) {
	baum := NewBaum()
	baum.Einfuegen(10)
	baum.Einfuegen(5)
	baum.Einfuegen(15)
	baum.Einfuegen(12)

	assert.Equal(t, []int{5, 10, 12, 15}, baum.AuslesenSortiert())
}

func TestBaum_Ausbalancieren(t *testing.T) {
	baum := NewBaum()
	for _, wert := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		baum.Einfuegen(wert)
	}

	baum.Ausbalancieren()

	assert.NotNil(t, baum.Wurzel)
	assert.Equal(t, 5, baum.Wurzel.Wert)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, baum.AuslesenSortiert())
}

func TestKnoten_IstBlatt(t *testing.T) {
	blatt := NewKnoten(7)
	assert.True(t, blatt.IstBlatt())

	blatt.Einfuegen(5)
	assert.False(t, blatt.IstBlatt())
}
