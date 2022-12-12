package main

import (
	"bufio"
	"fmt"
	"os"
)

var moves = []pos{
	{x: 0, y: 1},
	{x: 0, y: -1},
	{x: 1, y: 0},
	{x: -1, y: 0},
}

var visited = make(map[pos]bool)

type pos struct {
	x int
	y int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	m := [][]rune{}
	start := pos{}
	pStarts := []pos{}
	for s.Scan() {
		m = append(m, make([]rune, len(s.Text())))
		for i := range s.Text() {
			e := rune(s.Text()[i])
			if e == []rune("S")[0] {
				start.x = len(m) - 1
				start.y = i
				e = []rune("a")[0]
			}
			if e == []rune("a")[0] {
				s := pos{
					x: len(m) - 1,
					y: i,
				}
				pStarts = append(pStarts, s)
			}
			// we transform E to make it easier findable with our algorithm
			if e == []rune("E")[0] {
				e = []rune("z")[0] + 1
			}
			m[len(m)-1][i] = e
		}
	}
	fmt.Println(findSolution(m, start))
	r := len(m) * len(m[0])
	for i := range pStarts {
		x := findSolution(m, pStarts[i])
		if x < r {
			r = x
		}
	}
	fmt.Println(r)
}

func findSolution(m [][]rune, start pos) int {
	visited = map[pos]bool{}
	q := []pos{start}
	r := 0
	for len(q) > 0 {
		n := []pos{}
		for _, e := range q {
			if m[e.x][e.y] == []rune("z")[0]+1 {
				return r
			}
			n = append(n, explore(m, e)...)
		}
		r++
		q = n
	}
	return len(m) * len(m[0])
}

func explore(m [][]rune, e pos) []pos {
	r := []pos{}
	for i := range moves {
		n := pos{
			x: e.x + moves[i].x,
			y: e.y + moves[i].y,
		}
		if n.x >= 0 && n.x < len(m) && n.y >= 0 && n.y < len(m[0]) && !visited[n] {
			if m[n.x][n.y]-m[e.x][e.y] <= 1 {
				visited[n] = true
				r = append(r, n)
			}
		}
	}
	return r
}
