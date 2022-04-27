package main

import (
	"fmt"
)

func main() {
	var n int
	// считываем числа пока не будет введен 0
	fmt.Scan(&n)
	if n/100000+(n/10000)%10+(n%10000)/1000 == (n/100)%10+(n%100)/10+n%10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
