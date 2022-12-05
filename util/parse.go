package util

import (
	"log"
	"strconv"
)

func AtoiOrFatal(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
