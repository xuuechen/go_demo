package main

import "fmt"

//2020 06/02

func main() {
	//arr range
	arr := [3]int{3, 4, 5}
	for i, v := range arr {
		if i == 1 {
			v = 10
		}
		fmt.Printf("%d ", v)
	}
	fmt.Println(arr)

	//slice range
	// The loop we generate:
	//   for_temp := range
	//   len_temp := len(for_temp)
	//   for index_temp = 0; index_temp < len_temp; index_temp++ {
	//           value_temp = for_temp[index_temp]
	//           index = index_temp
	//           value = value_temp
	//           original body
	//   }

	s := []int{10, 11, 12}
	for i, v := range s {
		if i == 1 {
			v = 100
		}
		fmt.Printf("%d ", v)
	}
	fmt.Println(s)

	s1 := make([]*int, 3)
	for i, v := range s {
		s1[i] = &v
		//fmt.Printf("%v ",v)
	}

	for _, v := range s1 {
		fmt.Printf("%d ", *v)
	}
	fmt.Println()
	fmt.Println(s1)

	s2 := s
	for i, v := range s2 {
		fmt.Printf("%v %p ", i, &v)
	}
	fmt.Println()
	fmt.Println("-------------------------")
	for i, v := range s {
		fmt.Printf("%v %p ", i, &v)
	}

}
