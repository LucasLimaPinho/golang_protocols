package main

import "sync"

func dostuff(c1 chan int, c2 chan int, wg *sync.WaitGroup){
	<- c1
	c2 <- 1
	wg.Done()
}

func main(){

	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go dostuff(ch1,ch2,&wg)
	go dostuff(ch2,ch1,&wg) // swapped argument to cause deadlock
	wg.Wait()
}
