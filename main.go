package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	name string
}

func (person *Person) talk(event int) string {
	return fmt.Sprintf("%s says hi (%d).\n", (*person).name, event)
}

func (person *Person) like(event int) string {
	return fmt.Sprintf("%s likes that statement (%d).\n", (*person).name, event)
}

func (person *Person) share(event int) string {
	return fmt.Sprintf("%s shares that feeling (%d).\n", (*person).name, event)
}

func react(queue chan int, people *[]Person) {
	for {
		event := <-queue
		time.Sleep(time.Duration((rand.Intn(3))) * time.Second)
		person := (*people)[rand.Intn(len(*people))]
		coin := rand.Float32() < 0.5
		if coin {
			fmt.Println(person.like(event))
		} else {
			fmt.Println(person.share(event))
		}
	}
}

func main() {
	queue := make(chan int)
	rand.Seed(1337)
	people := []Person{
		Person{name: "Hans-Peter"},
		Person{name: "Karl-Heinz"},
		Person{name: "Mart-Lejeu"},
	}
	go react(queue, &people)

	for event := 0; ; event++ {
		person := people[rand.Intn(len(people))]
		fmt.Printf(person.talk(event))
		queue <- event
		time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)
	}

}
