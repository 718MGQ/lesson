package main

import (
	"fmt"
	"lesson/biz"
	"lesson/biz/ptr"
	"lesson/biz/student"
	"strconv"
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
	student1 := student.NewStudent("狗剩", 18, student.GenderMale)
	student2 := student.NewStudentV2("钢蛋", 20, student.GenderFemaleGenderTypeV2)

	fmt.Printf("student1 information: %+v\n", student1)
	fmt.Printf("student2 information: %+v\n", student2)

	// 这里想直接改名是不行的
	// student1.name

	fmt.Printf("\n===========================> 开始改名\n\n")

	student1 = student.SetStudentName(student1, "狗剩改名了")
	err := student2.SetStudentName("钢蛋也想改名")
	if err != nil {
		fmt.Println("钢蛋翻身失败了！")
	}

	fmt.Println("student1 name: ", student.GetStudentName(student1))
	fmt.Println("student2 name: ", student2.GetStudentName())

	// 狗剩成功改名重新做人，但是钢蛋失败了

	// 连续生成多个student V2测试一下
	for i := 0; i < 10; i++ {
		fmt.Println(student.NewStudentV2("test"+strconv.Itoa(i), 19, student.GenderFemaleGenderTypeV2))
	}

	p := 1
	pp := &p
	ptr.ChangeA(pp)
	fmt.Println(p)
}
