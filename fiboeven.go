package main

import "fmt"

func fiboeven() (int, int) {
	var x int = 1
	var y int = 2
	var z int
	var max int = 4000001
	var sum int = 2
	fmt.Printf("Start with %v --> %v", y, sum)
	for (x<max && y<max && z<max) {
		z = x+y
		fmt.Printf("x-->%d y-->%d z-->%d\n", x,y,z)
		if z % 2 == 0 {
			sum += z
			fmt.Printf("%v --> %v\n", x, sum)
		}
		x = y
		y = z
	}
	return sum, max
}

func main() {

	result, maximum := fiboeven()
	fmt.Printf("Total of all even Fibonacci numbers less than %v is %v.", maximum, result)
}

		
	
