package ejercicio

import (
	"errors"
	"sort"
	lista "github.com/untref-ayp2/data-structures/list"
)

// Ejercicio es una estructura con las características que tendrá su campo.
type Ejercicio struct {
	Nombre      string `csv:"Nombre"`      // Nombre del ejercicio
	Descripcion string `csv:"Descripcion"` // Descripción del ejercicio
	TiempoEnSegundos      int    `csv:"TiempoEnSegundos"`      // Tiempo estimado en segundos
	Calorias    int    `csv:"Calorias"`    // Calorías quemadas
	Tipo        string `csv:"Tipo"`        // Tipo de ejercicio
	Puntos      int    `csv:"Puntos"`      // Puntos por tipo de ejercicio
	Dificultad  string `csv:"Dificultad"`  // Nivel de dificultad
}

// GestorEjercicios es una estructura para gestionar los ejercicios.
type GestorEjercicios struct {
	ejercicios *lista.DoubleLinkedList[*Ejercicio]
}

// NuevoGestorEjercicios crea una nueva estructura 'GestorEjercicios'.
//
// Parámetros:
//   - Sin parámetros.
//
// Retorna:
//   - Un puntero a GestorEjercicios.
func NuevoGestorEjercicios() *GestorEjercicios {
	return &GestorEjercicios{
		ejercicios: lista.NewDoubleLinkedList[*Ejercicio](),
	}
}

// AgregarEjercicio añade un nuevo ejercicio a la lista en caso de no existir.
//
// Parámetros:
//   - 'ejercicio' será un puntero a una instancia de la estructura 'Ejercicio'.
//
// Retorna:
//   - Un nil en caso de que se haya agregado el ejercicio.
//   - Un error en caso de que ya exista el ejercicio.
func (gestor *GestorEjercicios) AgregarEjercicio(ejercicio *Ejercicio) error {
	for nodo := gestor.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
        if nodo.Data().Nombre == ejercicio.Nombre {
            return errors.New("el ejercicio ya existe")
        }
    }
    gestor.ejercicios.Append(ejercicio)
    return nil
}

// EliminarEjercicio busca un ejercicio por su nombre y lo elimina de la lista.
// Parámetros:
//   - 'nombre' será el nombre del ejercicio a buscar.
//
// Retorna:
//   - Un nil en caso de que se haya eliminado el ejercicio.
//   - Un error en caso de que no se encuentre en la lista.
func (g *GestorEjercicios) EliminarEjercicio(nombre string) error {
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == nombre {
			g.ejercicios.Remove(nodo.Data())
			return nil
		}
	}
	return errors.New("ejercicio no encontrado")
}

// ConsultarEjercicio busca un ejercicio por su nombre.
//
// Parámetros:
//   - 'nombre' será el nombre del ejercicio a buscar.
//
// Retorna:
//   - Un puntero al ejercicio en caso de encontrarlo.
//   - Un error en caso de que no se encuentre en la lista.
func (g *GestorEjercicios) ConsultarEjercicio(nombre string) (*Ejercicio, error) {
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == nombre {
			return nodo.Data(), nil
		}
	}
	return nil, errors.New("ejercicio no encontrado")
}

// ModificarEjercicio busca un ejercicio por su nombre y modifica sus campos.
//
// Parámetros:
//   - 'nombre' será el nombre del ejercicio a modificar.
//   - 'nuevoEjercicio' será un puntero a la estructura Ejercicio con los nuevos valores.
//
// Retorna:
//   - Un nil en caso de que se haya modificado el ejercicio correctamente.
//   - Un error en caso de que no se encuentre el ejercicio.
func (g *GestorEjercicios) ModificarEjercicio(nombre string, nuevoEjercicio *Ejercicio) error {
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == nombre {
			nodo.Data().Nombre = nuevoEjercicio.Nombre
			nodo.Data().Descripcion = nuevoEjercicio.Descripcion
			nodo.Data().TiempoEnSegundos = nuevoEjercicio.TiempoEnSegundos
			nodo.Data().Calorias = nuevoEjercicio.Calorias
			nodo.Data().Tipo = nuevoEjercicio.Tipo
			nodo.Data().Puntos = nuevoEjercicio.Puntos
			nodo.Data().Dificultad = nuevoEjercicio.Dificultad
			return nil
		}
	}
	return errors.New("ejercicio no encontrado")
}

