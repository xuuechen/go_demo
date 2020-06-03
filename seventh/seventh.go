package main

import (
	"fmt"
)

//2020 06/02
type student struct {
	age uint8
	name string
	hobby *[]string
}
func main() {

	s := []int{3,4,5}
	s1 := s
	fmt.Printf("%p %p \n",s,s1)
	map1 := make(map[int]*int,3)
	//var num int
	for i, v := range s{
		num := v
		map1[i] =  &num
		fmt.Printf("%p ",&num)
	}
	fmt.Println()
	for i,v:=range map1{
		fmt.Printf("%d--->%d ",i,*v)
	}
	fmt.Println()

	fmt.Printf("%T \n",s)

	fmt.Println("--------------------")
	m := map[int]*int{}

	for i ,v := range s{
		m[i] = &v
	}

	for _,v :=range m{
		fmt.Printf("%d ",*v) //5 5 5
	}
	fmt.Println()
	fmt.Println("-----------------")
	for i ,_ := range s{
		m[i] = &s[i]
	}

	for _,v :=range m{
		fmt.Printf("%d ",*v) //3 4 5
	}
	fmt.Println()
	fmt.Println("--------------")
	s2 := newStudent(10, "xuechen",&[]string{"ball","game"})
	fmt.Printf("s2 type: %T \n",s2)
	s2.setName("xc")
	fmt.Println(s2.name)

	s2.updateHobby(&[]string{"movies","sleep"})
	fmt.Println(*(s2.hobby))

	test := []int{1,2,3}
	updateSlice(test)
	fmt.Println(test)

}


func newStudent(age uint8,name string,hobby *[]string) *student{
	return &student{
		age,
		name,
		 hobby,
	}
}
func (s *student) setName(name string){
	s.name = name
}

func (s *student) updateHobby(hobbies *[]string){
	s.hobby = hobbies
}

func updateSlice(s []int){
	s[0] = 100
}