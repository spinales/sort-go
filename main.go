package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	flag.Parse()

	for _, f := range flag.Args() {
		sorting(f)
	}
}

func sorting(filename string) {
	data := openFile(filename)
	arr := strings.Split(string(data), "\n")
	sort.Strings(arr)
	result := strings.Join(arr[:], "\n")
	fmt.Println(result)
}

func openFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}
