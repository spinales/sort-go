package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var b bool // -b, --ignore-leading-blanks  ignore leading blanks

func main() {
	// -b, --ignore-leading-blanks  ignore leading blanks
	flag.BoolVar(&b, "b", false, "ignore leading blanks")
	flag.BoolVar(&b, "ignore-leading-blanks", false, "ignore leading blanks")
	flag.Parse()

	// recibo todos los los parametros
	for _, f := range flag.Args() {
		sorting(f, b)
	}
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
// ilb: ignorar espacios en blanco
func sorting(filename string, ilb bool) {
	data := openFile(filename)
	if ilb {
		// elimino espacios en blanco
		spaces := strings.ReplaceAll(string(data), "\n\n", "\n")
		arr := strings.Split(spaces, "\n")
		sort.Strings(arr)
		fmt.Println(strings.Join(arr[:], "\n"))
	} else {
		// divido el archivo por cada linea
		// organizo, e imprimo
		arr := strings.Split(string(data), "\n")
		sort.Strings(arr)
		fmt.Println(strings.Join(arr[:], "\n"))
	}
}

// abre el archivo por la ruta pasada
func openFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}
