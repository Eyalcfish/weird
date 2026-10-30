package main

import "fmt"

func main() {
	arr := make([]int, 1)  // create a new array(expected state: [])
	arr[0] = 1             // set the first array to 1(expected state: [1])
	arr2 := append(arr, 2) // append 2 and by that reallocate the buffer(expected state: [1,2]), by that making arr a reference to an old array
	arr[0] = 3             // set the first value of the reference to the old buffer to 3(expected state: [3,2])
	fmt.Println(arr2)      // print: actualy result [1,2]
}
