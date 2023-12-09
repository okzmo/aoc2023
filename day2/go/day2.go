package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Use this map for PART 2
// cubeMax := map[string]int{
// 	"red":   0,
// 	"green": 0,
// 	"blue":  0,
// }
//
// And change the return value from bool to int

func checkSets(sets []string) bool {
	cubeMax := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, set := range sets {
		cubes := strings.Split(strings.TrimSpace(set), ",")
		for _, cube := range cubes {
			infos := strings.Split(strings.TrimSpace(cube), " ")
			numOfCube, _ := strconv.Atoi(infos[0])
			if numOfCube > cubeMax[infos[1]] {
				//PART 1
				return false
				//PART 2
				// cubeMax[infos[1]] = numOfCube
			}
		}
	}

	// PART 2
	// res := 1
	// for _, color := range cubeMax {
	// 	res *= color
	// }

	//PART 1
	return true
	// PART 2
	//return res
}

func GetResult() int {
	file, err := os.Open("day2/go/day2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()
	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		game := strings.Split(line, ":")
		sets := strings.Split(strings.TrimSpace(game[1]), ";")

		// PART 1
		id, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		if checkSets(sets) {
			res += id
		}

		// PART 2
		// res += checkSets(sets)
	}

	return res
}
