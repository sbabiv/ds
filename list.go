package ds

type node struct {
	prev, next *node
	data       interface{}
}

type List struct {
	head, tail *node
	len        int
}

func NewList() *List {
	return new(List)
}

func (list *List) Add(value interface{}) {
	n := &node{
		data: value,
	}
	if list.len == 0 {
		list.head, list.tail = n, n
	} else {
		n.prev = list.tail
		list.tail = n
	}
	list.len++
}

func (list *List) Get() (bool, interface{}) {
	if list.len == 0 {
		return false, nil
	}
	list.len--
	n := list.tail
	if n.prev == nil {
		list.head, list.tail = nil, nil
		return true, n
	}

	list.tail = list.tail.prev

	return true, n.data
}

func (list *List) Len() int {
	return list.len
}

func (list *List) Clear() {
	list.tail, list.head = nil, nil
	list.len = 0
}
