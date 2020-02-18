package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// a = 0
	var a int
	// b = 10
	var b int = 10
	// same as b (type inference)
	c := 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// arrays
	// fixed length
	var myArray [5]int
	fmt.Println(myArray)
	// declaration with assignation
	myArrayB := [5]int{2, 3, 4}
	fmt.Println(myArrayB)

	// slices (array without fixed length)
	myArrayC := []int{2, 43, 13, 3, 3}
	fmt.Println(myArrayC)
	// non-destructive
	myArrayC = append(myArrayC, 12, 12, 234)
	fmt.Println(myArrayC)

	// map
	// map[keys-type]values-type
	mapA := make(map[string]int)
	mapA["age"] = 12
	mapA["number"] = 112
	fmt.Println(mapA["age"])
	delete(mapA, "age")
	// age is deleted, so mapA["age"] returns this map values default value (0)
	fmt.Println(mapA["age"])

	// loop
	// there is only for loops
	for i := 0; i < 8; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// but you can do what we usualy call while loop like that
	j := 0
	for j < 5 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	// we can also iterate an array with range
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Print("index = ", index, " - value = ", value, " | ")
	}
	fmt.Println()

	// same for map, with key instead of index
	maap := make(map[string]int)
	maap["asd"] = 12
	maap["mvl3"] = 2
	maap["hoho"] = 7

	for key, value := range maap {
		fmt.Print("key = ", key, " - value = ", value, " | ")
	}
	fmt.Println()

	// functions
	fmt.Println(sum(12, 3))
	rez, err := sqrt(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rez)
	}

	// structs
	p := person{name: "Hugo", age: 27}
	fmt.Println(p)
	fmt.Println(p.age)

	// pointers
	h := 12
	// print the adress of h
	fmt.Println(&h)
	// does not increment h because we pass it as a copy
	inc(h)
	fmt.Println(h)
	incAsPoint(&h)
	fmt.Println(h)
}

func inc(x int) {
	x++
}

func incAsPoint(x *int) {
	*x++
}

// 2 ints as params and returns int
func sum(n1 int, n2 int) int {
	return n1 + n2
}

// 1 float64 as param and 2 returns value
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("Underfined for negative numbers")
	}
	return math.Sqrt(n), nil
}

// structs must be declared outside a function
type person struct {
	name string
	age  int
}
