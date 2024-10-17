package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

var (
	allowedPattern    = "^[a-zA-Z0-9_,\\-=? ]+$"
	sanitizerInstance = bluemonday.StrictPolicy()
)

func SetLowerAndAddSpace(str string) string {
	lower := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	lower = matchAllCap.ReplaceAllString(lower, "${1} ${2}")
	return strings.ToLower(lower)
}

func ReplaceAllChar(text, char, replacement string) string {
	return strings.ReplaceAll(text, char, replacement)
}

func Sanitize(input string) (string, error) {
	sanitizedStr := sanitizerInstance.Sanitize(input)

	// Define the regular expression pattern to match allowed characters
	re := regexp.MustCompile(allowedPattern)
	if !re.MatchString(sanitizedStr) {
		return "", fmt.Errorf("bad query for %s", input)
	}

	return sanitizedStr, nil
}

func ByteToString(b []byte) string {
	return strings.Join(strings.Fields(string(b)), " ")
}

func MatchString(pattern string, s string) bool {
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

func CleanSpecialChars(s string) string {
	return specialChar.ReplaceAllString(s, "")
}

func RandomString(seq int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, seq)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
