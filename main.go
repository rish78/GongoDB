package main

import "fmt"

func main() {
	// initialize db
	dal, _ := NewDal("/mainTest")

	node, _ := dal.getNode(dal.root)
	node.Dal = dal
	index, containingNode, _ := node.findKey([]byte("Key1"))
	res := containingNode.items[index]

	fmt.Printf("key is: %s, value is: %s", res.key, res.value)
	// Close the db
	_ = dal.Close()
}
