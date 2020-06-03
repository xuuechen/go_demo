package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Id      string     `json:"id"`
	Student []*Student `json:"student"`
}

type Student struct {
	Name   string
	Age    string
	Gender string
}

func main() {
	//str := "hello,world"
	//bytes, err := json.Marshal(str)
	//if err != nil{
	//	return
	//}
	//fmt.Println(string(bytes))
	//
	//for _,v :=range bytes{
	//	fmt.Printf("%c ",v)
	//}
	//fmt.Println()
	//fmt.Println("--------------------")
	c := &Class{
		Id:      "101",
		Student: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%d", i),
			Age:    fmt.Sprintf("%d", i),
			Gender: "male",
		}
		c.Student = append(c.Student, stu)

	}

	//for _,v := range c.Student{
	//	fmt.Printf("%v ",*v)
	//}

	bytes, err := json.Marshal(c)
	if err != nil {
		panic("marshal failed!")
		return
	}
	fmt.Printf("%s ", bytes)

	fmt.Println()

}
