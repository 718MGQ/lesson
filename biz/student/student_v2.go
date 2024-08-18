package student

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GenderTypeV2 int

const (
	GenderMaleGenderTypeV2 GenderTypeV2 = iota + 1
	GenderFemaleGenderTypeV2
)

type StudentV2 struct {
	id     int    // id默认自增的，不允许修改和指定
	name   string // 姓名创建后不允许随意修改，只能通过SetStudentName方式修改
	Age    int    // 年龄公开
	gender GenderTypeV2
}

func (s *StudentV2) GetStudentName() string {
	return s.name
}

//go:embed student_num.txt
var studentNumV2 string

func NewStudentV2(n string, age int, g GenderTypeV2) *StudentV2 {
	i, err := strconv.Atoi(studentNumV2)
	if err != nil {
		panic(err)
	}
	i++

	// 每次生成的学号写回文件，保证了全局自增，进程的退出也不影响
	defer func() {
		studentNumV2 = strconv.Itoa(i)
		err := os.WriteFile("./biz/student/student_num.txt", []byte(studentNumV2), 0666)
		if err != nil {
			fmt.Println(err)
		}
	}()

	return &StudentV2{
		id:     i,
		name:   n,
		Age:    age,
		gender: g,
	}
}

func (s *StudentV2) SetStudentName(name string) error {
	// 名字字符长度不允许超过5，且不允许少于1
	// ！注意这里是字符的长度，并非字数，字符与字数的区别后续学习
	if len(name) < 1 || len(name) > 20 {
		// 不符合名称规范，不允许修改，直接忽略返回
		return errors.New("student name too long")
	}

	// 名字中不允许含有*号
	// strings.Contains用来判断一个字符串是否包含另一个子串，strings包相关函数后续学习
	if strings.Contains(name, "*") {
		// 不符合名称规范，不允许修改，直接忽略返回
		return errors.New("student name must not contain *")
	}

	s.name = name
	return nil
}

// ResetStudentNum 重置学号
func ResetStudentNum() error {
	return os.WriteFile("./biz/conf/student_num.txt", []byte("0"), 0666)
}
