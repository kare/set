// Code generated by "gends"; DO NOT EDIT.

package set // import "kkn.fi/set"

// Byte implements a set data structure that holds bytes.
type Byte struct {
	m map[byte]struct{}
}

// NewByte creates a set of bytes initialized with given values.
func NewByte(values ...byte) *Byte {
	s := &Byte{}
	s.m = make(map[byte]struct{})
	s.Add(values...)
	return s
}

// Add adds given values to the set.
func (s *Byte) Add(values ...byte) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		s.m[value] = valueExists
	}
}

// Contains returns true if set holds all values and false otherwise.
func (s *Byte) Contains(values ...byte) bool {
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
func (s *Byte) Remove(values ...byte) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		delete(s.m, value)
	}
}

// Clear removes all values from the set.
func (s *Byte) Clear() {
	s.m = make(map[byte]struct{})
}

// Visit iterates through the set and visits all values with function f.
// Iteration will stop if function f returns false.
// Visit will return false when set is visited partially.
func (s *Byte) Visit(f func(value byte) bool) bool {
	for v := range s.m {
		if !f(v) {
			return false
		}
	}
	return true
}

// Copy returns a copy of the set.
func (s *Byte) Copy() *Byte {
	return NewByte(s.Slice()...)
}

// Equals returns true if both sets are equal, but false otherwise.
func (s *Byte) Equals(other *Byte) bool {
	if s.Len() != other.Len() {
		return false
	}
	e := true
	other.Visit(func(value byte) bool {
		_, e = s.m[value]
		return e
	})
	return e
}

// IsSubset returns true if given set is a subset of s.
func (s *Byte) IsSubset(other *Byte) bool {
	result := true
	other.Visit(func(value byte) bool {
		_, result = s.m[value]
		return result
	})
	return result
}

// IsSuperset returns true if given set is a superset of s.
func (s *Byte) IsSuperset(other *Byte) bool {
	return other.IsSubset(s)
}

// Union returns a new union set of s and other.
func (s *Byte) Union(other *Byte) *Byte {
	result := s.Copy()
	other.Visit(func(value byte) bool {
		result.Add(value)
		return true
	})
	return result
}

// Slice returns a slice of bytes copied from the set contents.
func (s *Byte) Slice() []byte {
	result := make([]byte, 0, len(s.m))
	for value := range s.m {
		result = append(result, value)
	}
	return result
}

// Len returns the size of the set.
func (s *Byte) Len() int {
	return len(s.m)
}

// IsEmpty return true is set is empty, false otherwise.
func (s *Byte) IsEmpty() bool {
	return len(s.m) == 0
}
