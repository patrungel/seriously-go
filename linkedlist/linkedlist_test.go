package linkedlist

import (
	"testing"
)

func TestLinkedList_String(t *testing.T) {
	cases := []struct {
		Input    *LinkedList
		Expected string
	}{
		{New(), "LinkedList[]"},
		{New(1), "LinkedList[1]"},
		{New(1, 2, 3), "LinkedList[1 2 3]"},
		{New(3, "five", []uint8{3, 18, 4}), "LinkedList[3 five [3 18 4]]"},
	}
	for _, c := range cases {
		actual := c.Input.String()
		if actual != c.Expected {
			t.Errorf("Failed test: want %v, got %v instead", c.Expected, actual)
		} else {
			t.Logf("Passed for %v", actual)
		}

	}
}

func TestLinkedList_Reverse(t *testing.T) {
	cases := []struct {
		Input    *LinkedList
		Expected *LinkedList
	}{
		{New(), New()},
		{New(1), New(1)},
		{New(1, 2, 3), New(3, 2, 1)},
		{New(3, "five", byte('c')), New(byte('c'), "five", 3)},
	}
	for _, c := range cases {
		was := c.Input.String() // Keep the original list as we reverse in place
		actual := c.Input.Reverse()
		if !actual.Equal(c.Expected) {
			t.Errorf("Failed test for %v: want %v, got %v instead", was, c.Expected, actual)
		} else {
			t.Logf("Passed for %v => %v", was, actual)
		}
	}
}

func (l *LinkedList) Equal(that *LinkedList) bool {
	a := l
	b := that
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	for (a != nil) && (b != nil) {
		if a.Value != b.Value {
			return false
		}
		a, b = a.Next, b.Next
	}
	return true
}
