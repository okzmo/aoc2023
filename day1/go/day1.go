package day1

import (
	"os"
	"regexp"
	"slices"
	"strings"
)

func GetResult() int {
	input, _ := os.ReadFile("day1/go/day1.txt")
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	first := regexp.MustCompile(`(` + strings.Join(numbers, "|") + `)`)
	last := regexp.MustCompile(`.*` + first.String())
	res := 0

	for _, line := range strings.Fields(string(input)) {
		res += 10 * (slices.Index(numbers, first.FindStringSubmatch(line)[1])%9 + 1)
		res += (slices.Index(numbers, last.FindStringSubmatch(line)[1])%9 + 1)
	}

	return res
}
