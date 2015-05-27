// struts.go
package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
}

// We can put more structs inside an struct
type Developer struct {
	human Human
	yearsOfExperience int
}

// Create a main struct type
type Rectangle struct {
	width, height float64
}

// We can attach to that type specific methods to them
func (r Rectangle) area () float64 {
	return r.width * r.height
}

func main() {
	var me Developer
	
	me.human.name = "Sergio"
	me.human.age = 28
	
	fmt.Printf("My name is %s and I am %d years old", me.human.name, me.human.age)
	
	var rectangle Rectangle
	
	rectangle.height = 12.34
	rectangle.width = 20.45
	
	fmt.Println(rectangle.area())
}
