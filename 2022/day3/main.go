package main

import (
	"bufio"
	"fmt"
	"os"
)

func findItem(s string) int32 {
	left := s[:len(s)/2]
	right := s[len(s)/2:]
	for i := range left {
		for j := range right {
			if left[i] == right[j] {
				v := rune(left[i])
				if v < 91 {
					return v - 38
				} else {
					return v - 96
				}
			}
		}
	}
	return 1
}

func findBadge(s []string) int32 {
	for i := range s[0] {
		for j := range s[1] {
			for y := range s[2] {
				if s[0][i] == s[1][j] && s[1][j] == s[2][y] {
					v := rune(s[0][i])
					if v < 91 {
						return v - 38
					} else {
						return v - 96
					}
				}
			}
		}
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
	var r, r2 int32
	b := []string{}
	for s.Scan() {
		l := s.Text()
		r += findItem(l)
		b = append(b, l)
		if len(b) == 3 {
			r2 += findBadge(b)
			b = []string{}
		}
	}
	fmt.Println(r)
	fmt.Println(r2)
}
