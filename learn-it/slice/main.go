package main

import "fmt"

func main() {
	var empty_slice []int
	fmt.Println(empty_slice)

	slice_with_values := []int{1, 2, 3, 4, 5}
	fmt.Println(slice_with_values)

	array := [5]int{1, 2, 4, 6, 8}
	slice_from_array := array[1:4]
	fmt.Println(slice_from_array)

	create_slice_with_make := make([]int, 5)
	fmt.Println(create_slice_with_make)

	fmt.Println("append: ", append(empty_slice, 10, 11))

	fmt.Println("copy: ", copy(empty_slice, slice_with_values))

	fmt.Println("Length and Capabity of slice:", len(create_slice_with_make), cap(create_slice_with_make))

	s := []int{1, 2, 3, 4, 5}
	reslice := s[1:4]
	fmt.Println("reslice", reslice)

	s = []int{1, 2, 3, 4, 5}
	s = append(s[:2], s[3:]...)
	fmt.Println(s)

	s = []int{1, 2, 5}
	index := 2
	s = append(s[:index], append([]int{3, 4}, s[index:]...)...)
	fmt.Println(s)

	s = []int{1, 2, 3}
	for index, value := range s {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
