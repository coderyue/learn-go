package main

import (
	"fmt"
)

func main() {

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []bool{true, false, false, true}
	fmt.Println(q)

	s := []struct{
		a int
		b string
	}{
		{1, "sdfsd"},
		{2, "werew"},
		{3, "xxcv"},
	}
	fmt.Println(s)

	var ss []int
	fmt.Println(ss, len(ss), cap(ss))

	if ss == nil {
		fmt.Println("nil")
	}

	var sddsss []int
	fmt.Println(sddsss)



}

