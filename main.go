package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var c bool

func main() {
	flag.BoolVar(&c, "c", false, `Check	that the single input file is ordered as specified by the arguments and the collating sequence of the current locale. 
	Output shall not be sent to standard output. The exit code shall indicate whether or not disorder was detected or an error occurred.
	If disorder  (or,  with -u, a duplicate key) is detected, a warning message shall be sent to standard error indicating where the
	disorder or duplicate key was found.`)
	flag.Parse()

	// recibo todos los los parametros
	for _, file := range flag.Args() {
		sorting(file, c)
	}
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
// sorted: para determinar si el archivo esta ordenado
func sorting(filename string, sorted bool) {
	data := openFile(filename)
	if sorted {
		arr := strings.Split(string(data), "\n")
		if sort.StringsAreSorted(arr) {
			os.Exit(0)
		}
		fmt.Fprintln(os.Stderr, "el archivo no esta ordenado")
		os.Exit(1)
	}
	printDefault(string(data))
}

func printDefault(data string) {
	// divido el archivo por cada linea
	// organizo, e imprimo
	arr := strings.Split(data, "\n")
	sort.Strings(arr)
	fmt.Println(strings.Join(arr[:], "\n"))
}

// abre el archivo por la ruta pasada
func openFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}
