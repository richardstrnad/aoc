package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	c := 1
	r := 1
	result := 0
	result2 := []string{}
	for s.Scan() {
		var a string
		var x int
		fmt.Sscan(s.Text(), &a, &x)
		result2 = append(result2, drawPixel(r, c-1))
		c++
		result += checkCycle(r, c)
		switch a {
		case "addx":
			result2 = append(result2, drawPixel(r, c-1))
			c += 1
			r += x
			result += checkCycle(r, c)
		}
	}
	fmt.Println(result)
	for i := range result2 {
		if i%40 == 0 {
			fmt.Println()
		}
		fmt.Print(result2[i])
	}
}

func checkCycle(r, c int) int {
	if c-20 == 0 || (c-20)%40 == 0 {
		return c * r
	}
	return 0
}

func drawPixel(r, c int) string {
	c = c % 40
	fmt.Println(c)
	if c >= r-1 && c <= r+1 {
		return "#"
	}
	return "."
}
