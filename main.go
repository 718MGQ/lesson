package main

import (
	"fmt"
	"lesson/biz"
)

const (
	Pending  = 0
	Approved = 1
)

var name string

type InstanceStatus int

const (
	InstanceStatusPending  InstanceStatus = 0
	InstanceStatusApproved InstanceStatus = 1
)

type User struct {
	name string
	any  int
}

type inter interface {
}

func main() {
	fmt.Println("Add函数调用结果: ", biz.Add(1, 3))
	fmt.Println() // 空一行，为了好看，没有作用

	// new student1
	student1 := biz.NewStudent("狗剩", 18, biz.GenderMale)
	student2 := biz.NewStudent("钢蛋", 20, biz.GenderFemale)

	fmt.Printf("student1 information: %+v\n", student1)
	fmt.Printf("student2 information: %+v\n", student2)

	// 这里想直接改名是不行的
	// student1.name

	fmt.Printf("\n===========================> 开始改名\n\n")

	student1 = biz.SetStudentName(student1, "狗剩改名了")
	student2 = biz.SetStudentName(student2, "钢蛋*也想改名") // FAIL

	fmt.Println("student1 name: ", biz.GetStudentName(student1))
	fmt.Println("student2 name: ", biz.GetStudentName(student2))

	// 狗剩成功改名重新做人，但是钢蛋失败了
}

// study git
// repr
