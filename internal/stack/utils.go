package stack

// HasDuplicates vérifie s'il y a des doublons dans la slice.
func HasDuplicates(nums []int) bool {
	seen := map[int]bool{}
	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}

// IsSorted vérifie que la stack est triée en ordre croissant.
func IsSorted(s *Stack) bool {
	for i := 1; i < s.Len(); i++ {
		if s.Data[i-1] > s.Data[i] {
			return false
		}
	}
	return true
}
