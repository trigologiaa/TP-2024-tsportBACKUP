package main

//
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

func main() {
	gestorEjercicios := ejercicio.NuevoGestorEjercicios()
	ejercicios, err := almacenamiento.CargarEjerciciosDeUna()
	if err != nil {
		fmt.Println("Error al iniciar lista de ejercicios:", err)
	} else {
		for _, ejercicio := range ejercicios {
			gestorEjercicios.AgregarEjercicio(ejercicio)
		}
		fmt.Println("Lista de ejercicios iniciada correctamente.")
	}
	
	gestorRutinas := rutina.NuevoGestorRutinas(gestorEjercicios)
	rutinas, err := almacenamiento.CargarRutinasDeUna()
	if err != nil {
		fmt.Println("Error al iniciar lista de rutinas:", err)
	} else {
		for _, rutina := range rutinas {
			gestorRutinas.AgregarRutina(rutina)
		}
		fmt.Println("Lista de rutinas iniciada correctamente.")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("\n1. Gestionar Ejercicios")
		fmt.Println("2. Gestionar Rutinas")
		fmt.Println("3. Gestionar Archivos CSV")
		fmt.Println("4. Salir")
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
			GestionarCSV(gestorEjercicios, gestorRutinas, scanner)
		case "4":
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
		fmt.Println("\n1. Agregar Ejercicio Nuevo")
		fmt.Println("2. Eliminar Ejercicio por Nombre")
		fmt.Println("3. Consultar Ejercicio por Nombre")
		fmt.Println("4. Modificar Ejercicio por Nombre")
		fmt.Println("5. Listar todos los Ejercicios")
		fmt.Println("6. Volver al Menú Principal")
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
				TiempoEnSegundos:      tiempo,
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
			fmt.Print("Nombre del ejercicio a modificar: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Nueva descripción del ejercicio: ")
			scanner.Scan()
			descripcion := scanner.Text()
			fmt.Print("Nuevo tiempo en segundos: ")
			scanner.Scan()
			tiempo, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Nuevas calorías quemadas: ")
			scanner.Scan()
			calorias, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Nuevo tipo de ejercicio: ")
			scanner.Scan()
			tipo := scanner.Text()
			fmt.Print("Nuevos puntos por tipo de ejercicio: ")
			scanner.Scan()
			puntos, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Nueva dificultad del ejercicio: ")
			scanner.Scan()
			dificultad := scanner.Text()
			nuevoEjercicio := &ejercicio.Ejercicio{
				Nombre:           nombre,
				Descripcion:      descripcion,
				TiempoEnSegundos: tiempo,
				Calorias:         calorias,
				Tipo:             tipo,
				Puntos:           puntos,
				Dificultad:       dificultad,
			}
			err := gestor.ModificarEjercicio(nombre, nuevoEjercicio)
			if err != nil {
				fmt.Println("Error al modificar el ejercicio:", err)
			} else {
				fmt.Println("Ejercicio modificado correctamente.")
			}
		case "5":
			ejercicios := gestor.ListarEjercicios()
			for _, ej := range ejercicios {
				fmt.Printf("Ejercicio: %+v\n", ej)
			}
		case "6":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func gestionarRutinas(gestorEjercicios *ejercicio.GestorEjercicios, gestor *rutina.GestorRutinas, scanner *bufio.Scanner) {
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
		if !scanner.Scan() {
			break
		}
		opcion := scanner.Text()
		switch opcion {
		case "1":
			fmt.Print("\nNombre de la rutina: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Println("\n--- Lista de Ejercicios Disponibles ---")
			ejercicios := gestorEjercicios.ListarEjercicios()
			for indice, ejercicio := range ejercicios {
				fmt.Printf("%d. %s\n", indice + 1, ejercicio.Nombre)
			}
			fmt.Println("\nIngrese los números de los ejercicios a incluir en la rutina, separados por comas:")
			scanner.Scan()
			ejerciciosSeleccionadosStr := scanner.Text()
			ejerciciosSeleccionados := []ejercicio.Ejercicio{}
			for _, numStr := range strings.Split(ejerciciosSeleccionadosStr, ",") {
				num, err := strconv.Atoi(strings.TrimSpace(numStr))
				if err != nil || num < 1 || num > len(ejercicios) {
					fmt.Println("Número de ejercicio inválido:", numStr)
					continue
				}
				ejerciciosSeleccionados = append(ejerciciosSeleccionados, *ejercicios[num-1])
			}
			rut := &rutina.Rutina{
				Nombre:     nombre,
				CaracteristicasIndividuales: ejerciciosSeleccionados,
			}
			rut.CalcularPropiedades(gestorEjercicios)
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
			fmt.Println("¿Desea modificar los ejercicios de la rutina? (si/no): ")
			scanner.Scan()
			modificarEjercicios := scanner.Text() == "si"
			var ejerciciosSeleccionados []ejercicio.Ejercicio
			if modificarEjercicios {
				fmt.Println("\n--- Lista de Ejercicios Disponibles ---")
				ejercicios := gestorEjercicios.ListarEjercicios()
				for i, ej := range ejercicios {
					fmt.Printf("%d. %s\n", i+1, ej.Nombre)
				}
				fmt.Println("\nIngrese los números de los ejercicios a incluir en la rutina, separados por comas:")
				scanner.Scan()
				ejerciciosSeleccionadosStr := scanner.Text()
				ejerciciosSeleccionados = []ejercicio.Ejercicio{}
				for _, numStr := range strings.Split(ejerciciosSeleccionadosStr, ",") {
					num, err := strconv.Atoi(strings.TrimSpace(numStr))
					if err != nil || num < 1 || num > len(ejercicios) {
						fmt.Println("Número de ejercicio inválido:", numStr)
						continue
					}
					ejerciciosSeleccionados = append(ejerciciosSeleccionados, *ejercicios[num-1])
				}
			}
			nuevaRut := &rutina.Rutina{
				Nombre:     nuevoNombre,
				CaracteristicasIndividuales: ejerciciosSeleccionados,
			}
			nuevaRut.CalcularPropiedades(gestorEjercicios)
			err := gestor.ModificarRutina(nombre, nuevaRut)
			if err != nil {
				fmt.Println("Error al modificar la rutina:", err)
			} else {
				fmt.Println("Rutina modificada correctamente.")
			}
		case "5":
			fmt.Println("\n--- Listar Rutinas ---")
    		rutinas := gestor.ListarRutinas()
    		if len(rutinas) == 0 {
        		fmt.Println("No hay rutinas disponibles.")
    		} else {
        		fmt.Println("Rutinas disponibles:")
        		for i, rut := range rutinas {
            		fmt.Printf("%d. %s\n", i+1, rut.Nombre)
        		}
    		}
		case "6":
			fmt.Print("Dificultad de las rutinas a listar: ")
			scanner.Scan()
			dificultad := scanner.Text()
			rutinas := gestor.ListarRutinasPorDificultad(dificultad)
			for _, rut := range rutinas {
				fmt.Printf("Rutina: %+v\n", rut)
			}
		case "7":
			gestionarRutinasAutomagicas(gestorEjercicios, gestor, scanner)
		case "8":
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func gestionarRutinasAutomagicas(gestorEjercicios *ejercicio.GestorEjercicios, gestor *rutina.GestorRutinas, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Rutinas AUTOMÁGICAS ---")
		fmt.Println("\n1. RUTINA AUTOMÁGICA 1 (Máxima cantidad: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que maximicen la cantidad de ejercicios a realizar en el tiempo definido)")
		fmt.Println("2. RUTINA AUTOMÁGICA 2 (Mínima duración: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que minimicen la duración de la rutina)")
		fmt.Println("3. RUTINA AUTOMÁGICA 3 (Duración fija, Máximos puntos: Se seleccionarán los ejercicios que cumplan con los parámetros establecidos y que maximicen el puntaje de la rutina en la dimensión solicitada)")
		fmt.Println("4. Volver al menú anterior")
		fmt.Print("\nSeleccione una opción: ")
		if !scanner.Scan() {
			break
		}
		opcion := scanner.Text()
		switch opcion {
		case "1":
			fmt.Print("\nNombre de la rutina: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Duración total de la rutina (en minutos): ")
			scanner.Scan()
			duracionMinutos, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Tipo(s) de ejercicios a incluir (separados por comas): ")
			scanner.Scan()
			tiposInput := scanner.Text()
			tipos := strings.Split(tiposInput, ",")
			for i, tipo := range tipos {
				tipos[i] = strings.TrimSpace(tipo)
			}
			fmt.Print("Dificultad: ")
			scanner.Scan()
			dificultad := scanner.Text()
			ejerciciosFiltrados := gestorEjercicios.FiltrarPorTiposYDificultad(tipos, dificultad)
			if len(ejerciciosFiltrados) == 0 {
				fmt.Println("No se encontraron ejercicios válidos para esta rutina")
				return
			}
			ejerciciosOrdenados := gestorEjercicios.OrdenarTiempoMenorAMayor(ejerciciosFiltrados)
			duracionTotalSegundos := duracionMinutos * 60
			rutinaAutomagica := &rutina.Rutina{
				Nombre: nombre,
			}
			duracionAcumulada := 0
			for _, ej := range ejerciciosOrdenados {
				if duracionAcumulada+ej.TiempoEnSegundos <= duracionTotalSegundos {
					rutinaAutomagica.CaracteristicasIndividuales = append(rutinaAutomagica.CaracteristicasIndividuales, *ej)
					duracionAcumulada += ej.TiempoEnSegundos
				} else {
					break
				}
			}
			rutinaAutomagica.CalcularPropiedades(gestorEjercicios)
			err := gestor.AgregarRutina(rutinaAutomagica)
			if err != nil {
				fmt.Println("Error al generar la rutina automágica:", err)
			} else {
				fmt.Println("Rutina automágica generada correctamente.")
			}
		case "2":
			fmt.Print("\nNombre de la rutina: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Calorías totales a quemar: ")
			scanner.Scan()
			caloriasTotales, _ := strconv.Atoi(scanner.Text())
			ejercicios := gestorEjercicios.ListarEjercicios()
			ejerciciosFiltrados := []*ejercicio.Ejercicio{}
			for _, ej := range ejercicios {
				if ej.Calorias <= caloriasTotales {
					ejerciciosFiltrados = append(ejerciciosFiltrados, ej)
				}
			}
			ejerciciosOrdenados := gestorEjercicios.OrdenarTiempoMenorAMayor(ejerciciosFiltrados)
			
			rutinaAutomagica := &rutina.Rutina{
				Nombre: nombre,
			}
			caloriasAcumuladas := 0
			for _, ej := range ejerciciosOrdenados {
				if caloriasAcumuladas+ej.Calorias <= caloriasTotales {
					rutinaAutomagica.CaracteristicasIndividuales = append(rutinaAutomagica.CaracteristicasIndividuales, *ej)
					caloriasAcumuladas += ej.Calorias
				} else {
					break
				}
			}
			rutinaAutomagica.CalcularPropiedades(gestorEjercicios)
			err := gestor.AgregarRutina(rutinaAutomagica)
			if err != nil {
				fmt.Println("Error al generar la rutina automágica:", err)
			} else {
				fmt.Println("Rutina automágica generada correctamente.")
			}
		case "3":
			fmt.Print("\nNombre de la rutina: ")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Print("Tipo de puntos a maximizar (cardio, fuerza, flexibilidad): ")
			scanner.Scan()
			tipoPuntos := scanner.Text()
			fmt.Print("Duración máxima de la rutina (en minutos): ")
			scanner.Scan()
			duracionMaximaMinutos, _ := strconv.Atoi(scanner.Text())
			ejerciciosFiltrados := gestorEjercicios.FiltrarPorTipoPuntosYDuracionMaxima(tipoPuntos, duracionMaximaMinutos)
			if len(ejerciciosFiltrados) == 0 {
				fmt.Println("No se encontraron ejercicios válidos para esta rutina")
				return
			}
			ejerciciosOrdenados := gestorEjercicios.OrdenarPorPuntajeDescendente(ejerciciosFiltrados)
			duracionTotalSegundos := duracionMaximaMinutos * 60
			rutinaAutomagica := &rutina.Rutina{
				Nombre: nombre,
			}
			puntajeTotal := 0
			duracionTotal := 0
			ejerciciosUtilizados := make(map[string]bool)
			for _, ej := range ejerciciosOrdenados {
				if duracionTotal+ej.TiempoEnSegundos <= duracionTotalSegundos && !ejerciciosUtilizados[ej.Nombre] {
					rutinaAutomagica.CaracteristicasIndividuales = append(rutinaAutomagica.CaracteristicasIndividuales, *ej)
					puntajeTotal += ej.Puntos
					duracionTotal += ej.TiempoEnSegundos
					ejerciciosUtilizados[ej.Nombre] = true
				}
			}
			rutinaAutomagica.Puntos = puntajeTotal
			rutinaAutomagica.TiempoEnMinutos = duracionTotal
			err := gestor.AgregarRutina(rutinaAutomagica)
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

func GestionarCSV(gestorEjercicios *ejercicio.GestorEjercicios, gestorRutinas *rutina.GestorRutinas, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gestión de Archivos CSV ---")
		fmt.Println("\n1. Guardar Ejercicios en CSV")
		fmt.Println("2. Cargar Ejercicios desde CSV")
		fmt.Println("3. Guardar Rutinas en CSV")
		fmt.Println("4. Cargar Rutinas desde CSV")
		fmt.Println("5. Volver al Menú Principal")
		fmt.Print("Seleccione una opción: ")
		if !scanner.Scan() {
			break
		}
		opcion := scanner.Text()
		switch opcion {
		case "1":
			fmt.Print("Ingrese el nombre del archivo para guardar los ejercicios: ")
			if scanner.Scan() {
				nombreDeArchivo := scanner.Text()
				err := almacenamiento.GuardarEjercicios(gestorEjercicios.ListarEjercicios(), nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al guardar los ejercicios:", err)
				} else {
					fmt.Println("Ejercicios guardados correctamente.")
				}
			}
		case "2":
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
		case "3":
			fmt.Print("Ingrese el nombre del archivo para guardar las rutinas: ")
			if scanner.Scan() {
				nombreDeArchivo := scanner.Text()
				err := almacenamiento.GuardarRutinas(gestorRutinas.ListarRutinas(), gestorEjercicios, nombreDeArchivo)
				if err != nil {
					fmt.Println("Error al guardar las rutinas:", err)
				} else {
					fmt.Println("Rutinas guardadas correctamente.")
				}
			}
		case "4":
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
		case "5":
			fmt.Println("Volviendo al Menú Principal...")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}