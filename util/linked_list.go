package util

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Link[T constraints.Ordered] struct {
	v     T
	left  *Link[T]
	right *Link[T]
}

func (l *Link[T]) String() string {
	if l == nil {
		return "end"
	}
	return fmt.Sprintf("%v->%v", l.v, l.right)
}

func (l *Link[T]) Left() *Link[T] {
	if l == nil {
		return nil
	}
	return l.left
}

func (l *Link[T]) SetLeft(link *Link[T]) {
	if l == nil {
		return
	}
	l.left = link
}

func (l *Link[T]) Right() *Link[T] {
	if l == nil {
		return nil
	}
	return l.right
}

func (l *Link[T]) SetRight(link *Link[T]) {
	if l == nil {
		return
	}
	l.right = link
}

type LinkedList[T constraints.Ordered] struct {
	h *Link[T]
	t *Link[T]

	size     int
	capacity int
}

func (l *LinkedList[T]) String() string {
	return l.h.String()
}

func (l *LinkedList[T]) Index(i int) T {
	ii := 0
	link := l.h
	for ii < i {
		link = link.Right()
		ii += 1
	}
	return link.v
}

func NewLinkedList[T constraints.Ordered](capacity int) *LinkedList[T] {
	return &LinkedList[T]{capacity: capacity}
}

func (l *LinkedList[T]) Push(v T) {
	if l.h == nil {
		l.h = &Link[T]{v: v}
		l.t = l.h
	} else {
		var left *Link[T]
		right := l.h
		for right != nil && v < right.v {
			left = right
			right = right.Right()
		}

		link := &Link[T]{v: v, left: left, right: right}
		link.Left().SetRight(link)
		link.Right().SetLeft(link)
		if link.right == nil {
			l.t = link
		}
		if link.left == nil {
			l.h = link
		}
	}

	l.size += 1
	if l.size > l.capacity {
		l.Pop()
	}
}

func (l *LinkedList[T]) Pop() {
	l.t.Left().SetRight(nil)
	l.t = l.t.Left()
	l.size -= 1
}
