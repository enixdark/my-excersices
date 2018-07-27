// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"regexp"
	"fmt"
)
// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	r, _ := regexp.Compile("(what|how|when)")
	IsLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	q := "Whatever."
    
	if strings.TrimSpace(remark) == "" {
	    q = "Fine. Be that way!"
	} else if strings.Contains(remark, "?") {
		q = "Sure."
		if r.MatchString(strings.ToLower(remark[0:5])) {
			q = "Calm down, I know what I'm doing!"
		}
	} else if len(remark) > 0 && strings.ToUpper(remark) == remark {
		fmt.Print(remark)
		if IsLetter(remark) || regexp.MustCompile(`\!$`).MatchString(remark) {
			q = "Whoa, chill out!"
		}
	}
	return q
}
