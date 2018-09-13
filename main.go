package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	a := "ab"

	in := make([]byte, 4)
	n := utf8.EncodeRune(in, int32(a[0]))
	fmt.Println(string(in[:n]))
}
