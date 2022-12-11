package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Monkey struct {
	items []int
	op    func(int) int
	test  func(int) int
	t     int
	f     int
}

func main() {
	input, _ := os.ReadFile("input")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	ml := []Monkey{}
	// the ints are getting to big, we try to modulus them down
	mm := 1
	for _, s := range split {
		// got some help for the parsing from https://github.com/mnml/aoc/blob/main/2022/11/1.go
		ml = append(ml, Monkey{})
		var items, op string
		var i, v, test, t, f int
		fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(s),
			`Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`,
			&i, &items, &op, &v, &test, &t, &f)

		json.Unmarshal([]byte("["+items+"]"), &ml[i].items)
		ml[i].op = map[string]func(int) int{
			"+": func(o int) int { return o + v },
			"*": func(o int) int { return o * v },
			"^": func(o int) int { return o * o },
		}[op]
		ml[i].test = func(w int) int {
			if w%test == 0 {
				return t
			}
			return f
		}
		mm *= test
	}
	inspect(20, ml, len(split), func(i int) int { return i / 3 })
	inspect(10000, ml, len(split), func(i int) int { return i % mm })
}

func inspect(c int, ml []Monkey, le int, d func(int) int) {
	ml = append([]Monkey{}, ml...)
	res := make([]int, le)
	for i := 0; i < c; i++ {
		for m := range ml {
			for _, w := range ml[m].items {
				w = d(ml[m].op(w))
				nm := (ml[m].test(w))
				ml[nm].items = append(ml[nm].items, w)
				res[m]++
			}
			ml[m].items = []int{}
		}
	}
	sort.Ints(res)
	fmt.Println(res)
	fmt.Println(res[len(res)-2] * res[len(res)-1])
}
