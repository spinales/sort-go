package file

import (
	"io/ioutil"
	"strings"
)

// abre el archivo por la ruta pasada
func OpenFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}

func WriteFile(filepath string, data []string) {
	err := ioutil.WriteFile(filepath, []byte(strings.Join(data, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}
