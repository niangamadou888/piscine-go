package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	newHead := l.Head
	for newHead != nil {
		if comp(newHead.Data, ref) {
			return &newHead.Data
		}
		newHead = newHead.Next
	}

	return nil
}
