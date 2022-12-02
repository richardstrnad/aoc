package main

import (
	"bufio"
	"fmt"
	"os"
)

func eval(l string) int {
	switch l {
	case "A X":
		return 4
	case "A Y":
		return 8
	case "A Z":
		return 3
	case "B X":
		return 1
	case "B Y":
		return 5
	case "B Z":
		return 9
	case "C X":
		return 7
	case "C Y":
		return 2
	case "C Z":
		return 6
	}
	return 0
}

func eval2(l string) int {
	switch l {
	case "A X":
		return 3
	case "A Y":
		return 4
	case "A Z":
		return 8
	case "B X":
		return 1
	case "B Y":
		return 5
	case "B Z":
		return 9
	case "C X":
		return 2
	case "C Y":
		return 6
	case "C Z":
		return 7
	}
	return 0
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	r := 0
	r2 := 0
	for s.Scan() {
		r += eval(s.Text())
		r2 += eval2(s.Text())
	}
	fmt.Println(r)
	fmt.Println(r2)
}
