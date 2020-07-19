package main

import (
	"fmt"
	"sync"
)

func setup(){
	fmt.Println("Hello Once Setup")
}

func dostuff(wg *sync.WaitGroup, on *sync.Once){

	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}

func main(){

	var on sync.Once
	var wg sync.WaitGroup
	wg.Add(2)
	go dostuff(&wg, &on)
	go dostuff(&wg, &on)
	wg.Wait()
}
