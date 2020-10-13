package main

import "fmt"

func main()  {
	var n int
	var s [] uint16
	var total uint16
	fmt.Scanln(&n)

	for i := 0 ; i < n; i++ {
		var entrada uint16
		fmt.Scanln(&entrada)
		s = append(s, entrada)
	}
	
	for _, v := range s {
		total += v
	}

	fmt.Println(total)
 }