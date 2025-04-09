package id

import "math/rand/v2"

const uppercaseRange = ('Z' - 'A')
const lowercaseRange = ('z' - 'a')

func GenerateId(length int) string {
	id := ""

	for range length {
		char := rune(rand.IntN(uppercaseRange + lowercaseRange))

		if char > uppercaseRange {
			char += 'a' - uppercaseRange
		} else {
			char += 'A'
		}

		id += string(char)
	}

	return id
}
