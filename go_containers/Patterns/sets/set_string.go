package set
/*
**  This code was taken from https://pkg.go.dev/github.com/golang-collections/collections/set
**  Additional code was added
**
*/
type (
	Set struct {
		hash map[string]bool
	}

	
)

// Create a new set
func New(initial ...string) *Set {
	s := &Set{make(map[string]bool)}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

func (this *Set)Get_hash_map()*map[string]bool{
  return &this.hash
}

// Find the difference between two sets
func (this *Set) Difference(set *Set) *Set {
	n := make(map[string]bool)

	for k, _ := range this.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = true
		}
	}

	return &Set{n}
}

// Call f for each item in the set
func (this *Set) Do(f func(string)) {
	for k, _ := range this.hash {
		f(k)
	}
}

// Test to see whether or not the element is in the set
func (this *Set) Has(element string) bool {
	_, exists := this.hash[element]
	return exists
}

// Add an element to the set
func (this *Set) Insert(element string) {
	this.hash[element] = true
}

// Find the intersection of two sets
func (this *Set) Intersection(set *Set) *Set {
	n := make(map[string]bool)

	for k, _ := range this.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = true
		}
	}

	return &Set{n}
}

// Return the number of items in the set
func (this *Set) Len() int {
	return len(this.hash)
}

// Test whether or not this set is a proper subset of "set"
func (this *Set) ProperSubsetOf(set *Set) bool {
	return this.SubsetOf(set) && this.Len() < set.Len()
}

// Remove an element from the set
func (this *Set) Remove(element string) {
	delete(this.hash, element)
}

// Test whether or not this set is a subset of "set"
func (this *Set) SubsetOf(set *Set) bool {
	if this.Len() > set.Len() {
		return false
	}
	for k, _ := range this.hash {
		if _, exists := set.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Find the union of two sets
func (this *Set) Union(set *Set) *Set {
	n := make(map[string]bool)

	for k, _ := range this.hash {
		n[k] = true
	}
	for k, _ := range set.hash {
		n[k] = true
	}

	return &Set{n}
}