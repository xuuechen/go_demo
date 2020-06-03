////package main
////
////import (
////	"fmt"
////	"strings"
////)
////
//////2020 05/31
////
////func main() {
////	//切片
////	//切片底层就是数组
////	//切片有长度和容量的说法。容量指的是底层数组的cap
////
////	//基于数组创造切片
////	arr := [5]int{1, 2, 3, 4, 5}
////	s1 := arr[1:3]
////	fmt.Println(s1)
////	//因此对于切片的修改会引起底层数组的变化
////	s1[0] = 10
////	fmt.Println(arr)
////	fmt.Println("len: ", len(s1), "cap: ", cap(s1))
////
////	fmt.Println("------------")
////	//基于make()函数创造切片
////	s2 := make([]int, 2, 10) //len:2 cap:10
////	fmt.Println(s2)
////	fmt.Println(s2 == nil) //判断切片是否为空用len(),不能用nil
////	fmt.Println(len(s2) == 0)
////
////	s3 := make([]int, 0, 10)
////	fmt.Println(s3, s3 == nil, len(s3) == 0)
////
////	fmt.Println("----------")
////	//使用append()为切片添加元素
////	var s4 []int
////	for i := 0; i < 10; i++ {
////		s4 = append(s4, i)
////		fmt.Printf("len:%d cap:%d ptr:%p s4[0]ptr:%p value:%v\n", len(s4), cap(s4), s4, &s4[0], s4)
////	}
////	//以上结果表明，要是长度超过容量，那么扩容成原来的两倍
////	fmt.Println("-----------------------")
////	//copy()
////	s5 := make([]int, 5, 5)
////	s6 := []int{1, 2, 3, 4, 5}
////	copy(s5, s6)
////	s6[0] = 100
////	fmt.Println(s5, s6)
////	fmt.Println("-------------------------")
////	//Go语言中没有对切片的删除操作，可以使用切片本身的特性来完成
////	//删除index为3的元素
////	s6 = append(s6[:3], s6[4:]...)
////	fmt.Println(s6)
////
////	fmt.Println("--------------------------------")
////	m1 := make(map[int]string, 10)
////	m1[0] = "shanghai"
////	m1[3] = "beijing"
////	fmt.Printf("%p %p %v\n", m1, m1[0], m1)
////
////	m2 := map[int]string{
////		0: "xc",
////		1: "xuechen",
////	}
////	fmt.Println(m2)
////
////	if _, ok := m2[2]; ok {
////		fmt.Println("ok")
////	} else {
////		fmt.Println("no")
////	}
////	//注：map是无序的，插入与遍历顺序没多大关系
////	//map有delete函数
////	delete(m2, 0)
////
////	//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
////	str := "hello xc hello xuechen xuechen how do you do zw wjl zw"
////	split := strings.Split(str, " ")
////	//fmt.Printf("%T",split)
////	words := make(map[string]int, 10)
////	for _, v := range split {
////		if _, ok := words[v]; ok {
////			words[v] += 1
////		} else {
////			words[v] = 1
////		}
////	}
////	fmt.Println(words)
////
////	fmt.Println("--------------------")
////
////	intSum(1)
////	intSum(1, 2, 3, 4)
////
////	fmt.Println("---------------")
////
////	type myFunc func(x, y int) int
////	var m myFunc
////	m = add
////	fmt.Printf("m Type: %T,m: %v\n", m, m(1, 2))
////
////	fmt.Println(operator(2, 3, add))
////	//匿名函数
////	f := func(x, y int) {
////		fmt.Println(x + y)
////	}
////	f(10,20)
////
////	//自执行函数
////	func(x,y int){
////		fmt.Println(x * y)
////	}(10,20)
////
////	fmt.Println("-------------------------------")
////	f2 := join("test")
////	fmt.Printf("f2 type: %T,value: %v\n",f2,f2)
////	full1 := f2(".txt")
////	full2 := f2(".exe")
////	fmt.Println(full1,full2)
////}
////
//////函数
//////Go语言中，函数作为一等公民
//////1.当作参数传递 2.返回值
////func intSum(x ...int) {
////	for _, v := range x {
////		fmt.Println(v)
////	}
////}
////
////func operator(x, y int, op func(a, b int) int) int {
////	return add(x, y)
////}
////
////func add(x, y int) int {
////	return x + y
////}
//////闭包
////func join(suffix string) func(name string) string{
////	return func(name string) string {
////		return suffix + name
////	}
////
////}
//package main
//
//import "fmt"
//
//func f1() int {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//
//func f2() (x int) {
//	defer func() {
//		x++
//	}()
//	return 5
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//func f4() (x int) {
//	defer func(x int) {
//		x++
//	}(x)
//	return 5
//}
//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//}

package main

import "fmt"

//
//import "fmt"
//
//func calc(index string, a, b int) int {
//	ret := a + b
//	fmt.Println(index, a, b, ret)
//	return ret
//}
//
//func main() {
//	x := 1
//	y := 2
//	defer calc("AA", x, calc("A", x, y))
//	x = 10
//	defer calc("BB", x, calc("B", x, y))
//	y = 20
//}

func fn1(){
	//defer 必须定义在可能会发生panic的代码之上
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("recover")
		}

	}()
	panic("this is a panic")
}
func main() {
	fn1()
	fmt.Println("go on")


	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int)

	b["xc"] = 100
	fmt.Println(b,len(b))

}