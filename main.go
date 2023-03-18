package main

import (
	"github.com/btoll/howto-go/maze"
)

func main() {
	//	fmt.Println(str.Palindrome("o"))
	//	fmt.Println(str.Reverse("bentoll"))
	//	fmt.Println(str.CommonChars("bentoll", "gingertritch"))
	//	fmt.Println(str.Atoi("t0070042"))
	//	fmt.Println(str.Atoi("00t70042"))
	//	fmt.Println(str.Atoi("007t0042"))
	//	fmt.Println(str.Atoi("0070t042"))
	//	fmt.Println(str.Atoi("00700t42"))
	//	fmt.Println(str.Atoi("007004t2"))
	//	fmt.Println(str.Atoi("0070042t"))
	//	fmt.Println(str.Atoi("   -42"))
	// fmt.Println(str.Strstr("sadbutsad", "sad"))
	// fmt.Println(ints.SumTwoInts([]int{3, 5, 8, 9}, 12))
	// fmt.Println(ints.BinarySearch([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}, 19))
	// fmt.Println(ints.Reverse(-123))
	//	list := linkedlist.New(1)
	//	list.Append(2)
	//	list.Append(3)
	//	list.Append(4)
	//	list.Print()
	//	fmt.Println()
	//	list.Reverse()
	//	list.Print()
	//	s := stack.New(5)
	//	s.Print()
	//	fmt.Println()
	//	s.Push(7)
	//	s.Push(9)
	//	s.Print()
	//	nums := []int{4, 2, 6, 5, 1, 3}
	//	nums := []int{1000000, 10000, 1000, 100, 10, 4, 2, 6, -1, 5, 1, 3, 0}
	//	nums := []int{4, 1, 8, 6}
	//	fmt.Println(sort.BubbleSort(nums))
	//	fmt.Println(sort.SelectionSort(nums))
	//	fmt.Println(sort.MergeSort(nums))
	//               10
	//             /    \
	//            /      \
	//           7        25
	//          / \      / \
	//         /   \    /   \
	//        4     8  21    32
	//		   \			/  \
	//		    \		   /    \
	//		     5		 30		 40
	//	tree.CreateBST([]int{10, 7, 25, 8, 4, 21, 32, 5, 30, 40})
	//	fmt.Println(t.Contains(8))
	//	fmt.Println(t.Root.Right.Right)
	//	fmt.Println("preorder", tree.Preorder(t.Root))
	//	fmt.Println("inorder", tree.Inorder(t.Root))
	//	fmt.Println("postorder", tree.Postorder(t.Root))
	//	for _, n := range []int{10, 7, 25, 8, 4, 21, 32} {
	//		fmt.Printf("get %v, getdirectparent %v\n", t.Get(n), t.GetDirectParent(n))
	//	}
	//	fmt.Println("inorder", tree.Inorder(t.Root))
	//	t.Delete(10)
	//	fmt.Print(t.MinValue(t.Root.Left.Right))
	//	fmt.Println(t.Root)
	//	nums := []int{10, 7, 25, 4, 8, 21, 32}
	//	nums := []int{10, 7, 25, 4, 8, 21, 32, 2, 6}
	//	nums := []int{25, 15, 50, 10, 22, 35, 70, 4, 12, 18, 24, 31, 44, 66, 90}
	//	t := tree.CreateBST(nums)
	//	fmt.Println(t.Root.Left)
	//	fmt.Println(algorithms.ValidateBST(nums))
	//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//	t := tree.CreateBT(nums)
	//	fmt.Println("LevelOrder", tree.LevelOrder(t.Root))
	//	fmt.Println("min", algorithms.Min(nums))
	//	fmt.Println("         bfs", tree.Bfs(t.Root))
	//	fmt.Println()
	//	fmt.Println("    preorder", tree.Preorder(t.Root))
	//	fmt.Println("  r_preorder", tree.R_Preorder(t.Root))
	//	fmt.Println()
	//	fmt.Println("     inorder", tree.Inorder(t.Root))
	//	fmt.Println("   r_inorder", tree.R_Inorder(t.Root))
	//	fmt.Println()
	//	fmt.Println("   postorder", tree.Postorder(t.Root))
	//	fmt.Println(" r_postorder", tree.R_Postorder(t.Root))
	//	fmt.Println()
	//	fmt.Println("  levelorder", tree.Levelorder(t.Root))
	//	fmt.Println("r_levelorder", tree.Levelorder(t.Root))
	//	fmt.Println()
	// fmt.Println(tree.Height(t.Root))
	//	fmt.Println("max path", algorithms.MaxPath(nums))
	// fmt.Println("level order", algorithms.ListOfDepths(nums))
	m := maze.New(9, 9)
	// fmt.Printf("%+v\n", m.CreateGrid())
	m.Create()
	m.Generate()
	m.Draw()
}
