package main

import "fmt"

func main() {
	var (
		num int32
		str string
	)
	fmt.Scan( &num)
	fmt.Scan(&str)
	lcc1 := []rune(str)
	num = num % 26
	for _, v := range lcc1 {
		if v+num > 122 {
			fmt.Printf("%s",string(v + num - 26))
		} else {
			fmt.Printf("%s",string(v + num))
		}
	}
}
