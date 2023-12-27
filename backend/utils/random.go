package utils

import "math/rand"

var alphabets string = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	bits := []rune{}
	k := len(alphabets)

	for i := 0; i < n; i++ {
		index := rand.Intn(k)
		bits = append(bits, rune(alphabets[index]))
	}

	return string(bits)
}

func RandomEmail() string {
	// return RandomString(8) + "@" + RandomString(5) + ".com"
	return RandomString(8) + "@example.com"
}
