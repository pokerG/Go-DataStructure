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
	l.Print()
	err := l.Delete(2)
	if err != nil {
		t.Error("Error")
	}
	p, _ := l.Find(4)
	if p != 2 {
		t.Error("Error")
	}

}
