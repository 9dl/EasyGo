package easygo

import (
	"strings"
)

func ReplaceString(input string, oldSubstr string, newSubstr string) string {
	return strings.Replace(input, oldSubstr, newSubstr, -1)
}

func GetStringLength(input string) int {
	return len(input)
}

func ReverseString(input string) string {
	charArray := []rune(input)
	for i, j := 0, len(charArray)-1; i < j; i, j = i+1, j-1 {
		charArray[i], charArray[j] = charArray[j], charArray[i]
	}
	return string(charArray)
}

func SplitString(input, delimiter string) []string {
	return strings.Split(input, delimiter)
}

func JoinStrings(elements []string, separator string) string {
	return strings.Join(elements, separator)
}

func ContainsSubstring(input, substring string) bool {
	return strings.Contains(input, substring)
}

func ToLowerCase(input string) string {
	return strings.ToLower(input)
}

func ToUpperCase(input string) string {
	return strings.ToUpper(input)
}

func TrimSpaces(input string) string {
	return strings.TrimSpace(input)
}
