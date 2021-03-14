package linked_list

import (
	"bytes"
	"fmt"
)

type node struct {
	next  *node
	value interface{}
}

//LinkedList - simple go data structure
type LinkedList struct {
	head   *node
	length int
}

func NewLinkedList() *LinkedList {
	var ll LinkedList
	return &ll
}

func (l *LinkedList) InsertFirst(i interface{}) {
	newNode := node{
		l.head,
		i,
	}
	l.head = &newNode
	l.length++
}

func (l LinkedList) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for l.head != nil {
		buf.WriteString(fmt.Sprint(l.head.value))
		l.head = l.head.next
		if l.head != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
