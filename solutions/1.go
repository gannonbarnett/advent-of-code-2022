package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"gannonbarnett.com/advent-of-code-2022/util"
)

const (
	TopSnaxRequired = 3
)

func Run1(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	l := util.NewLinkedList[int](TopSnaxRequired)
	curr := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			l.Push(curr)
			curr = 0
			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		curr += val
	}

	fmt.Printf("Part 1: %v\n", l.Index(0))

	sum := 0
	for i := 0; i < TopSnaxRequired; i += 1 {
		sum += l.Index(i)
	}
	fmt.Printf("Part 2: %v\n", sum)
}
