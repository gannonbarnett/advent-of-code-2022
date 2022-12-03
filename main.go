package main

import (
	"fmt"

	"gannonbarnett.com/advent-of-code-2022/solutions"
	"github.com/jpillora/opts"
)

type Opts struct {
	Day       int    `opts:"help=day to solve"`
	InputFile string `opts:"help=input file"`
}

func main() {
	o := Opts{}
	opts.Parse(&o)

	switch o.Day {
	case 1:
		solutions.Run1(o.InputFile)
	case 2:
		solutions.Run2(o.InputFile)
	default:
		fmt.Printf("Unimplemented: %v\n", o.Day)
	}
}
