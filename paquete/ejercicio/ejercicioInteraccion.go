package ejercicio

import (
	"log"
	"os"
	"fmt"
	"github.com/eiannone/keyboard"
)

func EjercicioApp() {
	fmt.Printf("Pulsá 'a' para crear un ejercicio.\n")
	fmt.Printf("Pulsá 'b' para obtener una lista de todos los ejercicios.\n")
	fmt.Printf("Pulsá 'c' para eliminar un ejercicio por su nombre.\n")
	fmt.Printf("Pulsá 'd' para verificar si ya existe un ejercicio por su nombre.\n")
	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		switch char {
		case 'a':
			crearEjercicio()
			os.Exit(0)
		case 'b':
			listadoDeEjercicios()
			os.Exit(0)
		case 'c':
			eliminarEjercicio()
			os.Exit(0)
		case 'd':
			consultarEjercicio()
			os.Exit(0)
			return
		default:
			fmt.Printf("Tecla '%c' presionada (key code: %v)\n", char, key)
		}
	}
}

func crearEjercicio() {
	fmt.Printf("Se crea ejercicio")
}

func listadoDeEjercicios() {
	fmt.Printf("Se da lista de ejercicios")
}

func eliminarEjercicio() {
	fmt.Printf("Se elimina ejercicio")
}

func consultarEjercicio() {
	fmt.Printf("Se consulta ejercicio")
}