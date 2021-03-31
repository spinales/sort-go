package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"sort-go/internal/alphabet"
	"strings"
)

var (
	b bool // -b, --ignore-leading-blanks
	d bool // -d, --dictionary-order
	f bool // -f, --ignore-case
)

func main() {
	// -f, --ignore-case
	flag.BoolVar(&f, "f", false, "fold lower case to upper case characters")
	flag.BoolVar(&f, "ignore-case", false, "fold lower case to upper case characters")
	// -d, --dictionary-order      consider only blanks and alphanumeric characters
	flag.BoolVar(&d, "d", false, "consider only blanks and alphanumeric characters")
	flag.BoolVar(&d, "dictionary-order", false, "consider only blanks and alphanumeric characters")
	// -b, --ignore-leading-blanks  ignore leading blanks
	flag.BoolVar(&b, "b", false, "ignore leading blanks")
	flag.BoolVar(&b, "ignore-leading-blanks", false, "ignore leading blanks")
	flag.Parse()

	// recibo todos los los parametros
	for _, file := range flag.Args() {
		sorting(file, b, d, f)
	}
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
// ilb: ignorar espacios en blanco
// dict: ordenar por diccionario
// ign:
func sorting(filename string, ilb bool, dict bool, ign bool) {
	data := openFile(filename)
	if ilb {
		// elimino espacios en blanco
		spaces := strings.ReplaceAll(string(data), "\n\n", "\n")
		printDefault(spaces)
	} else if dict {
		printDefault(string(data))
	} else if ign {
		arr := strings.Split(string(data), "\n")
		sort.Sort(alphabet.Alphabetic(arr))
		fmt.Println(strings.Join(arr[:], "\n"))
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
