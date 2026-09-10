package main

import "fmt"


type HasName interface{
	GetName() string
}

func sayHello(value HasName){
	fmt.Println("hello ", value.GetName())
}

type Person struct{
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

type Animal struct{
	Name string
}

func (animal Animal) GetName() string{
	return animal.Name
}

func main() {
	fmt.Println("Hello, World!")
	person := Person{Name: "Tegar"}
	sayHello(person)

	animal := Animal{Name: "Dog"}
	sayHello(animal)

}