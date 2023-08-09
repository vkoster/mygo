package main

import "fmt"

func main() {
	// Base Array 1
	var stringArray1 = [3]string{"a", "b", "c"}

	// two Slices referring to it
	s11 := stringArray1[0:2]
	var s12 []string = stringArray1[1:3]

	// Slices look good
	fmt.Println("s11 s12: ", s11, s12)

	// Now I append to one Slice, extending the Array capacity, i.e. new Array will be created
	s13 := append(s11, "4", "5")
	// s13 is now pointing to a new array
	s13[0] = "!"
	fmt.Println("after extending the array capacity")
	fmt.Println("s13: ", s13)
	fmt.Println("s11 s12: ", s11, s12)
	fmt.Println("Array: ", stringArray1)
	fmt.Println("Summery: original slices are on original array - new slice is on new array")

	// another Array, sufficient for appending
	var stringArray2 = [4]string{"a", "b", "c", "x"}
	// two slices are on it
	s21 := stringArray2[0:2]
	var s22 []string = stringArray2[1:3]

	fmt.Println("s11 s12: ", s21, s22)
	// append while still fitting the array
	s23 := append(s22, "4")
	s23[0] = "!"
	fmt.Println("after append without extending the array capacity")
	fmt.Println("s23: ", s23)
	fmt.Println("s21 s22: ", s21, s22)
	fmt.Println("Array: ", stringArray2)
	fmt.Println("Summery: everyone is still on the original array and sees each other's changes")

}
