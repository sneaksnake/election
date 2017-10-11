package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	name string
}

func (person *Person) talk() string {
	return fmt.Sprintf("%s says hi.\n", (*person).name)
}

func main() {
	rand.Seed(1337)
	people := []Person{
		Person{name: "Hans-Peter"},
		Person{name: "Karl-Heinz"},
	}

	for {
		person := people[rand.Intn(len(people))]
		fmt.Printf(person.talk())
		time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)
	}

}
