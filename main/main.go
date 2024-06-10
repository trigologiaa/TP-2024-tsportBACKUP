package main

import (
	"TP-2024-TSPORT/paquete/almacenamiento"
	"TP-2024-TSPORT/paquete/ejercicio"
	"TP-2024-TSPORT/paquete/rutina"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// main es la función principal del programa. Inicializa los gestores de ejercicios y rutinas, y proporciona un menú interactivo para que el usuario gestione ejercicios y rutinas.
//
// Funcionamiento:
//
//	Se crea una instancia del gestor de ejercicios utilizando ejercicio.NuevoGestorEjercicios().
//	Se crea una instancia del gestor de rutinas utilizando rutina.NuevoGestorRutinas(gestorEjercicios).
//	Se inicializa un scanner para leer la entrada estándar (os.Stdin).
//	Se entra en un bucle infinito que muestra un menú y procesa las opciones seleccionadas por el usuario.
//	Dentro del bucle {
//		Se imprime el menú principal con las opciones disponibles.
//		Se lee la entrada del usuario utilizando scanner.Scan() y scanner.Text().
//		Se utiliza una estructura switch para manejar las diferentes opciones del menú {
//			- Caso "1": Llama a gestionarEjercicios(gestorEjercicios, scanner).
//			- Caso "2": Llama a gestionarRutinas(gestorRutinas, scanner).
//			- Caso "3": Solicita el nombre de un archivo y guarda los ejercicios en un CSV utilizando almacenamiento.GuardarEjercicios().
//			- Caso "4": Solicita el nombre de un archivo y carga los ejercicios desde un CSV utilizando almacenamiento.CargarEjercicios() y agrega cada ejercicio al gestor de ejercicios.
//			- Caso "5": Solicita el nombre de un archivo y guarda las rutinas en un CSV utilizando almacenamiento.GuardarRutinas().
//			- Caso "6": Solicita el nombre de un archivo y carga las rutinas desde un CSV utilizando almacenamiento.CargarRutinas() y agrega cada rutina al gestor de rutinas.
//			- Caso "7": Imprime un mensaje de salida y termina el programa.
//			- Default: Imprime un mensaje indicando que la opción no es válida.
//		}
//		El bucle se rompe si scanner.Scan() devuelve false, indicando el fin de la entrada o un error.
//	}
func main() {
	gestorEjercicios := ejercicio.NuevoGestorEjercicios()
	ejercicios, err := almacenamiento.CargarEjerciciosDeUna()
	if err != nil {
		fmt.Println("Error al iniciar los ejercicios:", err)
	} else {
		for _, ej := range ejercicios {
			gestorEjercicios.AgregarEjercicio(ej)
		}
		fmt.Println("Ejercicios iniciados correctamente.")
	}

	gestorRutinas := rutina.NuevoGestorRutinas(gestorEjercicios)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("\n1. Gestionar Ejercicios")
		fmt.Println("2. Gestionar Rutinas")
		fmt.Println("3. Guardar Ejercicios en CSV")
		fmt.Println("4. Cargar Ejercicios desde CSV")
		fmt.Println("5. Guardar Rutinas en CSV")
		fmt.Println("6. Cargar Rutinas desde CSV")
		fmt.Println("7. Salir")
		fmt.Print("Seleccione una opción: ")
		if !scanner.Scan() {
			break
		}
		opcion := scanner.Text()
		switch opcion {
		case "1":
			gestionarEjercicios(gestorEjercicios, scanner)
		case "2":
			gestionarRutinas(gestorEjercicios, gestorRutinas, scanner)
		case "3":
			fmt.Print("Ingrese el nombre del archivo para guardar los ejercicios: ")
			if scanner.Scan() {
				nombreDeArchivo := scanner.Text()
				err := almacenamiento.GuardarEjercicios(gestorEjercicios.ObtenerTodosLosEjercicios(), nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al guardar los ejercicios:", err)
				} else {
					fmt.Println("Ejercicios guardados correctamente.")
				}
			}
		case "4":
			fmt.Print("Ingrese el nombre del archivo desde donde cargar los ejercicios: ")
			if scanner.Scan() {
				nombreDeArchivo := scanner.Text()
				ejercicios, err := almacenamiento.CargarEjercicios(nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al cargar los ejercicios:", err)
				} else {
					for _, ej := range ejercicios {
						gestorEjercicios.AgregarEjercicio(ej)
					}
					fmt.Println("Ejercicios cargados correctamente.")
				}
			}
		case "5":
			fmt.Print("Ingrese el nombre del archivo para guardar las rutinas: ")
			if scanner.Scan() {
				nombreDeArchivo := scanner.Text()
				err := almacenamiento.GuardarRutinas(gestorRutinas.ObtenerTodasLasRutinas(), gestorEjercicios, nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al guardar las rutinas:", err)
				} else {
					fmt.Println("Rutinas guardadas correctamente.")
				}
			}
		case "6":
			fmt.Print("Ingrese el nombre del archivo desde donde cargar las rutinas: ")
			if scanner.Scan() {
				nombreDeArchivo := scanner.Text()
				rutinas, err := almacenamiento.CargarRutinas(nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al cargar las rutinas:", err)
				} else {
					for _, rt := range rutinas {
						gestorRutinas.AgregarRutina(rt)
					}
					fmt.Println("Rutinas cargadas correctamente.")
				}
			}
		case "7":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func gestionarEjercicios(gestor *ejercicio.GestorEjercicios, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Ejercicios ---")
		fmt.Println("\n1. Agregar Ejercicio")
		fmt.Println("2. Eliminar Ejercicio")
		fmt.Println("3. Consultar Ejercicio")
		fmt.Println("4. Listar Ejercicios")
		fmt.Println("5. Volver al Menú Principal")
		fmt.Print("Seleccione una opción: ")
		if !scanner.Scan() {
			break
		}
		opcion := scanner.Text()
		switch opcion {
		case "1":
			fmt.Print("Nombre del ejercicio: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Descripción del ejercicio: ")
			scanner.Scan()
			descripcion := scanner.Text()
			fmt.Print("Tiempo en segundos: ")
			scanner.Scan()
			tiempo, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Calorías quemadas: ")
			scanner.Scan()
			calorias, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Tipo de ejercicio: ")
			scanner.Scan()
			tipo := scanner.Text()
			fmt.Print("Puntos por tipo de ejercicio: ")
			scanner.Scan()
			puntos, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Dificultad del ejercicio: ")
			scanner.Scan()
			dificultad := scanner.Text()
			ej := &ejercicio.Ejercicio{
				Nombre:      nombre,
				Descripcion: descripcion,
				Tiempo:      tiempo,
				Calorias:    calorias,
				Tipo:        tipo,
				Puntos:      puntos,
				Dificultad:  dificultad,
			}
			err := gestor.AgregarEjercicio(ej)
			if err != nil {
				fmt.Println("Error al agregar el ejercicio:", err)
			} else {
				fmt.Println("Ejercicio agregado correctamente.")
			}
		case "2":
			fmt.Print("Nombre del ejercicio a eliminar: ")
			scanner.Scan()
			nombre := scanner.Text()
			err := gestor.EliminarEjercicio(nombre)
			if err != nil {
				fmt.Println("Error al eliminar el ejercicio:", err)
			} else {
				fmt.Println("Ejercicio eliminado correctamente.")
			}
		case "3":
			fmt.Print("Nombre del ejercicio a consultar: ")
			scanner.Scan()
			nombre := scanner.Text()
			ej, err := gestor.ConsultarEjercicio(nombre)
			if err != nil {
				fmt.Println("Error al consultar el ejercicio:", err)
			} else {
				fmt.Printf("Ejercicio: %+v\n", ej)
			}
		case "4":
			ejercicios := gestor.ObtenerTodosLosEjercicios()
			for _, ej := range ejercicios {
				fmt.Printf("Ejercicio: %+v\n", ej)
			}
		case "5":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func gestionarRutinas(gestorEjercicios *ejercicio.GestorEjercicios, gestor *rutina.GestorRutinas, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Rutinas ---")
		fmt.Println("\n1. Agregar Rutina")
		fmt.Println("2. Eliminar Rutina")
		fmt.Println("3. Consultar Rutina")
		fmt.Println("4. Modificar Rutina")
		fmt.Println("5. Listar Rutinas por Dificultad")
		fmt.Println("6. Generar Rutina Automágica")
		fmt.Println("7. Volver al Menú Principal")
		fmt.Print("Seleccione una opción: ")
		if !scanner.Scan() {
			break
		}
		opcion := scanner.Text()
		switch opcion {
		case "1":
			fmt.Print("\nNombre de la rutina: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Tiempo total (en segundos): ")
			scanner.Scan()
			tiempo, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Calorías quemadas: ")
			scanner.Scan()
			calorias, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Dificultad: ")
			scanner.Scan()
			dificultad := scanner.Text()
			fmt.Print("Tipo de rutina: ")
			scanner.Scan()
			tipo := scanner.Text()
			fmt.Print("Puntos por tipo: ")
			scanner.Scan()
			puntos, _ := strconv.Atoi(scanner.Text())
			rut := &rutina.Rutina{
				Nombre:        nombre,
				Tiempo:        tiempo,
				Calorias:      calorias,
				Dificultad:    dificultad,
				Tipos:         tipo,
				PuntosPorTipo: puntos,
			}
			err := gestor.AgregarRutina(rut)
			if err != nil {
				fmt.Println("Error al agregar la rutina:", err)
			} else {
				fmt.Println("Rutina agregada correctamente.")
			}
		case "2":
			fmt.Print("Nombre de la rutina a eliminar: ")
			scanner.Scan()
			nombre := scanner.Text()
			err := gestor.EliminarRutina(nombre)
			if err != nil {
				fmt.Println("Error al eliminar la rutina:", err)
			} else {
				fmt.Println("Rutina eliminada correctamente.")
			}
		case "3":
			fmt.Print("Nombre de la rutina a consultar: ")
			scanner.Scan()
			nombre := scanner.Text()
			rut, err := gestor.ConsultarRutina(nombre)
			if err != nil {
				fmt.Println("Error al consultar la rutina:", err)
			} else {
				fmt.Printf("Rutina: %+v\n", rut)
			}
		case "4":
			fmt.Print("Nombre de la rutina a modificar: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Nuevo nombre de la rutina: ")
			scanner.Scan()
			nuevoNombre := scanner.Text()
			fmt.Print("Nuevo tiempo total (en segundos): ")
			scanner.Scan()
			tiempo, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Nuevas calorías quemadas: ")
			scanner.Scan()
			calorias, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Nueva dificultad: ")
			scanner.Scan()
			dificultad := scanner.Text()
			fmt.Print("Nuevo tipo de rutina: ")
			scanner.Scan()
			tipo := scanner.Text()
			fmt.Print("Nuevos puntos por tipo: ")
			scanner.Scan()
			puntos, _ := strconv.Atoi(scanner.Text())
			nuevaRut := &rutina.Rutina{
				Nombre:        nuevoNombre,
				Tiempo:        tiempo,
				Calorias:      calorias,
				Dificultad:    dificultad,
				Tipos:         tipo,
				PuntosPorTipo: puntos,
			}
			err := gestor.ModificarRutina(nombre, nuevaRut)
			if err != nil {
				fmt.Println("Error al modificar la rutina:", err)
			} else {
				fmt.Println("Rutina modificada correctamente.")
			}
		case "5":
			fmt.Print("Dificultad de las rutinas a listar: ")
			scanner.Scan()
			dificultad := scanner.Text()
			rutinas := gestor.ListarRutinas(dificultad)
			for _, rut := range rutinas {
				fmt.Printf("Rutina: %+v\n", rut)
			}
		case "6":
			gestionarRutinasAutomagicas(gestorEjercicios, gestor, scanner)
		case "7":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func gestionarRutinasAutomagicas(gestorEjercicios *ejercicio.GestorEjercicios, gestor *rutina.GestorRutinas, scanner *bufio.Scanner) {
	todosLosEjercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
	for {
		fmt.Println("\n--- Gestión de Rutinas AUTOMÁGICAS ---")
		fmt.Println("\n1. RUTINA AUTOMÁGICA 1 (Máxima cantidad: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que maximicen la cantidad de ejercicios a realizar en el tiempo definido)")
		fmt.Println("2. RUTINA AUTOMÁGICA 2 (Mínima duración: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que minimicen la duración de la rutina)")
		fmt.Println("3. RUTINA AUTOMÁGICA 3 (Duración fija, Máximos puntos: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que maximicen el puntaje de la rutina en la dimensión solicitada)")
		fmt.Println("4. Volver al menú anterior")
		fmt.Print("Seleccione una opción: ")
		if !scanner.Scan() {
			break
		}
		opcion := scanner.Text()
		switch opcion {
		case "1":
			fmt.Print("\nNombre de la rutina: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Duración total en minutos: ")
			scanner.Scan()
			duracion, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Tipo: ")
			scanner.Scan()
			tipo := scanner.Text()
			fmt.Print("Dificultad: ")
			scanner.Scan()
			dificultad := scanner.Text()
			rutina, err := gestor.GenerarRutinaAutomagica1(todosLosEjercicios, nombre, duracion, tipo, dificultad)
			if err != nil {
				fmt.Println("Error al generar la rutina automágica:", err)
			} else {
				fmt.Println("Rutina automágica generada correctamente:")
				fmt.Printf("%+v\n", rutina)
			}
			gestor.AgregarRutina(rutina)
		case "2":
			fmt.Print("\nNombre de la rutina: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Calorías totales a quemar: ")
			scanner.Scan()
			calorias, _ := strconv.Atoi(scanner.Text())
			rutina, err := gestor.GenerarRutinaAutomagica2(todosLosEjercicios, nombre, calorias)
			if err != nil {
				fmt.Println("Error al generar la rutina automágica:", err)
			} else {
				fmt.Println("Rutina automágica generada correctamente:")
				fmt.Printf("%+v\n", rutina)
			}

		case "3":
			/*fmt.Print("\nNombre de la rutina: ")
						scanner.Scan()
			    		nombre := scanner.Text()
						fmt.Print("Tipos de puntos a maximizar: ")
						scanner.Scan()
						tipos := scanner.Text()
						fmt.Print("Duración máxima de la rutina: ")
						scanner.Scan()
						duracion := scanner.Text()*/
		case "4":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}
