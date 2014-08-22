package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Node struct {
	data int
	left *Node
	right *Node
}

type listInt []int

var sorted_list = make(listInt, 0, 10)

func (self listInt) exists(data int) bool {
	for _, v := range self {
		if data == v {
			return true
		}
	}

	return false
}

func (self *Node) add(data int) {
	if self.data == 0 {
		self.data = data
	}

	if data < self.data {
		if self.left == nil {
			node := Node {
				data: data,
				left: nil,
				right: nil,
			}
			self.left = &node
		} else {
			self.left.add(data)
		}

	} else if data > self.data {
		if self.right == nil {
			node := Node {
				data: data,
				left: nil,
				right: nil,
			}
			self.right = &node
		} else {
			self.right.add(data)
		}
	}
}


func (self *Node) sort() {
	if self.left != nil {
		self.left.sort()
	}

	sorted_list = append(sorted_list, self.data)

	if self.right != nil {
		self.right.sort()
	}

}


func main() {

	list := make(listInt, 10)
	var temp int

	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range list {
		for {
			temp = rand.Intn(100)
			if ! list.exists(temp) {
				list[i] = temp
				break
			}
		}
	}

	fmt.Printf("%v\n", list)

	tree := new(Node)
	for _, v := range list {
		tree.add(v)
	}
	tree.sort()

	fmt.Printf("%v\n", sorted_list)
}
