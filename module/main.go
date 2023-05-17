package main

import (
	"fmt"
	"github.com/agnivade/levenshtein"
)

func main() {
	s1 := "kitten"
	s2 := "sitting"

	distance := levenshtein.ComputeDistance(s1, s2)

	fmt.Printf("Levenshtein distance between %s and %s is %d\n", s1, s2, distance)
}



