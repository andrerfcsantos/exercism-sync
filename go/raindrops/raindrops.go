package raindrops

import "strconv"

type SoundFactor struct {
	Sound  string
	Factor int
}

var soundFactors = []SoundFactor{
	{Sound: "Pling", Factor: 3},
	{Sound: "Plang", Factor: 5},
	{Sound: "Plong", Factor: 7},
}

// Convert converts a number into a string that contains raindrop sounds
func Convert(input int) string {
	var res string
    
	for _, soundFactor := range soundFactors {
		if input%soundFactor.Factor == 0 {
			res += soundFactor.Sound
		}
	}

	if res == "" {
		res = strconv.Itoa(input)
	}
	return res
}
