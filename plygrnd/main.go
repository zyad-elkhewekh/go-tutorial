package main

import (
	"errors"
	"fmt"

	//"sync"
	"math"
	"math/rand/v2"
)

func main() {
	fmt.Println("hello world")
	var varName1 int = 0
	var attempts int = 9
	for attempts > 0 {
		//start program
		var n, m, x, y, a, b int
		n = 9
		m = 9
		x = rand.IntN(n)
		y = rand.IntN(m)

		//target point = (x, y)
		//user start point = (a, b)

		fmt.Print("Enter where you want to start (a, b): ")
		fmt.Scanln(&a, &b)
		var dist float64 = math.Abs(float64(a)-float64(x)) + math.Abs(float64(b)-float64(y))

		if a == x && y == b {
			y = (y + m) % (m - 1)
		}
		//reset target if coincides with user start

		var vi, vj rune
		fmt.Print("enter direction of motion (+ or -): ")
		fmt.Scanln(&vi, &vj)

		if vi == '+' {
			a = (a + 1 + n) % n
		} else {
			a = (a - 1 + n) % n
		}

		if vj == '+' {
			b = (b + 1 + n) % n
		} else {
			b = (b - 1 + n) % n
		}

		//user step

		var path float64 = math.Abs(float64(a)-float64(x)) + math.Abs(float64(b)-float64(y))
		if path-dist == 0 {
			fmt.Println("success!")
			break
		} else if path-dist > 0 {
			fmt.Println("colder")
		} else {
			fmt.Println("warmer")
		}
		attempts--
	}

	//fmt.Print("have to use variables so\n")
	//fmt.Println(varName1)
	varName1 = varName1 + 1
	//fmt.Println(varName1)

	//var redundency int = incr(varName1)
	//fmt.Println(redundency)

	//var ret int
	//var err error
	//ret, err = div(0)
	//fmt.Println(ret, err)

	// var c = make(chan int, 10)
	// go proccess(c)
	// for i := range c {
	// 	fmt.Println(i)
	// }
}

func incr(n int) int {
	fmt.Println(1 + n)
	return n
}
func div(n int) (int, error) {
	var err error
	if n == 0 {
		err = errors.New("no div by 0 pls")
		return n, err
	}
	fmt.Println(1 / n)
	return n, err
}
func proccess(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	fmt.Println("exit proc func")
}
