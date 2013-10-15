package main

import (
	"fmt"
	"os"
)

var next []int


func GetNext(t []byte) {
	i := 0
	j := -1
	next[0] = -1
	for i < len(t) - 1{
		if j == -1 || t[i] == t[j]{
			i += 1
			j += 1
			next[i] = j
			
		}else{
			j = next[j]
		}
	}
	fmt.Println(next[i])
}

func KMP(s, t []byte) int {
	i := 0
	j := 0
	for i < len(s) {
		//fmt.Println(j)
		if j == -1 || s[i] == t[j] {
			i += 1
			j += 1
		} else {
			j = next[j]
		}
		if j == len(t){
			return i - len(t)
		}
	}
		return -1
}

func main() {

	if len(os.Args) >= 3 {
		next = make([]int, len(os.Args[2]))
		//fmt.Println(os.Args[2])
		GetNext([]byte(os.Args[2]))
		k := KMP([]byte(os.Args[1]), []byte(os.Args[2]))
		fmt.Println("The position : ", k)
	}
}
