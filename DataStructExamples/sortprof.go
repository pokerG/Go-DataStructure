package main

import (
	"Coordinate"
	"GoSort"
	"fmt"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

const MAX int = 100000

var largest int //data <= largest

var (
	cx0 int = -100
	cy0 int = -100
	cx1 int = 1000
	cy1 int = 1000
)

var sortway map[string]func([]int)

func main() {
	sortway = make(map[string]func([]int), 8)

	{
		sortway["BubbleSort"] = GoSort.BubbleSort
		sortway["BucketSort"] = GoSort.BucketSort
		sortway["HeapSort"] = GoSort.HeapSort
		sortway["InsertSort"] = GoSort.InsertSort
		sortway["MergeSort"] = GoSort.MergeSort
		sortway["QuickSort"] = GoSort.QuickSort
		sortway["SelectSort"] = GoSort.SelectSort
		sortway["ShellSort"] = GoSort.ShellSort
	}

	SingleTest()
	BroadContrast(0)
	BroadContrast(1)
	BroadContrast(2)
}

//the argument is test way
func BroadContrast(ty int) {
	fn := testDataSorted
	rd := false
	var imgname string
	switch ty {
	case 0:
		fn = testDataSorted
		imgname = "testDataSorted"
	case 1:
		fn = testDataReversed
		imgname = "testDataReversed"
	case 2:
		fn = testDataRandom
		imgname = "testDataRandom"
		rd = true
	}

	c := Coordinate.NewCoordinateSize(cx0, cy0, cx1, cy1)
	var r Coordinate.Rgba
	for k, v := range sortway {

		node := testTime(v, fn, rd)
		switch k {
		case "BubbleSort":
			r = *Coordinate.NewRgba(0, 0, 255, 255) //blue
		case "BucketSort":
			r = *Coordinate.NewRgba(255, 0, 0, 255) //red
		case "HeapSort":
			r = *Coordinate.NewRgba(0, 255, 0, 255) //green
		case "InsertSort":
			r = *Coordinate.NewRgba(255, 255, 0, 255) //yellow
		case "MergeSort":
			r = *Coordinate.NewRgba(112, 128, 144, 255) //gray
		case "QuickSort":
			r = *Coordinate.NewRgba(255, 0, 255, 255) //purple
		case "SelectSort":
			r = *Coordinate.NewRgba(255, 105, 180, 255) //pink
		case "ShellSort":
			r = *Coordinate.NewRgba(0, 0, 0, 255) //green

		}

		c.FoldLine(node, r)
	}

	imgfile, _ := os.Create(fmt.Sprintf("%s.png", imgname))
	defer imgfile.Close()

	// 以PNG格式保存文件
	err := png.Encode(imgfile, c.Img)
	if err != nil {
		log.Fatal(err)
	}
}

func SingleTest() {

	for k, v := range sortway {
		c := Coordinate.NewCoordinateSize(cx0, cy0, cx1, cy1)
		node := testTime(v, testDataSorted, false)
		r := *Coordinate.NewRgba(0, 0, 255, 255) //blue
		c.FoldLine(node, r)
		node = testTime(v, testDataReversed, false)
		r = *Coordinate.NewRgba(255, 0, 0, 255) //red
		c.FoldLine(node, r)
		largest = 512
		node = testTime(v, testDataRandom, true)
		r = *Coordinate.NewRgba(0, 255, 0, 255) //green
		c.FoldLine(node, r)

		imgfile, _ := os.Create(fmt.Sprintf("%s.png", k))
		defer imgfile.Close()

		// 以PNG格式保存文件
		err := png.Encode(imgfile, c.Img)
		if err != nil {
			log.Fatal(err)
		}
	}

}

//firts argument is sort way second is createtest data rd is testDataRandom
//return is use to draw
func testTime(sort func([]int), td func(int) []int, rd bool) []Coordinate.Node {
	var node []Coordinate.Node
	// fmt.Println(math.Log10(float64(MAX)))
	num := int(math.Log10(float64(MAX)))
	node = make([]Coordinate.Node, num)

	// num = num / int(math.Log10(float64(cx1)))
	// MAX / cx1
	for i := 10; i <= MAX; i = i * 10 {
		if rd {
			// var tq float64
			// for j := 0; j < 5; j++ {
			// 	A := td(i)
			// 	t := time.Now()
			// 	sort(A)
			// 	tq += time.Since(t).Seconds()
			// }
			// tq = tq / 5
			// node[int(math.Log10(float64(i)))-1] = *Coordinate.NewNode(int(math.Log10(float64(i)))*cx1/num, int(tq)) // us
			A := td(i)

			t := time.Now()
			sort(A)
			tq := time.Since(t).Nanoseconds()
			fmt.Println(time.Since(t))
			node[int(math.Log10(float64(i)))-1] = *Coordinate.NewNode(int(math.Log10(float64(i)))*cx1/num, int(tq/100000)) // us
		} else {
			A := td(i)

			t := time.Now()
			sort(A)
			tq := time.Since(t).Nanoseconds()
			fmt.Println(time.Since(t))
			node[int(math.Log10(float64(i)))-1] = *Coordinate.NewNode(int(math.Log10(float64(i)))*cx1/num, int(tq/100000)) // us
		}

	}
	return node
}

func testDataSorted(size int) []int {
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = i
	}
	return input
}

func testDataReversed(size int) []int {
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = size - i - 1
	}
	return input
}

func testDataRandom(size int) []int {
	rand.Seed(time.Now().Unix())
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = rand.Int() % largest
	}
	return input
}
