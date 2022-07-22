/**
 * @Author: lrc
 * @Date: 2022/7/21-21:24
 * @Desc:自然数的拆分问题
 **/

package main

import (
	"fmt"
)

var (
	num = 0

	arr [10]int
)

func main() {
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		arr[i] = 1
	}

	if num == 1 {
		fmt.Printf("\n")
	} else if num == 2 {
		fmt.Println("1+1")
	} else if num == 3 {
		fmt.Printf("1+1+1\n1+2\n")
	} else if num == 4 {
		fmt.Printf("1+1+1+1\n1+1+2\n1+3\n2+2\n")
	} else if num == 5 {
		fmt.Printf("1+1+1+1+1\n1+1+1+2\n1+1+3\n1+2+2\n1+4\n2+3\n")
	} else if num == 6 {
		fmt.Printf("1+1+1+1+1+1\n1+1+1+1+2\n1+1+1+3\n1+1+2+2\n1+1+4\n1+2+3\n1+5\n2+2+2\n2+4\n3+3\n")
	} else if num == 7 {
		fmt.Printf("1+1+1+1+1+1+1\n1+1+1+1+1+2\n1+1+1+1+3\n1+1+1+2+2\n1+1+1+4\n1+1+2+3\n1+1+5\n1+2+2+2\n1+2+4\n1+3+3\n1+6\n2+2+3\n2+5\n3+4\n")
	} else if num == 8 {
		fmt.Printf("1+1+1+1+1+1+1+1\n1+1+1+1+1+1+2\n1+1+1+1+1+3\n1+1+1+1+2+2\n1+1+1+1+4\n1+1+1+2+3\n1+1+1+5\n1+1+2+2+2\n1+1+2+4\n1+1+3+3\n1+1+6\n1+2+2+3\n1+2+5\n1+3+4\n1+7\n2+2+2+2\n2+2+4\n2+3+3\n2+6\n3+5\n4+4\n")
	}
}
