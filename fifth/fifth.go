package main

import "fmt"

type person struct {
	name string
	age uint8
}
//2020 06/01
func main() {

	//实例化 1.
	var p person
	p.name = "xuechen"
	p.age = 13
	fmt.Println(p)

	//2.
	p1 := person{
		 "xc",
		  20,
	}
	fmt.Println(p1)

	//用new来实例化
	p2 := new(person)
	//因为new()函数返回的是指向person结构体的指针
	//但是Go语言为我们做了封装，比如p2.name其实应该要写(*p2).name
	fmt.Printf("p2 type: %T \n",p2)
	p2.name = "xchen"
	p2.age = 10
	fmt.Printf("p2.Name addr: %p  p2.Age addr: %p p2 addr: %p \n",&(p2.name),&(p2.age),p2)
	fmt.Println((*p2))
	//匿名结构体
	var user struct{
		name string;
		age uint8
	}
	user.name = "xuechen"
	user.age = 20
	fmt.Println(user)

	fmt.Println("------------------------")

	p3 := &person{
		name: "xuechen",
		age:  22,
	}

	fmt.Printf("p3 type: %T\n",p3) // *main.person
	fmt.Println(*p3)


	s := []string{"xc","zw"}
	fmt.Printf("%p %p %p %p \n",&s[0],&s[1],s,&s)
	for _,v:=range s{
		//?????为什么地址都是一样的？？
		fmt.Printf("%p %s ",&v,v)
	}
	fmt.Println()
	for i:=0;i<len(s);i++{
		fmt.Printf("%p %s ",&s[i],s[i])
	}

	var s1= "hello,薛琛"
	for _,v:=range s1{
		fmt.Printf("%c ",v)
	}

	fmt.Println()
	arr := [3]int{1,2,3}
	for i:=0;i<len(arr);i++{
		fmt.Printf("%p %v ",&arr[i],arr[i])
	}
	fmt.Println()
	for _,v:=range arr{
		fmt.Printf("%p %v ",&v,v)
	}
}
