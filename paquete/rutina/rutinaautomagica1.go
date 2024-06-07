package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"errors"
	"sort"
)

// GenerarRutinaAutomagica1 genera una rutina nueva basada en los parámetros especificados.
//
// Parámetros:
//   - nombre: Será el nombre de la rutina.
//   - duracion: Será la duración total de la rutina en minutos.
//   - tipo: Será el tipo de ejercicios a incluir.
//   - dificultad: Será el nivel de dificultad de los ejercicios a incluir.
//
// Retorna:
//   - Un puntero a Rutina.
//   - Un error en caso de que no se pueda generar la rutina.
//     Funcionamiento:
//     Se declara la variable 'todosLosEjercicios' de tipo slice de punteros a Ejercicio, que contiene el resultado del método 'ObtenerTodosLosEjercicios'
//     Se declara la variable 'ejerciciosFiltrados', que contiene el resultado del método 'filtrarEjerciciosRutinaAutomagica1' el cual tiene los parámetros 'todosLosEjercicios' para ejercicios, 'tipo' para tipo y 'dificultad' para dificultad
//     Se llama al método 'Slice' de 'sort' para ordenar los ejercicios por tiempo de duración ascendente,se le pasan como parámetros 'ejerciciosFiltrados', una función que tiene como parámetros 'i' e 'j' de tipo int, y que retorna un valor bool {
//     Se retorna el valor de 'i' en caso de ser menor que el valor de 'j'
//     }
//     Se declara la variable 'ejerciciosSeleccionados' de tipo slice de punteros a Ejercicio, que contiene el resultado del método 'seleccionarEjerciciosRutinaAutomagica1' el cual tendrá los parámetros 'ejerciciosFiltrados' y 'duración'
//     Si la longitud de 'ejerciciosSeleccionados' es 0 (no se seleccionaron ejercicios) {
//     Se retornan un nil y un error
//     }
//     Si no, se declara la variable 'nuevaRutina' que es una nueva instancia de 'Rutina' {
//     Se inicializa el campo Nombre, que es igual al parámetro 'nombre'
//     Se inicializa el campo Dificultad, que es la dificultad que más se repite en el slice de dificultades
//     }
//     Se llama al método calcularPropiedadesRutinaAutomagica1 pasándole como parámetros 'nuevaRutina' y 'ejerciciosSeleccionados'
//     Se llama al método 'AgregarRutina' pasándole como parámetro 'nuevaRutina', y si hay un error {
//     Se retorna nil y un error
//     }
//     Se retornan 'nuevaRutina' y nil
func (g *GestorRutinas) GenerarRutinaAutomagica1(todosLosEjercicios []*ejercicio.Ejercicio, nombre string, duracion int, tipo string, dificultad string) (*Rutina, error) {
	ejerciciosFiltrados := filtrarEjerciciosRutinaAutomagica1(todosLosEjercicios, tipo, dificultad)
	sort.Slice(ejerciciosFiltrados, func(i, j int) bool {
		return ejerciciosFiltrados[i].Tiempo < ejerciciosFiltrados[j].Tiempo
	})
	duracionSegundos := duracion * 60
	ejerciciosSeleccionados := seleccionarEjerciciosRutinaAutomagica1(ejerciciosFiltrados, duracionSegundos)
	if len(ejerciciosSeleccionados) == 0 {
		return nil, errors.New("no se encontraron ejercicios que cumplan con los criterios")
	}
	nuevaRutina := &Rutina{
		Nombre:     nombre,
		Dificultad: dificultad,
	}
	calcularPropiedadesRutinaAutomagica1(nuevaRutina, ejerciciosSeleccionados)
	if err := g.AgregarRutina(nuevaRutina); err != nil {
		return nil, err
	}
	return nuevaRutina, nil
}

// filtrarEjercicios filtra los ejercicios según el tipo y la dificultad.
//
// Parámetros:
//   - 'ejercicios' será un slice de punteros a Ejercicio.
//   - 'tipo' será un string.
//   - 'dificultad' será un string.
//
// Retorna:
//   - Un slice de punteros a Ejercicio (que contendrá los ejercicios filtrados).
//
// Funcionamiento:
//
//	Se declara la variable 'filtrados' de tipo slice de punteros a Ejercicio (vacío)
//	Se recorre la lista de 'ejercicios' {
//	    Si el Tipo de 'ejercicio' coincide con el tipo buscado y la Dificultad de 'ejercicio' coincide con la buscada {
//	        Se agrega el ejercicio a 'filtrados'
//	    }
//	}
//	Se retorna 'filtrados'
func filtrarEjerciciosRutinaAutomagica1(ejercicios []*ejercicio.Ejercicio, tipo string, dificultad string) []*ejercicio.Ejercicio {
	var filtrados []*ejercicio.Ejercicio
	for _, ejercicio := range ejercicios {
		if ejercicio.Tipo == tipo && ejercicio.Dificultad == dificultad {
			filtrados = append(filtrados, ejercicio)
		}
	}
	return filtrados
}

