package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	am := BTreeSearchItem(root, node.Data).Parent
	if am == nil {
		return rplc
	} else if node.Data < root.Data {
		root.Left = rplc
	} else {
		root.Right = rplc
	}
	rplc.Parent = am

	return root
}
