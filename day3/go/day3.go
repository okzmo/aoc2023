package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}

type Number struct {
	Val                int // 113
	Start, End         int // 0 2
	Line               int // 0
	SymboleY, SymboleX int
}

func getNumber(idx int, lineNum int, line string, SymboleX int, SymboleY int) Number {
	currentNum := ""
	out := Number{Line: lineNum}

	if isNumber(string(line[idx-1])) && isNumber(string(line[idx+1])) {
		for i := idx; i >= 0; i-- {
			if isDot(string(line[i])) {
				out.Start = i + 1
				break
			}
			currentNum += string(line[i])
		}
		currentNum = Reverse(currentNum)

		for i := idx + 1; i < len(line); i++ {
			if isDot(string(line[i])) {
				out.End = i - 1
				break
			}
			currentNum += string(line[i])
		}
	} else if isNumber(string(line[idx-1])) {
		for i := idx; i >= 0; i-- {
			if isDot(string(line[i])) {
				out.Start = i + 1
				break
			}
			currentNum += string(line[i])
		}
		currentNum = Reverse(currentNum)
		out.End = idx
	} else {
		for i := idx; i < len(line); i++ {
			if isDot(string(line[i])) {
				out.End = i - 1
				break
			}
			currentNum += string(line[i])
		}
		out.Start = idx
	}

	num, _ := strconv.Atoi(currentNum)
	out.Val = num
	out.SymboleY = SymboleY
	out.SymboleX = SymboleX
	return out
}

func isNumber(char string) bool {
	if _, err := strconv.Atoi(string(char)); err == nil {
		return true
	}
	return false
}

func isDot(char string) bool {
	return char == "."
}

func checkLines(lines []string) int {
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	res := 0

	numbers := []Number{}
	var seen = make(map[Number]bool)

	for i, line := range lines {
		for j, rowNum := range line {
			if !isNumber(string(rowNum)) && !isDot(string(rowNum)) {
				for _, dir := range dirs {
					y, x := i+dir[0], j+dir[1]
					if y >= 0 && y < len(lines) && x >= 0 && x < len(lines[y]) {
						char := lines[y][x]
						if isNumber(string(char)) && string(rowNum) == "*" {
							currentNum := getNumber(x, y, lines[y], i, j)
							if !seen[currentNum] {
								if len(numbers) > 0 && (numbers[0].SymboleY != currentNum.SymboleY || numbers[0].SymboleX != currentNum.SymboleX) {
									numbers = []Number{currentNum}
								} else {
									numbers = append(numbers, currentNum)
								}

								if len(numbers) == 2 {
									res += numbers[0].Val * numbers[1].Val
									numbers = []Number{}
								}
								seen[currentNum] = true
							}
						}
					}
				}
			}
		}
	}

	return res
}

func GetResult() {
	file, err := os.Open("day3/go/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	res := checkLines(lines)

	println(res)
}
