package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(foot string) {
	fmt.Printf("%v吃%v\n", c.name, foot)
}

type chicken struct {
	name string
	feet int8
}

func (c chicken) move() {
	fmt.Println("走鸡步")
}

func (c chicken) eat(foot string) {
	fmt.Printf("%v吃%v\n", c.name, foot)
}

func main() {
	var a1 animal
	var c1 cat
	c1.name = "Tom"
	c1.feet = 4
	a1 = c1
	a1.eat("小鱼干")
	var c2 chicken
	c2.name = "小G"
	c2.feet = 2
	a1 = c2
	a1.eat("小石头")
	fmt.Printf("%T", a1)
}
