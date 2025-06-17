package rle

import "fmt"

// Encode returns the run‑length encoding of UTF‑8 string s.
//
// "AAB" → "A2B1"
func Encode(s string) string {
	// TODO: implement

	collector := make(map[string]int)
	order := make([]string, 0)
	for _, c := range s {
		char := string(c)
		if val, ok := collector[char]; ok {
			collector[char] = val + 1
			continue
		}
		collector[char] = 1
		order = append(order, char)
	}

	ans := ""
	for _, w := range order {
		ans = fmt.Sprintf("%s%s%d", ans, w, collector[w])
	}

	return ans
}