// seleccionarEjercicios selecciona los ejercicios que maximicen la cantidad sin exceder la duración total.
// Previene la repetición de ejercicios.
//
// Parámetros:
//   - 'ejercicios' será un slice de punteros a Ejercicio.
//   - 'duracionMaxima' será un int.
//
// Retorna:
//   - Un slice de punteros a Ejercicio (que contendrá los ejercicios seleccionados).
//
// Funcionamiento:
//
//	Se declara la variable 'seleccionados' de tipo slice de punteros a Ejercicio (vacío)
//	Se declara la variable 'tiempoAcumulado' de tipo int con un valor de 0
//
// Se declara la variable 'ejerciciosUsados' de tipo map con claves de tipo string y valores bool para prevenir repeticiones
//
//	Se recorre la lista de 'ejercicios' {
//	    Si la suma de 'tiempoAcumulado' y 'Tiempo' del ejercicio es igual o menor a 'duracionMaxima' y el ejercicio no ha sido usado {
//	        Se agrega el ejercicio a 'seleccionados'
//	        Se suma 'Tiempo' del ejercicio actual a 'tiempoAcumulado'
//	        Se marca el ejercicio como usado en 'ejerciciosUsados'
//	    }
//	}
//	Se retorna 'seleccionados'
func seleccionarEjerciciosRutinaAutomagica1(ejercicios []*ejercicio.Ejercicio, duracionMaxima int) []*ejercicio.Ejercicio {
	var seleccionados []*ejercicio.Ejercicio
	tiempoAcumulado := 0
	ejerciciosUsados := make(map[string]bool)
	for _, ejercicio := range ejercicios {
		if tiempoAcumulado+ejercicio.Tiempo <= duracionMaxima && !ejerciciosUsados[ejercicio.Nombre] {
			seleccionados = append(seleccionados, ejercicio)
			tiempoAcumulado += ejercicio.Tiempo
			ejerciciosUsados[ejercicio.Nombre] = true
		}
	}
	return seleccionados
}

// calcularPropiedadesRutina calcula y asigna las propiedades de una rutina basada en los ejercicios seleccionados.
//
// Parámetros:
//   - 'rutina' será un puntero a 'Rutina'.
//   - 'ejercicios' será un slice de punteros a Ejercicio
//
// Retorna:
//   - Sin retorno.
//
// Funcionamiento:
//
//	Se declaran las variables 'tiempoTotal', 'caloriasTotales' y 'puntosTotales' de tipo int (vacías)
//	Se declara la variable 'nombresEjercicios' de tipo slice de string (vacío), el cual almacenará los nombres de ejercicios
//	Se declara la variable 'tiposSet' de tipo map con claves de tipo string y valores bool
//	Se recorre la lista de 'ejercicios' {
//	    Se suma 'Tiempo' de ejercicio a 'tiempoTotal'
//	    Se suma 'Calorias' de ejercicio a 'caloriasTotales'
//	    Se suma 'Puntos' de ejercicio a 'puntosTotales'
//	    Se agrega 'Nombre' de ejercicio a 'nombresEjercicios'
//	    Se agrega 'Tipo' de ejercicio a 'tiposSet'
//	}
//	Se le asigna a 'Tiempo' del parámetro 'rutina' el valor de 'tiempoTotal'
//	Se le asigna a 'Calorias' del parámetro 'rutina' el valor de 'caloriasTotales'
//	Se le asigna a 'Ejercicios' del parámetro 'rutina' el valor de 'nombresEjercicios' con el método 'unirKeys'
//	Se le asigna a 'Tipos' del parámetro 'rutina' el valor de 'tiposSet' pasado por el método 'mapKeyAStringSlice' con el método 'unirKeys'
//	Se le asigna a 'PuntosPorTipo' del parámetro 'rutina' el valor de 'puntosTotales'
func calcularPropiedadesRutinaAutomagica1(rutina *Rutina, ejercicios []*ejercicio.Ejercicio) {
	var tiempoTotal, caloriasTotales, puntosTotales int
	nombresEjercicios := make([]string, 0)
	tiposSet := make(map[string]bool)
	for _, ejercicio := range ejercicios {
		tiempoTotal += ejercicio.Tiempo
		caloriasTotales += ejercicio.Calorias
		puntosTotales += ejercicio.Puntos
		nombresEjercicios = append(nombresEjercicios, ejercicio.Nombre)
		tiposSet[ejercicio.Tipo] = true
	}
	rutina.Tiempo = tiempoTotal
	rutina.Calorias = caloriasTotales
	rutina.Ejercicios = unirKeys(nombresEjercicios, ", ")
	rutina.Tipos = unirKeys(mapKeysAStringSlice(tiposSet), ", ")
	rutina.PuntosPorTipo = puntosTotales
}
