package student

import (
	"strings"
)

type GenderType string

const (
	GenderMale   GenderType = "male"
	GenderFemale GenderType = "female"
)

type student struct {
	id     int    // id默认自增的，不允许修改和指定
	name   string // 姓名创建后不允许随意修改，只能通过SetStudentName方式修改
	Age    int    // 年龄公开
	gender GenderType
}

func GetStudentName(s student) string {
	return s.name
}

var studentNum int

// 以下方法初步学习，存在瑕疵，后续再回头优化

func NewStudent(n string, age int, g GenderType) student {
	studentNum = studentNum + 1
	return student{
		id:     studentNum,
		name:   n,
		Age:    age,
		gender: g,
	}
}

func SetStudentName(s student, name string) student {
	// 名字字符长度不允许超过5，且不允许少于1
	// ！注意这里是字符的长度，并非字数，字符与字数的区别后续学习
	if len(name) < 1 || len(name) > 20 {
		// 不符合名称规范，不允许修改，直接忽略返回
		return s
	}

	// 名字中不允许含有*号
	// strings.Contains用来判断一个字符串是否包含另一个子串，strings包相关函数后续学习
	if strings.Contains(name, "*") {
		// 不符合名称规范，不允许修改，直接忽略返回
		return s
	}

	s.name = name
	return s
}

/*
要点1 首字母小写的变量，不允许直接修改和访问。
要点2 go语言中通过自定义类型实现枚举，没有天生的枚举类型。
要点3 注意studentNum值的设置，实现了单次运行中的自增的效果（非最佳实践）。
*/
