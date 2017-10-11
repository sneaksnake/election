package main

import "fmt"

type Person struct {
	name string
}

func (person *Person) talk() {
	fmt.Printf("%s says hi.\n", (*person).name)
}

func main() {
	p1 := Person{name: "Hans-Peter"}
	p2 := Person{name: "Karl-Heinz"}
	p1.talk()
	p2.talk()

}
