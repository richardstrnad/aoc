package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name   string
	Parent *Node
	Child  []*Node
	Files  []int
	Size   int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	r := &Node{}
	current := r
	for s.Scan() {
		l := s.Text()
		cmd := strings.HasPrefix(l, "$")
		switch cmd {
		// We handle a command
		case true:
			c := strings.Split(l, " ")
			switch len(c) {
			case 3:
				switch c[2] {
				case "..":
					// we go one directory up
					current = current.Parent
				default:
					n := &Node{
						Name:   c[2],
						Parent: current,
					}
					current.Child = append(current.Child,
						n,
					)
					current = n
				}
			}
			// We handle ls output
		case false:
			ls := strings.Split(l, " ")
			size, err := strconv.Atoi(ls[0])
			switch err {
			case nil:
				t := current
				addUp := true
				for addUp {
					t.Size += size
					if t.Parent == nil {
						addUp = false
					} else {
						t = t.Parent
					}
				}

			default:
				// dir
			}
		}
	}
	result := 0
	free := 70000000 - r.Size
	update := 30000000
	q := []*Node{r}
	result2 := r.Size
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		if c.Size < 100000 {
			result += c.Size
		}
		if c.Size > update-free && c.Size < result2 {
			result2 = c.Size
		}
		q = append(q, c.Child...)
	}
	fmt.Println(result)
	fmt.Println(result2)
}
