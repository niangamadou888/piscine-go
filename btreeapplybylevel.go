package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	barham := BTreeLevelCount(root)
	for i := 1; i <= barham; i++ {
		BtreeCurrentLevel(root, i, f)
	}
}

func BtreeCurrentLevel(root *TreeNode, b int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if b == 1 {
		f(root.Data)
	} else if b > 1 {
		BtreeCurrentLevel(root.Left, b-1, f)
		BtreeCurrentLevel(root.Right, b-1, f)
	}
}
