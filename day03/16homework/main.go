package main

import "fmt"

/*
你有50枚金币，需要分配给一下几个人：Matthew，Sarah，Augustus，Heidi，Emilie，Peter，Giana，Adriano，Aaron，Elizabeth：
分配规则如下：
a. 名字中每包含一个'e'或'E'分1枚金币
b. 名字中每包含一个'i'或'I'分2枚金币
c. 名字中每包含一个'o'或'O'分3枚金币
d. 名字中每包含一个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现'dispatchCoin'函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, i := range users {
		for _, j := range i {
			// if string(j) == "e" || string(j) == "E" {
			// 	if coins >= 1 {
			// 		distribution[i] += 1
			// 		coins -= 1
			// 	}
			// } else if string(j) == "i" || string(j) == "I" {
			// 	if coins >= 2 {
			// 		distribution[i] += 2
			// 		coins -= 2
			// 	}
			// } else if string(j) == "o" || string(j) == "O" {
			// 	if coins >= 3 {
			// 		distribution[i] += 3
			// 		coins -= 3
			// 	}
			// } else if string(j) == "u" || string(j) == "U" {
			// 	if coins >= 4 {
			// 		distribution[i] += 4
			// 		coins -= 4
			// 	}
			// }
			switch j {
			case 'e', 'E':
				if coins >= 1 {
					distribution[i] += 1
					coins -= 1
				}
			case 'i', 'I':
				if coins >= 2 {
					distribution[i] += 2
					coins -= 2
				}
			case 'o', 'O':
				if coins >= 3 {
					distribution[i] += 3
					coins -= 3
				}
			case 'u', 'U':
				if coins >= 4 {
					distribution[i] += 4
					coins -= 4
				}
			}
		}
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}
