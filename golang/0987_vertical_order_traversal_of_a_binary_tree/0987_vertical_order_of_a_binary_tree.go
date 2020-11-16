package main

import (
	"fmt"
	"sort"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	var helper [][]int
	var tree func(*TreeNode, int, int)
	tree = func(t *TreeNode, x, y int) {
		if t == nil {
			return
		}
		helper = append(helper, []int{x, y, t.Val})
		tree(t.Left, x-1, y-1)
		tree(t.Right, x+1, y-1)
	}
	tree(root, 0, 0)
	sort.Slice(helper, func(i, j int) bool { // helper stores the sorted solutions
		return helper[i][0] < helper[j][0] || helper[i][0] == helper[j][0] && helper[i][1] > helper[j][1] || helper[i][0] == helper[j][0] && helper[i][1] == helper[j][1] && helper[i][2] < helper[j][2]
	})
	var res [][]int //to write the solutions in the required forms
	tmp := []int{helper[0][2]}
	x := helper[0][0]
	for i := 1; i < len(helper); i++ {
		if helper[i][0] != x {
			t := append([]int{}, tmp...)
			res = append(res, t)
			tmp = []int{helper[i][2]}
			x = helper[i][0]
		} else {
			tmp = append(tmp, helper[i][2])
		}
	}
	fmt.Println(helper)
	return append(res, tmp)
}
