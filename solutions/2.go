package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Move int

const (
	UnknownMove Move = iota
	Rock
	Paper
	Scissors
)

func toMove(m string) Move {
	switch m {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	log.Fatal("Unimplemented move: %v", m)
	return UnknownMove
}

type Result int

const (
	UnknownResult Result = iota
	Win
	Lose
	Tie
)

func toResult(r string) Result {
	switch r {
	case "X":
		return Lose
	case "Y":
		return Tie
	case "Z":
		return Win
	}
	log.Fatal("Unimplemented result: %v", r)
	return UnknownResult
}

var (
	// Map of self move, opp move, and result
	getScore = map[Move]map[Move]int{
		Rock:     map[Move]int{Scissors: 7, Paper: 1, Rock: 4},
		Paper:    map[Move]int{Rock: 8, Scissors: 2, Paper: 5},
		Scissors: map[Move]int{Paper: 9, Rock: 3, Scissors: 6},
	}

	// Map of opp move, result, self move
	getMove = map[Move]map[Result]Move{
		Rock:     map[Result]Move{Win: Paper, Lose: Scissors, Tie: Rock},
		Paper:    map[Result]Move{Win: Scissors, Lose: Rock, Tie: Paper},
		Scissors: map[Result]Move{Win: Rock, Lose: Paper, Tie: Scissors},
	}

	r = regexp.MustCompile(`[A-Z]`)
)

func Run2(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var part1Score int
	var part2Score int
	for scanner.Scan() {
		m := r.FindAllStringSubmatch(scanner.Text(), -1)
		opp := toMove(m[0][0])
		part1Move := toMove(m[1][0])
		part2Move := getMove[opp][toResult(m[1][0])]

		part1Score += getScore[part1Move][opp]

		part2Score += getScore[part2Move][opp]
	}

	fmt.Printf("Part 1: %v\n", part1Score)
	fmt.Printf("Part 2: %v\n", part2Score)
	return
}
