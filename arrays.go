package main

import "fmt"

func myfuncslice(myslice []int) {
	fmt.Println("This is a slice func: ", myslice[1])
}

func myfunc(a [6]int) {

	fmt.Println("This is an array func: ", a[1])

}

func main() {

	// Arrays

	// When initialized all members are zero in array
	var arr [10]int
	fmt.Println(arr)

	// Short hand declaration
	arrsh := [3]int{12, 11, 11}
	fmt.Println(arrsh)

	// Not necessary that all values should be sent in short hand
	arrpt := [10]int{12}
	fmt.Println(arrpt)

	// Ignore the length of arr
	arrig := [...]int{12, 22}
	fmt.Println(arrig)

	// Is zero indexed and can assign as follows
	arr[1] = 21
	arr[2] = 22
	fmt.Println(arr)

	// The size of the array is a part of the type. Hence [5]int and [25]int are distinct types.
	// Hence we use slices in golang
	// a := [3]int{5, 78, 8}
	// var b [5]int
	// b = a //not possible since [3]int and [5]int are distinct types

	// Arrays in golang are value types and not referenced type
	// Hence change in original array will not show in new arr
	// Similarly when arrays are passed to functions as parameters, they are passed by value and the original array is unchanged.

	// a := [...]string{"USA", "China", "India", "Germany", "France"}
	// b := a // a copy of a is assigned to b
	// b[0] = "Singapore"
	// fmt.Println("a is ", a)
	// fmt.Println("b is ", b)

	// Output:
	// a is [USA China India Germany France]
	// b is [Singapore China India Germany France]

	fmt.Println("Length: ", len(arr))

	// Printing arrays
	for i, v := range arr {
		fmt.Println("index: ", i, "value: ", v)
	}

	// MultiDem. Arrays
	arrstring := [3][2]string{
		{"idfc", "bank"},
		{"random", "line"},
		{"this is ", "a test"},
	}

	for _, v1 := range arrstring {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	// Passing array in a function
	myarray := [6]int{1, 2, 3, 4, 5, 6}
	myfunc(myarray)

	// Slices
	// Slices do not own any data -- they are just references to the original array.

	var slice1 []int = arr[1:3]
	slice2 := arr[1:3]
	fmt.Println(slice1)
	fmt.Println(slice2)

	// Different type of declaring slices
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)

	//Any modifications done to the slice will be reflected in the underlying array.

	//The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	// creating slices using make()
	// func make([]T, len, cap) []T
	i := make([]int, 5, 5)
	fmt.Println(i)

	// We can append elements in slice unlike arrays.
	// One question might be bothering you though.
	// If slices are backed by arrays and arrays themselves are of fixed length then how come a slice is of dynamic length.
	// Well what happens under the hood is, when new elements are appended to the slice, a new array is created.
	// The elements of the existing array are copied to this new array and a new slice reference for this new array is returned.
	// The capacity of the new slice is now twice that of the old slice.

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	//zero value of a slice is nil

	//Appending a slice to another slice
	// food := append(veggies, fruits...)

	// Multi Dem slices
	random := [][]string{
		{"C", "C++"},
		{"IDFC"},
		{"GOLANG", "MICROSERVICES"},
	}
	fmt.Println(random)

	//Passing a slice to a func
	myslice := []int{1, 2, 3}
	myfuncslice(myslice)

	//  Memory Optimisation
	// Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected. This might be of concern when it comes to memory management. Let's assume that we have a very large array and we are interested in processing only a small part of it. Henceforth we create a slice from that array and start processing the slice. The important thing to be noted here is that the array will still be in memory since the slice references it.

	// One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice. This way we can use the new slice and the original array can be garbage collected.

	// func countries() []string {
	//     countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	//     neededCountries := countries[:len(countries)-2]
	//     countriesCpy := make([]string, len(neededCountries))
	//     copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	//     return countriesCpy
	// }
	// func main() {
	//     countriesNeeded := countries()
	//     fmt.Println(countriesNeeded)
	// }

	// In line no. 9 of the above program, neededCountries := countries[:len(countries)-2] creates a slice of countries barring the last 2 elements. Line no. 11 of the above program copies neededCountries to countriesCpy and also returns it from the function in the next line. Now countries array can be garbage collected since neededCountries is no longer referenced.

}
