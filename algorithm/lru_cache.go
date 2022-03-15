package algorithm

type cacheNode struct {
	key  int
	val  int
	next *cacheNode
}

type LRUCache struct {
	head     *cacheNode
	tail     *cacheNode
	length   int
	capacity int
}

func Constructor(capacity int) LRUCache {
	cacheNode := &cacheNode{}
	return LRUCache{
		head:     cacheNode,
		tail:     cacheNode,
		length:   0,
		capacity: capacity,
	}
}

// 将listNode移动到链表的尾部
func (cache *LRUCache) toTail(node *cacheNode) {
	// 如果本身是链表的尾部就不用移动了
	if node == cache.tail {
		return
	}

	// 首先删除该链表
	key := node.key
	val := node.val
	node.val = node.next.val
	node.key = node.next.key
	if node.next == cache.tail {
		cache.tail = node
	}
	node.next = node.next.next

	// 然后插入
	cache.tail.next = &cacheNode{val: val, key: key}
	cache.tail = cache.tail.next
}

func (this *LRUCache) getAndPut(key int, val int) int {
	var p *cacheNode
	for p = this.head.next; p != nil; p = p.next {
		if p.key == key {
			if val >= 0 {
				p.val = val
			} else {
				val = p.val
			}
			this.toTail(p)
			return val
		}
	}
	return -1
}

func (this *LRUCache) Get(key int) int {
	return this.getAndPut(key, -1)
}

func (this *LRUCache) Put(key int, value int) {
	if this.getAndPut(key, value) < 0 {
		this.tail.next = &cacheNode{key: key, val: value}
		this.tail = this.tail.next
		if this.length++; this.length > this.capacity {
			this.head.next = this.head.next.next
			this.length--
		}
	}
	return
}
