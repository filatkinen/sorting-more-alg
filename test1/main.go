package main

import "fmt"

func SumOfSquares(c, quit chan int) {
	for{
		select {
		case :
			
		}
	}
}

func main() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 0; i < 6; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
	}()
	SumOfSquares(mychannel, quitchannel)
}
