package alphabet

import "unicode"

type Alphabetic []string

func (a Alphabetic) Len() int { return len(a) }

func (a Alphabetic) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// determino cual es menor en cada posicion
func (a Alphabetic) Less(i, j int) bool {
	// rune, son la representacion en unicode, en verdad es un int32
	// convierto a en la posicion de i y j en rune
	iRunes, jRunes := []rune(a[i]), []rune(a[j])

	// determino cual es mayor de los dos arreglos
	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for i := 0; i < max; i++ {
		// determino cual es menor en cada posicion
		ir, jr := iRunes[i], jRunes[i]

		// lo llevo a minusculas, para que sean iguales
		lir, ljr := unicode.ToLower(ir), unicode.ToLower(jr)

		// comparo si son diferentes
		if lir != ljr {
			return lir < ljr
		}

		// en caso de que las rune sean las mismas al ser minusculas ambas
		// comparo el original
		if ir != jr {
			return ir < jr
		}
	}

	return len(iRunes) < len(jRunes)
}
