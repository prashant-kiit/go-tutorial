package main

import "fmt"

func main() {
	var arr1 [3]int64 // default {0,0,0}
	fmt.Println(arr1)
	var arr2 [3]byte // default {0,0,0}
	fmt.Println(arr2)
	var arr3 [3]rune // default {0,0,0}
	fmt.Println(arr3)
	var arr4 [3]bool // default {0,0,0}
	fmt.Println(arr4)
	var arr5 [3]string     // default {'','',''}
	fmt.Println(len(arr5)) // 3

	var arr6 [3]int64 = [3]int64{1, 2} // default {1,2,0}
	fmt.Println(arr6)
	fmt.Println(arr6[2])
}
