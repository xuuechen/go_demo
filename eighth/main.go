package main

import "fmt"

//2020 06/02
//自定义类型
type myInt int
//结构体匿名字段

type person struct{
	uint8
	int
}

//结构体的嵌套
type address struct{
	name string
	province string
	city	string
}

type user struct{
	name string
	address address
}
//结构体嵌套
type animal struct{
	feet uint8
}

type dog struct{
	name string
	*animal
}

func (a *animal) run(){
	fmt.Printf("动物用 %d 只脚跑步\n",a.feet)
}

func (d *dog) wang(){
	fmt.Printf("我家的 %s 会汪汪汪\n",d.name)
}
func main() {
	var i myInt
	fmt.Printf("type: %T\n", i)
	result := i.add(10,20)
	fmt.Println(result)
	fmt.Println(person{
		10,
		 20, //由于匿名字段采用类型名作为字段名，因此同一个结构体中不能有重复的类型
	})

	var u user
	u.address.city = "shanghai"
	u.address.name = "test"
	fmt.Println(u)

	d := dog{
		name:   "ww",
		animal: &animal{feet: 4},
	}

	d.run()
	d.wang()

}

//为自定义类型添加方法
func (m myInt) add(x, y int) int {
	return x + y
}
