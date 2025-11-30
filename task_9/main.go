package main

import (
	"fmt"
)

func main() {
	exampleStr1 := "a4bc2d5e"
	exampleStr2 := "abcd"
	exampleStr3 := "45"
	exampleStr4 := ""

	res1, err := unpack(exampleStr1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(exampleStr1, ": ", res1)
	}
	res2, err := unpack(exampleStr2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(exampleStr2, ": ", res2)
	}
	res3, err := unpack(exampleStr3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(exampleStr3, ": ", res3)
	}
	res4, err := unpack(exampleStr4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(exampleStr4, ": ", res4)
	}
}
