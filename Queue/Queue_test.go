package Queue 

import "testing"

func Test(t * testing.T){
	queue := NewGoQueue()
	queue.Enter(1)
	queue.Enter(2)
	queue.Enter(3)
	queue.Enter(4)
	queue.Enter(5)
	queue.Enter(6)
	v, err := queue.Delete()
	if err != nil || v != 1{
		t.Error("Error")
	}
	queue.Print()
}