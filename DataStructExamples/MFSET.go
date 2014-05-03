//Disjoint Set
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com
package main

import (
	"fmt"
)

/*
 test data:
2 4
2 6
4 8
1 3
3 7
1 5
5 9
*/

const n int = 10

type node struct {
	father int
	count  int //weight
}
type Mfset struct {
	element [n + 1]node
}

func (this *Mfset) Union(a, b int) {
	a = this.Find(a)
	b = this.Find(b)
	if this.element[a].count > this.element[b].count { //path compression
		this.element[b].father = a
		this.element[a].count += this.element[b].count
	} else {
		this.element[a].father = b
		this.element[b].count += this.element[a].count
	}
}

func (this *Mfset) Find(x int) int {
	tmp := x
	for this.element[tmp].father != tmp {
		tmp = this.element[tmp].father
	}
	return tmp
}

func (this *Mfset) Init(x int) {
	this.element[x].father = x
	this.element[x].count = 1
}

func Equivalence(s *Mfset) {
	var i, j, k, m int
	for i = 1; i < n+1; i++ {
		s.Init(i)
	}
	// fmt.Println(s)
	fmt.Scanf("%d %d\n", &i, &j)
	for !(i == 0 && j == 0) {
		// fmt.Println(i, j)
		k = s.Find(i)
		m = s.Find(j)
		if k != m {
			s.Union(i, j)
		}
		// fmt.Println(s)
		fmt.Scanf("%d %d\n", &i, &j)
	}

}

func main() {
	m := &Mfset{}
	fmt.Println("Please input the equivalence(0 0 is end)")
	Equivalence(m)
	// for _, v := range m.element {
	// 	fmt.Print(v, " ")
	// }
	fmt.Println("Please input two number of you want to know have relate(0 0 is end)")
	for {
		var i, j int
		fmt.Scanf("%d %d\n", &i, &j)
		if i == 0 && j == 0 {
			break
		}
		if m.Find(i) == m.Find(j) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
