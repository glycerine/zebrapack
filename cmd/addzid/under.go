package main

import "unicode"

func underToCamelCase(s string) string {
	ru := []rune(s)
	n := len(ru)
	last := n - 1
	for i := 0; i < n; i++ {
		if ru[i] == '_' && i < last {
			if unicode.IsLower(ru[i+1]) {
				ru[i+1] = unicode.ToUpper(ru[i+1])
			}
			copy(ru[i:], ru[i+1:])
			ru = ru[:last]

			last--
			n--
			i--
		}
	}
	return string(ru)
}
