package main

import (
	"fmt"
	"os"
)

// func f1() {
// 	fileObj, err := os.Open("./main.go")
// 	defer fileObj.Close() //不能这么写，因为如果err有值的话（读取文件出错），fileObj就为nil  而nil是没有close()的，这样就会报错
// 	if err != nil {
// 		fmt.Printf("open file failed ,err:%v", err)
// 		return
// 	}
// }

func f2() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed ,err:%v", err)
		return
	}
	defer fileObj.Close()
}

func main() {
	// f1()
	f2()
}
