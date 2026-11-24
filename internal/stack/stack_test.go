package stack

import "testing"

func TestPush(t *testing.T) {
	a := New([]int{1, 2, 3})
	b := New(nil)
	Push(a, b)
	if len(a.Data) != 2 || a.Data[0] != 2 {
		t.Errorf("a devrait être [2 3], obtenu %v", a.Data)
	}
	if len(b.Data) != 1 || b.Data[0] != 1 {
		t.Errorf("b devrait être [1], obtenu %v", b.Data)
	}
}

func TestSwap(t *testing.T) {
	a := New([]int{1, 2, 3})
	Swap(a)
	if a.Data[0] != 2 || a.Data[1] != 1 {
		t.Errorf("swap raté : %v", a.Data)
	}
}

func TestRotate(t *testing.T) {
	a := New([]int{1, 2, 3})
	Rotate(a)
	if a.Data[0] != 2 || a.Data[2] != 1 {
		t.Errorf("rotate raté : %v", a.Data)
	}
}

func TestReverseRotate(t *testing.T) {
	a := New([]int{1, 2, 3})
	ReverseRotate(a)
	if a.Data[0] != 3 || a.Data[1] != 1 {
		t.Errorf("reverse rotate raté : %v", a.Data)
	}
}

func TestHasDuplicates(t *testing.T) {
	if !HasDuplicates([]int{1, 1, 2}) {
		t.Error("aurait dû détecter les doublons")
	}
	if HasDuplicates([]int{1, 2, 3}) {
		t.Error("ne devrait pas détecter de doublons")
	}
}

func TestIsSorted(t *testing.T) {
	if !IsSorted(New([]int{1, 2, 3})) {
		t.Error("devrait être trié")
	}
	if IsSorted(New([]int{3, 1, 2})) {
		t.Error("ne devrait pas être trié")
	}
}
