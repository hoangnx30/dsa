package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) <= numRows {
		return s
	}

	rows := make([]string, numRows)

	currentRow, step := 0, -1

	for _, char := range s {
		rows[currentRow] += string(char)

		if currentRow == 0 || currentRow == numRows-1 {
			step = -step
		}
		currentRow += step
	}

	return strings.Join(rows, "")
}

/**
P   A   H   N
A P L S I I G
Y   I   R
*/

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows)) // Output: "PAHNAPLSIIGYIR"
}
