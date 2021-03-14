package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList_String(t *testing.T) {
	ll := LinkedList{}
	if ll.String() != "[]" {
		t.Errorf(fmt.Sprintf(`Empty Linked List should be represented as "[]", got "%s"`, ll.String()))
	}
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	if ll.String() != "[]" {
		t.Errorf(fmt.Sprintf(`Empty Linked List should be represented as "[]", got "%s"`, ll.String()))
	}
}

func TestLinkedList_InsertFirst(t *testing.T) {
	ll := LinkedList{}
	ll.InsertFirst(10)
	ll.InsertFirst("20")
	ll.InsertFirst(30)
	got := ll.String()
	want := `[30, "20", 10]`
	if got != want {
		t.Errorf(`Insetring 10, "20", 30: got %s, want %s`, got, want)
	}
}
