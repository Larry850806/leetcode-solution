package main

import (
	"fmt"
	"strconv"
)

type bigInt []int

func newBigInt(n int) (x bigInt) {
	// 321 -> [1, 2, 3]
	for n > 0 {
		x = append(x, n%10)
		n /= 10
	}
	return
}

func (n bigInt) String() (s string) {
	// [1, 2, 3] -> 321
	for _, e := range n {
		s = strconv.Itoa(e) + s
	}
	return
}

func (n bigInt) mul(x bigInt) bigInt {
	res := make([]int, len(n)+len(x))

	// 18 x 18 -> [8, 1] x [1, 8]
	// [8, 1] x [8, 1] -> [64, 16, 1, 0]
	for i := range n {
		for j := range x {
			res[i+j] += n[i] * x[j]
		}
	}

	// [64, 16, 1, 0] -> [4, 22, 1, 0]
	// [4, 22, 1, 0] -> [4, 2, 3, 0](0324)
	for i := range res {
		if res[i] > 9 {
			res[i+1] += res[i] / 10
			res[i] %= 10
		}
	}

	// [4, 2, 3, 0](0324) -> [4, 2, 3](324)
	for res[len(res)-1] == 0 {
		res = res[0 : len(res)-1]
	}

	return bigInt(res)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	res := newBigInt(1)
	for i := 2; i <= n; i++ {
		x := newBigInt(i)
		res = res.mul(x)
	}

	fmt.Println(res)
}
