package raindrops

import "fmt"

func Convert(inp int) string {
	res := ""

	if inp%3 == 0 {
		res += "Pling"
	}

	if inp%5 == 0 {
		res += "Plang"
	}

	if inp%7 == 0 {
		res += "Plong"
	}

	if res == "" {
		res = fmt.Sprintf("%d", inp)
	}

	return res
}
