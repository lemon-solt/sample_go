package main

import (
	"fmt"
	"sample_go/mylib"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// short
	short_cut_sample := "short_cut_sample"
	fmt.Println(short_cut_sample)

	// normal
	var normal_sample string = "normal_sample"
	fmt.Println(normal_sample)

	result := mylib.SampleFunc()
	fmt.Printf("result value is type=%T, value=%v", result)
}
