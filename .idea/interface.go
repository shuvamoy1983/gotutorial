package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}


type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}


func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}


type Men interface {
	SayHi()

}


func main() {
		mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	        paul := Employee{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}

	// define interface i
	var i Men
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()


	i=paul
	fmt.Println("This is Paul, an Employee:")
	i.SayHi()


	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 2)

	x[0], x[1] = paul, mike

	for _, value := range x {
		value.SayHi()
	}
}
