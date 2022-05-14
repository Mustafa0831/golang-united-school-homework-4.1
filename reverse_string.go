package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	reverseString := strings.Split(input, "")

	for i, j := 0, len(reverseString)-1; i < j; i, j = i+1, j-1 {
		reverseString[i], reverseString[j] = reverseString[j], reverseString[i]
	}
	return strings.Join(reverseString, "")
	// solution goes here
	return output
}
