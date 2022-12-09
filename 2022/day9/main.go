package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func direction(from, to int) int {
	if to == from {
		return 0
	} else if to > from {
		return 1
	}
	return -1
}

type p struct {
	x, y int
}

type rope struct {
	knots  []p
	fields map[p]bool
}

func (r rope) move(d string) {
	switch d {
	case "R":
		r.knots[0].x += 1
	case "L":
		r.knots[0].x -= 1
	case "D":
		r.knots[0].y -= 1
	case "U":
		r.knots[0].y += 1
	}
	for i := 1; i < len(r.knots); i++ {
		if abs(r.knots[i-1].x-r.knots[i].x) > 1 || abs(r.knots[i-1].y-r.knots[i].y) > 1 {
			r.knots[i] = p{
				r.knots[i].x + direction(r.knots[i].x, r.knots[i-1].x),
				r.knots[i].y + direction(r.knots[i].y, r.knots[i-1].y),
			}
		}
	}
	r.fields[r.knots[len(r.knots)-1]] = true
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	rope1 := rope{[]p{{0, 0}, {0, 0}}, map[p]bool{}}
	rope2 := rope{[]p{}, map[p]bool{}}
	for i := 0; i < 10; i++ {
		rope2.knots = append(rope2.knots, p{0, 0})
	}
	for s.Scan() {
		var d string
		var n int
		fmt.Sscanf(s.Text(), "%s %d", &d, &n)
		for i := 0; i < n; i++ {
			rope1.move(d)
			rope2.move(d)
		}
	}
	fmt.Println(len(rope1.fields))
	fmt.Println(len(rope2.fields))
}
