package set // import "kkn.fi/set"

// IntSet implements a set data structure that holds ints.
type IntSet struct {
	m map[int]struct{}
}

// NewInt creates a set of ints initialized with given values.
func NewInt(values ...int) *IntSet {
	s := &IntSet{}
	s.m = make(map[int]struct{})
	s.Add(values...)
	return s
}

// Add adds given values to the set.
func (s *IntSet) Add(values ...int) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		s.m[value] = valueExists
	}
}

// Contains returns true if set holds all values and false otherwise.
func (s IntSet) Contains(values ...int) bool {
	if len(values) == 0 {
		return false
	}
	for _, value := range values {
		if _, has := s.m[value]; !has {
			return false
		}
	}
	return true
}

// Remove removes given values from the set.
func (s *IntSet) Remove(values ...int) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		delete(s.m, value)
	}
}

// Clear removes all values from the set.
func (s *IntSet) Clear() {
	s.m = make(map[int]struct{})
}

// Visit iterates through the set and visits all values with function f.
// Iteration will stop if function f returns false.
// Visit will return false when set is visited partially.
func (s IntSet) Visit(f func(value int) bool) bool {
	for v := range s.m {
		if !f(v) {
			return false
		}
	}
	return true
}

// Copy returns a copy of the set.
func (s IntSet) Copy() *IntSet {
	return NewInt(s.Slice()...)
}

// Equals returns true if both sets are equal, but false otherwise.
func (s IntSet) Equals(other *IntSet) bool {
	if s.Len() != other.Len() {
		return false
	}
	e := true
	other.Visit(func(value int) bool {
		_, e = s.m[value]
		return e
	})
	return e
}

// IsSubset returns true if given set is a subset of s.
func (s IntSet) IsSubset(other *IntSet) bool {
	result := true
	other.Visit(func(value int) bool {
		_, result = s.m[value]
		return result
	})
	return result
}

// IsSuperset returns true if given set is a superset of s.
func (s IntSet) IsSuperset(other *IntSet) bool {
	return other.IsSubset(&s)
}

// Slice returns a slice of ints copied from the set contents.
func (s IntSet) Slice() []int {
	result := make([]int, 0, len(s.m))
	for value := range s.m {
		result = append(result, value)
	}
	return result
}

// Len returns the size of the set.
func (s IntSet) Len() int {
	return len(s.m)
}

// IsEmpty return true is set is empty, false otherwise.
func (s IntSet) IsEmpty() bool {
	return len(s.m) == 0
}
