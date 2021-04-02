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
	c  bool   // -c
	cs bool   // -C
	m  bool   // -m
	o  string // -o
)
var data []string

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
	// -o
	flag.StringVar(&o, "o", "", "Specify  the  name  of  an  output  file to be used instead of the standard output. This file can be the same as one of the input files.")
	flag.Parse()

	// recibo todos los los parametros
	for _, file := range flag.Args() {
		fileData := openFile(file)
		switch {
		case m:
			data = append(data, strings.Split(string(fileData), "\n")...)
		case c || cs:
			sorting(file)
		default:
			data = append(data, strings.Split(string(fileData), "\n")...)
			sort.Strings(data)
		}
	}

	if o != "" {
		writeFile(o, data)
		os.Exit(0)
	}

	fmt.Println(strings.Join(data[:], "\n"))
	os.Exit(0)
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

// abre el archivo por la ruta pasada
func openFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}

func writeFile(filepath string, data []string) {
	err := ioutil.WriteFile(filepath, []byte(strings.Join(data, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}
