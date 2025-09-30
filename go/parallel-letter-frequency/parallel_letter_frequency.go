package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {

	res := make(FreqMap)
	ch := make(chan FreqMap, 10)

	for _, text := range texts {
		go func(t string) {
			ch <- Frequency(t)
		}(text)
	}

	for i := 0; i < len(texts); i++ {
        freqmap := <-ch
		for letter, freq := range freqmap {
			res[letter] += freq
		}
	}

	return res
}
