package luhn

import (
	"strings"
	// "errors"
	"strconv"
	"regexp"
)

func Valid(num string) bool {
	sum := 0
    re := regexp.MustCompile(" ")
	list_num := re.Split(num, -1)
	_num := strings.Join(list_num, "")

	for _, n := range list_num {
		if _, err := strconv.Atoi(n); err != nil {
			return false
		}
	}

	digits := len(_num)
	parity := digits % 2
    
	for i := 0; i < digits; i++ {
		var digit = int(_num[i] - 48)
		if i % 2 == parity {
			digit *= 2
			if digit > 9 {
			    digit -= 9
		    }
		}

		sum += digit
	}

	return len(num) > 1 && sum % 10 == 0
}

