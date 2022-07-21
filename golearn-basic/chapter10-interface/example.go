package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Run()
}

type Cat struct {
	Name string
	Sex  bool
}

type Dog struct {
	Name string
}

func (d Dog) Run() {
	fmt.Println(d.Name, "start to run")
}
func (d Dog) Eat() {
	fmt.Println(d.Name, "start to eat")
}

func (c Cat) Run() {
	fmt.Println(c.Name, "start to run")
}
func (c Cat) Eat() {
	fmt.Println(c.Name, "start to eat")
}

func main() {
	var a Animal
	c := Cat{
		Name: "tom",
		Sex:  false,
	}
	a = c
	a.Run()
	a.Eat()
	var b Animal
	b = Dog{
		Name: "google",
	}
	b.Run()
	b.Eat()

	generalAct(a)
	generalAct(b)
}

func generalAct(a Animal) {
	a.Run()
	a.Eat()
}
