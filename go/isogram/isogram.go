package isogram

import (
	"strings"
	"strconv"
	"regexp"
)

func IsIsogram(word string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	l_word := strings.ToLower(word)
	var temp []rune

	for _, w := range reg.ReplaceAllString(l_word, "") {
		for _, t := range temp {
			if strconv.QuoteRune(t) == strconv.QuoteRune(w) {
				return false
			}
		}
		temp = append(temp, w)
	}

	return true

}