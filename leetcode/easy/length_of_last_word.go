package easy

import (
	"regexp"
)

func LengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	reg := regexp.MustCompile(`\s+`)
	words := reg.Split(s, -1)
	if len(words) == 0{
		return 0
	}
	result := words[len(words)-1:][0]
	if result == "" {
		return len(words[len(words)-2 : len(words)-1][0])
	}
	return len(result)
}
