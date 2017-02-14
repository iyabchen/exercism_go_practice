package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// map does not allowed concurent write, even if it is
// monotonously increasing
func ConcurrentFrequency(s []string) FreqMap {
	c := make(chan FreqMap, 3)
	result := FreqMap{}

	for _, v := range s {
		go func(str string) {
			c <- Frequency(str)
		}(v)
	}
	var f FreqMap
	for i := 0; i < 3; i++ {
		f = <-c
		for key, value := range f {
			result[key] += value

		}
	}
	return result
}
