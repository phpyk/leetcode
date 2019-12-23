package queue

import "errors"

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
func (l *LinkedList) Append(node *Node) *Node {
	var last *Node
	i := &(l.root)
	for {
		if i.next == nil {
			last = i
			break
		}
		i = i.next
	}
	last.next = node
	l.length++
	return node
}
//取头部结点
func (l *LinkedList) Pop() interface{} {
	//空队列
	if l.length == 0 {
		return nil
	}
	first := l.root.next
	l.root.next = first.next
	l.length--
	return first.value
}

//基于链表实现的队列
type LinkedListQueue struct {
	list *LinkedList
	n int	//队列长度
	head,tail *Node
}
//队列初始化
func (q *LinkedListQueue) Init(capacity int) *LinkedListQueue {
	q.list = new(LinkedList).Init()
	q.n = capacity
	q.head = q.list.root.next
	q.tail = nil
	return q
}
//入队
func (q *LinkedListQueue) Enqueue(node *Node) (*Node, error) {
	//队列已满
	if q.n == q.list.length {
		return nil,errors.New("not enough space")
	}
	//旧tail设置为新结点
	q.tail.next = node
	//新tail后移一位
	q.tail = q.tail.next
	q.n++
	return q.tail,nil
}

//出队
func (q *LinkedListQueue) Dequeue() (*Node, error) {
	//队列已空
	if q.n == 0 || q.list.length == 0 {
		return nil,errors.New("empty list")
	}
	n := q.head
	q.head = q.head.next
	q.n--
	return n,nil
}