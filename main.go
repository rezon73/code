package main

import "strconv"

func main() {
	var arr []int
	arr = append(arr, 1, 2, 3, 4, 5)

	for key, val := range arr {
		//print(strconv.Itoa(key) + " ", strconv.Itoa(val) + "\n")
		defer print(strconv.Itoa(key) + " ", strconv.Itoa(val) + "\n")

		go func(key int, val int) {
			print(strconv.Itoa(key) + " ", strconv.Itoa(val) + "\n")
		}(key, val)
	}
}
