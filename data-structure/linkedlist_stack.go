package main

import (
	"errors"
)

type Node struct {
	next *Node
	value interface{}
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	root Node
	length int
}

func (l *LinkedList) Init() *LinkedList {
	l.root.next = nil
	l.length = 0
	return l
}
//添加到最后
func (l *LinkedList) Append(n *Node) (bool, *Node) {
	var lastNode Node
	i := &(l.root)
	for {
		if i.next == nil {
			lastNode = *i
			break
		}
		i = i.next
	}
	lastNode.next = n
	l.length++
	return true,n
}
//删除最后结点，并返回
func (l *LinkedList) DeleteLast() (bool, *Node) {
	if l.length == 0{
		return false,nil
	}
	var daoShuDiEr Node //倒数第二个结点
	i := &(l.root)
	for {
		if i.next.next == nil {
			daoShuDiEr = *i
			break
		}
		i = i.next
	}
	last := daoShuDiEr.next
	daoShuDiEr.next = nil
	l.length--
	return true,last
}

type LinkedListStack struct {
	items *LinkedList
	count int //栈中元素个数
	n int //栈的大小
}

func (s *LinkedListStack) Init(length int) *LinkedListStack {
	s.items = new(LinkedList).Init()
	s.count = 0
	s.n = length
	return s
}

func (s *LinkedListStack) Push(val interface{}) (bool,error) {
	if s.count == s.n {
		return false,errors.New("not enough space")
	}
	node := Node{nil,val}
	s.items.Append(&node)
	s.count++
	return true,nil
}

func (s *LinkedListStack) Pop() (interface{}, error) {
	if s.count == 0{
		return nil,errors.New("empty stack")
	}
	res,node := s.items.DeleteLast()
	if res == false {
		return nil,errors.New("empty stack")
	}
	s.count--
	return node.value,nil
}

