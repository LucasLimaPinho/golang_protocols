package main

import (
	"fmt"
	"math"
)

type Point struct{
	x float64
	y float64
}

func (p *Point) DistToOrigin() float64 {
	t := math.Pow(p.x,2) + math.Pow(p.y,2)
	return math.Sqrt(t)
}

func main(){
	p1 := Point{3,4}
	fmt.Println(p1)
	fmt.Println(p1.DistToOrigin())
	fmt.Println(p1)
}

// With Pointer Reference, podemos contornar o problema das limitações dos métodos

// 1. Muito tempo para copiar a CHAMA POR VALOR quando temos valores a passar muito grande para os métodos;
// 2. Métodos não podem transformar dados dentro do receiver por se tratarem de cópias
// Podemos passar os ponteiros em uma CHAMADA POR REFERÊNCIA que estas limitações dos métodos são contornadas.

// Dereferencing is automatic with dot (.) operator 
// DO not need to reference when calling the method 
// Good Programming Practice
//////////1. All methods for a type have pointer receivers;
//////////2. All methods for a type have non-pointer receivers;
///////// Mixing pointer/non-pointer receivers will get confusing - POinter receivers allows modifications.
