package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"errors"
	"strings"
	list "github.com/untref-ayp2/data-structures/list"
)

// Rutina es una estructura que representa una rutina de ejercicios.
type Rutina struct {
	NombreDeRutina        				string 					`csv:"Nombre"`							//Nombre de la rutina
	ListaDeEjerciciosDeRutina 			string 					`csv:"ListaDeEjercicios"`				//Ejercicios que contiene la rutina
	TiempoEnMinutosDeRutina        		int    					`csv:"TiempoEnMinutos"`					//Tiempo total en segundos
	CaloriasDeRutina      				int    					`csv:"Calorias"`						//Calorías quemadas en total
	DificultadDeRutina    				string 					`csv:"Dificultad"`						//Dificultad de la rutina
	TiposDeRutina         				string 					`csv:"Tipos"`							//Tipo de rutina
	PuntosDeRutina 						int    					`csv:"Puntos"`							//Puntos totales por tipo
	CaracteristicasIndividualesDeRutina	[]ejercicio.Ejercicio	`csv:"CaracteristicasIndividuales"`		//Caracteristicas individuales de cada ejercicio
}

// GestorDeRutinas es una estructura para gestionar las rutinas de ejercicios.
type GestorDeRutinas struct {
	rutinas          *list.DoubleLinkedList[*Rutina]
	gestorDeEjercicios *ejercicio.GestorDeEjercicios
}

// NuevoGestorDeRutinas crea una nueva instancia de GestorDeRutinas.
//
// Parámetros:
//   - gestorDeEjercicios: Puntero a GestorDeEjercicios, la estructura que maneja los ejercicios.
//
// Retorna:
//   - Un puntero a GestorDeRutinas.
func NuevoGestorDeRutinas(gestorDeEjercicios *ejercicio.GestorDeEjercicios) *GestorDeRutinas {
	return &GestorDeRutinas {
		rutinas:          	list.NewDoubleLinkedList[*Rutina](),
		gestorDeEjercicios:	gestorDeEjercicios,
	}
}

// AgregarRutina añade una nueva rutina a la lista en caso de no existir.
//
// Parámetros:
//   - rutina: Puntero a una instancia de la estructura Rutina.
//
// Retorna:
//   - nil si la rutina se agregó correctamente.
//   - Un error si la rutina ya existe o si hubo un problema calculando sus propiedades.
func (gestorDeRutinas *GestorDeRutinas) AgregarRutina(rutina *Rutina) error {
	for nodo := gestorDeRutinas.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeRutina == rutina.NombreDeRutina {
			return errors.New("la rutina ya existe")
		}
	}
	err := rutina.CalcularPropiedades(gestorDeRutinas.gestorDeEjercicios)
	if err != nil {
		return err
	}
	var nombresEjercicios []string
	for _, ejercicio := range rutina.CaracteristicasIndividualesDeRutina {
		nombresEjercicios = append(nombresEjercicios, ejercicio.NombreDeEjercicio)
	}
	rutina.ListaDeEjerciciosDeRutina = "\"" + strings.Join(nombresEjercicios, "\", \"") + "\""
	gestorDeRutinas.rutinas.Append(rutina)
	return nil
}

// EliminarRutina elimina una rutina de la lista por su nombre.
//
// Parámetros:
//   - nombreDeRutina: String con el nombre de la rutina a eliminar.
//
// Retorna:
//   - nil si la rutina se eliminó correctamente.
//   - Un error si no se encontró la rutina.
func (gestorDeRutinas *GestorDeRutinas) EliminarRutina(nombreDeRutina string) error {
	for nodo := gestorDeRutinas.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeRutina == nombreDeRutina {
			gestorDeRutinas.rutinas.Remove(nodo.Data())
			return nil
		}
	}
	return errors.New("rutina no encontrada")
}

// ConsultarRutina busca una rutina por su nombre.
//
// Parámetros:
//   - nombreDeRutina: String con el nombre de la rutina a buscar.
//
// Retorna:
//   - Un puntero a la rutina si se encuentra.
//   - Un error si no se encontró la rutina.
func (gestorDeRutinas *GestorDeRutinas) ConsultarRutina(nombreDeRutina string) (*Rutina, error) {
	for nodo := gestorDeRutinas.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeRutina == nombreDeRutina {
			return nodo.Data(), nil
		}
	}
	return nil, errors.New("rutina no encontrada")
}

// ModificarRutina modifica una rutina existente.
//
// Parámetros:
//   - nombreDeRutina: String con el nombre de la rutina a modificar.
//   - nuevaRutina: Puntero a la estructura Rutina con los nuevos valores.
//
// Retorna:
//   - nil si la rutina se modificó correctamente.
//   - Un error si no se encontró la rutina.
func (gestorDeRutinas *GestorDeRutinas) ModificarRutina(nombreDeRutina string, nuevaRutina *Rutina) error {
	for nodo := gestorDeRutinas.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().NombreDeRutina == nombreDeRutina {
			// Recalcular propiedades para la nueva rutina
			nuevaRutina.CalcularPropiedades(gestorDeRutinas.gestorDeEjercicios)
			// Establecer los nuevos datos en el nodo correspondiente
			nodo.SetData(nuevaRutina)
			return nil
		}
	}
	return errors.New("rutina no encontrada")
}

