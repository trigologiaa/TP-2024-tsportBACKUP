package interaccion

import (
	"TP-2024-TSPORT/extras/tiempo"
	"TP-2024-TSPORT/paquete/ejercicio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"bufio"
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
			if i == 0 || i == alto - 1 {
				fmt.Print("*")
			} else if j == 0 || j == ancho - 1 {
				fmt.Print("*")
			} else if i == inicioY && j >= inicioX && j < inicioX + longitudTexto {
				fmt.Print(string(texto[j - inicioX]))
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
        caracter, _, _ := keyboard.GetKey()
        if caracter == 'a' {
            fmt.Printf("Elegiste la sección de Rutina.\n")
            CrearEjercicio()
        }
		if caracter == 'b' {
            fmt.Printf("Elegiste la sección de Ejercicio.\n")
            ejercicio.EjercicioApp()
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

func CrearEjercicio() {
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
	nuevoEjercicio := &ejercicio.Ejercicio {
		Nombre:      nombre,
        Descripcion: descripcion,
        Tiempo:      tiempo, // Tiempo estimado en minutos
        Calorias:    calorias,
        Tipo:       tipos,
        Puntos:      puntos,
        Dificultad:  dificultad,
	}
	err := gestor.AgregarEjercicio(nuevoEjercicio)
	if err != nil {
		log.Fatalf("Error al agregar el ejercicio: %v", err)
	}
}