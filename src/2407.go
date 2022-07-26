package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BigNum struct {
	high, low string
}

var dp [][]string

func Add(a, b string) string {
	// 상위 15자리, 하위 15자리 나눠서 더하기
	lenA, lenB := len(a), len(b)
	if lenA <= 15 && lenB <= 15 {
		aNum, _ := strconv.Atoi(a)
		bNum, _ := strconv.Atoi(b)
		temp := strconv.Itoa(aNum + bNum)
		if len(temp) <= 15 {
			zeroString := ""
			for i := 0; i < (15 - len(temp)); i++ {
				zeroString += "0"
			}
			temp = zeroString + temp
		}
		return temp
	} else if lenA > 15 && lenB <= 15 {
		high := a[:(lenA - 15)]
		low := a[(lenA - 15):]
		aNum, _ := strconv.Atoi(low)
		bNum, _ := strconv.Atoi(b)
		temp := strconv.Itoa(aNum + bNum)
		if len(temp) > 15 {
			highNum, _ := strconv.Atoi(high)
			highNum++
			high = strconv.Itoa(highNum)
			temp = temp[(len(temp) - 15):]
		} else {
			zeroString := ""
			for i := 0; i < (15 - len(temp)); i++ {
				zeroString += "0"
			}
			temp = zeroString + temp
		}
		return high+temp
	} else if lenA <= 15 && lenB > 15 {
		high := b[:(lenB - 15)]
		low := b[(lenB - 15):]
		aNum, _ := strconv.Atoi(low)
		bNum, _ := strconv.Atoi(a)
		temp := strconv.Itoa(aNum + bNum)
		if len(temp) > 15 {
			highNum, _ := strconv.Atoi(high)
			highNum++
			high = strconv.Itoa(highNum)
			temp = temp[(len(temp) - 15):]
		} else {
			zeroString := ""
			for i := 0; i < (15 - len(temp)); i++ {
				zeroString += "0"
			}
			temp = zeroString + temp
		}
		return high+temp
	} else {
		aHigh, aLow, bHigh, bLow := a[:(lenA - 15)], a[(lenA - 15):], b[:(lenB - 15)], b[(lenB - 15):]
		aLowNum, _ := strconv.Atoi(aLow)
		bLowNum, _ := strconv.Atoi(bLow)
		temp := strconv.Itoa(aLowNum + bLowNum)
		carry := 0
		if len(temp) > 15 {
			carry++
			temp = temp[(len(temp) - 15):]
		} else {
			zeroString := ""
			for i := 0; i < (15 - len(temp)); i++ {
				zeroString += "0"
			}
			temp = zeroString + temp
		}
		aHighNum, _ := strconv.Atoi(aHigh)
		bHighNum, _ := strconv.Atoi(bHigh)
		high := strconv.Itoa(aHighNum + bHighNum + carry)
		return high+temp
	}
}

func nCm(N, M int) string {
	if M == N || M == 0 {
		return "1"
	}
	if dp[N][M] != "" {
		return dp[N][M]
	}

	dp[N][M] = Add(nCm(N-1, M-1), nCm(N-1, M))

	return dp[N][M]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	
	var N, M int
	
	fmt.Fscanln(r, &N, &M)
	dp = make([][]string, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]string, M+1)
	}

	result := nCm(N, M)
	startIndex := 0
	for i, c := range result {
		if c != '0' {
			startIndex = i
			break
		}
	}
	for i := startIndex; i < len(result); i++ {
		fmt.Fprintf(w, "%c", result[i])
	}
	fmt.Fprintf(w, "\n")
}