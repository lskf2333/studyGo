package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := 10
	b := 10.0
	c := true
	d := "xixi"
	fmt.Printf("a_type:%T,a_value:%v\n b_type:%T,b_value:%v\n c_type:%T,c_value:%v\n d_type:%T,d_value:%v \n", a, a, b, b, c, c, d, d)
	e := "hello沙河小丸子"
	count := 0
	for _, c := range e {
		//2.判断字符中是否为汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Printf("字符串%#v含有%v个汉字\n", e, count)
	f := "哈" //string
	g := '哈' //int32(rune)
	fmt.Printf("f_type:%T,g_type:%T", f, g)
}
