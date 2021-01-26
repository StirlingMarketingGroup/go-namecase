package namecase_test

import (
	"strings"
	"testing"

	"github.com/StirlingMarketingGroup/go-namecase"
	. "github.com/stretchr/testify/assert"
)

func TestLazy(t *testing.T) {
	formatter := namecase.New()
	Equal(t, "Dougal MACDonald", formatter.NameCase("Dougal MACDonald"))

	formatter.SetLazy(false)
	Equal(t, "Dougal MacDonald", formatter.NameCase("Dougal MACDonald"))

	formatter.SetLazy(true)
	Equal(t, "Dougal MACDonald", formatter.NameCase("Dougal MACDonald"))
}

func TestIrish(t *testing.T) {
	formatter := namecase.New()
	Equal(t, "Macmurdo", formatter.NameCase("Macmurdo"))

	formatter.SetIrish(false)
	Equal(t, "Macmurdo", formatter.NameCase("Macmurdo"))

	formatter.SetIrish(true)
	Equal(t, "Macmurdo", formatter.NameCase("Macmurdo"))
}

func TestSpanish(t *testing.T) {
	names := []string{"Ruiz y Picasso", "Dato e Iradier", "Mas i Gavarró"}

	formatter := namecase.New()
	formatter.SetSpanish(true)
	Equal(t, "La Luna", formatter.NameCase(strings.ToLower("La Luna")))
	Equal(t, "El Paso", formatter.NameCase(strings.ToLower("El Paso")))
	for _, name := range names {
		Equal(t, name, formatter.NameCase(strings.ToLower(name)))
	}

	formatter.SetSpanish(false)
	Equal(t, "la Luna", formatter.NameCase(strings.ToLower("La Luna")))
	Equal(t, "el Paso", formatter.NameCase(strings.ToLower("El Paso")))
	for _, name := range names {
		NotEqual(t, name, formatter.NameCase(strings.ToLower(name)))
	}
}

func TestRoman(t *testing.T) {
	formatter := namecase.New()

	formatter.SetRoman(false)
	Equal(t, "Na Li", formatter.NameCase(strings.ToLower("Na Li")))

	formatter.SetRoman(true)
	Equal(t, "Na LI", formatter.NameCase(strings.ToLower("Na Li")))
}

func TestHebrew(t *testing.T) {
	formatter := namecase.New()

	formatter.SetHebrew(false)
	Equal(t, "Aharon Ben Amram Ha-Kohein", formatter.NameCase(strings.ToLower("Aharon BEN Amram Ha-Kohein")))
	Equal(t, "Ben Gurion", formatter.NameCase(strings.ToLower("ben Gurion")))

	formatter.SetHebrew(true)
	Equal(t, "Aharon ben Amram Ha-Kohein", formatter.NameCase(strings.ToLower("Aharon BEN Amram Ha-Kohein")))
	Equal(t, "ben Gurion", formatter.NameCase(strings.ToLower("ben Gurion")))
}

func TestPostNominal(t *testing.T) {
	formatter := namecase.New()

	formatter.SetPostNominal(false)
	Equal(t, "Tam Phd", formatter.NameCase(strings.ToLower("Tam PHD")))

	formatter.SetPostNominal(true)
	Equal(t, "Tam PhD", formatter.NameCase(strings.ToLower("Tam PHD")))
}
