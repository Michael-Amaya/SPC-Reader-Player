package utils

const MINIMUM_PRINTABLE_CHAR = 0x20
const MAXIMUM_PRINTABLE_CHAR = 0x7E
const VERIFICATION = 0x1A

func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
