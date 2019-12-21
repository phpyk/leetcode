package main

import (
	"fmt"
)

func main() {
	l1 := NewList()
	for i:= 9;i>0;i-=2 {
		e := Element{nil,i}
		l1.Insert(&e,&l1.root)
	}
	fmt.Println("--------l1---------")
	l1.Print()
	l2 := NewList()
	for i:=10;i>0;i-=2 {
		e := Element{nil,i}
		l2.Insert(&e,&l2.root)
	}
	fmt.Println("--------l2---------")
	l2.Print()
	l3 := NewList()
	l3 = SortedLinkedListMerge(l1,l2)
	fmt.Println("--------l3---------")
	l3.Print()
}

type Element struct {
	next  *Element
	Value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

// ----------------------------------------------------------------------------
type List struct {
	root Element //root 作为哨兵节点
	len  int
}

func NewList() *List {
	return new(List).Init()
}

//初始化list的时候，root的下一节点指向它自身
func (l *List) Init() *List {
	l.root.next = nil
	l.len = 0
	return l
}

//------------------------ 链表基本操作 ------------------------
//将e插入到at后面，返回插入的结点
func (l *List) Insert(e *Element, at *Element) *Element {
	n := at.next
	at.next = e
	e.next = n
	l.len += 1
	return e
}

func (l *List) Delete(e *Element) *Element {
	prev := l.root.next
	for ; prev.next != nil; prev = prev.next {
		//prev为空说明遍历完了，e不在list中
		if prev == nil {
			return nil
		}
		if prev.next == e {
			break
		}
	}

	prev.next = e.next
	l.len -= 1
	return prev
}

func (l *List) Print() {
	cur := l.root.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("[val:%v, addr:%p, next_addr:%p]",cur.Value,cur,cur.next)
		cur = cur.next
		if cur != nil {
			format += " --> \n"
		}
	}
	fmt.Println(format)
}
//------------------------ 链表相关算法 ------------------------
//单链表反转
func (l *List) Reverse() {
	if (*l).len <= 1 {
		return
	}
	//用lastNode记录
	var lastNode *Element
	var n *Element

	for p := l.root.next; p != nil; {
		//n记录下一个结点
		n = p.next
		//将当前结点的下一个结点设置为上一个
		p.next = lastNode
		//反转完毕，lastNode和p向后移动一位
		l.root.next = p //注意这一句：设置l的头结点
		lastNode = p
		p = n
	}
}
//单链表转为循环链表
func (l *List) TransToCircle()  {
	head := l.root.next
	for p := l.root.next;p != nil;p = p.next {
		if p.next == nil {
			(*p).next = head
			break
		}
	}
}

//检测是否是循环链表
func (l *List) IsCircleList() bool {
	if (*l).len == 0 {
		return false
	}
	if (*l).len == 1 && l.root.next == &l.root {
		return true
	}

	isCircle := false
	i := 1
	head := l.root.next
	for p := l.root.next; p != nil; p = p.next {
		//遍历l.len次之后，p的下一个==头结点，即为循环链表
		if i == (*l).len && p.next == head {
			isCircle = true
			break
		}
		i++
	}
	return isCircle
}

//两个有序链表合并
func SortedLinkedListMerge(l1 *List, l2 *List) *List {
	nl := NewList()
	insertAt := nl.root
	n := l1.root.next
	for p := l2.root.next;p != nil; p = p.next {
		t := n.next
		if n.Value.(int) < p.Value.(int) {
			res := nl.Insert(n,&insertAt)
			insertAt = *res
		}else {
			res := nl.Insert(p,&insertAt)
			insertAt = *res
		}
		n = t
		if n == nil {
			break
		}
	}
	return nl
}