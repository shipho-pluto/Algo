package Language

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnglish(t *testing.T) {
	lang := "English"
	exp := "Hello"
	actual, err := Grating(lang)

	assert.Nil(t, err)
	assert.Equal(t, exp, actual)
}

func TestSpanish(t *testing.T) {
	lang := "Spanish"
	exp := "Hola"
	actual, err := Grating(lang)

	assert.Nil(t, err)
	assert.Equal(t, exp, actual)
}

func TestGerman(t *testing.T) {
	lang := "German"
	exp := "Hallo"
	actual, err := Grating(lang)

	assert.Nil(t, err)
	assert.Equal(t, exp, actual)
}

func TestFrench(t *testing.T) {
	lang := "French"
	exp := "Sova"
	actual, err := Grating(lang)

	assert.Nil(t, err)
	assert.Equal(t, exp, actual)
}

func TestErr(t *testing.T) {
	lang := "Russian"
	exp := ""
	errExpected := errors.New("unsupported language")
	actual, errActual := Grating(lang)

	assert.Equal(t, errExpected, errActual)
	assert.Equal(t, exp, actual)
}
