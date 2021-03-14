package linked_list

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

type Value interface {
	String() string
}

type node struct {
	next *node
	value Value
}

func newValue(v interface{}) (Value, error) {
	var val Value
	switch t := v.(type) {
	case int, int8, int16, int32, int64:
		val = &intValue{value: reflect.ValueOf(v).Int()}
	case uint, uint8, uint16, uint32, uint64:
		val = &uIntValue{value: reflect.ValueOf(v).Uint()}
	default:
		return nil, errors.New(fmt.Sprintf("unsupported type: %T", t))
	}
	return val, nil
}

type uIntValue struct {
	value uint64
}

func (i *uIntValue) String() string {
	return fmt.Sprintf("%d", i.value)
}

type intValue struct {
	value int64
}

func (i *intValue) String() string {
	return fmt.Sprintf("%d", i.value)
}

//LinkedList - simple go data structure
type LinkedList struct {
	head *node
	length int
}

func NewLinkedList() *LinkedList {
	var ll LinkedList
	return &ll
}

func (l *LinkedList) InsertFirst(i interface{}) error {
	value, err := newValue(i)
	if err != nil {
		return err
	}
	newNode := node{
		l.head,
		value,
	}
	l.head = &newNode
	l.length++
	return nil
}

func (l LinkedList) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for l.head != nil {
		buf.WriteString(l.head.value.String())
		l.head = l.head.next
		if l.head != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
