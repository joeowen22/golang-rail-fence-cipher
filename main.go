package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(encrypt(3, "WE ARE DISCOVERED FLEE AT ONCE"))
}

func encrypt(rails int, data string) string {
	cleansedData := removeSpace(data)
	chars := []rune(cleansedData)
	lines := make([][]rune, rails)
	direction := 0
	currentRail := 0
	for j := 0; j < len(chars); j = j + 1 {
		lines[currentRail] = append(lines[currentRail], chars[j])
		fmt.Printf("%s: %d \n", string(chars[j]), currentRail)
		direction = updateDirection(direction, currentRail, rails-1)
		currentRail = direction + currentRail
	}

	var sb strings.Builder
	for _, line := range lines {
		sb.WriteString(string(line))
	}
	return sb.String()
}

func updateDirection(direction int, current int, maxRail int) int {
	if current == 0 {
		return 1
	} else if current == maxRail {
		return -1
	}
	return direction
}

func removeSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
