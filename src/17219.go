package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M int

	fmt.Fscanln(r, &N, &M)

	password := make(map[string]string)

	for i := 0; i < N; i++ {
		S, _ := r.ReadString('\n')
		S = strings.TrimSpace(S)
		site_pass := strings.Split(S, " ")
		password[site_pass[0]] = site_pass[1]
	}


	for i := 0; i < M; i++ {
		S, _ := r.ReadString('\n')
		S = strings.TrimSpace(S)
		fmt.Fprintln(w, password[S])
	}
}