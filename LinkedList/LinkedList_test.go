package LinkedList

import "testing"

func TestLinkList(t *testing.T) {
	l := NewLinkedList()
	n := Node{1, nil}
	l.Insert(n)
	n = Node{2, nil}
	l.Insert(n)
	n = Node{3, nil}
	l.Insert(n)
	_, b := l.Find(1)

	if b == false {
		t.Error("Error")
	}
	err := l.Delete(2)
	if err != nil {
		//log.Fatal(err)
		t.Error("Error")
	}

}
