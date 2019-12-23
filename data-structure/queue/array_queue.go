package queue
//基于数组实现的顺序队列
type ArrayQueue struct {
	//数组：items，数组大小：n
	items []interface{}
	n     int
	// head表示队头下标，tail表示队尾下标
	head int
	tail int //tail指向的位置其实是个空位置
}
// 申请一个大小为capacity的数组
func (a *ArrayQueue) Init(capacity int) *ArrayQueue {
	a.n = capacity
	return a
}
//入队
func (a *ArrayQueue) Enqueue(item interface{}) bool {
	//tail==n时，表示队尾没有空间了
	if a.tail == a.n {
		//head==0,表示整个队列都满了
		if a.head == 0 {
			return false
		}
		//数据搬移
		for i:=a.head;i<a.n;i++ {
			a.items[i-a.head] = a.items[i]
		}
		//充值head，tail
		a.tail = a.tail - a.head
		a.head = 0
	}
	a.items[a.tail] = item
	a.tail++
	return true
}
//出队
func (a *ArrayQueue) Dequeue() interface{} {
	//head==tail时，队列已空
	if a.head == a.tail {
		return nil
	}
	item := a.items[a.head]
	//出队时，items不能变，否则你数组的大小和n不相等了，只将head指针+1，
	//a.items = a.items[a.head+1:]
	a.head++
	return item
}