package LinearList

import "testing"

func Test(t *testing.T) {
	l := NewArrayList()
	l.Insert(1, 0)
	l.Insert(2, 1)
	l.Insert(3, 2)
	l.Insert(4, 3)
	l.Insert(5, 4)
	l.Insert(6, 5)
	v, err := l.Delete(2)
	if err != nil || v != 1 {
		t.Error("Error")
	}
	l.Print()
}
