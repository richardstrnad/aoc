package main

import (
	"bufio"
	"fmt"
	"os"
)

func unique(s string, l int) bool {
	set := make(map[byte]bool, l)
	for c := range s {
		set[s[c]] = true
	}
	return len(set) == l
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s1, s2 := 0, 0
	for s.Scan() {
		t := s.Text()
		for i, j := 0, 4; j < len(t)+1; i, j = i+1, j+1 {
			if unique(t[i:j], 4) {
				s1 = j
				break
			}
		}
		for i, j := 0, 14; j < len(t)+1; i, j = i+1, j+1 {
			if unique(t[i:j], 14) {
				s2 = j
				break
			}
		}
	}
	fmt.Println(s1, s2)
}
