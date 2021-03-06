package set_test

import (
	"testing"

	"kkn.fi/set"
)

func TestIntNew(t *testing.T) {
	s := set.NewInt()
	if !s.IsEmpty() {
		t.Error("expected empty set")
	}
}

func TestIntAddDuplicate(t *testing.T) {
	s := set.NewInt(1)
	if !s.Contains(1) {
		t.Error("expected set to contain 1")
	}
	s.Add(1)
	if s.Len() != 1 {
		t.Errorf("expected len 1, but got %d", s.Len())
	}
}

func TestIntNewWithDuplicates(t *testing.T) {
	s := set.NewInt(1, 2, 3, 1, 2)
	if s.Len() != 3 {
		t.Errorf("expected len 3, but got %d", s.Len())
	}
}

func TestIntContains(t *testing.T) {
	s := set.NewInt(1, 2, 3)
	if !s.Contains(1, 2, 3) {
		t.Error("expected set to contain 1, 2, 3")
	}
	s = set.NewInt(4, 5, 6)
	if s.Contains(1) {
		t.Error("set contains 4, 5 and 6, but not 1")
	}
}

func TestIntRemove(t *testing.T) {
	s := set.NewInt(1, 2, 3)
	s.Remove(1)
	if s.Contains(1) {
		t.Error("set should not contain 1")
	}
	if s.Len() != 2 {
		t.Errorf("expected len 2, but got %d", s.Len())
	}
	s.Remove(2, 3)
	if s.Contains(2) {
		t.Error("set should not contain 2")
	}
	if s.Contains(3) {
		t.Error("set should not contain 3")
	}
}

func TestIntClear(t *testing.T) {
	s := set.NewInt(1, 2, 3)
	s.Clear()
	if !s.IsEmpty() {
		t.Errorf("expected empty set, but got len %d", s.Len())
	}
}

func TestIntVisit(t *testing.T) {
	s := set.NewInt(1, 2, 3)
	result := make([]int, 0)
	s.Visit(func(i int) bool {
		result = append(result, i)
		return true
	})
	if len(result) != 3 {
		t.Errorf("expected to visit 3 values, but got %d", len(result))
	}
}

func TestIntCopy(t *testing.T) {
	s := set.NewInt(1, 2, 3)
	copy := s.Copy()
	if copy.Len() != 3 {
		t.Errorf("expected set len 3, but got %d", copy.Len())
	}
}

func TestIntEquals(t *testing.T) {
	a := set.NewInt(1, 2, 3)
	b := set.NewInt(1, 2, 3)
	if !a.Equals(b) {
		t.Error("expected equality")
	}
	c := set.NewInt(4, 5, 6)
	if a.Equals(c) {
		t.Error("sets are not equal")
	}
}

func TestIntIsSubset(t *testing.T) {
	a := set.NewInt(1, 2, 3, 4)
	b := set.NewInt(1, 2, 3)
	if !b.IsSuperset(a) {
		t.Error("a is a superset of b")
	}
	if a.IsSuperset(b) {
		t.Error("b is not a superset of a")
	}
}

func TestIntIsSuperset(t *testing.T) {
	a := set.NewInt(1, 2, 3, 4)
	b := set.NewInt(1, 2, 3)
	if !a.IsSubset(b) {
		t.Error("a is a subset of b")
	}
	if b.IsSubset(a) {
		t.Error("b is not a subset of a")
	}
}

func TestIntUnion(t *testing.T) {
	a := set.NewInt(1, 3, 5, 7)
	b := set.NewInt(1, 2, 4, 6)
	c := a.Union(b)
	expected := set.NewInt(1, 2, 3, 4, 5, 6, 7)
	if !c.Equals(expected) {
		t.Errorf("expected set to equal %v", expected)
	}
}

func TestIntSlice(t *testing.T) {
	s := set.NewInt(1, 2, 3)
	slice := s.Slice()
	if len(slice) != 3 {
		t.Errorf("expected len 3, but got %d", len(slice))
	}
	expected := []int{1, 2, 3}
	for _, v := range expected {
		if !containsInt(v, slice) {
			t.Errorf("expected slice to contain %d", v)
		}
	}
}

func containsInt(value int, slice []int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
