package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	start := 50       // initial position
	passcode := 0     // count times we land exactly on 0
	const size = 100  // positions 0–99

	for _, line := range lines {
		dir := line[0]
		move, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("invalid number:", line)
			continue
		}

		// LEFT MOVE
		if dir == 'L' {
			start = (start - move + size) % size
		}

		// RIGHT MOVE
		if dir == 'R' {
			start = (start + move) % size
		}

		// ✅ Count ONLY when we land exactly on 0
		if start == 0 {
			passcode++
		}

		fmt.Printf("location %v, zeros %v\n", start, passcode)
	}

	fmt.Println("Final count:", passcode)
}
