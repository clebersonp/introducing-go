package main

import "fmt"

type Prepareble interface {
	Prepare()
}

type Tea struct{}

func (t *Tea) addWater() {
	println("Adding Water")
}

func (t *Tea) addTeaPowder() {
	println("Adding Tea Powder")
}

func (t *Tea) addSugar() {
	println("Adding Sugar")
}

func (t *Tea) addMilk() {
	println("Adding Milk")
}

func (t *Tea) heatUp() {
	println("Heating up")
}

type GreenTea struct {
	Tea
}

func (t *GreenTea) addTeaPowder() {
	println("Adding Green Tea Powder")
}

func (t *GreenTea) Prepare() {
	t.addWater()
	t.addTeaPowder()
	t.addSugar()
	t.heatUp()
	fmt.Println("--------------------")
	fmt.Println("--------------------")
}

func (t *Tea) Prepare() {
	t.addWater()
	t.addTeaPowder()
	t.addSugar()
	t.addMilk()
	t.heatUp()
	fmt.Println("--------------------")
	fmt.Println("--------------------")
}

func main() {
	var tea Prepareble = &GreenTea{}
	prepare(tea)

	tea = &Tea{}
	prepare(tea)
}

func prepare(p Prepareble) {
	fmt.Printf("Typeof: %T\n", p)
	p.Prepare()
}
