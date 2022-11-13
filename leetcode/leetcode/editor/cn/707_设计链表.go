package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

type MyLinkedNode struct {
	val  int
	next *MyLinkedNode
}

type MyLinkedList struct {
	head *MyLinkedNode
	tail *MyLinkedNode
	len  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	for i, p := 0, this.head; p != nil; p = p.next {
		if i == index {
			return p.val
		}
		i++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.len++
	node := &MyLinkedNode{val: val, next: this.head}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.len++
	node := &MyLinkedNode{val: val}
	if this.tail != nil {
		this.tail.next = node
		this.tail = node
	} else {
		this.tail = node
		this.head = this.tail
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	}
	if index <= 0 || this.head == nil {
		this.AddAtHead(val)
		return
	}
	if index == this.len {
		this.AddAtTail(val)
		return
	}
	this.len++
	for i, p := 0, this.head; p != nil; p = p.next {
		if i == index-1 {
			node := &MyLinkedNode{val: val, next: p.next}
			p.next = node
			if node.next == nil {
				this.tail = node
			}
			return
		}
		i++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}
	this.len--
	if index == 0 {
		this.head = this.head.next
		if this.head == nil {
			this.tail = nil
		}
		return
	}

	for i, p := 0, this.head; p != nil; p = p.next {
		if i == index-1 {
			p.next = p.next.next
			if p.next == nil {
				this.tail = p
			}
			return
		}
		i++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := OrderedStreamConstructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
