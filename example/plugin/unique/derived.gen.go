// Code generated by goderive DO NOT EDIT.

package unique

func deriveUnique(list []*Visitor) []*Visitor {
	if len(list) == 0 {
		return nil
	}
	u := 1
	for i := 1; i < len(list); i++ {
		if !deriveContainsSliceOfPtrToVisitor(list[:u], list[i]) {
			if i != u {
				list[u] = list[i]
			}
			u++
		}
	}
	return list[:u]
}

func deriveContainsSliceOfPtrToVisitor(list []*Visitor, item *Visitor) bool {
	for _, v := range list {
		if deriveEqualPtrToVisitor(v, item) {
			return true
		}
	}
	return false
}

func deriveEqualPtrToVisitor(this, that *Visitor) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			((this.UserName == nil && that.UserName == nil) || (this.UserName != nil && that.UserName != nil && *(this.UserName) == *(that.UserName))) &&
			this.RemoteAddr == that.RemoteAddr
}
