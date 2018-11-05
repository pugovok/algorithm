//   Problem B. Infinite House of Pancakes
//   url:  https://code.google.com/codejam/contest/6224486/dashboard#s=p1
//   tags: greedy
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type scanner struct {
	s *bufio.Scanner
}

func initScanner(filename string) scanner {
	f, _ := os.Open(filename)
	return scanner{s: bufio.NewScanner(f)}
}

func (s scanner) ReadInt() int {
	s.s.Scan()
	val, _ := strconv.Atoi(s.s.Text())
	return val
}

func main() {
	fo, _ := os.Create("output.txt")
	writer := bufio.NewWriter(fo)
	defer fo.Close()

	defer writer.Flush()

	//scanner := initScanner("B-small-practice.in")
	scanner := initScanner("B-large-practice.in")
	//scanner := initScanner("input.txt")
	scanner.s.Split(bufio.ScanWords)

	countTests := scanner.ReadInt()

	for currentTest := 0; currentTest < countTests; currentTest++ {
		plates := scanner.ReadInt()

		var p []int
		maxPanCake := 0
		for i := 0; i < plates; i++ {
			v := scanner.ReadInt()
			if v > maxPanCake {
				maxPanCake = v
			}
			p = append(p, v)
		}

		ret := maxPanCake
		for x := 1; x < maxPanCake; {

			moves := 0
			for _, Pi := range p {
				moves += (Pi - 1) / x
			}

			if ret > moves+x {
				ret = moves + x
			}
			x++
		}
		writer.WriteString(fmt.Sprintf("Case #%d: %d\n", currentTest+1, ret))
	}

}
