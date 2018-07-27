package luhn

import (
	"strings"
	// "errors"
	"strconv"
)

func Valid(num string) bool {
	_num := strings.Trim(num, "")
	if _, err := strconv.Atoi(_num); err != nil {
		return false
	}

	if len(_num) == 3 || len(_num) == 2 {
		return true
	}
	// var nums []int
	total := 0
	for i, n := range _num {
		var temp int
		if(i % 2 == 0){
			temp = int(n) * 2 
			if temp > 9 {
				temp = temp - 9
			}
		} else {
			temp = int(n)
		}
		
		total += temp
        // nums = append(nums, temp)
	}
	
	return total % 10 == 0

}