package main

import "fmt"

//用面向对象的方法来写一个简单的学生管理系统
//主要功能：增删改查
//2020/0604
var studentMan StudentMan

type StudentMan struct {
	student []*Student
}

type Student struct {
	Id     int
	Name   string
	Age    int
	Gender string
}

//构造方法
func NewStudent(id int, name string, age int, gender string) *Student {
	return &Student{
		Id:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}

//查看学生信息
func (s *StudentMan) GetStudentInfo() {
	for _, v := range studentMan.student {
		fmt.Printf("stu info: %v ", *v)
	}
	fmt.Println()
}

//添加学生
func (s *StudentMan) addStudent(stu *Student) {
	s.student = append(s.student, stu)
}

//删除学生信息
func (s *StudentMan) deleteStu(index int) {
	s.student = append(s.student[:index], s.student[index+1:]...)
}

//修改学生信息
func (s *StudentMan) updateStu(stuId int, stu *Student) {
	for i, v := range s.student {
		if stuId == v.Id {
			s.student[i] = stu
		}
	}
}

func main() {

	studentMan = StudentMan{[]*Student{}}
	//添加学生信息
	fmt.Println("添加学生信息")
	studentMan.addStudent(NewStudent(1, "xiaojia", 24, "male"))
	studentMan.addStudent(NewStudent(2, "xioaoyi", 15, "female"))
	studentMan.addStudent(NewStudent(3, "xiaobin", 23, "male"))
	studentMan.addStudent(NewStudent(4, "xiaoding", 24, "female"))

	//查看学生信息
	studentMan.GetStudentInfo()
	//删除学生信息
	//studentMan.deleteStu(0)
	//studentMan.GetStudentInfo()
	//修改学生信息
	studentMan.updateStu(1, &Student{
		Id:     1,
		Name:   "xuechen",
		Age:    22,
		Gender: "female",
	})

	studentMan.GetStudentInfo()
}
