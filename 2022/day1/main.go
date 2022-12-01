package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	filePath := "input"
	readFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	var c int
	r := []int{}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		n, err := strconv.Atoi(fileScanner.Text())
		if err == nil {
			c += n
		} else {
			r = append(r, c)
			c = 0
		}
	}
	sort.Ints(r)
	fmt.Println(r[len(r)-1])
	fmt.Println(sum(r[len(r)-3:]))
}
