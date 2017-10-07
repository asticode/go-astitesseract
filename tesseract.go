package astitesseract

/*
#cgo LDFLAGS: -ltesseract -llept
#include "tesseract.h"
*/
import "C"
import (
	"strings"

	"github.com/pkg/errors"
)

// Tesseract represents a Tesseract client
type Tesseract struct {
	w *C.TesseractWrapper
}

// Options represents options
type Options struct {
	DataPath  string
	Languages []string
}

// New creates a new Tesseract based on options
func New(o Options) (t *Tesseract, err error) {
	t = &Tesseract{
		w: C.New(),
	}
	if C.Init(t.w, C.CString(o.DataPath), C.CString(strings.Join(o.Languages, "+"))) > 0 {
		err = errors.New("astitesseract: C.Init failed")
		return
	}
	return
}

// Close closes Tesseract properly
func (t *Tesseract) Close() error {
	C.End(t.w)
	return nil
}

// GetUTF8Text retrieves UTF8 text from an image path
func (t *Tesseract) GetUTF8Text(path string) string {
	return C.GoString(C.GetUTF8Text(t.w, C.CString(path)))
}
