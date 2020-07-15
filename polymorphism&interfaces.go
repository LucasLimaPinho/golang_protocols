// Ability for a object to have different forms depending on the context.
// Example: one method does one thing for one kind of object and another thing for another kind of object. 
// Example: Area
// Area -> Rectangle -> Area = base * height
// Area -> Triangle -> Area = 0.5 * base * height
// Identical at a HIGH LEVEL of abastraction - they compute the area, don't matter the object. 
// Different at a low level of abstraction

// GOLANG DOES NOT INHERITANCE - it is used to implement polymorphism in other POO languages.
// In inherticance, subclasses inherits the methods/data of the superclass. Subclass extends superclass.
// OVERRIDING: Subclass redefines a method inherited from the superclass. 

// INTERFACES helps us et POLYMORPHISMI in GO - we dont have inheritance and overriding in GO. We need to have interfaces.
// INTERFACE is a SET OF METHOD SIGNATURES! There is no implementation of methods in interfaces. Just say that the method has this names, this parameters, this returns.
// Used to express conceptual similarity between types.
// Type satisfies an interface if type defines all methods specificied in the interface.

package main

import "fmt"

type Speaker interface {
	Speak()
}
type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name)
}

func main(){
	var s1 Speaker
	d1 := Dog{"Brian"}
	s1 = d1 // is legal because Dog satisfies the Speaker interface. Interface type assigned to a concrete type
	s1.Speak()
}

