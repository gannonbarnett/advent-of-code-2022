package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"gannonbarnett.com/advent-of-code-2022/util"
)

var (
	numRegex = regexp.MustCompile(`[0-9]+`)
)

func Run4(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	encap := 0
	overlap := 0
	for scanner.Scan() {
		n := numRegex.FindAllString(scanner.Text(), -1)
		if len(n) != 4 {
			log.Fatal("not enough numbers in line %v", scanner.Text())
		}
		a1 := util.AtoiOrFatal(n[0])
		a2 := util.AtoiOrFatal(n[1])
		b1 := util.AtoiOrFatal(n[2])
		b2 := util.AtoiOrFatal(n[3])

		var aMin int
		var aMax int
		if a1 < a2 {
			aMin = a1
			aMax = a2
		} else {
			aMin = a2
			aMax = a1
		}
		var bMin int
		var bMax int
		if b1 < b2 {
			bMin = b1
			bMax = b2
		} else {
			bMin = b2
			bMax = b1
		}

		if aMin <= bMin && aMax >= bMax || bMin <= aMin && bMax >= aMax {
			encap += 1
		}

		if aMax <= bMax {
			if aMax >= bMin {
				overlap += 1
			}
		} else {
			if bMax >= aMin {
				overlap += 1
			}
		}
	}

	fmt.Printf("Part 1: %v\n", encap)
	fmt.Printf("Part 2: %v\n", overlap)
}
