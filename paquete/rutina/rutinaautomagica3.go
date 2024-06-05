package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"errors"

	"github.com/untref-ayp2/data-structures/heap"
)

func GenerarRutinaPorPuntosPorTipoYDuracion(nombre string, duracionTotal int, tipoAMaximizar string) error {

	// Busco todos los ejercicios almacenados, si no hay ejercicios retorno un error
	ejerciciosAlmacenados, err := ejercicio.CargarEjercicios("Donde está la rutina?")
	if err != nil {
		return err
	}

	// Ordeno los ejercicios de mayor a menor cantidad de puntos por el tipo a maximizar
	// para respetar el principio de máximos puntos de la rutina.
	ejerciciosOrdenados := heap.NewGenericHeap[*ejercicio.Ejercicio](ejerciciosDeMayorAMenorPuntos)
	for _, valor := range ejerciciosAlmacenados {
		ejerciciosOrdenados.Insert(valor)
	}

	// Itero por los ejercicios ordenados,
	// si al sumar la duración del ejercicio actual con la duración
	// de los ejercicios anteriores no supero la duración total
	// agrego el ejercicio actual a la rutina
	var excedeDuracion bool
	var duracionNuevaRutina int
	conjuntoEjercicios := ejercicio.NuevoGestorEjercicios()
	for i := 0; i < ejerciciosOrdenados.Size() && !excedeDuracion; i++ {
		ejercicio, _ := ejerciciosOrdenados.Remove()
		respetaDuracion := ejercicio.Tiempo+duracionNuevaRutina < duracionTotal
		if respetaDuracion {
			conjuntoEjercicios.AgregarEjercicio(ejercicio)
			duracionNuevaRutina += ejercicio.Tiempo
		} else {
			excedeDuracion = true
		}
	}

	if len(conjuntoEjercicios.ObtenerTodosLosEjercicios()) < 1 {
		return errors.New("no hay suficiente ejercicios en la rutina")
	}

	// Falta meter los ejercicios a una rutina en si y para eso hay que hacer el func NuevaRutina()
	// Después falta el agregar la rutina al gestor de rutinas
	// Y si devuelvo el puntero a la rutina y que interacción se encargue de agregarla al gestor de rutinas?
	return nil
}

// Esta función compara los puntos de dos ejercicios y devuelve cual contiene más
// es para que el heap generico que ordena los ejercicios funcione
// COMPLETAR esta documentación
func ejerciciosDeMayorAMenorPuntos(a *ejercicio.Ejercicio, b *ejercicio.Ejercicio) int {
	if a.Puntos < b.Puntos {
		return -1
	} else if b.Puntos < a.Puntos {
		return 1
	}
	return 0
}
