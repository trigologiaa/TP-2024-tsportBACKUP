package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"errors"
)

// GenerarRutinaAutomagica2 genera automáticamente una rutina nueva que cumple con los parámetros especificados.
//
// Parámetros:
//   - nombre: Nombre de la rutina.
//   - caloriasTotales: Calorías totales a quemar.
//   - gestor: GestorEjercicios que proporciona acceso a los ejercicios disponibles.
//
// Retorna:
//   - Un puntero a Rutina que representa la rutina generada.
//   - Un error en caso de que no se puedan encontrar ejercicios que cumplan con los requisitos.
func (g *GestorRutinas) GenerarRutinaAutomagica2(todosLosEjercicios []*ejercicio.Ejercicio, nombre string, caloriasTotales int) (*Rutina, error) {
	ejerciciosFiltrados := filtrarEjerciciosPorCalorias(todosLosEjercicios, caloriasTotales)
	if len(ejerciciosFiltrados) == 0 {
		return nil, errors.New("no hay ejercicios disponibles que puedan quemar las calorías totales especificadas")
	}
	ejerciciosSeleccionados := seleccionarEjercicios(ejerciciosFiltrados, caloriasTotales)
	rutina := &Rutina{
		Nombre:        nombre,
		Ejercicios:    unirNombresEjercicios(ejerciciosSeleccionados),
		Tiempo:        calcularDuracionTotal(ejerciciosSeleccionados),
		Calorias:      caloriasTotales,
		Dificultad:    calcularDificultad(ejerciciosSeleccionados),
		Tipos:         tiposDeEjercicios(ejerciciosSeleccionados),
		PuntosPorTipo: calcularPuntosTotales(ejerciciosSeleccionados),
	}
	return rutina, nil
}

// Filtra los ejercicios que puedan quemar las calorías totales especificadas.
func filtrarEjerciciosPorCalorias(ejercicios []*ejercicio.Ejercicio, caloriasTotales int) []*ejercicio.Ejercicio {
	var ejerciciosFiltrados []*ejercicio.Ejercicio
	for _, ejercicio := range ejercicios {
		if ejercicio.Calorias <= caloriasTotales {
			ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
		}
	}
	return ejerciciosFiltrados
}

// Selecciona los ejercicios que minimicen la duración total de la rutina.
func seleccionarEjercicios(ejercicios []*ejercicio.Ejercicio, caloriasTotales int) []*ejercicio.Ejercicio {
	ordenarEjerciciosPorTiempo(ejercicios)
	var ejerciciosSeleccionados []*ejercicio.Ejercicio
	caloriasRestantes := caloriasTotales
	for _, ejercicio := range ejercicios {
		if ejercicio.Calorias <= caloriasRestantes {
			ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejercicio)
			caloriasRestantes -= ejercicio.Calorias
		}
	}
	return ejerciciosSeleccionados
}

// Ordena los ejercicios por tiempo ascendente.
func ordenarEjerciciosPorTiempo(ejercicios []*ejercicio.Ejercicio) {
	for i := 0; i < len(ejercicios)-1; i++ {
		for j := 0; j < len(ejercicios)-i-1; j++ {
			if ejercicios[j].Tiempo > ejercicios[j+1].Tiempo {
				ejercicios[j], ejercicios[j+1] = ejercicios[j+1], ejercicios[j]
			}
		}
	}
}

// Calcula la duración total de los ejercicios.
func calcularDuracionTotal(ejercicios []*ejercicio.Ejercicio) int {
	var duracionTotal int
	for _, ejercicio := range ejercicios {
		duracionTotal += ejercicio.Tiempo
	}
	return duracionTotal
}

// Calcula la dificultad de la rutina.
func calcularDificultad(ejercicios []*ejercicio.Ejercicio) string {
	// Se cuentan las frecuencias de los niveles de dificultad
	dificultades := make(map[string]int)
	for _, ejercicio := range ejercicios {
		dificultades[ejercicio.Dificultad]++
	}
	maxFrecuencia := 0
	dificultadMasFrecuente := ""
	for dificultad, frecuencia := range dificultades {
		if frecuencia > maxFrecuencia {
			maxFrecuencia = frecuencia
			dificultadMasFrecuente = dificultad
		}
	}
	return dificultadMasFrecuente
}

// Retorna una cadena que contiene los nombres de los ejercicios separados por coma.
func unirNombresEjercicios(ejercicios []*ejercicio.Ejercicio) string {
	var nombres []string
	for _, ejercicio := range ejercicios {
		nombres = append(nombres, ejercicio.Nombre)
	}
	return unirStrings(nombres, ", ")
}

// Retorna una cadena que contiene los tipos de ejercicios separados por coma.
func tiposDeEjercicios(ejercicios []*ejercicio.Ejercicio) string {
	var tipos []string
	tipoMap := make(map[string]bool)
	for _, ejercicio := range ejercicios {
		if _, exists := tipoMap[ejercicio.Tipo]; !exists {
			tipos = append(tipos, ejercicio.Tipo)
			tipoMap[ejercicio.Tipo] = true
		}
	}
	return unirStrings(tipos, ", ")
}

// Calcula los puntos totales de la rutina.
func calcularPuntosTotales(ejercicios []*ejercicio.Ejercicio) int {
	var puntosTotales int
	for _, ejercicio := range ejercicios {
		puntosTotales += ejercicio.Puntos
	}
	return puntosTotales
}

// Une los elementos de un slice en una cadena usando un separador.
func unirStrings(strings []string, separador string) string {
	var resultado string
	for i, s := range strings {
		if i > 0 {
			resultado += separador
		}
		resultado += s
	}
	return resultado
}
