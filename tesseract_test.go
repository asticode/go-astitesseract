package astitesseract_test

import (
	"testing"

	"github.com/asticode/go-astitesseract"
	"github.com/stretchr/testify/assert"
)

func TestTesseract_GetUTF8Text(t *testing.T) {
	var tst, err = astitesseract.New(astitesseract.Options{Languages: []string{"eng"}})
	assert.NoError(t, err)
	defer tst.Close()
	assert.Equal(t, `The (quick) [brown] {fox} jumps!
Over the $43,456.78 <lazy> #90 dog
& duck/goose, as 12.5% of E-mail
from aspammer@website.com is spam.
Der ,,schnelle” braune Fuchs springt
ﬁber den faulen Hund. Le renard brun
«rapide» saute par-dessus le chien
paresseux. La volpe marrone rapida
salta sopra i] cane pigro. El zorro
marrén répido salta sobre el perro
perezoso. A raposa marrom répida
salta sobre 0 C50 preguieoso.

`, tst.GetUTF8Text("./testdata/test.png"))
}
