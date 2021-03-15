package linked_list

import (
	"errors"
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
	ll2 := LinkedList{}
	ll2.InsertFirst(30)
	ll2.InsertFirst(50)
	ll.InsertFirst(10)
	ll.InsertFirst("20")
	ll.InsertFirst(ll2)
	got := ll.String()
	want := `[[50, 30], "20", 10]`
	if got != want {
		t.Errorf(`Insetring 10, "20", [50, 30]: got %s, want %s`, got, want)
	}
}

func TestLinkedList_Length(t *testing.T) {
	ll := LinkedList{}
	if ll.Length() != 0 {
		t.Errorf("Length of empty LinkedList should be 0, got %d", ll.Length())
	}
	ll.InsertFirst(1)
	if ll.Length() != 1 {
		t.Errorf("Length of LinkedList should be 1, got %d", ll.Length())
	}
}

func TestLinkedList_Get(t *testing.T) {
	ll := LinkedList{}
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	testCases := []struct {
		i   int
		exp interface{}
		err error
	}{
		{i: 0, exp: 3, err: nil},
		{i: 1, exp: 2, err: nil},
		{i: 2, exp: 1, err: nil},
		{i: 3, exp: nil, err: errors.New("invalid index: 3")},
		{i: -1, exp: nil, err: errors.New("invalid index: -1")},
	}
	for _, testCase := range testCases {
		val, err := ll.Get(testCase.i)
		if err != nil && err.Error() != testCase.err.Error() {
			t.Errorf("expected error %v, got %v", err, testCase.err)
		}
		if val != testCase.exp {
			t.Errorf("expectex %v, got %v", val, testCase.exp)
		}
	}

}
