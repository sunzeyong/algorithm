package queue

type CycleQueue struct {
	item       []int
	capacity   int
	head, tail int
}

func NewCycleQueue(cap int) *CycleQueue {
	return &CycleQueue{
		item:     make([]int, cap),
		capacity: cap,
	}
}

func (c *CycleQueue) EnQueue(input int) bool {
	// tail位置为空 但下一个位置是head位置 则表示队列满了
	// 当队列满时，图中的 tail 指向的位置实际上是没有存储数据的。所以，循环队列会浪费一个数组的存储空间。
	if (c.tail+1)%c.capacity == c.head {
		return false
	}

	c.item[c.tail] = input
	c.tail = (c.tail + 1) % c.capacity

	return true
}

func (c *CycleQueue) DeQueue() (int, bool) {
	if c.head == c.tail {
		return 0, false
	}
	output := c.item[c.head]
	c.head = (c.head + 1) % c.capacity
	return output, true
}
