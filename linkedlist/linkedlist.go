package linkedlist

import (
	"fmt"
	"strings"
)

// Singly-linked list
type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

// Construct a new LinkedList from a number of elements
func New(a ...interface{}) *LinkedList {
	var l *LinkedList
	for i := len(a) - 1; i >= 0; i-- {
		l = &LinkedList{Value: a[i], Next: l}
	}
	return l
}

// Provide String representation on a LinkedList
func (l *LinkedList) String() string {
	var sb strings.Builder
	sb.WriteString("LinkedList[")
	for p := l; p != nil; p = p.Next {
		sb.WriteString(fmt.Sprint(p.Value))
		if p.Next != nil {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// Reverse given LinkedList in place
func (l *LinkedList) Reverse() *LinkedList {
	if l == nil {
		return nil
	}

	var tail *LinkedList
	for el, next := l, l.Next; el != nil; {
		el.Next = tail
		tail = el
		el = next
		if el != nil {
			next = el.Next
		}
	}
	return tail
}
