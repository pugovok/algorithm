//   C. Успеть все
//   https://codeforces.com/contest/1031/problem/C
//   tags: greedy
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scanner struct {
	s *bufio.Scanner
}

func initScanner(filename string) scanner {
	//f, _ := os.Open(filename)
	f := os.Stdin
	return scanner{s: bufio.NewScanner(f)}
}

func (s scanner) ReadInt() int {
	s.s.Scan()
	val, _ := strconv.Atoi(s.s.Text())
	return val
}

func sum(x float64) float64 {
	return x * (x + 1) / 2
}

func Implode(glue string, args []int) string {
	data := make([]string, len(args))
	for i, s := range args {
		data[i] = fmt.Sprint(s)
	}
	return strings.Join(data, glue)
}

func main() {

	scanner := initScanner("input.txt")
	scanner.s.Split(bufio.ScanWords)

	a := scanner.ReadInt()
	b := scanner.ReadInt()

	x := 1
	for ; sum(float64(x)) < float64(a+b); x++ {
	}

	var outA []int
	var outB []int

	for limit := 0; x != 0 && limit < 2; x-- {
		aX := a
		aY := b
		limit++

		var aRes []int
		var bRes []int

		for tempX := x; tempX > 0; tempX-- {
			if aX >= tempX {
				aX -= tempX
				aRes = append(aRes, tempX)
			} else {
				if aY >= tempX {
					aY -= tempX
					bRes = append(bRes, tempX)
				}
			}
		}

		if len(aRes)+len(bRes) > len(outA)+len(outB) {
			outA = aRes
			outB = bRes
		}
	}

	fmt.Println(len(outA))
	fmt.Println(Implode(" ", outA))
	fmt.Println(len(outB))
	fmt.Println(Implode(" ", outB))
}
