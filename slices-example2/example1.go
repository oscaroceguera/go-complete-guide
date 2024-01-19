package main

import "fmt"

func slices1() {
	myslice1 := []int{}
	// In Go, there are two functions that can be used to return the length and capacity of a slice:
	fmt.Println(len(myslice1)) // returns the length of the slice (the number of elements in the slice)
	fmt.Println(cap(myslice1)) // returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	// output:
	// 0
	// 0

	fmt.Println("---")

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	// output: 4
	fmt.Println(cap(myslice2))
	// output: 4
	fmt.Println(myslice2)
	// output: [Go Slices Are Powerful]

	fmt.Println("---")

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	mysliceArr := arr1[2:4]

	fmt.Printf("mysliceArr = %v\n", mysliceArr)
	// output: mysliceArr = [12 13]
	fmt.Printf("length = %d\n", len(mysliceArr))
	// output: length = 2
	fmt.Printf("capacity = %d\n", cap(mysliceArr))
	// output: capacity = 4

	fmt.Println("---")

	// The make() function can also be used to create a slice.
	myslicemake := make([]int, 5, 10) // slice_name := make([]type, length, capacity)
	// Note: If the capacity parameter is not defined, it will be equal to length.
	fmt.Printf("myslicemake = %v\n", myslicemake)
	// output: myslicemake = [0 0 0 0 0]
	fmt.Printf("length = %d\n", len(myslicemake))
	// output = length = 5
	fmt.Printf("capacity = %d\n", cap(myslicemake))
	// output = capacity = 10

	fmt.Println("---")

	myslicemake2 := make([]int, 5)
	fmt.Printf("myslicemake2 = %v\n", myslicemake2)
	// output: myslicemake2 = [0 0 0 0 0]
	fmt.Printf("length = %d\n", len(myslicemake2))
	// output = length = 5
	fmt.Printf("capacity = %d\n", cap(myslicemake2))
	// output = capacity = 5
}