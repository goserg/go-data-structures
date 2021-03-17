// Реализация Связного списка на языке Go
package linked_list

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

type node struct {
	next  *node
	value interface{}
}

type LinkedList struct {
	head   *node
	length int
}

//NewLinkedList - простой конструктор
func NewLinkedList() *LinkedList {
	var ll LinkedList
	return &ll
}

//InsertFirst вставляет элемент в начало списка.
func (l *LinkedList) InsertFirst(i interface{}) {
	newNode := node{
		l.head,
		i,
	}
	l.head = &newNode
	l.length++
}

//Get принимает int индекс, возвращает value из списка под этим индексом.
//При некорректном индексе возвращает nil и ошибку.
func (l LinkedList) Get(i int) (interface{}, error) {
	if i < 0 || i >= l.length {
		return nil, errors.New(fmt.Sprintf("invalid index: %d", i))
	}
	current := l.head
	for i > 0 {
		i--
		current = current.next
	}
	return current.value, nil
}

//FromList генерирует Linked List из слийса или массива.
//Возвращает ошибку, если в аргумент подан неверный тип.
func FromList(list interface{}) (*LinkedList, error) {
	var ll LinkedList
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(list)
		for i := s.Len() - 1; i >= 0; i-- {
			ll.InsertFirst(s.Index(i))
		}
	default:
		return nil, errors.New("can generate list only from slices or arrays")
	}
	return &ll, nil
}

//Pop удаляет элемент под индексом i из списка. Возвращает удаленный элемент и ошибку.
func (l *LinkedList) Pop(i int) (interface{}, error) {
	if i < 0 || i >= l.length {
		return nil, errors.New(fmt.Sprintf("invalid index: %d", i))
	}
	current := l.head
	last := current
	for i > 0 {
		i--
		last = current
		current = current.next
	}
	last.next = current.next
	l.length--
	return current.value, nil
}

//String реализация интерфейса fmt.Stringer.
func (l LinkedList) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for l.head != nil {
		val := l.head.value
		_, ok := val.(string)
		if ok {
			buf.WriteString(fmt.Sprint(`"`, val, `"`))
		} else {
			buf.WriteString(fmt.Sprint(val))
		}
		l.head = l.head.next
		if l.head != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

//Length возвращает длину списка.
func (l LinkedList) Length() int {
	return l.length
}

//ToSlice возвращает срез []interface{}, сгенерированный на основе списка.
func (l LinkedList) ToSlice() []interface{} {
	sl := make([]interface{}, 0, l.length)

	current := l.head
	i := 0
	for current != nil {
		sl = append(sl, current.value)
		current = current.next
		i++
	}

	return sl
}
