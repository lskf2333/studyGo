package main

import "fmt"

func main() {
	//求数组[1, 3, 5, 7, 8]所有元素的和
	a := [5]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)

	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("%v %v\n", i, j)
			}
		}
	}
}
