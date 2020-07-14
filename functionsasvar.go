package main

import "fmt"

var funcVar func(int) int

func incFn(x int) int {
	return x+1
}

// Function can be passed to another function as argument.
func mul (x int) int {
	return x*3
}
func applyIt(afunct func (int) int, val int) int {
	return afunct(val)
}

func incrFn (x int) int {return x+1}
func decrFn (x int) int {return x-1}

func main(){
	funcVar = incFn
	fmt.Println(funcVar(1))
	fmt.Println(funcVar(3))
	fmt.Println(applyIt(mul,3))
	fmt.Println(applyIt(incrFn,3))
	fmt.Println(applyIt(decrFn,3))
	v := applyIt(func(x int) int  {return x+1 },2) //lambda function
	fmt.Println(v)
}

