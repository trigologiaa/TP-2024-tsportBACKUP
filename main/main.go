package main

import (
	"TP-2024-TSPORT/paquete/almacenamiento"
	"TP-2024-TSPORT/paquete/ejercicio"
	"TP-2024-TSPORT/paquete/rutina"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main es la función principal del programa.
func main() {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	ejercicios, err := almacenamiento.CargarEjerciciosDeUna()
	if err != nil {
		fmt.Println("Error al iniciar lista de ejercicios:", err)
	} else {
		for _, ejercicio := range ejercicios {
			gestorDeEjercicios.AgregarEjercicio(ejercicio)
		}
		fmt.Println("Lista de ejercicios iniciada correctamente.")
	}
	gestorDeRutinas := rutina.NuevoGestorDeRutinas(gestorDeEjercicios)
	rutinas, err := almacenamiento.CargarRutinasDeUna()
	if err != nil {
		fmt.Println("Error al iniciar lista de rutinas:", err)
	} else {
		for _, rutina := range rutinas {
			gestorDeRutinas.AgregarRutina(rutina)
		}
		fmt.Println("Lista de rutinas iniciada correctamente.")
	}
	escaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("\n1. Gestionar Ejercicios")
		fmt.Println("2. Gestionar Rutinas")
		fmt.Println("3. Gestionar Archivos CSV")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			gestionarEjercicios(gestorDeEjercicios, escaner)
		case "2":
			gestionarRutinas(gestorDeEjercicios, gestorDeRutinas, escaner)
		case "3":
			GestionarCSV(gestorDeEjercicios, gestorDeRutinas, escaner)
		case "4":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

// gestionarEjercicios maneja las operaciones relacionadas con la gestión de ejercicios.
//
// Parámetros:
//   - gestorDeEjercicios: Puntero a GestorDeEjercicios, la estructura que maneja los ejercicios.
//   - escaner: Puntero a bufio.Scanner, usado para leer la entrada del usuario.
func gestionarEjercicios(gestorDeEjercicios *ejercicio.GestorDeEjercicios, escaner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Ejercicios ---")
		fmt.Println("\n1. Agregar Ejercicio Nuevo")
		fmt.Println("2. Eliminar Ejercicio por Nombre")
		fmt.Println("3. Consultar Ejercicio por Nombre")
		fmt.Println("4. Modificar Ejercicio por Nombre")
		fmt.Println("5. Listar todos los Ejercicios")
		fmt.Println("6. Volver al Menú Principal")
		fmt.Print("Seleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			fmt.Print("Nombre del ejercicio: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Print("Descripción del ejercicio: ")
			escaner.Scan()
			descripcion := escaner.Text()
			fmt.Print("Tiempo en segundos: ")
			escaner.Scan()
			tiempo, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Calorías quemadas: ")
			escaner.Scan()
			calorias, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Tipo de ejercicio: ")
			escaner.Scan()
			tipo := escaner.Text()
			fmt.Print("Puntos por tipo de ejercicio: ")
			escaner.Scan()
			puntos, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Dificultad del ejercicio: ")
			escaner.Scan()
			dificultad := escaner.Text()
			ejercicio := &ejercicio.Ejercicio {
				NombreDeEjercicio:      		nombre,
				DescripcionDeEjercicio: 		descripcion,
				TiempoEnSegundosDeEjercicio:	tiempo,
				CaloriasDeEjercicio:    		calorias,
				TipoDeEjercicio:        		tipo,
				PuntosPorTipoDeEjercicio:     	puntos,
				DificultadDeEjercicio:  		dificultad,
			}
			err := gestorDeEjercicios.AgregarEjercicio(ejercicio)
			if err != nil {
				fmt.Println("Error al agregar el ejercicio:", err)
			} else {
				fmt.Println("Ejercicio agregado correctamente.")
			}
		case "2":
			fmt.Print("Nombre del ejercicio a eliminar: ")
			escaner.Scan()
			nombre := escaner.Text()
			err := gestorDeEjercicios.EliminarEjercicio(nombre)
			if err != nil {
				fmt.Println("Error al eliminar el ejercicio:", err)
			} else {
				fmt.Println("Ejercicio eliminado correctamente.")
			}
		case "3":
			fmt.Print("Nombre del ejercicio a consultar: ")
			escaner.Scan()
			nombre := escaner.Text()
			ejercicio, err := gestorDeEjercicios.ConsultarEjercicio(nombre)
			if err != nil {
				fmt.Println("Error al consultar el ejercicio:", err)
			} else {
				fmt.Printf("Ejercicio: %+v\n", ejercicio)
			}
		case "4":
			fmt.Print("Nombre del ejercicio a modificar: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Print("Nueva descripción del ejercicio: ")
			escaner.Scan()
			descripcion := escaner.Text()
			fmt.Print("Nuevo tiempo en segundos: ")
			escaner.Scan()
			tiempo, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Nuevas calorías quemadas: ")
			escaner.Scan()
			calorias, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Nuevo tipo de ejercicio: ")
			escaner.Scan()
			tipo := escaner.Text()
			fmt.Print("Nuevos puntos por tipo de ejercicio: ")
			escaner.Scan()
			puntos, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Nueva dificultad del ejercicio: ")
			escaner.Scan()
			dificultad := escaner.Text()
			nuevoEjercicio := &ejercicio.Ejercicio {
				NombreDeEjercicio:           	nombre,
				DescripcionDeEjercicio:      	descripcion,
				TiempoEnSegundosDeEjercicio:	tiempo,
				CaloriasDeEjercicio:         	calorias,
				TipoDeEjercicio:             	tipo,
				PuntosPorTipoDeEjercicio:		puntos,
				DificultadDeEjercicio:       	dificultad,
			}
			err := gestorDeEjercicios.ModificarEjercicio(nombre, nuevoEjercicio)
			if err != nil {
				fmt.Println("Error al modificar el ejercicio:", err)
			} else {
				fmt.Println("Ejercicio modificado correctamente.")
			}
		case "5":
			ejercicios := gestorDeEjercicios.ListarEjercicios()
			for _, ejercicio := range ejercicios {
				fmt.Printf("Ejercicio: %+v\n", ejercicio)
			}
		case "6":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

// gestionarRutinas maneja las operaciones relacionadas con la gestión de rutinas.
//
// Parámetros:
//   - gestorDeEjercicios: Puntero a GestorDeEjercicios, la estructura que maneja los ejercicios.
//   - gestorDeRutinas: Puntero a GestorDeRutinas, la estructura que maneja las rutinas.
//   - escaner: Puntero a bufio.Scanner, usado para leer la entrada del usuario.
func gestionarRutinas(gestorDeEjercicios *ejercicio.GestorDeEjercicios, gestorDeRutinas *rutina.GestorDeRutinas, escaner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Rutinas ---")
		fmt.Println("\n1. Agregar Rutina Nueva")
		fmt.Println("2. Eliminar Rutina por Nombre")
		fmt.Println("3. Consultar Rutina por Nombre")
		fmt.Println("4. Modificar Rutina por Nombre")
		fmt.Println("5. Listar todas las Rutinas")
		fmt.Println("6. Listar Rutinas por Dificultad")
		fmt.Println("7. Generar Rutina Automágica")
		fmt.Println("8. Volver al Menú Principal")
		fmt.Print("Seleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			fmt.Print("\nNombre de la rutina: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Println("\n--- Lista de Ejercicios Disponibles ---")
			ejercicios := gestorDeEjercicios.ListarEjercicios()
			for indice, ejercicio := range ejercicios {
				fmt.Printf("%d. %s\n", indice + 1, ejercicio.NombreDeEjercicio)
			}
			fmt.Println("\nIngrese los números de los ejercicios a incluir en la rutina, separados por comas:")
			escaner.Scan()
			ejerciciosSeleccionadosString := escaner.Text()
			ejerciciosSeleccionados := []ejercicio.Ejercicio{}
			for _, numeroString := range strings.Split(ejerciciosSeleccionadosString, ",") {
				numero, err := strconv.Atoi(strings.TrimSpace(numeroString))
				if err != nil || numero < 1 || numero > len(ejercicios) {
					fmt.Println("Número de ejercicio inválido:", numeroString)
					continue
				}
				ejerciciosSeleccionados = append(ejerciciosSeleccionados, *ejercicios[numero - 1])
			}
			rutina := &rutina.Rutina {
				NombreDeRutina:     					nombre,
				CaracteristicasIndividualesDeRutina:	ejerciciosSeleccionados,
			}
			rutina.CalcularPropiedades(gestorDeEjercicios)
			err := gestorDeRutinas.AgregarRutina(rutina)
			if err != nil {
				fmt.Println("Error al agregar la rutina:", err)
			} else {
				fmt.Println("Rutina agregada correctamente.")
			}
		case "2":
			fmt.Print("Nombre de la rutina a eliminar: ")
			escaner.Scan()
			nombre := escaner.Text()
			err := gestorDeRutinas.EliminarRutina(nombre)
			if err != nil {
				fmt.Println("Error al eliminar la rutina:", err)
			} else {
				fmt.Println("Rutina eliminada correctamente.")
			}
		case "3":
			fmt.Print("Nombre de la rutina a consultar: ")
			escaner.Scan()
			nombre := escaner.Text()
			rutina, err := gestorDeRutinas.ConsultarRutina(nombre)
			if err != nil {
				fmt.Println("Error al consultar la rutina:", err)
			} else {
				fmt.Printf("Rutina: %+v\n", rutina)
			}
		case "4":
			fmt.Print("Nombre de la rutina a modificar: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Print("Nuevo nombre de la rutina: ")
			escaner.Scan()
			nuevoNombre := escaner.Text()
			fmt.Println("¿Desea modificar los ejercicios de la rutina? (si/no): ")
			escaner.Scan()
			modificarEjercicios := escaner.Text() == "si"
			var ejerciciosSeleccionados []ejercicio.Ejercicio
			if modificarEjercicios {
				fmt.Println("\n--- Lista de Ejercicios Disponibles ---")
				ejercicios := gestorDeEjercicios.ListarEjercicios()
				for indice, ejercicio := range ejercicios {
					fmt.Printf("%d. %s\n", indice + 1, ejercicio.NombreDeEjercicio)
				}
				fmt.Println("\nIngrese los números de los ejercicios a incluir en la rutina, separados por comas:")
				escaner.Scan()
				ejerciciosSeleccionadosString := escaner.Text()
				ejerciciosSeleccionados = []ejercicio.Ejercicio{}
				for _, numeroString := range strings.Split(ejerciciosSeleccionadosString, ",") {
					numero, err := strconv.Atoi(strings.TrimSpace(numeroString))
					if err != nil || numero < 1 || numero > len(ejercicios) {
						fmt.Println("Número de ejercicio inválido:", numeroString)
						continue
					}
					ejerciciosSeleccionados = append(ejerciciosSeleccionados, *ejercicios[numero - 1])
				}
			}
			nuevaRutina := &rutina.Rutina {
				NombreDeRutina:     					nuevoNombre,
				CaracteristicasIndividualesDeRutina:	ejerciciosSeleccionados,
			}
			nuevaRutina.CalcularPropiedades(gestorDeEjercicios)
			err := gestorDeRutinas.ModificarRutina(nombre, nuevaRutina)
			if err != nil {
				fmt.Println("Error al modificar la rutina:", err)
			} else {
				fmt.Println("Rutina modificada correctamente.")
			}
		case "5":
			fmt.Println("\n--- Listar Rutinas ---")
    		rutinas := gestorDeRutinas.ListarRutinas()
    		if len(rutinas) == 0 {
        		fmt.Println("No hay rutinas disponibles.")
    		} else {
        		fmt.Println("Rutinas disponibles:")
        		for indice, rutina := range rutinas {
            		fmt.Printf("%d. %s\n", indice + 1, rutina.NombreDeRutina)
        		}
    		}
		case "6":
			fmt.Print("Dificultad de las rutinas a listar: ")
			escaner.Scan()
			dificultad := escaner.Text()
			rutinas := gestorDeRutinas.ListarRutinasPorDificultad(dificultad)
			for _, rutina := range rutinas {
				fmt.Printf("Rutina: %+v\n", rutina)
			}
		case "7":
			gestionarRutinasAutomagicas(gestorDeEjercicios, gestorDeRutinas, escaner)
		case "8":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

// gestionarRutinasAutomagicas permite la gestión de la creación automática de rutinas de ejercicio a través de una interfaz de consola.
//
// Parámetros:
//   - gestorDeEjercicios: puntero a una instancia de GestorDeEjercicios que maneja los ejercicios disponibles.
//   - gestorDeRutinas: puntero a una instancia de GestorDeRutinas que maneja las rutinas creadas.
//   - escaner: puntero a un bufio.Scanner utilizado para la entrada de datos por el usuario.
func gestionarRutinasAutomagicas(gestorDeEjercicios *ejercicio.GestorDeEjercicios, gestorDeRutinas *rutina.GestorDeRutinas, escaner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Rutinas AUTOMÁGICAS ---")
		fmt.Println("\n1. RUTINA AUTOMÁGICA 1 (Máxima cantidad: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que maximicen la cantidad de ejercicios a realizar en el tiempo definido)")
		fmt.Println("2. RUTINA AUTOMÁGICA 2 (Mínima duración: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que minimicen la duración de la rutina)")
		fmt.Println("3. RUTINA AUTOMÁGICA 3 (Duración fija, Máximos puntos: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que maximicen el puntaje de la rutina en la dimensión solicitada)")
		fmt.Println("4. Volver al menú anterior")
		fmt.Print("\nSeleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			fmt.Print("\nNombre de la rutina: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Print("Duración total de la rutina (en minutos): ")
			escaner.Scan()
			duracionEnMinutos, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Tipo(s) de ejercicios a incluir (separados por comas): ")
			escaner.Scan()
			entradaDeTipos := escaner.Text()
			tipos := strings.Split(entradaDeTipos, ",")
			for indice, tipo := range tipos {
				tipos[indice] = strings.TrimSpace(tipo)
			}
			fmt.Print("Dificultad: ")
			escaner.Scan()
			dificultad := escaner.Text()
			ejerciciosFiltrados := gestorDeEjercicios.FiltrarPorTiposYDificultad(tipos, dificultad)
			if len(ejerciciosFiltrados) == 0 {
				fmt.Println("No se encontraron ejercicios válidos para esta rutina")
				return
			}
			ejerciciosOrdenados := gestorDeEjercicios.OrdenarTiempoMenorAMayor(ejerciciosFiltrados)
			duracionTotalEnSegundos := duracionEnMinutos * 60
			rutinaAutomagica := &rutina.Rutina {
				NombreDeRutina:	nombre,
			}
			duracionAcumulada := 0
			for _, ejercicio := range ejerciciosOrdenados {
				if duracionAcumulada + ejercicio.TiempoEnSegundosDeEjercicio <= duracionTotalEnSegundos {
					rutinaAutomagica.CaracteristicasIndividualesDeRutina = append(rutinaAutomagica.CaracteristicasIndividualesDeRutina, *ejercicio)
					duracionAcumulada += ejercicio.TiempoEnSegundosDeEjercicio
				} else {
					break
				}
			}
			rutinaAutomagica.CalcularPropiedades(gestorDeEjercicios)
			err := gestorDeRutinas.AgregarRutina(rutinaAutomagica)
			if err != nil {
				fmt.Println("Error al generar la rutina automágica:", err)
			} else {
				fmt.Println("Rutina automágica generada correctamente.")
			}
		case "2":
			fmt.Print("\nNombre de la rutina: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Print("Calorías totales a quemar: ")
			escaner.Scan()
			caloriasTotales, _ := strconv.Atoi(escaner.Text())
			ejercicios := gestorDeEjercicios.ListarEjercicios()
			ejerciciosFiltrados := []*ejercicio.Ejercicio{}
			for _, ejercicio := range ejercicios {
				if ejercicio.CaloriasDeEjercicio <= caloriasTotales {
					ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
				}
			}
			ejerciciosOrdenados := gestorDeEjercicios.OrdenarTiempoMenorAMayor(ejerciciosFiltrados)
			rutinaAutomagica := &rutina.Rutina {
				NombreDeRutina:	nombre,
			}
			caloriasAcumuladas := 0
			for _, ejercicio := range ejerciciosOrdenados {
				if caloriasAcumuladas + ejercicio.CaloriasDeEjercicio <= caloriasTotales {
					rutinaAutomagica.CaracteristicasIndividualesDeRutina = append(rutinaAutomagica.CaracteristicasIndividualesDeRutina, *ejercicio)
					caloriasAcumuladas += ejercicio.CaloriasDeEjercicio
				} else {
					break
				}
			}
			rutinaAutomagica.CalcularPropiedades(gestorDeEjercicios)
			err := gestorDeRutinas.AgregarRutina(rutinaAutomagica)
			if err != nil {
				fmt.Println("Error al generar la rutina automágica:", err)
			} else {
				fmt.Println("Rutina automágica generada correctamente.")
			}
		case "3":
			fmt.Print("\nNombre de la rutina: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Print("Tipo de puntos a maximizar (cardio, fuerza, flexibilidad): ")
			escaner.Scan()
			puntosPorTipo := escaner.Text()
			fmt.Print("Duración máxima de la rutina (en minutos): ")
			escaner.Scan()
			duracionMaximaEnMinutos, _ := strconv.Atoi(escaner.Text())
			ejerciciosFiltrados := gestorDeEjercicios.FiltrarPorTipoPuntosYDuracionMaxima(puntosPorTipo, duracionMaximaEnMinutos)
			if len(ejerciciosFiltrados) == 0 {
				fmt.Println("No se encontraron ejercicios válidos para esta rutina")
				return
			}
			ejerciciosOrdenados := gestorDeEjercicios.OrdenarPorPuntajeDescendente(ejerciciosFiltrados)
			duracionTotalEnSegundos := duracionMaximaEnMinutos * 60
			rutinaAutomagica := &rutina.Rutina {
				NombreDeRutina:	nombre,
			}
			puntajeTotal := 0
			duracionTotal := 0
			ejerciciosUtilizados := make(map[string]bool)
			for _, ejercicio := range ejerciciosOrdenados {
				if duracionTotal + ejercicio.TiempoEnSegundosDeEjercicio <= duracionTotalEnSegundos && !ejerciciosUtilizados[ejercicio.NombreDeEjercicio] {
					rutinaAutomagica.CaracteristicasIndividualesDeRutina = append(rutinaAutomagica.CaracteristicasIndividualesDeRutina, *ejercicio)
					puntajeTotal += ejercicio.PuntosPorTipoDeEjercicio
					duracionTotal += ejercicio.TiempoEnSegundosDeEjercicio
					ejerciciosUtilizados[ejercicio.NombreDeEjercicio] = true
				}
			}
			rutinaAutomagica.PuntosDeRutina = puntajeTotal
			rutinaAutomagica.TiempoEnMinutosDeRutina = duracionTotal
			err := gestorDeRutinas.AgregarRutina(rutinaAutomagica)
			if err != nil {
				fmt.Println("Error al generar la rutina automágica:", err)
			} else {
				fmt.Println("Rutina automágica generada correctamente.")
			}
		case "4":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

// GestionarCSV maneja la importación y exportación de datos en formato CSV.
//
// Parámetros:
//   - gestorDeEjercicios: Puntero a GestorDeEjercicios, la estructura que maneja los ejercicios.
//   - gestorDeRutinas: Puntero a GestorDeRutinas, la estructura que maneja las rutinas.
//   - escaner: Puntero a bufio.Scanner, usado para leer la entrada del usuario.
func GestionarCSV(gestorDeEjercicios *ejercicio.GestorDeEjercicios, gestorDeRutinas *rutina.GestorDeRutinas, escaner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Archivos CSV ---")
		fmt.Println("\n1. Guardar Ejercicios en CSV")
		fmt.Println("2. Cargar Ejercicios desde CSV")
		fmt.Println("3. Guardar Rutinas en CSV")
		fmt.Println("4. Cargar Rutinas desde CSV")
		fmt.Println("5. Volver al Menú Principal")
		fmt.Print("Seleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			fmt.Print("Ingrese el nombre del archivo para guardar los ejercicios: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				err := almacenamiento.GuardarEjercicios(gestorDeEjercicios.ListarEjercicios(), nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al guardar los ejercicios:", err)
				} else {
					fmt.Println("Ejercicios guardados correctamente.")
				}
			}
		case "2":
			fmt.Print("Ingrese el nombre del archivo desde donde cargar los ejercicios: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				ejercicios, err := almacenamiento.CargarEjercicios(nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al cargar los ejercicios:", err)
				} else {
					for _, ejercicio := range ejercicios {
						gestorDeEjercicios.AgregarEjercicio(ejercicio)
					}
					fmt.Println("Ejercicios cargados correctamente.")
				}
			}
		case "3":
			fmt.Print("Ingrese el nombre del archivo para guardar las rutinas: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				err := almacenamiento.GuardarRutinas(gestorDeRutinas.ListarRutinas(), gestorDeEjercicios, nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al guardar las rutinas:", err)
				} else {
					fmt.Println("Rutinas guardadas correctamente.")
				}
			}
		case "4":
			fmt.Print("Ingrese el nombre del archivo desde donde cargar las rutinas: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				rutinas, err := almacenamiento.CargarRutinas(nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al cargar las rutinas:", err)
				} else {
					for _, rutina := range rutinas {
						gestorDeRutinas.AgregarRutina(rutina)
					}
					fmt.Println("Rutinas cargadas correctamente.")
				}
			}
		case "5":
			fmt.Println("Volviendo al Menú Principal...")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}