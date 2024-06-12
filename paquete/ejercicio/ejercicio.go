package ejercicio

import (
	"errors"
	"sort"

	lista "github.com/untref-ayp2/data-structures/list"
)

// Ejercicio es una estructura con las características que tendrá su campo.
type Ejercicio struct {
	NombreDeEjercicio           string `csv:"Nombre"`           // Nombre del ejercicio
	DescripcionDeEjercicio      string `csv:"Descripcion"`      // Descripción del ejercicio
	TiempoEnSegundosDeEjercicio int    `csv:"TiempoEnSegundos"` // Tiempo estimado en segundos
	CaloriasDeEjercicio         int    `csv:"Calorias"`         // Calorías quemadas
	TipoDeEjercicio             string `csv:"Tipo"`             // Tipo de ejercicio
	PuntosPorTipoDeEjercicio    int    `csv:"Puntos"`           // Puntos por tipo de ejercicio
	DificultadDeEjercicio       string `csv:"Dificultad"`       // Nivel de dificultad
}

// GestorDeEjercicios es una estructura para gestionar los ejercicios.
type GestorDeEjercicios struct {
	ejercicios *lista.DoubleLinkedList[*Ejercicio]
}

// NuevoGestorDeEjercicios crea una nueva estructura 'GestorDeEjercicios'.
//
// Retorna:
//   - Un puntero a GestorDeEjercicios.
func NuevoGestorDeEjercicios() *GestorDeEjercicios {
	return &GestorDeEjercicios{
		ejercicios: lista.NewDoubleLinkedList[*Ejercicio](),
	}
}

// AgregarEjercicio añade un nuevo ejercicio a la lista en caso de no existir.
//
// Parámetros:
//   - ejercicio: Puntero a Ejercicio, el ejercicio a añadir.
//
// Retorna:
//   - nil si el ejercicio se agregó correctamente.
//   - error si el ejercicio ya existe.
func (gestorDeEjercicios *GestorDeEjercicios) AgregarEjercicio(ejercicio *Ejercicio) error {
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeEjercicio == ejercicio.NombreDeEjercicio {
			return errors.New("el ejercicio ya existe")
		}
	}
	gestorDeEjercicios.ejercicios.Append(ejercicio)
	return nil
}

// EliminarEjercicio busca un ejercicio por su nombre y lo elimina de la lista.
//
// Parámetros:
//   - nombreDeEjercicio: string, el nombre del ejercicio a eliminar.
//
// Retorna:
//   - nil si el ejercicio se eliminó correctamente.
//   - error si el ejercicio no se encuentra en la lista.
func (gestorDeEjercicios *GestorDeEjercicios) EliminarEjercicio(nombreDeEjercicio string) error {
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeEjercicio == nombreDeEjercicio {
			gestorDeEjercicios.ejercicios.Remove(nodo.Data())
			return nil
		}
	}
	return errors.New("ejercicio no encontrado")
}

// ConsultarEjercicio busca un ejercicio por su nombre.
//
// Parámetros:
//   - nombreDeEjercicio: string, el nombre del ejercicio a buscar.
//
// Retorna:
//   - Puntero a Ejercicio si el ejercicio se encuentra.
//   - error si el ejercicio no se encuentra en la lista.
func (gestorDeEjercicios *GestorDeEjercicios) ConsultarEjercicio(nombreDeEjercicio string) (*Ejercicio, error) {
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeEjercicio == nombreDeEjercicio {
			return nodo.Data(), nil
		}
	}
	return nil, errors.New("ejercicio no encontrado")
}

// ModificarEjercicio busca un ejercicio por su nombre y modifica sus campos.
//
// Parámetros:
//   - nombreDeEjercicio: string, el nombre del ejercicio a modificar.
//   - nuevoEjercicio: Puntero a Ejercicio, con los nuevos valores del ejercicio.
//
// Retorna:
//   - nil si el ejercicio se modificó correctamente.
//   - error si el ejercicio no se encuentra en la lista.
func (gestorDeEjercicios *GestorDeEjercicios) ModificarEjercicio(nombreDeEjercicio string, nuevoEjercicio *Ejercicio) error {
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeEjercicio == nombreDeEjercicio {
			nodo.Data().NombreDeEjercicio = nuevoEjercicio.NombreDeEjercicio
			nodo.Data().DescripcionDeEjercicio = nuevoEjercicio.DescripcionDeEjercicio
			nodo.Data().TiempoEnSegundosDeEjercicio = nuevoEjercicio.TiempoEnSegundosDeEjercicio
			nodo.Data().CaloriasDeEjercicio = nuevoEjercicio.CaloriasDeEjercicio
			nodo.Data().TipoDeEjercicio = nuevoEjercicio.TipoDeEjercicio
			nodo.Data().PuntosPorTipoDeEjercicio = nuevoEjercicio.PuntosPorTipoDeEjercicio
			nodo.Data().DificultadDeEjercicio = nuevoEjercicio.DificultadDeEjercicio
			return nil
		}
	}
	return errors.New("ejercicio no encontrado")
}

