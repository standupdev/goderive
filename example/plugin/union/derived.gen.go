// Code generated by goderive DO NOT EDIT.

package union

// deriveUnique returns a list containing only the unique items from the input list.
// It does this by reusing the input list.
func deriveUnique(list []*Person) []*Person {
	if len(list) == 0 {
		return nil
	}
	table := make(map[uint64][]int)
	u := 0
	for i := 0; i < len(list); i++ {
		contains := false
		hash := deriveHash(list[i])
		indexes := table[hash]
		for _, index := range indexes {
			if deriveEqual(list[index], list[i]) {
				contains = true
				break
			}
		}
		if contains {
			continue
		}
		if i != u {
			list[u] = list[i]
		}
		table[hash] = append(table[hash], u)
		u++
	}
	return list[:u]
}

// deriveFilter returns a list of all items in the list that matches the predicate.
func deriveFilter(predicate func(*Person) bool, list []*Person) []*Person {
	out := make([]*Person, 0, len(list))
	for i, elem := range list {
		if predicate(elem) {
			out = append(out, list[i])
		}
	}
	return out
}

// deriveUnion returns the union of the items of the two input lists.
// It does this by append items to the first list.
func deriveUnion(this, that []*Person) []*Person {
	for i, v := range that {
		if !deriveContains(this, v) {
			this = append(this, that[i])
		}
	}
	return this
}

// deriveEqual returns whether this and that are equal.
func deriveEqual(this, that *Person) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			this.Name == that.Name &&
			((this.Vote == nil && that.Vote == nil) || (this.Vote != nil && that.Vote != nil && *(this.Vote) == *(that.Vote)))
}

// deriveHash returns the hash of the object.
func deriveHash(object *Person) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	h = 31*h + deriveHash_(object.Name)
	h = 31*h + deriveHash_1(object.Vote)
	return h
}

// deriveContains returns whether the item is contained in the list.
func deriveContains(list []*Person, item *Person) bool {
	for _, v := range list {
		if deriveEqual(v, item) {
			return true
		}
	}
	return false
}

// deriveHash_ returns the hash of the object.
func deriveHash_(object string) uint64 {
	var h uint64
	for _, c := range object {
		h = 31*h + uint64(c)
	}
	return h
}

// deriveHash_1 returns the hash of the object.
func deriveHash_1(object *string) uint64 {
	if object == nil {
		return 0
	}
	return (31 * 17) + deriveHash_(*object)
}
