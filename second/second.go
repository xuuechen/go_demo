package main

import (
	"fmt"
	"math"
	"strings"
)

// 5.30日
//常量在声明的时候就必须得定义
const name = "xc"

const (
	age    = 10
	heitht = 178
)

//iota 只能在const中使用：1.首次出现被置为0 2.每增加一行iota便会+1.对于枚举很好用
const (
	a = iota
	b
	_
	c = 100
	d
)

var x = `Hello,
World!
`

func main() {
	fmt.Println(a, b, c, d)
	fmt.Printf("%.10f\n", math.Pi)
	c1 := 1 + 2i
	fmt.Println(c1)
	var a bool
	fmt.Println(a)
	fmt.Println("'sdf'")
	fmt.Println(x)

	//	String: the number of bytes in v.
	//对于字符串，len()返回的是字符串的字节数
	fmt.Println(len("hello,薛琛"))

	//fmt.Sprint()拼接字符串
	fmt.Println(fmt.Sprint(1, "xc"))
	fmt.Println(1, "xc")
	//strings.Split()分割字符串
	fmt.Println(strings.Split("www.baidu.com", "."))
	//strings.Contains()判断是否包含
	fmt.Println(strings.Contains("xue[chen]", "chen"))
	//strings.HasPrefix()判断前缀
	fmt.Println(strings.HasPrefix("38138515052391845", "3"))
	//strings.HasSuffix()判断后缀
	fmt.Println(strings.HasSuffix("12345", "45"))
	fmt.Println(strings.Index("Hello,薛琛", "Hello"))
	//
	fmt.Println(strings.Join([]string{"78", "56"}, "34"))

	//在Go语言中，字符有两种：1.uint8 = byte 2.rune=int32
	//字符串底层其实是数组
	str := "hello,薛琛"
	fmt.Printf("%v\n", []byte(str))

	for i := 0; i < len(str); i++ { //byte
		fmt.Printf("%c ", str[i])
	}

	for _, v := range str { //rune
		fmt.Printf("%c ", v)
	}

	//string类型转换
	str1 := "hello,薛琛"
	//b1 := []byte(str1)
	//
	//b1[6] = 'a'
	//b1[7] = 'b'

	b1 := []rune(str1)
	b1[6] = '世'
	b1[7] = '界'
	fmt.Println(string(b1))

	fmt.Println("-------------")

	fmt.Println(2 << 3)
	fmt.Println(3 << 2)
	fmt.Println(12 >> 2)

	fmt.Println("-------------")

	height := 178
	if height > 190 {
		fmt.Println("too tall")
	} else if height > 175 {
		fmt.Println("just right")
	} else {
		fmt.Println("too short")
	}

	//这种写法的变量作用域只在当前的if作用域中
	if grade := 60; grade > 80 {
		fmt.Println("A")
	} else if grade > 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	fmt.Println("-------------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//Go语言中没有while循环
	//无限循环
	cnt := 0
	for {
		fmt.Println(1)
		cnt++
		if cnt == 10 {
			break
		}
	}
	fmt.Println("-------------")

	value := 10
	switch {
	case value > 10:
		fmt.Println(1)
	case value < 10:
		fmt.Println(2)
	default:
		fmt.Println("not found")
	}

	fmt.Println("-------------")
	//因为break只会跳出当前循环,因此用如下方法能够退出整个循环
	var flag bool
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				flag = true
				break
			}

		}
		if flag {
			fmt.Println("exit")
			break
		}

	}
	//使用goto方法跳出
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				goto breakTag
			}
			fmt.Println(i, j)
		}
	}
breakTag:
	fmt.Println("exit")

forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

	fmt.Println("----------")

	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("跳过", i)
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("-----99乘法表------")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", i, j, i*j)
		}
		fmt.Println()
	}

	fmt.Println("----------")

	//数组
	var arr = [5]int{1,2,3}
	var arr1 = [5]int{1,2,3,0,0}

	fmt.Println(arr == arr1)//true  在Go语言中,长度也是类型的一部分.

	//数组的定义
	//1.
	var array [3]int
	fmt.Println(array)
	//2.
	var array1 = [...]string{"xue","chen"}
	fmt.Println(array1)
	//3.
	array2 := [...]int{1:10,0:3}
	fmt.Println(array2)

	//数组遍历
	//1.
	for i:=0;i<len(array1);i++{
		fmt.Printf("%v ",array1[i])
	}
	fmt.Println()
	//2.
	for _,v := range array1{
		fmt.Printf("%v ",v)
	}
	fmt.Println("-------------")
	//数组是值类型,因此赋值,传参都是 复制 整个数组,因此改变的是副本的值,不会改变本身
	modifyArr := [3]int{10,20,30}
	modifyArray(modifyArr)
	fmt.Println(modifyArr)

	var test *int
	tt := 10
	test = &tt
	*test = 100

	var pa [3]*int
	pa[0] = test

	fmt.Println(tt)
}

func modifyArray(x [3]int)  {
	x[0] = 100
}