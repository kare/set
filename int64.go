// Code generated by "gends"; DO NOT EDIT.

package set // import "kkn.fi/set"

// Int64Set implements a set data structure that holds int64s.
type Int64Set struct {
	m map[int64]struct{}
}

// NewInt64 creates a set of int64s initialized with given values.
func NewInt64(values ...int64) *Int64Set {
	s := &Int64Set{}
	s.m = make(map[int64]struct{})
	s.Add(values...)
	return s
}

// Add adds given values to the set.
func (s *Int64Set) Add(values ...int64) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		s.m[value] = valueExists
	}
}

// Contains returns true if set holds all values and false otherwise.
func (s Int64Set) Contains(values ...int64) bool {
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
func (s *Int64Set) Remove(values ...int64) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		delete(s.m, value)
	}
}

// Clear removes all values from the set.
func (s *Int64Set) Clear() {
	s.m = make(map[int64]struct{})
}

// Visit iterates through the set and visits all values with function f.
// Iteration will stop if function f returns false.
// Visit will return false when set is visited partially.
func (s Int64Set) Visit(f func(value int64) bool) bool {
	for v := range s.m {
		if !f(v) {
			return false
		}
	}
	return true
}

// Copy returns a copy of the set.
func (s Int64Set) Copy() *Int64Set {
	return NewInt64(s.Slice()...)
}

// Equals returns true if both sets are equal, but false otherwise.
func (s Int64Set) Equals(other *Int64Set) bool {
	if s.Len() != other.Len() {
		return false
	}
	e := true
	other.Visit(func(value int64) bool {
		_, e = s.m[value]
		return e
	})
	return e
}

// IsSubset returns true if given set is a subset of s.
func (s Int64Set) IsSubset(other *Int64Set) bool {
	result := true
	other.Visit(func(value int64) bool {
		_, result = s.m[value]
		return result
	})
	return result
}

// IsSuperset returns true if given set is a superset of s.
func (s Int64Set) IsSuperset(other *Int64Set) bool {
	return other.IsSubset(&s)
}

// Union returns a new union set of s and other.
func (s Int64Set) Union(other *Int64Set) *Int64Set {
	result := s.Copy()
	other.Visit(func(value int64) bool {
		result.Add(value)
		return true
	})
	return result
}

// Slice returns a slice of int64s copied from the set contents.
func (s Int64Set) Slice() []int64 {
	result := make([]int64, 0, len(s.m))
	for value := range s.m {
		result = append(result, value)
	}
	return result
}

// Len returns the size of the set.
func (s Int64Set) Len() int {
	return len(s.m)
}

// IsEmpty return true is set is empty, false otherwise.
func (s Int64Set) IsEmpty() bool {
	return len(s.m) == 0
}
