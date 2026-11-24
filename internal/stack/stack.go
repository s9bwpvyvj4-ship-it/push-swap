package stack

// Stack représente une pile d'entiers.
// L'élément du haut est à l'index 0.
type Stack struct {
	Data []int
}

// New crée une stack à partir d'une slice d'int.
func New(values []int) *Stack {
	return &Stack{Data: append([]int{}, values...)}
}

// Len retourne la taille de la stack.
func (s *Stack) Len() int {
	return len(s.Data)
}

// IsEmpty dit si la stack est vide.
func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

// Push enlève le haut de src et le met en haut de dst.
func Push(src, dst *Stack) {
	if src.IsEmpty() {
		return
	}
	top := src.Data[0]
	src.Data = src.Data[1:]                    // retirer le haut
	dst.Data = append([]int{top}, dst.Data...) // ajouter en haut
}

// Swap échange les 2 premiers éléments.
func Swap(s *Stack) {
	if s.Len() < 2 {
		return
	}
	s.Data[0], s.Data[1] = s.Data[1], s.Data[0]
}

// Rotate : le premier devient le dernier.
func Rotate(s *Stack) {
	if s.Len() < 2 {
		return
	}
	first := s.Data[0]
	s.Data = append(s.Data[1:], first)
}

// ReverseRotate : le dernier devient le premier.
func ReverseRotate(s *Stack) {
	if s.Len() < 2 {
		return
	}
	last := s.Data[s.Len()-1]
	s.Data = append([]int{last}, s.Data[:s.Len()-1]...)
}
