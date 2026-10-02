package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	defer fmt.Println("【程序结束】")

	for i := 1; i <= 3; i++ {
		if i%2 == 0 {
			fmt.Println(i, "是偶数")
		} else {
			fmt.Println(i, "是奇数")
		}
	}

	fmt.Println("3 + 5 =", add(3, 5))
}
