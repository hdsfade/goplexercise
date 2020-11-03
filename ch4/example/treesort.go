//@author: hdsfade
//@date: 2020-11-02-10:32
package sort

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value) //此处一定要给t.left赋值，因为如果t.left==nil
		//add中修改了指针，而函数只是值传递
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendvalues(values []int, t *tree) []int {
	if t != nil {
		values = appendvalues(values, t.left)
		values = append(values, t.value)
		values = appendvalues(values, t.right)
	}
	return values
}

func treesort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendvalues(values[:0], root)
}