// ListarEjercicios devuelve una lista de todos los ejercicios almacenados en la gestión.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que contiene todos los ejercicios.
func (gestorDeEjercicios *GestorDeEjercicios) ListarEjercicios() []*Ejercicio {
	var ejercicios []*Ejercicio
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicios = append(ejercicios, nodo.Data())
	}
	return ejercicios
}

// ObtenerEjercicioPorNombre devuelve una lista de ejercicios que coinciden con los nombres proporcionados.
//
// Parámetros:
//   - nombresDeEjercicios: slice de strings, con los nombres de los ejercicios a buscar.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que coinciden con los nombres especificados.
func (gestorDeEjercicios *GestorDeEjercicios) ObtenerEjercicioPorNombre(nombresDeEjercicios []string) []*Ejercicio {
	var ejercicios []*Ejercicio
	for _, nombreDeEjercicio := range nombresDeEjercicios {
		ejercicio, err := gestorDeEjercicios.ConsultarEjercicio(nombreDeEjercicio)
		if err == nil {
			ejercicios = append(ejercicios, ejercicio)
		}
	}
	return ejercicios
}

// FiltrarPorTiposYDificultad devuelve una lista de ejercicios que coinciden con el tipo y la dificultad especificados.
//
// Parámetros:
//   - tipos: slice de strings, que indica los tipos de ejercicio a buscar.
//   - dificultad: string, que indica el nivel de dificultad de los ejercicios a buscar.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que coinciden con el tipo y la dificultad especificados.
func (gestorDeEjercicios *GestorDeEjercicios) FiltrarPorTiposYDificultad(tipos []string, dificultad string) []*Ejercicio {
	var ejerciciosFiltrados []*Ejercicio
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		for _, tipo := range tipos {
			if ejercicio.TipoDeEjercicio == tipo && ejercicio.DificultadDeEjercicio == dificultad {
				ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
				break
			}
		}
	}
	return ejerciciosFiltrados
}

// OrdenarTiempoMenorAMayor ordena los ejercicios filtrados por tiempo en segundos de menor a mayor.
//
// Parámetros:
//   - ejercicios: slice de punteros a Ejercicio, que se desea ordenar.
//
// Retorna:
//   - Un slice de punteros a Ejercicio ordenado por tiempo en segundos de menor a mayor.
func (gestorDeEjercicios *GestorDeEjercicios) OrdenarTiempoMenorAMayor(ejercicios []*Ejercicio) []*Ejercicio {
	sort.Slice(ejercicios, func(i, j int) bool {
		return ejercicios[i].TiempoEnSegundosDeEjercicio < ejercicios[j].TiempoEnSegundosDeEjercicio
	})
	return ejercicios
}

// FiltrarPorTipoPuntosYDuracionMaxima devuelve una lista de ejercicios que coinciden con el tipo de puntos y la duración máxima especificados.
//
// Parámetros:
//   - puntosPorTipo: string, que indica el tipo de puntos a maximizar (cardio, fuerza, flexibilidad).
//   - duracionMaximaEnMinutos: int, la duración máxima de la rutina en minutos.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que coinciden con los criterios especificados.
func (gestorDeEjercicios *GestorDeEjercicios) FiltrarPorTipoPuntosYDuracionMaxima(puntosPorTipo string, duracionMaximaEnMinutos int) []*Ejercicio {
	var ejerciciosFiltrados []*Ejercicio
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		if ejercicio.TipoDeEjercicio == puntosPorTipo && ejercicio.TiempoEnSegundosDeEjercicio <= duracionMaximaEnMinutos*60 {
			ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
		}
	}
	return ejerciciosFiltrados
}

// FiltrarPorTipoPuntosYDuracionMaxima devuelve una lista de ejercicios que coinciden con el tipo de puntos y la duración máxima especificados.
//
// Parámetros:
//   - puntosPorTipo: string, que indica el tipo de puntos a maximizar (cardio, fuerza, flexibilidad).
//   - duracionMaximaEnMinutos: int, la duración máxima de la rutina en minutos.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que coinciden con los criterios especificados.
func (gestorDeEjercicios *GestorDeEjercicios) FiltrarPorCaloriasQuemadas(calorias int) []*Ejercicio {
	var ejerciciosFiltrados []*Ejercicio
	for nodo := gestorDeEjercicios.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		if ejercicio.CaloriasDeEjercicio >= calorias {
			ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
		}
	}
	return ejerciciosFiltrados
}

// OrdenarPorPuntajeDescendente ordena los ejercicios por puntaje de manera descendente.
//
// Parámetros:
//   - ejercicios: slice de punteros a Ejercicio, que se desea ordenar.
//
// Retorna:
//   - Un slice de punteros a Ejercicio ordenado por puntaje de manera descendente.
func (gestorDeEjercicios *GestorDeEjercicios) OrdenarPorPuntajeDescendente(ejercicios []*Ejercicio) []*Ejercicio {
	sort.Slice(ejercicios, func(i, j int) bool {
		return ejercicios[i].PuntosPorTipoDeEjercicio > ejercicios[j].PuntosPorTipoDeEjercicio
	})
	return ejercicios
}
