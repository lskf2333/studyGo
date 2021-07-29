package main

import "fmt"

type student struct {
	id   int
	name string
}

type students map[int]*student

// var allStudents students

func (s students) showAllStudent() {
	for k, v := range s {
		fmt.Printf("学号：%d，姓名：%v\n", k, v.name)
	}
}

func (s students) addStudent() {
	var (
		id   int
		name string
		s1   student
	)
	fmt.Print("请输入学生的学号：")
	// inputId:
	fmt.Scanln(&id)
	// if _, ok := s[id]; ok {
	// 	fmt.Printf("学号%d已经存在，请重新输入：", id)
	// 	goto inputId
	// }
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	s1.id = id
	s1.name = name
	s[id] = &s1
}

func (s students) deleStudent() {
	var deleId int
	fmt.Print("请输入需要删除学生的学号：")
inputDeleId:
	fmt.Scanln(&deleId)
	if _, ok := s[deleId]; ok {
		delete(s, deleId)
	} else {
		fmt.Print("该学号不存在，请重新输入需要删除学生的学号：")
		goto inputDeleId
	}
}

func main() {
	var choice int //操作的编号
	var s1 = make(students, 20)
	fmt.Println("欢迎使用学生管理系统")
	fmt.Println(`
	1.查看所有学生
	2.新增学生
	3.删除学生
	4.退出
	`)
inputChoice:
	fmt.Print("请输入您选择操作的编号：")
	fmt.Scanln(&choice)
	fmt.Printf("您选择的是%d!!!\n", choice)
	switch choice {
	case 1:
		s1.showAllStudent()
		goto inputChoice
	case 2:
		s1.addStudent()
		goto inputChoice
	case 3:
		s1.deleStudent()
		goto inputChoice
	case 4:
		fmt.Println("退出")
	default:
		fmt.Println("gck")
		goto inputChoice
	}

}
