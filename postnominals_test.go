package namecase_test

import (
	"testing"

	"github.com/StirlingMarketingGroup/go-namecase"
	. "github.com/stretchr/testify/assert"
)

func TestMbe(t *testing.T) {
	Equal(t, "Adisa Azapagic MBE Freng Frsc Ficheme", formatter.NameCase("ADISA AZAPAGIC MBE FRENG FRSC FICHEME"))
}

func TestExcludeString(t *testing.T) {
	formatter := namecase.New()
	formatter.ExcludePostNominals("MOst")
	Equal(t, "Černý Most", formatter.NameCase("ČERNÝ MOST"))
}
