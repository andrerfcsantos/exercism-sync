// Package raindrops is a small library for converting a number into pitterpatter.
package raindrops
import (
	"strconv"
)
// BenchmarkConvert-12    	 3233023	       389.5 ns/op	      64 B/op	       4 allocs/op
func dripping(wet *uint32, sound string, num int, divisor int) (drop string) {
	if num%divisor == 0 {
		*wet |= 0b1
		drop = sound
	}
	return drop
}
// Convert takes a number and converts it into pitterpatter.
func Convert(num int) string {
	var wet uint32 = 0
	pling, plang, plong := dripping(&wet, "Pling", num, 3), dripping(&wet, "Plang", num, 5), dripping(&wet, "Plong", num, 7)
	if wet != 0 {
		return pling + plang + plong
	}
	return strconv.Itoa(num)
}
// // Package raindrops is a small library for converting a number into pitterpatter.
// package raindrops
// import "strconv"
// type raindrop struct {
// 	divisor int
// 	drop    string
// }
// var drops = []raindrop{
// 	{3, "Pling"},
// 	{5, "Plang"},
// 	{7, "Plong"},
// }
// // Convert takes a number and converts it into pitterpatter.
// func Convert(num int) (pitterpatter string) {
// 	// I'll go with concatenation
// 	// https://dev.to/pmalhaire/concatenate-strings-in-golang-a-quick-benchmark-4ahh
// 	// BenchmarkConvert-12    	 1317374	       956.5 ns/op	      80 B/op	       5 allocs/op
// 	for _, drip := range drops {
// 		if num%drip.divisor == 0 {
// 			pitterpatter += drip.drop
// 		}
// 	}
// 	if pitterpatter == "" {
// 		pitterpatter = strconv.Itoa(num)
// 	}
// 	return pitterpatter
// }
