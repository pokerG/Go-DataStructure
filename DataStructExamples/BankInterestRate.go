package main

import (
	"ArrayList"
	"fmt"
)

func main() {
	al := ArrayList.NewArrayList()
	var index int
	var s1, s2 string
	p := 0
	menu()
	for {
		fmt.Scanf("%d\n", &index)
		switch index {
		case 1:
			fmt.Scanf("%s %s", &s1, &s2)
			err := al.Insert(s1+" "+s2, p)
			if err != nil {
				fmt.Println(err)
			} else {
				p++
			}
		case 2:
			var q int
			fmt.Scanf("%d\n", &q)
			err := al.Delete(q)
			if err != nil {
				fmt.Println(err)
			} else {
				p--
			}
		case 3:
			al.Print()
		case 4:
			menu()
		case 5:
			return
		}
	}
}

func menu() {
	fmt.Println("***Welcome Bank Interest Rate Manager***")
	fmt.Println("1.Iuput a new rate")
	fmt.Println("2.Delete a old rate")
	fmt.Println("3.Print the rate list")
	fmt.Println("4.Show the menu")
	fmt.Println("5.Exit")
}
