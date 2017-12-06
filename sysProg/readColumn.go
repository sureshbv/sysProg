package main

import (
	"fmt"
	"strings"
)


func main() {
	var s[3] string
	s[0] = "1 2 3"
	s[1] = "a b c d e f g h"
	s[2] = "abcd bcde efgh"

	column := 5
	
	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		if len(data) >= column {
			fmt.Println(data[column-1])
		}
	}
	
}
