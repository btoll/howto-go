package tree

// R_Bfs(*Node) Visited
// R_Inorder(*Node) Visited
// R_Levelorder(node *Node) []Visited
// R_Postorder(*Node) Visited
// R_Preorder(*Node) Visited

//               10
//             /    \
//            /      \
//           7        25
//          / \      / \
//         /   \    /   \
//        4     8  21    32
//
// bfs       [10, 7, 25, 4, 8, 21, 32]
// preorder  [10, 7, 4, 8, 25, 21, 32]
// inorder   [4, 7, 8, 10, 21, 25, 32]
// postorder [4, 8, 7, 21, 32, 25, 10]

func R_Bfs(node *Node) Visited {
	return nil
}
func __r_inorder(node *Node, visited []int) Visited {
	if node == nil {
		return visited
	}
	visited = __r_inorder(node.Left, visited)
	visited = append(visited, node.Value)
	visited = __r_inorder(node.Right, visited)
	return visited
}
func R_Inorder(node *Node) Visited {
	return __r_inorder(node, []int{})
}
func __r_levelorder(node *Node, levels [][]int, level int) [][]int {
	if node == nil {
		return levels
	}
	levels[level] = append(levels[level], node.Value)
	levels = __r_levelorder(node.Left, levels, level+1)
	levels = __r_levelorder(node.Right, levels, level+1)
	return levels
}
func R_Levelorder(node *Node) [][]int {
	return __r_levelorder(node, [][]int{}, 0)
}
func __r_postorder(node *Node, visited []int) Visited {
	if node == nil {
		return visited
	}
	visited = __r_postorder(node.Left, visited)
	visited = __r_postorder(node.Right, visited)
	visited = append(visited, node.Value)
	return visited
}
func R_Postorder(node *Node) Visited {
	return __r_postorder(node, []int{})
}
func __r_preorder(node *Node, visited []int) Visited {
	if node == nil {
		return visited
	}
	visited = append(visited, node.Value)
	visited = __r_preorder(node.Left, visited)
	visited = __r_preorder(node.Right, visited)
	return visited
}
func R_Preorder(node *Node) Visited {
	return __r_preorder(node, []int{})
}
