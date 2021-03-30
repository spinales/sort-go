package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var (
	b bool // -b, --ignore-leading-blanks
	d bool // -d, --dictionary-order
)

func main() {
	// -d, --dictionary-order      consider only blanks and alphanumeric characters
	flag.BoolVar(&d, "d", false, "consider only blanks and alphanumeric characters")
	flag.BoolVar(&d, "dictionary-order", false, "consider only blanks and alphanumeric characters")
	// -b, --ignore-leading-blanks  ignore leading blanks
	flag.BoolVar(&b, "b", false, "ignore leading blanks")
	flag.BoolVar(&b, "ignore-leading-blanks", false, "ignore leading blanks")
	flag.Parse()

	// recibo todos los los parametros
	for _, f := range flag.Args() {
		sorting(f, b, d)
	}
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
// ilb: ignorar espacios en blanco
// dict: ordenar por diccionario
func sorting(filename string, ilb bool, dict bool) {
	data := openFile(filename)
	if ilb {
		// elimino espacios en blanco
		spaces := strings.ReplaceAll(string(data), "\n\n", "\n")
		printDefault(spaces)
	} else if dict {
		printDefault(string(data))
	} else {
		printDefault(string(data))
	}
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
