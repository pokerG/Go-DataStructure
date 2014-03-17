package main

import (
	"CursorList"
	"fmt"
)

func main() {
	mNew := CursorList.NewSpace()
	mNew.Init()
	mNew.Insert(33, 5)
	mNew.Insert(23, 10)
	mNew.Insert(107, 7)
	mNew.Print()
	fmt.Println("******")
	mNew.Delete(5)
	mNew.Print()
	fmt.Println("******")
	mNew.Delete(10)
	mNew.Print()
	fmt.Println("******")
}
