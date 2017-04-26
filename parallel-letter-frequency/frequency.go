// Package letter has functions for counting frecuency of letters in text.
package letter

const testVersion = 1

type FreqMap map[rune]int

// Frequency counts frequency of letters in string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts frequency of letters in many string using concurrency.
func ConcurrentFrequency(ss []string) FreqMap {
	messages := make(chan FreqMap, len(ss))
	for _, str := range ss {
		go func(s string) {
			messages <- Frequency(s)
		}(str)
	}

	m := FreqMap{}
	for range ss {
		msg := <-messages
		for k, v := range msg {
			m[k] += v
		}
	}

	return m
}
