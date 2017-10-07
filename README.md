This is a Golang library to manipulate [tesseract-ocr](https://github.com/tesseract-ocr/tesseract).

# Installation
## Install tesseract-ocr
### Linux

Run the following command:

    $ apt-get install libtesseract-dev libleptonica-dev tesseract-ocr-eng

Then it's up to you languages you want to install.

## Install go-astitesseract

Run the following command:

    $ go get -u github.com/asticode/go-astitesseract/...

# Use the library in your code

WARNING: the code below doesn't handle errors for readibility purposes. However you SHOULD!

```go
t, _ := astitesseract.New(astitesseract.Options{Languages: []string{"eng"}})
defer t.Close()
fmt.Println(tst.GetUTF8Text("./testdata/test.png"))
```

will print

```
The (quick) [brown] {fox} jumps!
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
```