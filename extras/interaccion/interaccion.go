package interaccion

import (
	"TP-2024-TSPORT/extras/tiempo"
	"TP-2024-TSPORT/paquete/ejercicio"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/eiannone/keyboard"
)

func Teclado() {
	if error := keyboard.Open(); error != nil {
		log.Fatal(error)
	}
	defer keyboard.Close()
}

func Portada() {
	texto := "TRABAJO PRACTICO ALGORITMOS Y PROGRAMACION II"
	ancho := 70
	alto := 10
	longitudTexto := len(texto)
	inicioX := (ancho - longitudTexto) / 2
	inicioY := alto / 2
	for i := 0; i < alto; i++ {
		for j := 0; j < ancho; j++ {
			if i == 0 || i == alto-1 {
				fmt.Print("*")
			} else if j == 0 || j == ancho-1 {
				fmt.Print("*")
			} else if i == inicioY && j >= inicioX && j < inicioX+longitudTexto {
				fmt.Print(string(texto[j-inicioX]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Inicio() {
	error := keyboard.Open()
	if error != nil {
		panic(error)
	}
	defer keyboard.Close()
	momentoDelDia := tiempo.Tiempo()
	fmt.Printf("\n----------------------------------------------------------------------\n\n")
	fmt.Printf("¡Hola%s Soy una aplicación para la gestión de ejercicios.\n", momentoDelDia)
	fmt.Printf("Tené en cuenta que si querés finalizar esta aplicación en cualquier momento pulsá 'esc'.\n")
	fmt.Printf("Para comenzar a utilizar mis funciones pulsá 'Enter'\n")
	for {
		_, k, _ := keyboard.GetKey()
		if k == '\x0D' {
			fmt.Printf("\n¡Bien, comencemos con los ejercicios!\n")
			RutinaOEjercicio()
		}
		if k == '\x1B' {
			fmt.Printf("\n¡Gracias por utilizarnos!\n")
			os.Exit(0)
		}
	}

}

func RutinaOEjercicio() {
	error := keyboard.Open()
	if error != nil {
		panic(error)
	}
	defer keyboard.Close()
	fmt.Printf("Elegí entre estas 2 secciones para seguir:\n")
	fmt.Printf("Pulsá 'a' para elegir la sección de Rutina\n")
	fmt.Printf("Pulsá 'b' para elegir la sección de Ejercicio\n")
	for {
		caracter, k, _ := keyboard.GetKey()
		if caracter == 'a' {
			fmt.Printf("Elegiste la sección de Rutina.\n")
			rutinaApp()
		}
		if caracter == 'b' {
			fmt.Printf("Elegiste la sección de Ejercicio.\n")
			ejercicioApp()
		}
		if k == '\x1B' {
			fmt.Printf("\n¡Gracias por utilizarnos!\n")
			os.Exit(0)
		}
		fmt.Printf("Pulsaste una tecla errónea.\n")
	}
}

func FinalizarMain() {
	_, k, _ := keyboard.GetKey()
	//Si se presiona esc
	if k == '\x1B' {
		fmt.Printf("Programa finalizado.")
		os.Exit(0)
	}
}

func ejercicioApp() {
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

		if key == '\x1B' {
			fmt.Printf("\n¡Gracias por utilizarnos!\n")
			os.Exit(0)
		}
	}
}

func crearEjercicio() {
	gestor := ejercicio.NuevoGestorEjercicios()
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Bien, necesito que introduzcas la siguiente información:")

	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Descripción: ")
	descripcion, _ := reader.ReadString('\n')
	descripcion = strings.TrimSpace(descripcion)

	fmt.Print("Tiempo estimado (en segundos): ")
	tiempoStr, _ := reader.ReadString('\n')
	tiempoStr = strings.TrimSpace(tiempoStr)
	tiempo, _ := strconv.Atoi(tiempoStr)

	fmt.Print("Calorías quemadas: ")
	caloriasStr, _ := reader.ReadString('\n')
	caloriasStr = strings.TrimSpace(caloriasStr)
	calorias, _ := strconv.Atoi(caloriasStr)

	fmt.Print("Tipo de ejercicio: ")
	tipos, _ := reader.ReadString('\n')
	tipos = strings.TrimSpace(tipos)

	fmt.Print("Puntos por tipo de ejercicio: ")
	puntosStr, _ := reader.ReadString('\n')
	puntosStr = strings.TrimSpace(puntosStr)
	puntos, _ := strconv.Atoi(puntosStr)

	fmt.Print("Nivel de dificultad: ")
	dificultad, _ := reader.ReadString('\n')
	dificultad = strings.TrimSpace(dificultad)
	nuevoEjercicio := &ejercicio.Ejercicio{
		Nombre:      nombre,
		Descripcion: descripcion,
		TiempoEnSegundos:      tiempo,
		Calorias:    calorias,
		Tipo:        tipos,
		Puntos:      puntos,
		Dificultad:  dificultad,
	}
	err := gestor.AgregarEjercicio(nuevoEjercicio)
	if err != nil {
		log.Fatalf("Error al agregar el ejercicio: %v", err)
	}
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

func rutinaApp() {
	fmt.Printf("Pulsá 'a' para crear una rutina.\n")
	fmt.Printf("Pulsá 'b' para obtener una lista de todas las rutinas.\n")
	fmt.Printf("Pulsá 'c' para eliminar una rutina por su nombre.\n")
	fmt.Printf("Pulsá 'd' para verificar si ya existe una rutina por su nombre.\n")
	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		switch char {
		case 'a':
			crearRutina()
			os.Exit(0)
		case 'b':
			listadoDeRutinas()
			os.Exit(0)
		case 'c':
			eliminarRutina()
			os.Exit(0)
		case 'd':
			consultarRutina()
			os.Exit(0)
			return
		default:
			fmt.Printf("Tecla '%c' presionada (key code: %v)\n", char, key)
		}

		if key == '\x1B' {
			fmt.Printf("\n¡Gracias por utilizarnos!\n")
			os.Exit(0)
		}
	}
}

func crearRutina() {
	gestor := ejercicio.NuevoGestorEjercicios()
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Bien, necesito que introduzcas la siguiente información:")

	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Descripción: ")
	descripcion, _ := reader.ReadString('\n')
	descripcion = strings.TrimSpace(descripcion)

	fmt.Print("Tiempo estimado (en segundos): ")
	tiempoStr, _ := reader.ReadString('\n')
	tiempoStr = strings.TrimSpace(tiempoStr)
	tiempo, _ := strconv.Atoi(tiempoStr)

	fmt.Print("Calorías quemadas: ")
	caloriasStr, _ := reader.ReadString('\n')
	caloriasStr = strings.TrimSpace(caloriasStr)
	calorias, _ := strconv.Atoi(caloriasStr)

	fmt.Print("Tipo de ejercicio: ")
	tipos, _ := reader.ReadString('\n')
	tipos = strings.TrimSpace(tipos)

	fmt.Print("Puntos por tipo de ejercicio: ")
	puntosStr, _ := reader.ReadString('\n')
	puntosStr = strings.TrimSpace(puntosStr)
	puntos, _ := strconv.Atoi(puntosStr)

	fmt.Print("Nivel de dificultad: ")
	dificultad, _ := reader.ReadString('\n')
	dificultad = strings.TrimSpace(dificultad)
	nuevoEjercicio := &ejercicio.Ejercicio{
		Nombre:      nombre,
		Descripcion: descripcion,
		TiempoEnSegundos:      tiempo, // Tiempo estimado en minutos
		Calorias:    calorias,
		Tipo:        tipos,
		Puntos:      puntos,
		Dificultad:  dificultad,
	}
	err := gestor.AgregarEjercicio(nuevoEjercicio)
	if err != nil {
		log.Fatalf("Error al agregar la rutina: %v", err)
	}
}

func listadoDeRutinas() {
	fmt.Printf("Se da lista de rutinas")
}

func eliminarRutina() {
	fmt.Printf("Se elimina rutina")
}

func consultarRutina() {
	fmt.Printf("Se consulta rutina")
}
