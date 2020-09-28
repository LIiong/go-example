package main

import "math"

/**
* https://leetcode-cn.com/problems/two-sum/
* 两数之和
 */
func twoSum(nums []int, target int) []int {

    m := make(map[int]int)
    for i := 0;i < len(nums);i++{
        m[nums[i]] = i
    }
    for i := 0;i < len(nums);i++{
        s := target - nums[i]
        if v,ok:=m[s];ok && v != i {
            return []int{i, m[s]}   
        }
    }
    return []int{}
}
//TreeNode Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
/**
* https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
* 二叉树的前序遍历
*/
func preorderTraversal(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    result := []int{root.Val}
    result = append(result, preorderTraversal(root.Left)...)
    result = append(result, preorderTraversal(root.Right)...)
    return result
}
/**
* https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
* 二叉搜索树的最近公共祖先
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    n := root
    l := n.Val > p.Val && n.Val > q.Val
    r := n.Val < p.Val && n.Val < q.Val
    for l || r {
        if l {
            n = n.Left
        }
        if r {
            n = n.Right
        }
        if n == nil {
            break
        }
        l = n.Val > p.Val && n.Val > q.Val
        r = n.Val < p.Val && n.Val < q.Val
    }
    return n
}
/**
* https://leetcode-cn.com/problems/path-sum-ii/
* 路径总和 II
* 深度优先算法
*/
func pathSum(root *TreeNode, sum int) [][]int {
    s := [][]int{}
    path := []int{}
    var dfs func(node  *TreeNode, less int)
    dfs = func(node *TreeNode, less int) {
        if node == nil {
            return
        }
        less -= node.Val
        path = append(path, node.Val)
        defer func() {path = path[:len(path) - 1]}()
        if less == 0 && node.Left == nil && node.Right == nil{ 
            s = append(s, append([]int(nil), path...))
            return
        }
        dfs(node.Left, less)
        dfs(node.Right, less)
    }
    dfs(root, sum)
    return s
}
/**
* https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
* 二叉树的最大深度
*/
func maxDepth(root *TreeNode) int {
    d := 0
    s := 0
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        d++
        if d > s {
            s = d
        }
        defer func() {d--}()
        if node.Left == nil && node.Right == nil{
            return
        }
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return s
}
/**
* 广度优先算法
*/
func maxDepthBFS(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    ans := 0
    for len(queue) > 0 {
        sz := len(queue)
        for sz > 0 {
            node := queue[0]
            queue = queue[1:]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            sz--
        }
        ans++
    }
    return ans
}
/**
* https://leetcode-cn.com/problems/validate-binary-search-tree/submissions/
* 验证二叉搜索树
* 递归
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := isValidBSTDP(root.Left, math.MinInt64, root.Val)
	r := isValidBSTDP(root.Right, root.Val, math.MaxInt64)
	return l && r
}

func isValidBSTDP(node *TreeNode, min, max int) bool {

	if node == nil {
		return true
	}
	if node.Val >= max || node.Val <= min{
		return false
	}
	l := isValidBSTDP(node.Left, min, node.Val)
	r := isValidBSTDP(node.Right, node.Val, max)
	return l && r
}
/**
* 迭代
*/
func isValidBSTFOR(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	minValue := math.MinInt64
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= minValue {
			return false
		}
		minValue = root.Val
		root = root.Right
	}
	return true
}
//Node Definition for a Node.
type Node struct {
    Val int
     Left *Node
     Right *Node
     Next *Node
}
/**
* https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
* 填充每个节点的下一个右侧节点指针 II
* 广度优先算法
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for  len(queue) > 0 {
		sz := len(queue)
		for  sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if  node.Left != nil {
				queue = append(queue, node.Left)
			}
			if  node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
			if sz > 0 {
				node.Next = queue[0]
			} else {
				node.Next = nil
			}
		}
	}
	return root
}