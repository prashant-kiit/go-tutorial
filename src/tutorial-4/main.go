package main

import (
	"fmt"
	"slices"
)

func main() {
	// array
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
	fmt.Println(arr6[0:1]) // subarray {1}

	// slice
	var slice1 []int64 = []int64{1, 2, 5, 4, 3} // default {1, 2, 5, 4, 3}
	fmt.Println(slice1)
	fmt.Println(slice1[3])
	fmt.Println(slice1[3:5])

	slices.Sort(slice1)
	slices.Reverse(slice1)
	fmt.Println(slice1)

	slice1 = append(slice1, 10)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1 = append(slice1, slice1...)
	fmt.Println(slice1)

	var slice2 []int64 = make([]int64, 3, 5)
	fmt.Println((slice2))
	// fmt.Println((slice2[4])) // index of bound error

	// map
	var table1 map[string]int64
	fmt.Println(table1)
	fmt.Println(len(table1)) // no key so 0

	var table2 map[string]int64 = map[string]int64{"abc": 100, "pqr": -10}
	table2["xyz"] = 300
	fmt.Println(table2)
	table2["abc"] = 200
	fmt.Println(table2["abc"]) // 100
	fmt.Println(table2["ac"]) // defaulted to 0 (b/c value is int64)
	fmt.Println(len(table2)) // 2 keys so 2
	// fmt.Println(cap(table2)) // no capacity method for map

	var value, isPresent = table2["lmn"]
	fmt.Println(value, isPresent)

}
