package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

const (
	RucksacksPerGroup = 3
)

func priority(r rune) int {
	p := int(r) % 32
	if unicode.IsUpper(r) {
		p += 26
	}
	return p
}

func Run3(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1Sum := 0
	part2Sum := 0
	badges := make(map[rune]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sack := scanner.Text()
		rucksackSet := make(map[rune]bool)
		firstCompartmentSet := make(map[rune]bool)
		part1Done := false
		part2Done := false
		for i, c := range scanner.Text() {
			if _, seen := rucksackSet[c]; !part2Done && !seen {
				if _, ok := badges[c]; !ok {
					badges[c] = 0
				}
				badges[c] += 1
				if badges[c] == RucksacksPerGroup {
					part2Sum += priority(c)

					// reset for next badge group
					part2Done = true
					badges = make(map[rune]int)
				}
			}
			rucksackSet[c] = true

			if !part1Done {
				if i < len(sack)/2 {
					firstCompartmentSet[c] = true
				} else {
					// in the second compartment
					if _, ok := firstCompartmentSet[c]; ok {
						part1Sum += priority(c)
						part1Done = true
					}
				}
			}
		}
	}
	fmt.Printf("Part 1: %v\n", part1Sum)
	fmt.Printf("Part 2: %v\n", part2Sum)
}
