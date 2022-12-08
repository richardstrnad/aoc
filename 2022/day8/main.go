package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	m := [][]int{}
	r := 0
	r2 := 0
	for s.Scan() {
		m = append(m, []int{})
		l := s.Text()
		for i := range l {
			x, _ := strconv.Atoi(string(l[i]))
			m[len(m)-1] = append(m[len(m)-1], x)
		}
	}
	for i := range m[1 : len(m)-1] {
		for j := range m[i+1][1 : len(m[i+1])-1] {
			c := m[i+1][j+1]
			downs := []int{}
			for ii := i + 2; ii <= len(m)-1; ii++ {
				downs = append(downs, m[ii][j+1])
			}
			ups := []int{}
			for ii := i; ii >= 0; ii-- {
				ups = append(ups, m[ii][j+1])
			}
			lefts := []int{}
			for jj := j; jj >= 0; jj-- {
				lefts = append(lefts, m[i+1][jj])
			}
			cs := score(c, m[i+1][j+2:])
			cs *= score(c, lefts)
			cs *= score(c, downs)
			cs *= score(c, ups)
			if cs > r2 {
				r2 = cs
			}
			right := visible(c, m[i+1][j+2:])
			if right {
				r++
				continue
			}
			left := visible(c, m[i+1][:j+1])
			if left {
				r++
				continue
			}
			down := visible(c, downs)
			if down {
				r++
				continue
			}
			up := visible(c, ups)
			if up {
				r++
				continue
			}
		}
	}
	fmt.Println(r+(len(m)*4-4), r2)
}

func score(c int, s []int) int {
	r := 0
	for i := range s {
		r++
		switch s[i] >= c {
		case true:
			return r
		}
	}
	return r
}
func visible(c int, s []int) bool {
	for i := range s {
		if s[i] >= c {
			return false
		}
	}
	return true
}
