package sandbox

import "fmt"

func Main(){
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	// s = s[:6] 	// slice bounds out of range [:6] with capacity 4
	// printSlice(s)

	var ss []int
	printSlice(ss)

	a := make([]int, 5)
	printSlice(a)
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func Main_Slices(){
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct{
		i int 
		b bool 
	}{
		{2,true},
		{3,false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}
func Main_ArraysSlices(){
	fmt.Println("this should work")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])

	primes := [6]int {2,3,5,7,11,13}
	fmt.Println(primes)

	// a slice
	var s []int = primes[1:4]
	fmt.Println(s)
}