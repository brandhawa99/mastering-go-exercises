package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	struct where under each node there are at most 2 other nodes
	Depth of tree - (height) longest path from root to node
	depth of node - number of edges from the node to the root of the tree
	leaf - node with no children
	balanced - longest distance from root to a leaf is at most one more than the shortest such distance
	- balancing tree is expensive operation so balance from the beginning
*/

/*
	BT best for representing hierarchical data
	they are ordered by design
	deleting element is not trivial
	on balanced tree insert search and delete take log(n)steps
	height of balanced BT is ~log2(n) 10k elements will have height of 14

*/

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}
func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
