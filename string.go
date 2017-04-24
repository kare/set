package set // import "kkn.fi/set"

// StringSet implements a set data structure that holds strings.
type StringSet struct {
	m map[string]struct{}
}

// NewString creates a set of strings initialized with given values.
func NewString(values ...string) *StringSet {
	s := &StringSet{}
	s.m = make(map[string]struct{})
	s.Add(values...)
	return s
}

// Add adds given values to the set.
func (s *StringSet) Add(values ...string) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		s.m[value] = valueExists
	}
}

// Contains returns true if set holds all values and false otherwise.
func (s StringSet) Contains(values ...string) bool {
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
func (s *StringSet) Remove(values ...string) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		delete(s.m, value)
	}
}

// Clear removes all values from the set.
func (s *StringSet) Clear() {
	s.m = make(map[string]struct{})
}

// Visit iterates through the set and visits all values with function f.
// Iteration will stop if function f returns false.
// Visit will return false when set is visited partially.
func (s StringSet) Visit(f func(value string) bool) bool {
	for v := range s.m {
		if !f(v) {
			return false
		}
	}
	return true
}

// Copy returns a copy of the set.
func (s StringSet) Copy() *StringSet {
	return NewString(s.Slice()...)
}

// Equals returns true if both sets are equal, but false otherwise.
func (s StringSet) Equals(other *StringSet) bool {
	if s.Len() != other.Len() {
		return false
	}
	e := true
	other.Visit(func(value string) bool {
		_, e = s.m[value]
		return e
	})
	return e
}

// IsSubset returns true if given set is a subset of s.
func (s StringSet) IsSubset(other *StringSet) bool {
	result := true
	other.Visit(func(value string) bool {
		_, result = s.m[value]
		return result
	})
	return result
}

// IsSuperset returns true if given set is a superset of s.
func (s StringSet) IsSuperset(other *StringSet) bool {
	return other.IsSubset(&s)
}

// Union returns an union set of s and other.
func (s StringSet) Union(other *StringSet) *StringSet {
	result := s.Copy()
	other.Visit(func(value string) bool {
		result.Add(value)
		return true
	})
	return result
}

// Slice returns a slice of strings copied from the set contents.
func (s StringSet) Slice() []string {
	result := make([]string, 0, len(s.m))
	for value := range s.m {
		result = append(result, value)
	}
	return result
}

// Len returns the size of the set.
func (s StringSet) Len() int {
	return len(s.m)
}

// IsEmpty return true is set is empty, false otherwise.
func (s StringSet) IsEmpty() bool {
	return len(s.m) == 0
}
