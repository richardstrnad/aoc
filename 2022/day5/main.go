package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func inverse(m [][]string) [][]string {
	for l := range m {
		for i, j := 0, len(m[l])-1; i < j; i, j = i+1, j-1 {
			m[l][i], m[l][j] = m[l][j], m[l][i]
		}
	}
	return m
}

func move1(count, from, to int, m [][]string) [][]string {
	for i := 0; i < count; i++ {
		e := m[from][len(m[from])-1]
		m[from] = m[from][:len(m[from])-1]
		m[to] = append(m[to], e)
	}
	return m
}

func move2(count, from, to int, m [][]string) [][]string {
	m[to] = append(m[to], m[from][len(m[from])-count:]...)
	m[from] = m[from][:len(m[from])-count]
	return m
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	m := [][]string{}
	for i := 0; i < 9; i++ {
		m = append(m, []string{})
	}
	m2 := make([][]string, len(m))
	moves := false
	for s.Scan() {
		l := s.Text()
		// change from arrays to moves
		if l == "" {
			moves = true
			m = inverse(m)
			for i := range m {
				m2[i] = make([]string, len(m[i]))
				copy(m2[i], m[i])
			}
			continue
		}
		switch moves {
		case false:
			c := 1
			for i := 0; i < 9; i++ {
				v := string(l[i+c])
				// we exlude the numbers from being append
				if _, err := strconv.Atoi(v); err == nil {
					//
				} else if v != " " {
					m[i] = append(m[i], v)
				}
				c += 3
			}
		case true:
			r, err := regexp.Compile(`(\d+).*(\d+).*(\d+)`)
			if err != nil {
				panic(err)
			}
			res := r.FindStringSubmatch(l)
			count, _ := strconv.Atoi(res[1])
			from, _ := strconv.Atoi(res[2])
			to, _ := strconv.Atoi(res[3])
			// account for 0 indexed slices
			from--
			to--
			m = move1(count, from, to, m)
			m2 = move2(count, from, to, m2)
		}
	}
	result := ""
	result2 := ""
	for i := range m {
		result += m[i][len(m[i])-1]
		result2 += m2[i][len(m2[i])-1]
	}
	fmt.Println(result)
	fmt.Println(result2)
}
