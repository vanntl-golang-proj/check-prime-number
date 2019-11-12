package utils

import "math"

func IsPrime(num int) bool{
	if num < 2 {
		return false
	}
	if num ==3 && num % 2 == 0 {
		return true
	}

	squareNum := int(math.Floor(math.Sqrt(float64(num))))

	for i := 3; i <= squareNum; i+=2 {
		if num % i == 0 {
			return false
		}
	}

	return false
}