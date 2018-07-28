package grains

import (
	"errors"
	"math"
)

func Total() uint64 {
	var total uint64 = 0
	for i := 1; i <= 64 ; i++ {
		n, _ := Square(i)
		total += n
	}
	return total
}

func Square(num int) (uint64, error) {
	if num <= 0 || num > 64  {
		return 0, errors.New("")
	}
	return uint64(math.Exp2(float64(num - 1))), nil
}