// ListarRutinas devuelve una lista de todas las rutinas.
//
// Retorna:
//   - Un slice de punteros a Rutina con todas las rutinas.
func (gestorDeRutinas *GestorDeRutinas) ListarRutinas() []*Rutina {
	resultado := []*Rutina{}
	for nodo := gestorDeRutinas.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		resultado = append(resultado, nodo.Data())
	}
	return resultado
}

// ListarRutinasPorDificultad devuelve una lista de rutinas que coinciden con la dificultad especificada.
//
// Parámetros:
//   - dificultadDeRutina: String que indica la dificultad de las rutinas a buscar.
//
// Retorna:
//   - Un slice de punteros a Rutina que coinciden con la dificultad especificada.
func (gestorDeRutinas *GestorDeRutinas) ListarRutinasPorDificultad(dificultadDeRutina string) []*Rutina {
	resultado := []*Rutina{}
	for nodo := gestorDeRutinas.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().DificultadDeRutina == dificultadDeRutina {
			resultado = append(resultado, nodo.Data())
		}
	}
	return resultado
}

// CalcularPropiedades calcula las propiedades de una rutina basado en sus ejercicios.
//
// Parámetros:
//   - gestorDeEjercicios: Puntero a GestorDeEjercicios, usado para obtener los ejercicios por nombre.
//
// Retorna:
//   - nil si las propiedades se calcularon correctamente.
//   - Un error si no se encontraron ejercicios válidos.
func (rutina *Rutina) CalcularPropiedades(gestorDeEjercicios *ejercicio.GestorDeEjercicios) error {
	var nombresDeEjercicios []string
	for _, ejercicio := range rutina.CaracteristicasIndividualesDeRutina {
		nombresDeEjercicios = append(nombresDeEjercicios, ejercicio.NombreDeEjercicio)
	}
	ejercicios := gestorDeEjercicios.ObtenerEjercicioPorNombre(nombresDeEjercicios)
	if len(ejercicios) == 0 {
		return errors.New("no se encontraron ejercicios válidos para esta rutina")
	}
	var tiempoTotalEnSegundos, caloriasTotales, puntosTotales int
	tiposSet := make(map[string]bool)
	dificultades := make(map[string]int)
	for _, ejercicio := range ejercicios {
		tiempoTotalEnSegundos += ejercicio.TiempoEnSegundosDeEjercicio
		caloriasTotales += ejercicio.CaloriasDeEjercicio
		puntosTotales += ejercicio.PuntosPorTipoDeEjercicio
		tiposSet[ejercicio.TipoDeEjercicio] = true
		dificultades[ejercicio.DificultadDeEjercicio]++
	}
	rutina.TiempoEnMinutosDeRutina = tiempoTotalEnSegundos / 60
	rutina.CaloriasDeRutina = caloriasTotales
	rutina.TiposDeRutina = unirKeys(mapKeysAStringSlice(tiposSet), ", ")
	rutina.PuntosDeRutina = puntosTotales
	if rutina.DificultadDeRutina == "" {
		rutina.DificultadDeRutina = keyMaxima(dificultades)
	}
	return nil
}

// unirKeys une los elementos de un slice en una sola cadena, separados por el delimitador especificado.
//
// Parámetros:
//   - elementos: Slice de strings que se unirá.
//   - separador: Cadena que se utilizará como separador entre los elementos.
//
// Retorna:
//   - Una cadena que contiene todos los elementos del slice unidos, separados por el delimitador especificado.
func unirKeys(elementos []string, separador string) string {
	resultado := ""
	for indice, elemento := range elementos {
		if indice > 0 {
			resultado += separador
		}
		resultado += elemento
	}
	return resultado
}

// mapKeysAStringSlice convierte las claves de un mapa en un slice de strings.
//
// Parámetros:
//   - mapa: Mapa del que se extraerán las claves.
//
// Retorna:
//   - Un slice de strings con las claves del mapa.
func mapKeysAStringSlice(mapa map[string]bool) []string {
	claves := make([]string, 0, len(mapa))
	for clave := range mapa {
		claves = append(claves, clave)
	}
	return claves
}

// keyMaxima encuentra y devuelve la clave con el valor máximo en un mapa de enteros.
//
// Parámetros:
//   - mapa: Mapa del que se buscará la clave máxima.
//
// Retorna:
//   - La clave con el valor máximo en el mapa.
func keyMaxima(mapa map[string]int) string {
	recuentoMaximo := -1
	claveMaxima := ""
	for clave, cuenta := range mapa {
		if cuenta > recuentoMaximo {
			recuentoMaximo = cuenta
			claveMaxima = clave
		}
	}
	return claveMaxima
}