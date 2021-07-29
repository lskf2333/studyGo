package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hoddy  []string
}

func main() {
	var xixi person
	xixi.name = "xixi"
	xixi.age = 18
	xixi.gender = "男"
	xixi.hoddy = []string{"篮球", "游戏", "鼓掌"}
	fmt.Println(xixi)
}
