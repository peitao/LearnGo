package main

import "fmt"
//这个程序的执行流程是什么啊？主线程不会在select阻塞吗？
//擦，make第2个参数是缓冲区！！！！如果是0，就阻塞了。
func main(){

	ch := make(chan int, 123)
	for {
		select {
			case ch <-0:
			case ch <-1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}
