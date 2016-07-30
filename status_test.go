package status

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFound(t *testing.T) {
	// given
	httpCode := http.StatusFound

	// when
	textCode := TextCode(httpCode)

	// then
	assert.Equal(t, "FOUND", textCode)
}

func TestNotFound(t *testing.T) {
	// given
	httpCode := http.StatusNotFound

	// when
	textCode := TextCode(httpCode)

	// then
	assert.Equal(t, "NOT_FOUND", textCode)
}

func TestTeapot(t *testing.T) {
	// given
	httpCode := http.StatusTeapot

	// when
	textCode := TextCode(httpCode)

	// then
	assert.Equal(t, "I_M_A_TEAPOT", textCode)
}

func TestNonAuthoritativeInformation(t *testing.T) {
	// given
	httpCode := http.StatusNonAuthoritativeInfo

	// when
	textCode := TextCode(httpCode)

	// then
	assert.Equal(t, "NON_AUTHORITATIVE_INFORMATION", textCode)
}

func TestUnknown(t *testing.T) {
	// given
	httpCode := 399

	// when
	textCode := TextCode(httpCode)

	// then
	assert.Empty(t, textCode)
}