// ObtenerTodosLosEjercicios devuelve una lista de todos los ejercicios almacenados en la gestión.
//
// Parámetros:
//   - Sin parámetros.
//
// Retorna:
//   - Un slice con todos los ejercicios.
func (g *GestorEjercicios) ListarEjercicios() []*Ejercicio {
	var ejercicios []*Ejercicio
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicios = append(ejercicios, nodo.Data())
	}
	return ejercicios
}

//Métodos que no van al menú sino que se llaman desde otros métodos

// ObtenerEjerciciosPorNombres devuelve una lista de ejercicios que coinciden con los nombres proporcionados.
//
// Parámetros:
//   - 'nombres' será un slice de strings con los nombres de los ejercicios a buscar.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que coinciden con los nombres especificados.
func (g *GestorEjercicios) ObtenerEjercicioPorNombre(nombres []string) []*Ejercicio {
	var ejercicios []*Ejercicio
	for _, nombre := range nombres {
		ejercicio, err := g.ConsultarEjercicio(nombre)
		if err == nil {
			ejercicios = append(ejercicios, ejercicio)
		}
	}
	return ejercicios
}

// FiltrarEjerciciosAutomagica1 devuelve una lista de ejercicios que coinciden con el tipo y la dificultad especificados.
//
// Parámetros:
//   - 'tipo' será un string que indica el tipo de ejercicio a buscar.
//   - 'dificultad' será un string que indica el nivel de dificultad de los ejercicios a buscar.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que coinciden con el tipo y la dificultad especificados.
func (g *GestorEjercicios) FiltrarPorTiposYDificultad(tipos []string, dificultad string) []*Ejercicio {
	var ejerciciosFiltrados []*Ejercicio
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		for _, tipo := range tipos {
			if ejercicio.Tipo == tipo && ejercicio.Dificultad == dificultad {
				ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
				break
			}
		}
	}
	return ejerciciosFiltrados
}

// OrdenamientoAutomagica1 ordena los ejercicios filtrados por tiempo en segundos de menor a mayor.
func (g *GestorEjercicios) OrdenarTiempoMenorAMayor(ejercicios []*Ejercicio) []*Ejercicio {
	sort.Slice(ejercicios, func(i, j int) bool {
		return ejercicios[i].TiempoEnSegundos < ejercicios[j].TiempoEnSegundos
	})
	return ejercicios
}

// FiltrarPorTipoPuntosYDuracionMaxima devuelve una lista de ejercicios que coinciden con el tipo de puntos y la duración máxima especificados.
//
// Parámetros:
//   - 'tipoPuntos' es un string que indica el tipo de puntos a maximizar (cardio, fuerza, flexibilidad).
//   - 'duracionMaximaMinutos' es la duración máxima de la rutina en minutos.
//
// Retorna:
//   - Un slice de punteros a Ejercicio que coinciden con los criterios especificados.
func (g *GestorEjercicios) FiltrarPorTipoPuntosYDuracionMaxima(tipoPuntos string, duracionMaximaMinutos int) []*Ejercicio {
	var ejerciciosFiltrados []*Ejercicio
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicio := nodo.Data()
		if ejercicio.Tipo == tipoPuntos && ejercicio.TiempoEnSegundos <= duracionMaximaMinutos*60 {
			ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
		}
	}
	return ejerciciosFiltrados
}

// OrdenarPorPuntajeDescendente ordena los ejercicios por puntaje de manera descendente.
//
// Parámetros:
//   - 'ejercicios' es un slice de punteros a Ejercicio.
//
// Retorna:
//   - Un slice de punteros a Ejercicio ordenado por puntaje de manera descendente.
func (g *GestorEjercicios) OrdenarPorPuntajeDescendente(ejercicios []*Ejercicio) []*Ejercicio {
	sort.Slice(ejercicios, func(i, j int) bool {
		return ejercicios[i].Puntos > ejercicios[j].Puntos
	})
	return ejercicios
}