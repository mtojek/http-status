package status

import (
	"net/http"
	"regexp"
	"strings"
)

const (
	httpStatusCodeRangeStart = 100
	httpStatusCodeRangeEnd   = 600
	replacingCharacter       = "_"
)

var nonAlphaCharacters = regexp.MustCompile("[^A-Za-z]")

var textCodes = createTextCodes()

// TextCode method returns an API friendly HTTP status text code.
func TextCode(code int) string {
	return textCodes[code]
}

func createTextCodes() map[int]string {
	codes := map[int]string{}
	for i := httpStatusCodeRangeStart; i < httpStatusCodeRangeEnd; i++ {
		text := http.StatusText(i)
		if text != "" {
			codes[i] = transformToTextCode(text)
		}
	}
	return codes
}

func transformToTextCode(text string) string {
	return strings.ToUpper(nonAlphaCharacters.ReplaceAllString(text, replacingCharacter))
}
