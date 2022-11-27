package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	size := 0
	var current *NodeL
	if l == nil {
		return nil
	} else {
		for l != nil {
			if size == pos {
				current = l
			}
			size++
			l = l.Next
		}
	}

	return current
}
