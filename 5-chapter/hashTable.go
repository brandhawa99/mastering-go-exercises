package main

const SIZE = 15

/*
DS that stores one or more key and value pairs and uses a hash function to compute an index into an array of buckets or slots from which the correct value can be discovered
Ideally hash function assigns each key a unique bucket
*/
type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}
