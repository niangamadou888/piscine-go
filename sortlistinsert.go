package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	insert := &NodeI{Data: data_ref}
	current := l
	str := current
	for current != nil {
		str = current
		current = current.Next
	}
	str.Next = insert

	return ListSort(l)
}
