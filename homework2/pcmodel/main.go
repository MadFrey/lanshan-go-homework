/**
 * @Author: lrc
 * @Date: 2022/7/21-14:44
 * @Desc:生产者消费者模型
 **/

package main

import (
	"fmt"
)

var channel = make(chan int)

func main() {
	go func() {
		for i := 0; i < 1000; i++ {
			channel <- i
		}
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}
}
