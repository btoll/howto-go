// https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/880/
package ints

import (
	"fmt"
	"math"
)

func determinePowerOf10(n int) int {
	i := 1
	p := 0
	for i <= n {
		i *= 10
		p += 1
	}
	return p
}

func Reverse(n int) int {
	// if given 123, return 321
	fmt.Printf("Reverse the integer `%d`.\n", n)
	isNegative := false

	if n < 0 {
		isNegative = true
		n = int(math.Abs(float64(n)))
	}

	var reversed float64 = 0
	q := float64(determinePowerOf10(n)) - 1
	for n > 0 {
		reversed += float64(n%10) * math.Pow(10, q)
		n /= 10
		q -= 1
	}

	if isNegative {
		return int(reversed) * -1
	}

	if reversed > math.Pow(2, 31)-1 || reversed < -math.Pow(2, 31) {
		fmt.Println("got here")
	}
	return int(reversed)
}
