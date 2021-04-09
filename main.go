package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"sort-go/internal/alphabet"
	"sort-go/internal/file"
	"sort-go/internal/reverse"
	"strconv"
	"strings"
)

var (
	data   []string
	values []int
	c      bool   // -c
	cs     bool   // -C
	m      bool   // -m
	o      string // -o
	u      bool   // -u
	h      bool   // -h --help
	d      bool   // -d
	i      bool   // -i
	f      bool   // -f
	n      bool   // -n
	r      bool   // -r
	t      string // -t
	b      bool   // -b
	k      string // -k
)

func main() {
	// -c
	flag.BoolVar(&c, "c", false, `Check that the single input file is ordered as specified by the arguments and 
	the collating sequence of the current locale. Output shall not be sent to standard output. The exit code 
	shall indicate whether or not disorder was detected or an error occurred. If disorder (or, with -u, a duplicate key) 
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
	// -i
	flag.BoolVar(&i, "i", false, `Ignore all characters that are non-printable, according to the current 
	setting of LC_CTYPE. The behavior is undefined for a sort key for which -n also applies.`)
	// -f
	flag.BoolVar(&f, "f", false, `Consider all lowercase characters that have uppercase equivalents, 
	according to the current setting of LC_CTYPE, to be the upper-case equivalent for the 
	purposes of comparison.`)
	// -n
	flag.BoolVar(&n, "n", false, `Restrict the sort key to an initial numeric string, consisting of optional <blank> 
	characters, optional minus-sign, and  zero	or more  digits  with an optional radix character and thousands 
	separators (as defined in the current locale), which shall be sorted by arithmetic value. An empty digit string 
	shall be treated as zero. Leading zeros and signs on zeros shall not affect ordering`)
	// -r
	flag.BoolVar(&r, "r", false, "Reverse the sense of comparisons.")
	// -t
	flag.StringVar(&t, "t", "", `Use char as the field separator character; char shall not be considered to be part of 
	a field (although it can be included in a sort  key). Each occurrence of char shall be significant 
	(for example, <char><char> delimits an empty field). If -t is not specified, <blank> characters 
	shall be used as default field separators; each maximal non-empty sequence of  <blank>  characters that 
	follows a non-<blank> shall be a field separator.`)
	// -b
	flag.BoolVar(&b, "b", false, `Ignore leading <blank> characters when determining the starting and ending 
	positions of a restricted sort key. If the	-b  option is specified before the first -k option, it shall 
	be applied to all -k options. Otherwise, the -b option can be attached independently to each -k 
	field_start or field_end option-argument (see below).`)
	// -k
	flag.StringVar(&k, "k", "", `keydef The keydef argument is a restricted sort key field definition. The format of this definition is:
	where field_start and field_end define a key field restricted to a portion of the line (see the EXTENDED DESCRIPTION section), and 
	type is a modifier from the list of characters 'b', 'd', 'f', 'i', 'n', 'r'. The 'b' modifier shall behave like the -b option, but 
	shall apply only to the field_start or field_end to which it is attached. The other modifiers shall behave like the corresponding 
	options, but shall apply only to the key field to which they are attached; they shall have this effect if specified with field_start, 
	field_end, or both. If any modifier is attached to a field_start or to a field_end, no option shall apply to either. 
	Implementations shall support at least nine occurrences of the -k option, which shall be significant in command line order. 
	If no -k option is specified, a default sort key of the entire line shall be used.
	When  there  are  multiple key fields, later keys shall be compared only after all earlier keys compare equal. 
	Except when the -u option is specified, lines that otherwise compare equal shall be ordered as if none of the 
	options -d, -f, -i, -n, or -k were present (but	with -r still in effect, if it was specified) and with all 
	bytes in the lines significant to the comparison. The order in which lines that still compare equal are 
	written is unspecified.`)
	// -h --help
	flag.BoolVar(&h, "h", false, "help command.")
	flag.BoolVar(&h, "help", false, "help command.")
	flag.Parse()

	if h {
		flag.PrintDefaults()
	}

	// recibo todos los los parametros
	for _, fle := range flag.Args() {
		fileData := file.OpenFile(fle)
		var arr []string
		if t != "" {
			arr = strings.Split(string(fileData), t)
		} else {
			arr = strings.Split(string(fileData), "\n")
		}

		switch {
		case m:
			data = append(data, arr...)
		case c && u:
			if len(arr) != len(suprimeDuplicates(arr)) {
				fmt.Fprintln(os.Stderr, "el archivo tiene duplicados")
			}
			sorting(fileData)
		case c || cs:
			sorting(fileData)
		case u:
			data = append(data, suprimeDuplicates(arr)...)
		case d:
			sort.Strings(arr)
			data = append(data, arr...)
		case f:
			sort.Sort(alphabet.Alphabetic(arr))
			data = append(data, arr...)
		case i:
			res := invisibleChar(fileData)
			arrsplit := strings.Split(string(res), "\n")
			sort.StringsAreSorted(arrsplit)
			data = append(data, arrsplit...)
		case b:
			spaces := strings.ReplaceAll(string(fileData), "\n\n", "\n")
			for strings.Contains(string(fileData), "\n\n") {
				spaces = strings.ReplaceAll(spaces, "\n\n", "\n")
				if !strings.Contains(spaces, "\n\n") {
					break
				}
			}
			arr := strings.Split(spaces, "\n")
			data = append(data, arr...)
		case n:
			for _, v := range arr {
				num, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					panic(err)
				}
				values = append(values, int(num))
			}
			sort.Ints(values)
		case k != "":
			i, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}

			var sorting []string // columna por la que se ordenara
			for _, v := range arr {
				part := strings.Split(v, " ")
				sorting = append(sorting, part[i-1])
			}

			sort.Strings(sorting)

			for _, v := range sorting {
				for _, s := range arr {
					if v == strings.Split(s, " ")[i-1] {
						data = append(data, s)
					}
				}
			}

			// fmt.Println(strings.Join(result[:], "\n"))
		default:
			data = append(data, arr...)
		}
	}

	if r {
		reverse.StringArray(data)
	} else if n && r {
		reverse.IntArray(values)
	}

	if n {
		res := strings.ReplaceAll(fmt.Sprint(values), "[", "")
		res = strings.ReplaceAll(res, "]", "")
		// res, err := fmt.Printf("%q", fmt.Sprint(values))
		if o != "" {
			file.WriteFile(o, strings.Split(res, " "))
			os.Exit(0)
		}

		fmt.Println(strings.Join(strings.Split(res, " "), "\n"))
		os.Exit(0)
	}

	if o != "" {
		file.WriteFile(o, data)
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

func compare(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// funcionamiento de variables por orden:
// filename: nombre del archivo
func sorting(data []byte) {
	arr := strings.Split(string(data), "\n")

	// en caso de que este ordenado
	switch {
	case i:
		res := invisibleChar(data)
		arrsplit := strings.Split(string(res), "\n")
		if sort.StringsAreSorted(arrsplit) {
			os.Exit(0)
		}
	case d:
		if sort.StringsAreSorted(arr) {
			os.Exit(0)
		}
	case f:
		arrSorted := strings.Split(string(data), "\n")
		sort.Sort(alphabet.Alphabetic(arrSorted))
		if compare(arr, arrSorted) {
			os.Exit(0)
		}
	case n:
		var values []int
		for _, v := range arr {
			num, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			values = append(values, int(num))
		}
		if sort.IntsAreSorted(values) {
			os.Exit(0)
		}
		// sort.Ints(values)
	}

	switch {
	case c:
		fmt.Fprintln(os.Stderr, "El archivo no esta ordenado")
	case cs:
		fmt.Println("El archivo no esta ordenado")
	}

	os.Exit(1)
}
