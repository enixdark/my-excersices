package raindrops

import "strconv"

func Convert(num int) string {
	result := ""
    for i := 1; i <= num; i++ {
		if i == 3 && num % i == 0 {
			result += "Pling"
		} else if i == 5 && num % i == 0 {
			result += "Plang"
		} else if i == 7 && num % i == 0 {
			result += "Plong"
		}
	}
	
	if result == "" {
		return strconv.Itoa(num)
	}
	return result
}