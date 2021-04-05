package main

import (
	"bytes"
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
	u  bool   // -u
	h  bool   // -h --help
	d  bool   //
)
var data []string

func main() {
	// -c
	flag.BoolVar(&c, "c", false, `Check	that the single input file is ordered as specified by the arguments and 
	the collating sequence of the current locale. Output shall not be sent to standard output. The exit code 
	shall indicate whether or not disorder was detected or an error occurred. If disorder (or,  with -u, a duplicate key) 
	is detected, a warning message shall be sent to standard error indicating where the disorder or duplicate key was found.`)
	// -C
	flag.BoolVar(&cs, "C", false, `Same as -c, except that a warning message shall not be sent to standard error 
		if disorder or, with -u, a duplicate key is detected.`)
	// -m
	flag.BoolVar(&m, "m", false, "Merge only; the input file shall be assumed to be already sorted.")
	// -o
	flag.StringVar(&o, "o", "", `Specify the name of an output file to be used instead of the standard output. 
		This file can be the same as one of the input files.`)
	// -u
	flag.BoolVar(&u, "u", false, `Unique: suppress all but one in each set of lines having equal keys. If used with the -c option, 
		check that there are  no  lines with duplicate keys, in addition to checking that the input file is sorted.`)
	// -d
	flag.BoolVar(&d, "d", false, `Specify that only <blank> characters and alphanumeric characters, according to the 
		current setting of LC_CTYPE, shall be significant in comparisons. The behavior is undefined for a sort key to 
		which -i or -n also applies.`)
	// -h --help
	flag.BoolVar(&h, "h", false, "help command.")
	flag.BoolVar(&h, "help", false, "help command.")
	flag.Parse()

	if h {
		flag.PrintDefaults()
	}

	// recibo todos los los parametros
	for _, file := range flag.Args() {
		fileData := openFile(file)
		arr := strings.Split(string(fileData), "\n")
		switch {
		case m:
			data = append(data, arr...)
		case c && u:
			if len(arr) != len(suprimeDuplicates(arr)) {
				fmt.Fprintln(os.Stderr, "el archivo tiene duplicados")
			}
			sorting(file)
		case c || cs:
			sorting(file)
		case u:
			data = append(data, suprimeDuplicates(arr)...)
		}
	}

	if o != "" {
		writeFile(o, data)
		os.Exit(0)
	}

	fmt.Println(strings.Join(data[:], "\n"))
	os.Exit(0)
}

func suprimeDuplicates(src []string) []string {
	seen := make(map[string]bool, len(src))
	j := 0

	for _, v := range src {
		if _, ok := seen[v]; !ok {
			seen[v] = true
			src[j] = v
			j++
		}
	}
	return src[:j]
}

func invisibleChar(arr []byte) []byte {
	for v := range invisible {
		arr = bytes.ReplaceAll(arr, []byte{v}, []byte(invisible[v]))
	}
	return arr
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
func sorting(filename string) {
	data := openFile(filename)
	arr := strings.Split(string(data), "\n")

	// en caso de que este ordenado
	switch {
	// case i:
	// 	res := invisibleChar(data)
	// 	arrsplit := strings.Split(string(res), "\n")
	// 	if sort.StringsAreSorted(arrsplit) {
	// 		os.Exit(0)
	// 	}
	case d:
		if sort.StringsAreSorted(arr) {
			os.Exit(0)
		}
	}

	switch {
	case c:
		fmt.Fprintln(os.Stderr, "El archivo no esta ordenado")
	case cs:
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
