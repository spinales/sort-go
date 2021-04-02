package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var (
	c  bool // -c
	cs bool // -C
	m  bool // -m
)

func main() {
	// -c
	flag.BoolVar(&c, "c", false, `Check	that the single input file is ordered as specified by the arguments and the collating sequence of the current locale. 
	Output shall not be sent to standard output. The exit code shall indicate whether or not disorder was detected or an error occurred.
	If disorder  (or,  with -u, a duplicate key) is detected, a warning message shall be sent to standard error indicating where the
	disorder or duplicate key was found.`)
	// -C
	flag.BoolVar(&cs, "C", false, `Same as -c, except that a warning message shall not be sent to standard error if  disorder  or,  with	-u,  a	duplicate  key	is
	detected.`)
	// -m
	flag.BoolVar(&m, "m", false, "Merge only; the input file shall be assumed to be already sorted.")
	flag.Parse()

	// recibo todos los los parametros
	for _, file := range flag.Args() {
		switch {
		case m:
			printFile(file)
		case c || cs:
			sorting(file)
		default:
			printFile(file)
		}
	}
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
func sorting(filename string) {
	data := openFile(filename)
	arr := strings.Split(string(data), "\n")

	// en caso de que este ordenado
	if sort.StringsAreSorted(arr) {
		os.Exit(0)
		return
	}

	switch {
	case c:
		fmt.Fprintln(os.Stderr, "El archivo no esta ordenado")
	default:
		fmt.Println("El archivo no esta ordenado")
	}

	os.Exit(1)
}

func printFile(filename string) {
	data := openFile(filename)
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
