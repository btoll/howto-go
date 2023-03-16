package str

import (
	"fmt"
)

func getMultiplier(n int) int {
	i, j := 1, 1
	for i <= n {
		j *= 10
		i += 1
	}
	return j
}

func Atoi(s string) int {
	fmt.Printf("If possible, turn `%s` into an integer.\n", s)
	//ascii
	//48 0
	//49 1
	//50 2
	//51 3
	//52 4
	//53 5
	//54 6
	//55 7
	//56 8
	//57 9
	allowZero := false
	res := []rune{}
	for _, char := range s {
		fmt.Println("char ", char)
		if char == 9 || char == 32 {
			fmt.Println("got here")
			continue
		}
		if 48 <= char && char <= 57 {
			if char != 48 {
				allowZero = true
			}
			if char == 48 && !allowZero {
				continue
			}
			res = append(res, char)
		} else {
			break
		}
	}
	fmt.Println("res", res)
	d := 0
	for i, j := 0, len(res)-1; i < len(res); i, j = i+1, j-1 {
		d += int(res[i]-'0') * getMultiplier(j)
	}
	return d
}
