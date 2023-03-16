package str

import (
	"fmt"
)

func CommonChars(p, q string) string {
	fmt.Printf("What characters do strings `%s` and `%s` share?\n", p, q)
	aa := make(map[byte]bool)
	for i := 0; i < len(p); i += 1 {
		aa[p[i]] = true
	}
	res := []byte{}
	for i := 0; i < len(q); i += 1 {
		_, ok := aa[q[i]]
		if ok {
			res = append(res, q[i])
		}
	}
	return string(res[:])
}
