package main

import (
	"TP-2024-TSPORT/paquete/almacenamiento"
	"TP-2024-TSPORT/paquete/ejercicio"
	"TP-2024-TSPORT/paquete/rutina"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// main es la función principal del programa.
func main() {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	ejercicios, err := almacenamiento.CargarEjerciciosDeUna()
	if err != nil {
		fmt.Println("\nError al iniciar lista de ejercicios:", err)
	} else {
		for _, ejercicio := range ejercicios {
			gestorDeEjercicios.AgregarEjercicio(ejercicio)
		}
		fmt.Println("\nSe inicia una lista predeterminada de ejercicios.")
	}
	gestorDeRutinas := rutina.NuevoGestorDeRutinas(gestorDeEjercicios)
	rutinas, err := almacenamiento.CargarRutinasDeUna()
	if err != nil {
		fmt.Println("Error al iniciar lista de rutinas:", err)
	} else {
		for _, rutina := range rutinas {
			gestorDeRutinas.AgregarRutina(rutina)
		}
		fmt.Println("Se inicia una lista predeterminada de rutinas.")
	}
	escaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("\n1. Gestionar Ejercicios")
		fmt.Println("2. Gestionar Rutinas")
		fmt.Println("3. Gestionar Archivos CSV")
		fmt.Println("4. Salir")
		fmt.Print("\nSeleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			ClearScreen()
			gestionarEjercicios(gestorDeEjercicios, escaner)
		case "2":
			ClearScreen()
			gestionarRutinas(gestorDeEjercicios, gestorDeRutinas, escaner)
		case "3":
			ClearScreen()
			GestionarCSV(gestorDeEjercicios, gestorDeRutinas, escaner)
		case "4":
			ClearScreen()
			fmt.Println("\nSaliendo...")
			return
		default:
			fmt.Println("\nOpción no válida. Intente de nuevo.")
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
		fmt.Println("6. Listar Ejercicios por Tipo y/o Dificultad")
		fmt.Println("7. Listar Ejercicios por Calorias")
		fmt.Println("8. Volver al Menú Principal")
		fmt.Print("\nSeleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			ClearScreen()
			fmt.Print("\nNombre del ejercicio: ")
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
			ejercicio := &ejercicio.Ejercicio{
				NombreDeEjercicio:           nombre,
				DescripcionDeEjercicio:      descripcion,
				TiempoEnSegundosDeEjercicio: tiempo,
				CaloriasDeEjercicio:         calorias,
				TipoDeEjercicio:             tipo,
				PuntosPorTipoDeEjercicio:    puntos,
				DificultadDeEjercicio:       dificultad,
			}
			err := gestorDeEjercicios.AgregarEjercicio(ejercicio)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al agregar el ejercicio:", err)
			} else {
				fmt.Println("\nEjercicio agregado correctamente.")
				fmt.Printf("\nNombre: %+v\n", ejercicio.NombreDeEjercicio)
				fmt.Printf("Descripción: %+v\n", ejercicio.DescripcionDeEjercicio)
				fmt.Printf("Tiempo en segundos: %+v\n", ejercicio.TiempoEnSegundosDeEjercicio)
				fmt.Printf("Calorías quemadas: %+v\n", ejercicio.CaloriasDeEjercicio)
				fmt.Printf("Tipo de ejercicio: %+v\n", ejercicio.TipoDeEjercicio)
				fmt.Printf("Puntos por tipo de ejercicio: %+v\n", ejercicio.PuntosPorTipoDeEjercicio)
				fmt.Printf("Dificultad: %+v\n", ejercicio.DificultadDeEjercicio)
			}
		case "2":
			ClearScreen()
			fmt.Print("\nNombre del ejercicio a eliminar: ")
			escaner.Scan()
			nombre := escaner.Text()
			err := gestorDeEjercicios.EliminarEjercicio(nombre)
			if err != nil {
				fmt.Println("\nError al eliminar el ejercicio:", err)
			} else {
				fmt.Println("\nEjercicio eliminado correctamente.")
			}
		case "3":
			ClearScreen()
			fmt.Print("\nNombre del ejercicio a consultar: ")
			escaner.Scan()
			nombre := escaner.Text()
			ejercicio, err := gestorDeEjercicios.ConsultarEjercicio(nombre)
			if err != nil {
				fmt.Println("\nError al consultar el ejercicio:", err)
			} else {
				fmt.Printf("\nNombre: %+v\n", ejercicio.NombreDeEjercicio)
				fmt.Printf("Descripción: %+v\n", ejercicio.DescripcionDeEjercicio)
				fmt.Printf("Tiempo en segundos: %+v\n", ejercicio.TiempoEnSegundosDeEjercicio)
				fmt.Printf("Calorías quemadas: %+v\n", ejercicio.CaloriasDeEjercicio)
				fmt.Printf("Tipo de ejercicio: %+v\n", ejercicio.TipoDeEjercicio)
				fmt.Printf("Puntos por tipo de ejercicio: %+v\n", ejercicio.PuntosPorTipoDeEjercicio)
				fmt.Printf("Dificultad: %+v\n", ejercicio.DificultadDeEjercicio)
			}
		case "4":
			ClearScreen()
			fmt.Print("Nombre del ejercicio a modificar: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Print("Nuevo nombre del ejercicio: ")
			escaner.Scan()
			nuevoNombre := escaner.Text()
			fmt.Print("Nueva descripción del ejercicio: ")
			escaner.Scan()
			nuevaDescripcion := escaner.Text()
			fmt.Print("Nuevo tiempo en segundos: ")
			escaner.Scan()
			nuevoTiempo, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Nuevas calorías quemadas: ")
			escaner.Scan()
			nuevasCalorias, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Nuevo tipo de ejercicio: ")
			escaner.Scan()
			nuevoTipo := escaner.Text()
			fmt.Print("Nuevos puntos por tipo de ejercicio: ")
			escaner.Scan()
			nuevosPuntos, _ := strconv.Atoi(escaner.Text())
			fmt.Print("Nueva dificultad del ejercicio: ")
			escaner.Scan()
			nuevaDificultad := escaner.Text()
			nuevoEjercicio := &ejercicio.Ejercicio{
				NombreDeEjercicio:           nuevoNombre,
				DescripcionDeEjercicio:      nuevaDescripcion,
				TiempoEnSegundosDeEjercicio: nuevoTiempo,
				CaloriasDeEjercicio:         nuevasCalorias,
				TipoDeEjercicio:             nuevoTipo,
				PuntosPorTipoDeEjercicio:    nuevosPuntos,
				DificultadDeEjercicio:       nuevaDificultad,
			}
			err := gestorDeEjercicios.ModificarEjercicio(nombre, nuevoEjercicio)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al modificar el ejercicio:", err)
			} else {
				fmt.Println("\nEjercicio modificado correctamente.")
			}
		case "5":
			ClearScreen()
			fmt.Println("\n--- Listar Ejercicios ---")
			ejercicios := gestorDeEjercicios.ListarEjercicios()
			if len(ejercicios) == 0 {
				fmt.Println("\nNo hay ejercicios disponibles.")
			} else {
				for indice, ejercicio := range ejercicios {
					fmt.Printf("%d. %s\n", indice+1, ejercicio.NombreDeEjercicio)
				}
			}
		case "6":
			ClearScreen()
			fmt.Println("\n--- Filtrar Ejercicios ---")
			fmt.Println("1. Filtrar por Tipo")
			fmt.Println("2. Filtrar por Dificultad")
			fmt.Println("3. Filtrar por Tipo y Dificultad")
			fmt.Print("\nSeleccione una opción: ")
			escaner.Scan()
			opcionFiltro := escaner.Text()
			var ejerciciosFiltrados []*ejercicio.Ejercicio
			switch opcionFiltro {
			case "1":
				ClearScreen()
				fmt.Print("\nTipo(s) de ejercicios a listar (separados por comas): ")
				escaner.Scan()
				entradaDeTipos := escaner.Text()
				tipos := strings.Split(entradaDeTipos, ",")
				for indice, tipo := range tipos {
					tipos[indice] = strings.TrimSpace(tipo)
				}
				ejerciciosFiltrados = gestorDeEjercicios.FiltrarPorTipos(tipos)
			case "2":
				ClearScreen()
				fmt.Print("Dificultad: ")
				escaner.Scan()
				dificultad := escaner.Text()
				ejerciciosFiltrados = gestorDeEjercicios.FiltrarPorDificultad(dificultad)
			case "3":
				ClearScreen()
				fmt.Print("\nTipo(s) de ejercicios a listar (separados por comas): ")
				escaner.Scan()
				entradaDeTipos := escaner.Text()
				tipos := strings.Split(entradaDeTipos, ",")
				for indice, tipo := range tipos {
					tipos[indice] = strings.TrimSpace(tipo)
				}
				fmt.Print("Dificultad: ")
				escaner.Scan()
				dificultad := escaner.Text()
				ejerciciosFiltradosPorTipo := gestorDeEjercicios.FiltrarPorTipos(tipos)
				for _, ejercicio := range ejerciciosFiltradosPorTipo {
					if ejercicio.DificultadDeEjercicio == dificultad {
						ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
					}
				}
			default:
				ClearScreen()
				fmt.Println("\nOpción no válida. Intente de nuevo.")
				return
			}
			ClearScreen()
			if len(ejerciciosFiltrados) == 0 {
				fmt.Println("\nNo hay ejercicios disponibles.")
			} else {
				fmt.Printf("\nEjercicios disponibles:\n\n")
				for indice, ejercicio := range ejerciciosFiltrados {
					fmt.Printf("%d. %s\n", indice + 1, ejercicio.NombreDeEjercicio)
				}
			}
		case "7":
			ClearScreen()
			fmt.Print("\nCantidad de calorias quemadas por ejercicio a listar: ")
			escaner.Scan()
			calorias, _ := strconv.Atoi(escaner.Text())
			ejerciciosFiltrados := gestorDeEjercicios.FiltrarPorCaloriasQuemadas(calorias)
			ClearScreen()
			if len(ejerciciosFiltrados) == 0 {
				fmt.Println("\nNo hay ejercicios disponibles.")
				return
			} else {
				fmt.Printf("\nEjercicios disponibles:\n\n")
				for indice, ejercicio := range ejerciciosFiltrados {
					fmt.Printf("%d. %s\n", indice + 1, ejercicio.NombreDeEjercicio)
				}
			}
		case "8":
			ClearScreen()
			return
		default:
			ClearScreen()
			fmt.Println("\nOpción no válida. Intente de nuevo.")
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
		fmt.Print("\nSeleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			ClearScreen()
			fmt.Print("\nNombre de la rutina: ")
			escaner.Scan()
			nombre := escaner.Text()
			fmt.Println("\n--- Lista de Ejercicios Disponibles ---")
			ejercicios := gestorDeEjercicios.ListarEjercicios()
			for indice, ejercicio := range ejercicios {
				fmt.Printf("%d. %s\n", indice+1, ejercicio.NombreDeEjercicio)
			}
			fmt.Println("\nIngrese los números de los ejercicios a incluir en la rutina, separados por comas:")
			escaner.Scan()
			ejerciciosSeleccionadosString := escaner.Text()
			ejerciciosSeleccionados := []ejercicio.Ejercicio{}
			for _, numeroString := range strings.Split(ejerciciosSeleccionadosString, ",") {
				numero, err := strconv.Atoi(strings.TrimSpace(numeroString))
				if err != nil || numero < 1 || numero > len(ejercicios) {
					fmt.Println("\nNúmero de ejercicio inválido:", numeroString)
					continue
				}
				ejerciciosSeleccionados = append(ejerciciosSeleccionados, *ejercicios[numero-1])
			}
			rutina := &rutina.Rutina{
				NombreDeRutina:                      nombre,
				CaracteristicasIndividualesDeRutina: ejerciciosSeleccionados,
			}
			rutina.CalcularPropiedades(gestorDeEjercicios)
			err := gestorDeRutinas.AgregarRutina(rutina)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al agregar la rutina:", err)
			} else {
				fmt.Println("\nRutina agregada correctamente.")
				fmt.Printf("\nNombre de rutina: %+v\n", rutina.NombreDeRutina)
				fmt.Printf("Lista de ejercicios: %+v\n", rutina.ListaDeEjerciciosDeRutina)
				fmt.Printf("Tiempo total en minutos: %+v\n", rutina.TiempoEnMinutosDeRutina)
				fmt.Printf("Calorías totales: %+v\n", rutina.CaloriasDeRutina)
				fmt.Printf("Dificultad: %+v\n", rutina.DificultadDeRutina)
				fmt.Printf("Tipos de la rutina: %+v\n", rutina.TiposDeRutina)
				fmt.Printf("Puntos por tipos de la rutina: %+v\n", rutina.PuntosDeRutina)
			}
		case "2":
			ClearScreen()
			fmt.Print("\nNombre de la rutina a eliminar: ")
			escaner.Scan()
			nombre := escaner.Text()
			err := gestorDeRutinas.EliminarRutina(nombre)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al eliminar la rutina:", err)
			} else {
				fmt.Println("\nRutina eliminada correctamente.")
			}
		case "3":
			ClearScreen()
			fmt.Print("\nNombre de la rutina a consultar: ")
			escaner.Scan()
			nombre := escaner.Text()
			rutina, err := gestorDeRutinas.ConsultarRutina(nombre)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al consultar la rutina:", err)
			} else {
				var nombresEjercicios []string
				for _, ejercicio := range rutina.CaracteristicasIndividualesDeRutina {
					nombresEjercicios = append(nombresEjercicios, ejercicio.NombreDeEjercicio)
				}
				rutina.ListaDeEjerciciosDeRutina = "\"" + strings.Join(nombresEjercicios, "\", \"") + "\""
				fmt.Printf("\nNombre de rutina: %+v\n", rutina.NombreDeRutina)
				fmt.Printf("Lista de ejercicios: %+v\n", rutina.ListaDeEjerciciosDeRutina)
				fmt.Printf("Tiempo total en minutos: %+v\n", rutina.TiempoEnMinutosDeRutina)
				fmt.Printf("Calorías totales: %+v\n", rutina.CaloriasDeRutina)
				fmt.Printf("Dificultad: %+v\n", rutina.DificultadDeRutina)
				fmt.Printf("Tipos de la rutina: %+v\n", rutina.TiposDeRutina)
				fmt.Printf("Puntos por tipos de la rutina: %+v\n", rutina.PuntosDeRutina)
			}
		case "4":
			ClearScreen()
			fmt.Print("\nNombre de la rutina a modificar: ")
			escaner.Scan()
			nombre := escaner.Text()
			rutinaOriginal, err := gestorDeRutinas.ConsultarRutina(nombre)
			if err != nil {
				fmt.Println("\nError al consultar la rutina:", err)
				return
			} else {
				fmt.Print("¿Desea modificar el nombre de la rutina? (si/no): ")
				escaner.Scan()
				modificarNombre := strings.ToLower(escaner.Text()) == "si"
				var nuevoNombre string
				if modificarNombre {
					fmt.Print("Ingrese el nuevo nombre de la rutina: ")
					escaner.Scan()
					nuevoNombre = escaner.Text()
				} else {
					nuevoNombre = nombre
				}
				fmt.Println("\n¿Desea modificar los ejercicios de la rutina? (si/no): ")
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
							fmt.Println("\nNúmero de ejercicio inválido:", numeroString)
							continue
						}
						ejerciciosSeleccionados = append(ejerciciosSeleccionados, *ejercicios[numero-1])
					}
				} else {
					ejerciciosSeleccionados = rutinaOriginal.CaracteristicasIndividualesDeRutina
				}
				nuevaRutina := &rutina.Rutina {
					NombreDeRutina:                      nuevoNombre,
					CaracteristicasIndividualesDeRutina: ejerciciosSeleccionados,
				}
				nuevaRutina.CalcularPropiedades(gestorDeEjercicios)
				err = gestorDeRutinas.ModificarRutina(nombre, nuevaRutina)
				ClearScreen()
				if err != nil {
					fmt.Println("\nError al modificar la rutina:", err)
				} else {
					var nombresEjercicios []string
					for _, ejercicio := range nuevaRutina.CaracteristicasIndividualesDeRutina {
						nombresEjercicios = append(nombresEjercicios, ejercicio.NombreDeEjercicio)
					}
					nuevaRutina.ListaDeEjerciciosDeRutina = "\"" + strings.Join(nombresEjercicios, "\", \"") + "\""
					fmt.Println("\nRutina modificada correctamente.")
					fmt.Printf("\nNombre de rutina: %+v\n", nuevaRutina.NombreDeRutina)
					fmt.Printf("Lista de ejercicios: %+v\n", nuevaRutina.ListaDeEjerciciosDeRutina)
					fmt.Printf("Tiempo total en minutos: %+v\n", nuevaRutina.TiempoEnMinutosDeRutina)
					fmt.Printf("Calorías totales: %+v\n", nuevaRutina.CaloriasDeRutina)
					fmt.Printf("Dificultad: %+v\n", nuevaRutina.DificultadDeRutina)
					fmt.Printf("Tipos de la rutina: %+v\n", nuevaRutina.TiposDeRutina)
					fmt.Printf("Puntos por tipos de la rutina: %+v\n", nuevaRutina.PuntosDeRutina)
				}
			}
		case "5":
			ClearScreen()
			fmt.Println("\n--- Listar Rutinas ---")
			rutinas := gestorDeRutinas.ListarRutinas()
			if len(rutinas) == 0 {
				fmt.Println("\nNo hay rutinas disponibles.")
			} else {
				fmt.Println("\nRutinas disponibles:")
				for indice, rutina := range rutinas {
					fmt.Printf("%d. %s\n", indice+1, rutina.NombreDeRutina)
				}
			}
		case "6":
			ClearScreen()
			fmt.Print("\nDificultad de las rutinas a listar: ")
			escaner.Scan()
			dificultad := escaner.Text()
			rutinas := gestorDeRutinas.ListarRutinasPorDificultad(dificultad)
			if len(rutinas) == 0 {
				fmt.Println("\nNo hay rutinas disponibles.")
			} else {
				fmt.Println("\nRutinas disponibles:")
				for indice, rutina := range rutinas {
					fmt.Printf("%d. %s\n", indice+1, rutina.NombreDeRutina)
				}
			}
		case "7":
			ClearScreen()
			gestionarRutinasAutomagicas(gestorDeEjercicios, gestorDeRutinas, escaner)
		case "8":
			ClearScreen()
			return
		default:
			ClearScreen()
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
		fmt.Println("\n1. Generar Rutina Automágica por Duración, Tipos y Dificultad.")
		fmt.Println("2. Generar Rutina Automágica por Calorías a quemar.")
		fmt.Println("3. Generar Rutina Automágica por Puntos y Duración Máxima")
		fmt.Println("4. Volver al menú anterior")
		fmt.Print("\nSeleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			ClearScreen()
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
			rutinaAutomagica := &rutina.Rutina{
				NombreDeRutina: nombre,
			}
			duracionAcumulada := 0
			for _, ejercicio := range ejerciciosOrdenados {
				if duracionAcumulada+ejercicio.TiempoEnSegundosDeEjercicio <= duracionTotalEnSegundos {
					rutinaAutomagica.CaracteristicasIndividualesDeRutina = append(rutinaAutomagica.CaracteristicasIndividualesDeRutina, *ejercicio)
					duracionAcumulada += ejercicio.TiempoEnSegundosDeEjercicio
				} else {
					break
				}
			}
			rutinaAutomagica.CalcularPropiedades(gestorDeEjercicios)
			err := gestorDeRutinas.AgregarRutina(rutinaAutomagica)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al generar la rutina automágica:", err)
			} else {
				fmt.Println("\nRutina automágica generada correctamente.")
				var nombresEjercicios []string
				for _, ejercicio := range rutinaAutomagica.CaracteristicasIndividualesDeRutina {
					nombresEjercicios = append(nombresEjercicios, ejercicio.NombreDeEjercicio)
				}
				rutinaAutomagica.ListaDeEjerciciosDeRutina = "\"" + strings.Join(nombresEjercicios, "\", \"") + "\""
				fmt.Printf("\nNombre de rutina: %+v\n", rutinaAutomagica.NombreDeRutina)
				fmt.Printf("Lista de ejercicios: %+v\n", rutinaAutomagica.ListaDeEjerciciosDeRutina)
				fmt.Printf("Tiempo total en minutos: %+v\n", rutinaAutomagica.TiempoEnMinutosDeRutina)
				fmt.Printf("Calorías totales: %+v\n", rutinaAutomagica.CaloriasDeRutina)
				fmt.Printf("Dificultad: %+v\n", rutinaAutomagica.DificultadDeRutina)
				fmt.Printf("Tipos de la rutina: %+v\n", rutinaAutomagica.TiposDeRutina)
				fmt.Printf("Puntos por tipos de la rutina: %+v\n", rutinaAutomagica.PuntosDeRutina)
			}
		case "2":
			ClearScreen()
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
			rutinaAutomagica := &rutina.Rutina{
				NombreDeRutina: nombre,
			}
			caloriasAcumuladas := 0
			for _, ejercicio := range ejerciciosOrdenados {
				if caloriasAcumuladas+ejercicio.CaloriasDeEjercicio <= caloriasTotales {
					rutinaAutomagica.CaracteristicasIndividualesDeRutina = append(rutinaAutomagica.CaracteristicasIndividualesDeRutina, *ejercicio)
					caloriasAcumuladas += ejercicio.CaloriasDeEjercicio
				} else {
					break
				}
			}
			rutinaAutomagica.CalcularPropiedades(gestorDeEjercicios)
			err := gestorDeRutinas.AgregarRutina(rutinaAutomagica)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al generar la rutina automágica:", err)
			} else {
				fmt.Println("\nRutina automágica generada correctamente.")
				var nombresEjercicios []string
				for _, ejercicio := range rutinaAutomagica.CaracteristicasIndividualesDeRutina {
					nombresEjercicios = append(nombresEjercicios, ejercicio.NombreDeEjercicio)
				}
				rutinaAutomagica.ListaDeEjerciciosDeRutina = "\"" + strings.Join(nombresEjercicios, "\", \"") + "\""
				fmt.Printf("\nNombre de rutina: %+v\n", rutinaAutomagica.NombreDeRutina)
				fmt.Printf("Lista de ejercicios: %+v\n", rutinaAutomagica.ListaDeEjerciciosDeRutina)
				fmt.Printf("Tiempo total en minutos: %+v\n", rutinaAutomagica.TiempoEnMinutosDeRutina)
				fmt.Printf("Calorías totales: %+v\n", rutinaAutomagica.CaloriasDeRutina)
				fmt.Printf("Dificultad: %+v\n", rutinaAutomagica.DificultadDeRutina)
				fmt.Printf("Tipos de la rutina: %+v\n", rutinaAutomagica.TiposDeRutina)
				fmt.Printf("Puntos por tipos de la rutina: %+v\n", rutinaAutomagica.PuntosDeRutina)
			}
		case "3":
			ClearScreen()
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
				fmt.Println("\nNo se encontraron ejercicios válidos para esta rutina")
				return
			}
			ejerciciosOrdenados := gestorDeEjercicios.OrdenarPorPuntajeDescendente(ejerciciosFiltrados)
			duracionTotalEnSegundos := duracionMaximaEnMinutos * 60
			rutinaAutomagica := &rutina.Rutina{
				NombreDeRutina: nombre,
			}
			puntajeTotal := 0
			duracionTotal := 0
			ejerciciosUtilizados := make(map[string]bool)
			for _, ejercicio := range ejerciciosOrdenados {
				if duracionTotal+ejercicio.TiempoEnSegundosDeEjercicio <= duracionTotalEnSegundos && !ejerciciosUtilizados[ejercicio.NombreDeEjercicio] {
					rutinaAutomagica.CaracteristicasIndividualesDeRutina = append(rutinaAutomagica.CaracteristicasIndividualesDeRutina, *ejercicio)
					puntajeTotal += ejercicio.PuntosPorTipoDeEjercicio
					duracionTotal += ejercicio.TiempoEnSegundosDeEjercicio
					ejerciciosUtilizados[ejercicio.NombreDeEjercicio] = true
				}
			}
			rutinaAutomagica.PuntosDeRutina = puntajeTotal
			rutinaAutomagica.TiempoEnMinutosDeRutina = duracionTotal
			err := gestorDeRutinas.AgregarRutina(rutinaAutomagica)
			ClearScreen()
			if err != nil {
				fmt.Println("\nError al generar la rutina automágica:", err)
			} else {
				fmt.Println("\nRutina automágica generada correctamente.")
				var nombresEjercicios []string
				for _, ejercicio := range rutinaAutomagica.CaracteristicasIndividualesDeRutina {
					nombresEjercicios = append(nombresEjercicios, ejercicio.NombreDeEjercicio)
				}
				rutinaAutomagica.ListaDeEjerciciosDeRutina = "\"" + strings.Join(nombresEjercicios, "\", \"") + "\""
				fmt.Printf("\nNombre de rutina: %+v\n", rutinaAutomagica.NombreDeRutina)
				fmt.Printf("Lista de ejercicios: %+v\n", rutinaAutomagica.ListaDeEjerciciosDeRutina)
				fmt.Printf("Tiempo total en minutos: %+v\n", rutinaAutomagica.TiempoEnMinutosDeRutina)
				fmt.Printf("Calorías totales: %+v\n", rutinaAutomagica.CaloriasDeRutina)
				fmt.Printf("Dificultad: %+v\n", rutinaAutomagica.DificultadDeRutina)
				fmt.Printf("Tipos de la rutina: %+v\n", rutinaAutomagica.TiposDeRutina)
				fmt.Printf("Puntos por tipos de la rutina: %+v\n", rutinaAutomagica.PuntosDeRutina)
			}
		case "4":
			ClearScreen()
			return
		default:
			ClearScreen()
			fmt.Println("\nOpción no válida. Intente de nuevo.")
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
		fmt.Print("\nSeleccione una opción: ")
		if !escaner.Scan() {
			break
		}
		opcion := escaner.Text()
		switch opcion {
		case "1":
			ClearScreen()
			fmt.Print("\nIngrese el nombre del archivo para guardar los ejercicios: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				err := almacenamiento.GuardarEjercicios(gestorDeEjercicios.ListarEjercicios(), nombreDeArchivo)
				ClearScreen()
				if err != nil {
					fmt.Println("\nError al guardar los ejercicios:", err)
				} else {
					fmt.Println("\nEjercicios guardados correctamente.")
				}
			}
		case "2":
			ClearScreen()
			fmt.Print("\nIngrese el nombre del archivo desde donde cargar los ejercicios: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				ejercicios, err := almacenamiento.CargarEjercicios(nombreDeArchivo)
				ClearScreen()
				if err != nil {
					fmt.Println("\nError al cargar los ejercicios:", err)
				} else {
					for _, ejercicio := range ejercicios {
						gestorDeEjercicios.AgregarEjercicio(ejercicio)
					}
					fmt.Println("\nEjercicios cargados correctamente.")
				}
			}
		case "3":
			ClearScreen()
			fmt.Print("\nIngrese el nombre del archivo para guardar las rutinas: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				err := almacenamiento.GuardarRutinas(gestorDeRutinas.ListarRutinas(), gestorDeEjercicios, nombreDeArchivo)
				ClearScreen()
				if err != nil {
					fmt.Println("\nError al guardar las rutinas:", err)
				} else {
					fmt.Println("\nRutinas guardadas correctamente.")
				}
			}
		case "4":
			ClearScreen()
			fmt.Print("\nIngrese el nombre del archivo desde donde cargar las rutinas: ")
			if escaner.Scan() {
				nombreDeArchivo := escaner.Text()
				rutinas, err := almacenamiento.CargarRutinas(nombreDeArchivo)
				ClearScreen()
				if err != nil {
					fmt.Println("\nError al cargar las rutinas:", err)
				} else {
					for _, rutina := range rutinas {
						gestorDeRutinas.AgregarRutina(rutina)
					}
					fmt.Println("\nRutinas cargadas correctamente.")
				}
			}
		case "5":
			ClearScreen()
			return
		default:
			ClearScreen()
			fmt.Println("\nOpción no válida. Intente de nuevo.")
		}
	}
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}