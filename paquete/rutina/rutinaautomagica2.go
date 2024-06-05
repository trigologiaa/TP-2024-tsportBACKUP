package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"errors"

	"github.com/untref-ayp2/data-structures/heap"
)

func GenerarRutinaPorCalorias(nombre string, caloriasTotales int) error {

	// Busco todos los ejercicios almacenados, si no hay ejercicios retorno un error
	ejerciciosAlmacenados, err := ejercicio.CargarEjercicios("Donde está la rutina?")
	if err != nil {
		return err
	}

	// Ordeno los ejercicios de mayor a menor cantidad de calorias
	// para respetar el principio de minima duración de la rutina.
	ejerciciosOrdenados := heap.NewGenericHeap[*ejercicio.Ejercicio](ejerciciosDeMayorAMenorCalorias)
	for _, valor := range ejerciciosAlmacenados {
		ejerciciosOrdenados.Insert(valor)
	}

	// Itero por los ejercicios ordenados,
	// si al sumar las calorias del ejercicio actual con las
	// calorias de los ejercicios anteriores no supero las calorias totales,
	// agrego el ejercicio actual a la rutina
	var excedeCalorias bool
	var caloriasNuevaRutina int
	conjuntoEjercicios := ejercicio.NuevoGestorEjercicios()
	for i := 0; i < ejerciciosOrdenados.Size() && !excedeCalorias; i++ {
		ejercicio, _ := ejerciciosOrdenados.Remove()
		respetaCalorias := ejercicio.Calorias+caloriasNuevaRutina < caloriasTotales
		if respetaCalorias {
			conjuntoEjercicios.AgregarEjercicio(ejercicio)
			caloriasNuevaRutina += ejercicio.Calorias
		} else {
			excedeCalorias = true
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

// Esta función compara las calorias de dos ejercicios y devuelve cual contiene más
// es para que el heap generico que ordena los ejercicios funcione
// COMPLETAR esta documentación
func ejerciciosDeMayorAMenorCalorias(a *ejercicio.Ejercicio, b *ejercicio.Ejercicio) int {
	if a.Calorias < b.Calorias {
		return -1
	} else if b.Calorias < a.Calorias {
		return 1
	}
	return 0
}
