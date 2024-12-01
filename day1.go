package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day1_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)

		leftNumber, _ := strconv.Atoi(parts[0])
		rightNumber, _ := strconv.Atoi(parts[1])

		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]

		if distance < 0 {
			distance *= -1
		}

		totalDistance += distance
	}

	fmt.Println("Day 1 Part 1", totalDistance)
}
