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
	if err := ll.InsertFirst(10); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if err := ll.InsertFirst(20); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if err := ll.InsertFirst(30); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	got := ll.String()
	want := "[30, 20, 10]"
	if got != want {
		t.Errorf("Insetring 10, 20, 30: got %s, want %s", got, want)
	}
}

func TestNewValue(t *testing.T) {
	var i = 2
	var i8 int8 = 2
	var i16 int16 = 2
	var i32 int32 = 2
	var i64 int64 = 2
	var ui uint = 2
	var ui8 uint8 = 2
	var ui16 uint16 = 2
	var ui32 uint32 = 2
	var ui64 uint64 = 2
	testCases := []struct{
		value interface{}
	}{
		{i}, {i8}, {i16}, {i32}, {i64},
		{ui}, {ui8}, {ui16}, {ui32}, {ui64},
	}
	for _, testCase := range testCases {
		val, err := newValue(testCase.value)
		if err != nil {
			t.Errorf(fmt.Sprintf("should support %T value, returning error: %s", testCase.value, err))
		} else if val.String() != "2" {
			t.Errorf(fmt.Sprintf(`String() method of Value object of %T 10 should return "2", got "%s"`, testCase.value, val.String()))
		}
	}
}

func TestNewValueUnsupportedTypeError(t *testing.T) {
	ll := LinkedList{}
	err := ll.InsertFirst(struct {}{})
	want := "unsupported type: struct {}"
	if err == nil || err.Error() != want {
		t.Errorf("expected  error %s, got %s", want, err)
	}
}
