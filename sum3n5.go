// Sum of all numbers divisable by 3 or 5 
// Solves https://projecteuler.net/problem=1
package main

import "fmt"

func main() {
	var sum int
	var flag bool
	for x:=3; x<1000; x++ {
		flag = false
		switch {
			case x % 3 == 0:
				flag = true
			case x % 5 == 0:
				flag = true
		}
		if flag == true {
			sum += x
			fmt.Printf("%d --> %d\n", x, sum)
		}
	}
	
	fmt.Printf("The total amount of those numbers is %d.\n", sum)
}
