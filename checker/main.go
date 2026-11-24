package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/internal/stack"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return // pas d'arguments → rien
	}

	// 1. Parser les nombres
	numbers, err := stack.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// 2. Vérifier les doublons
	if stack.HasDuplicates(numbers) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	a := stack.New(numbers)
	b := stack.New(nil)

	// 3. Lire les instructions sur stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if !execute(a, b, line) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// 4. Vérifier le résultat
	if stack.IsSorted(a) && b.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// execute applique une instruction. Retourne false si elle est invalide.
func execute(a, b *stack.Stack, instr string) bool {
	switch instr {
	case "pa":
		stack.Push(b, a)
	case "pb":
		stack.Push(a, b)
	case "sa":
		stack.Swap(a)
	case "sb":
		stack.Swap(b)
	case "ss":
		stack.Swap(a)
		stack.Swap(b)
	case "ra":
		stack.Rotate(a)
	case "rb":
		stack.Rotate(b)
	case "rr":
		stack.Rotate(a)
		stack.Rotate(b)
	case "rra":
		stack.ReverseRotate(a)
	case "rrb":
		stack.ReverseRotate(b)
	case "rrr":
		stack.ReverseRotate(a)
		stack.ReverseRotate(b)
	default:
		return false // instruction inconnue
	}
	return true
}
