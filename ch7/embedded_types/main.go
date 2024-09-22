package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	// relationships like this by using embedded types
	// sometimes also referred to as anonymous fields
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	android := new(Android)
	// we can access the Person's Talk() method using the type name:
	android.Person.Talk()

	// but we can also call any Person methods directly on the Android:
	// The is-a relationship works this way intuitively: people can talk, and android is a person,
	// therefore an android can talk.
	android.Talk()
}
