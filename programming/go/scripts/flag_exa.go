package main

import (
	"flag"    // Paquete para el análisis de banderas de línea de comandos
	"fmt"     // Paquete para la entrada/salida formateada
	"strings" // Paquete para la manipulación de cadenas de texto
)

// Definir banderas de línea de comandos
var n = flag.Bool("n", false, "omitir la nueva línea al final") // Bandera booleana -n para omitir la nueva línea al final
var sep = flag.String("s", " ", "separador")                    // Bandera de cadena -s para especificar el separador (el valor predeterminado es un espacio)

func main() {
	flag.Parse() // Analizar los argumentos de la línea de comandos

	// Imprimir el resultado de unir los argumentos de la línea de comandos usando el separador especificado
	fmt.Print(strings.Join(flag.Args(), *sep))

	// Si la bandera -n no se proporciona, imprimir una nueva línea al final
	if !*n {
		fmt.Println()
	}
}
