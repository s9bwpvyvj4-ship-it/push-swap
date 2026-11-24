package main

import (
	"fmt"
	"os"

	"push-swap/internal/stack"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	numbers, err := stack.ParseArgs(args)
	if err != nil || stack.HasDuplicates(numbers) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	a := stack.New(numbers)
	b := stack.New(nil)

	if stack.IsSorted(a) {
		return
	}

	// On choisit l'algorithme selon la taille de la liste
	if a.Len() <= 5 {
		sortSmall(a, b)
	} else {
		sortBig(a, b)
	}
}

// --- POUR LES PETITES LISTES (3, 4, 5 éléments) ---
func sortSmall(a, b *stack.Stack) {
	// On envoie le minimum dans b jusqu'à ce qu'il ne reste que 3 éléments dans a
	for a.Len() > 3 {
		minIdx := indexOfMin(a)
		if minIdx <= a.Len()/2 {
			for minIdx > 0 {
				stack.Rotate(a)
				fmt.Println("ra")
				minIdx--
			}
		} else {
			for minIdx < a.Len() {
				stack.ReverseRotate(a)
				fmt.Println("rra")
				minIdx++
			}
		}
		stack.Push(a, b)
		fmt.Println("pb")
	}

	// On trie les 3 éléments restants dans a
	sortThree(a)

	// On remet les éléments de b dans a
	for !b.IsEmpty() {
		stack.Push(b, a)
		fmt.Println("pa")
	}
}

// Tri spécifique pour 3 éléments
func sortThree(a *stack.Stack) {
	if a.Len() < 2 {
		return
	}
	if a.Data[0] > a.Data[1] && a.Data[1] < a.Data[2] && a.Data[0] < a.Data[2] {
		stack.Swap(a)
		fmt.Println("sa")
	} else if a.Data[0] > a.Data[1] && a.Data[1] > a.Data[2] {
		stack.Swap(a)
		fmt.Println("sa")
		stack.ReverseRotate(a)
		fmt.Println("rra")
	} else if a.Data[0] > a.Data[1] && a.Data[1] < a.Data[2] && a.Data[0] > a.Data[2] {
		stack.Rotate(a)
		fmt.Println("ra")
	} else if a.Data[0] < a.Data[1] && a.Data[1] > a.Data[2] && a.Data[0] < a.Data[2] {
		stack.Swap(a)
		fmt.Println("sa")
		stack.Rotate(a)
		fmt.Println("ra")
	} else if a.Data[0] < a.Data[1] && a.Data[1] > a.Data[2] && a.Data[0] > a.Data[2] {
		stack.ReverseRotate(a)
		fmt.Println("rra")
	}
}

// --- POUR LES GRANDES LISTES (Optimisé) ---
func sortBig(a, b *stack.Stack) {
	// Phase 1 : On envoie les éléments de a vers b
	// On cherche le minimum, on le rapproche du haut (ou du bas selon ce qui est plus court), et on l'envoie
	for a.Len() > 3 {
		minIdx := indexOfMin(a)
		if minIdx <= a.Len()/2 {
			for minIdx > 0 {
				stack.Rotate(a)
				fmt.Println("ra")
				minIdx--
			}
		} else {
			for minIdx < a.Len() {
				stack.ReverseRotate(a)
				fmt.Println("rra")
				minIdx++
			}
		}
		stack.Push(a, b)
		fmt.Println("pb")
	}

	sortThree(a)

	// Phase 2 : On remet tout dans a
	// Comme b contient les plus petits éléments en haut, on les renvoie dans l'ordre
	for !b.IsEmpty() {
		stack.Push(b, a)
		fmt.Println("pa")
	}
}

func indexOfMin(s *stack.Stack) int {
	minIdx := 0
	for i := 1; i < s.Len(); i++ {
		if s.Data[i] < s.Data[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}
