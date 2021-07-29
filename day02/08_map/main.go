package main

import (
	"fmt"
	"strings"
)

//统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func main() {
	// s := "how do you do"
	// m := make(map[string]int, 20)
	// sS := []rune(s)
	// for _, v := range sS {
	// 	m[string(v)]++
	// 	// if _, ok := m[string(v)]; ok {
	// 	// 	m[string(v)]++
	// 	// }else{
	// 	// 	m[string(v)]=1
	// 	// }
	// }
	// fmt.Println(m)
	s := "how do you do"
	m := make(map[string]int, 10)
	sS := strings.Split(s, " ")
	for _, v := range sS {
		m[string(v)]++
	}
	fmt.Println(m)
}
