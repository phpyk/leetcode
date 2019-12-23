package queue
//基于数组的循环队列
type CircleQueue struct {
	//数组：items，数组大小：n
	items []interface{}
	n int
	head,tail int	// head表示队头下标，tail表示队尾下标
}

func (c *CircleQueue) Init(capacity int) *CircleQueue {
	c.n = capacity
	return c
}
//入队
func (c *CircleQueue) Enqueue(item interface{}) bool {
	//队列已满
	if (c.tail+1)%c.n == c.head {
		return false
	}
	c.items[c.tail] = item
	c.tail = (c.tail+1)%c.n
	return true
}
//出队
func (c *CircleQueue) Dequeue() interface{} {
	if c.head == c.tail {
		return nil
	}
	item := c.items[c.head]
	c.head = (c.head+1)%c.n
	return item
}