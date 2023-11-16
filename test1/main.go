package main

import (
	"fmt"
	"sync"
)

func workerSumSquire(slice []int) int {
	wg := sync.WaitGroup{}
	taskChan := make(chan int)
	countChan := make(chan int)

	go func() {
		for i := 0; i < len(slice); i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				value := <-taskChan
				countChan <- value * value
			}()
		}
		wg.Wait()
		close(countChan)
	}()

	for _, v := range slice {
		taskChan <- v
	}

	sum := 0
	for value := range countChan {
		sum += value
	}

	return sum
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sum := workerSumSquire(slice)
	fmt.Println(sum)
}
