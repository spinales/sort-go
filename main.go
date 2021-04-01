package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"sort-go/internal/alphabet"
	"strconv"
	"strings"
)

var (
	b bool // -b, --ignore-leading-blanks
	d bool // -d, --dictionary-order
	f bool // -f, --ignore-case
	g bool // -g, --general-numeric-sort
)

func main() {
	// -g, --general-numeric-sort
	flag.BoolVar(&g, "g", false, "compare according to general numerical value")
	flag.BoolVar(&g, "general-numeric-sort", false, "compare according to general numerical value")
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
		sorting(file, b, d, f, g)
	}
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
// ilb: ignorar espacios en blanco
// dict: ordenar por diccionario
// ign: ordenar por orden alphanumerico
// nums: si son numeros a ordenar
func sorting(filename string, ilb bool, dict bool, ign bool, nums bool) {
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
	} else if nums {
		numbs(data)
	} else {
		printDefault(string(data))
	}
}

func numbs(data []byte) {
	arr := toNumbers(data)
	sort.Ints(arr)
	printNumbers(arr)
}

func printNumbers(data []int) {
	for _, n := range data {
		fmt.Printf("%v \n", n)
	}
}

func toNumbers(data []byte) []int {
	strs := strings.Split(string(data), "\n")
	var nums []int
	for _, s := range strs {
		result, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(result))
	}
	return nums
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
