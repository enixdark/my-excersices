package reverse

import (
	"strings"
)

func String(word string) string {
	var reverse []string

	for _, w := range word {
		reverse = append([]string{string(w)}, reverse...)
	}

	return strings.Join(reverse, "")
}