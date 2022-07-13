package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		num, oper, result, temp int
		isInParent bool = false
	)

	fmt.Fscanf(r, "%d", &num)
	result = num

	for {
		fmt.Fscanf(r, "%c", &oper)
		if (oper != 43) && (oper != 45) {
			result -= temp
			break
		}

		fmt.Fscanf(r, "%d", &num)

		if !isInParent && oper == 45 {
			isInParent = true
			temp = num
		} else if !isInParent && oper == 43 {
			result += num
		} else if isInParent && oper == 43 {
			temp += num
		} else if isInParent && oper == 45 {
			result -= temp
			temp = num
		}

	}

	fmt.Fprintln(w, result)
}