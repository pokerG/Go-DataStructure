package Stack

import "testing"

func Test(t *testing.T) {
	aStack := NewGoStack()
	if aStack.top != 20 {
		t.Error("...")
	}
	aStack.Push("1234")
	if aStack.Empty() == nil {
		t.Error("....")
	}
	if aStack.top != 19 {
		t.Error("...")
	}
	aStack.Push("hell")
	if aStack.top != 18 {
		t.Error("...")
	}
	aStack.Push("uuu")
	if aStack.top != 17 {
		t.Error("...")
	}
	a, _ := aStack.Top()
	if a != "uuu" {
		t.Error("Error")
	}
	if aStack.top != 17 {
		t.Error("...")
	}
	aStack.Pop()
	a, _ = aStack.Pop()
	if aStack.top != 19 {
		t.Error("...")
	}
	if a != "hell" {
		t.Error("Error")
	}

	aStack.MakeNull()
	_, err := aStack.Pop()
	if err == nil {
		t.Error("Error")
	}

	bStack := NewGoStackSize(3)
	bStack.Push(1)
	bStack.Push(1)
	bStack.Push(1)
	err = bStack.Push(1)
	if err == nil {
		t.Error("Error")
	}

}
