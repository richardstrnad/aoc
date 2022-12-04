package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func h(l, r string) []int {
	lt := strings.Split(l, "-")
	rt := strings.Split(r, "-")
	res := make([]int, 4)
	res[0], _ = strconv.Atoi(lt[0])
	res[1], _ = strconv.Atoi(lt[1])
	res[2], _ = strconv.Atoi(rt[0])
	res[3], _ = strconv.Atoi(rt[1])
	return res

}
func fullyContains(l, r string) int {
	m := h(l, r)
	if (m[0] >= m[2] && m[1] <= m[3]) || (m[2] >= m[0] && m[3] <= m[1]) {
		return 1
	}
	return 0
}

func overlap(l, r string) int {
	m := h(l, r)
	if m[1] < m[2] || m[3] < m[0] {
		return 0
	}
	return 1
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	res := 0
	res2 := 0
	for s.Scan() {
		sp := strings.Split(s.Text(), ",")
		l := sp[0]
		r := sp[1]
		res += fullyContains(l, r)
		res2 += overlap(l, r)
	}
	fmt.Println(res)
	fmt.Println(res2)
}
