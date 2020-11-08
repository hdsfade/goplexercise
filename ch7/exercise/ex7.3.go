//@author: hdsfade
//@date: 2020-11-06-09:04
package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	vaule       int
	left, right *tree
}

func adds(t *tree, values []int) *tree {
	for _, v := range values {
		t = add(t, v)
	}
	return t
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.vaule = v
		return t
	}
	if t.vaule > v {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

func (t *tree) String() string {
	ret := ""
	var inorderTraver func(*tree)
	inorderTraver = func(t *tree) {
		if t.left != nil {
			inorderTraver(t.left)
		}
		if ret == "" {
			ret += strconv.Itoa(t.vaule)
		} else {
			ret += "," + strconv.Itoa(t.vaule)
		}
		if t.right != nil {
			inorderTraver(t.right)
		}
	}
	inorderTraver(t)
	return ret
}

func main() {
	var root *tree
	values := []int{1, 2, 3, 4, 5, 6}
	root = adds(root, values)
	fmt.Println(root)
}
