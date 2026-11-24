package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Serialize encodes the tree as comma-separated preorder with "#" marking
// nil children — the markers are what make one traversal sufficient.
func Serialize(root *TreeNode) string {
	var b strings.Builder
	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			b.WriteString("#,")
			return
		}
		b.WriteString(strconv.Itoa(n.Val))
		b.WriteByte(',')
		walk(n.Left)
		walk(n.Right)
	}
	walk(root)
	return strings.TrimSuffix(b.String(), ",")
}

// Deserialize rebuilds the tree from Serialize's output: consume one token,
// then recursively rebuild left and right — the cursor self-synchronizes.
func Deserialize(s string) *TreeNode {
	tokens := strings.Split(s, ",")
	i := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if i >= len(tokens) || tokens[i] == "#" {
			i++
			return nil
		}
		v, _ := strconv.Atoi(tokens[i])
		i++
		n := &TreeNode{Val: v}
		n.Left = build()
		n.Right = build()
		return n
	}
	return build()
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, nil, nil), tn(3, tn(4, nil, nil), tn(5, nil, nil)))
	s := Serialize(root)
	fmt.Println(s)
	fmt.Println(Serialize(Deserialize(s)) == s)
}
