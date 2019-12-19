package main

import (
	"fmt"
)

func main() {
	myList := NewList()
	fmt.Printf("myList: %+v \n", myList)
	fmt.Printf("myList.root address: %p \n", &myList.root)
	fmt.Printf("myList.root.next address: %p \n", &myList.root.next)
	fmt.Println("-----Insert e1,e2------")
	e1 := Element{nil, 100}
	myList.Insert(&e1, &myList.root)
	e2 := Element{nil, 200}
	myList.Insert(&e2, &e1)

	fmt.Println("-----Insert 5 elements------")
	//此处p作为哨兵节点，作为下一次插入时的位置
	p := &e2
	for i := 0; i < 3; i++ {
		e := Element{nil, i}
		myList.Insert(&e, p)
		p = &e
	}
	fmt.Println("-----Insert e3------")
	e3 := Element{nil, 300}
	myList.Insert(&e3, p)

	fmt.Println("-----Print all Elements------")
	myList.Print()

	fmt.Println("-----Delete e2------")
	myList.Delete(&e2)
	myList.Print()

	fmt.Println("-----Reverse list------")
	myList.Reverse()
	myList.Print()

	fmt.Println("-----Trans to circle list------")
	//myList.TransToCircle()
	//myList.Print()

	fmt.Println("-----Check if is circle------")
	fmt.Printf("list is Circle ? %v\n", myList.IsCircleList())
}

type Element struct {
	next  *Element
	Value interface{}
}

//func (e *Element) Next() *Element {
//	//处理边界
//	//当e.next 不为空且不为list的根节点时，返回next
//	if p := e.next; e.list != nil && p != &e.list.root {
//		return p
//	}
//	return nil
//}

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

//单链表反转
func (l *List) Reverse() {
	if (*l).len <= 1 {
		return
	}
	//用lastNode记录
	var lastNode *Element
	for p := l.root.next; p.next != nil; {
		//n记录下一个结点
		n := p.next
		//将当前结点的下一个结点设置为上一个
		p.next = lastNode
		//反转完毕，lastNode和p向后移动一位
		lastNode = p
		p = n
	}
}
//单链表转为循环链表
//func (l *List) TransToCircle()  {
//	head := l.root.next
//	for p := l.root.next;p != nil;p = p.next {
//		if p.list == nil {
//			(*p).next = head
//			break
//		}
//	}
//}

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
	for p := l.root.next; p != nil; p = p.next {
		if i == (*l).len {
			break
		}
		if p.next == nil {
			isCircle = true
			break
		}
		i++
	}
	return isCircle
}
