package stack

import (
	"errors"
	"strconv"
	"strings"
)

// ParseArgs transforme les arguments du programme en []int.
// Accepte "3 2 1 0" (un seul argument) ou "3" "2" "1" "0" (plusieurs).
// Retourne une erreur si un argument n'est pas un entier.
func ParseArgs(args []string) ([]int, error) {
	var numbers []int
	for _, a := range args {
		for _, field := range strings.Fields(a) {
			n, err := strconv.Atoi(field)
			if err != nil {
				return nil, errors.New("not an integer")
			}
			numbers = append(numbers, n)
		}
	}
	return numbers, nil
}
