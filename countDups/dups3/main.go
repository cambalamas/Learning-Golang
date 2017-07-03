package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	counts = make(map[string]int)
)

// chker recibe un error y lo muestra por LOG
func chker(err error) {
	if err != nil {
		log.Fatalf("[!]: %v", err)
	}
}

// main es el punto de entrada del programa
func main() {
	// Para cada uno de los archivos que se pasan por parametro
	// os.Args[1:] porque os.Args[0] es el nombre del ejecutable
	for _, filename := range os.Args[1:] {
		// ioutil.Readfile() lee en bytes el contenido del archivo
		data, err := ioutil.ReadFile(filename)
		chker(err)
		// Para cada linea (separada por \n) una vez convertido 'data' a cadena
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	// Con lo leido de todos los archivos, imprime si hay mas de una ocurrencia
	// en base al mapa 'counts' (k=texto, v=numero de ocurrencias)
	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
