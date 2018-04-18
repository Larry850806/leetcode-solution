package main

import (
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var x1, x2, v1, v2 int32
	fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)

	dx := x2 - x1
	dv := v1 - v2

	if dv <= 0 {
		fmt.Println("NO")
		return
	}

	if dx%dv == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
