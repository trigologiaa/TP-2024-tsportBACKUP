package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"errors"
	//"sort"
	"strings"

	list "github.com/untref-ayp2/data-structures/list"
)

// Rutina es una estructura con las características que tendrá su campo.
type Rutina struct {
	Nombre        					string 					`csv:"Nombre"`							//Nombre de la rutina
	ListaDeEjercicios 				string 					`csv:"ListaDeEjercicios"`				//Ejercicios que contiene la rutina
	TiempoEnMinutos        			int    					`csv:"TiempoEnMinutos"`					//Tiempo total en segundos
	Calorias      					int    					`csv:"Calorias"`						//Calorías quemadas en total
	Dificultad    					string 					`csv:"Dificultad"`						//Dificultad de la rutina
	Tipos         					string 					`csv:"Tipos"`							//Tipo de rutina
	Puntos 							int    					`csv:"Puntos"`							//Puntos totales por tipo
	CaracteristicasIndividuales    	[]ejercicio.Ejercicio	`csv:"CaracteristicasIndividuales"`		//Caracteristicas individuales de cada ejercicio
}

// GestorRutinas es una estructura para gestionar las rutinas.
type GestorRutinas struct {
	rutinas          *list.DoubleLinkedList[*Rutina]
	gestorEjercicios *ejercicio.GestorEjercicios
}

// NuevoGestorRutinas crea e inicializa un nuevo gestor de rutinas con acceso a un gestor de ejercicios existente.
//
// Parámetros:
//   - 'gestorEj' será un puntero al gestor de ejercicios con el que se integrará este gestor de rutinas.
//
// Retorna:
//   - Un puntero a un nuevo GestorRutinas inicializado.
func NuevoGestorRutinas(gestorEj *ejercicio.GestorEjercicios) *GestorRutinas {
	return &GestorRutinas{
		rutinas:          list.NewDoubleLinkedList[*Rutina](),
		gestorEjercicios: gestorEj,
	}
}

// AgregarRutina añade una nueva rutina a la lista, evitando duplicados.
//
// Parámetros:
//   - 'rutina' será un puntero a Rutina.
//
// Retorna:
//   - Un nil si la rutina se agregó correctamente.
//   - Un error si la rutina ya existe en la lista.
func (g *GestorRutinas) AgregarRutina(rutina *Rutina) error {
	for nodo := g.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == rutina.Nombre {
			return errors.New("la rutina ya existe")
		}
	}
	err := rutina.CalcularPropiedades(g.gestorEjercicios)
	if err != nil {
		return err
	}
	var nombresEjercicios []string
	for _, ejercicio := range rutina.CaracteristicasIndividuales {
		nombresEjercicios = append(nombresEjercicios, ejercicio.Nombre)
	}
	rutina.ListaDeEjercicios = "\"" + strings.Join(nombresEjercicios, "\", \"") + "\""
	g.rutinas.Append(rutina)
	return nil
}

// EliminarRutina busca y elimina una rutina por su nombre.
//
// Parámetros:
//   - 'nombre' será el nombre de la rutina a buscar.
//
// Retorna:
//   - Un nil en caso de que se haya elimiado la rutina.
//   - Un error en caso de que no se encuentre en la lista.
func (g *GestorRutinas) EliminarRutina(nombre string) error {
	for nodo := g.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == nombre {
			g.rutinas.Remove(nodo.Data())
			return nil
		}
	}
	return errors.New("rutina no encontrada")
}

// ConsultarRutina devuelve los datos de la rutina buscando por su nombre.
//
// Parámetros:
//   - 'nombre' será el nombre de la rutina a buscar.
//
// Retorna:
//   - Un puntero a la rutina si se encuentra.
//   - Un error si la rutina no se encuentra.
func (g *GestorRutinas) ConsultarRutina(nombre string) (*Rutina, error) {
	for nodo := g.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == nombre {
			return nodo.Data(), nil
		}
	}
	return nil, errors.New("rutina no encontrada")
}

// ModificarRutina busca una rutina por su nombre y reemplaza sus datos con los de una nueva rutina.
//
// Parámetros:
//   - 'nombre' será el nombre de la rutina a modificar.
//   - 'nuevaRutina' será un puntero a la nueva rutina que se utilizará para reemplazar los datos de la rutina existente.
//
// Retorna:
//   - Un nil si se modifica la rutina correctamente.
//   - Un error si la rutina no se encuentra en la lista.
func (g *GestorRutinas) ModificarRutina(nombre string, nuevaRutina *Rutina) error {
	for nodo := g.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == nombre {
			nuevaRutina.CalcularPropiedades(g.gestorEjercicios)
			nodo.SetData(nuevaRutina)
			return nil
		}
	}
	return errors.New("rutina no encontrada")
}

// ObtenerTodasLasRutinas devuelve una lista de todas las rutinas almacenadas.
//
// Retorna:
//   - Un slice de punteros a Rutina que contiene todas las rutinas almacenadas.
func (g *GestorRutinas) ListarRutinas() []*Rutina {
	resultado := []*Rutina{}
	for nodo := g.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		resultado = append(resultado, nodo.Data())
	}
	return resultado
}

// ListarRutinas devuelve una lista de rutinas que coinciden con una dificultad específica.
//
// Parámetros:
//   - 'dificultad' será la dificultad de la rutina a buscar.
//
// Retorna:
//   - Un slice de punteros a Rutina que coinciden con la dificultad especificada.
func (g *GestorRutinas) ListarRutinasPorDificultad(dificultad string) []*Rutina {
	resultado := []*Rutina{}
	for nodo := g.rutinas.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Dificultad == dificultad {
			resultado = append(resultado, nodo.Data())
		}
	}
	return resultado
}

//Métodos que no van al menú sino que se llaman desde otros métodos

// CalcularPropiedades calcula y actualiza las propiedades de una rutina basadas en los ejercicios disponibles.
//
// Parámetros:
//   - 'gestor' como puntero al gestorEjercicios que proporciona acceso a los ejercicios disponibles.
func (r *Rutina) CalcularPropiedades(gestor *ejercicio.GestorEjercicios) error {
	var nombresEjercicios []string
	for _, ejercicio := range r.CaracteristicasIndividuales {
		nombresEjercicios = append(nombresEjercicios, ejercicio.Nombre)
	}
	ejercicios := gestor.ObtenerEjercicioPorNombre(nombresEjercicios)
	if len(ejercicios) == 0 {
		return errors.New("no se encontraron ejercicios válidos para esta rutina")
	}
	var totalTiempoEnSegundos, totalCalorias, totalPuntos int
	tiposSet := make(map[string]bool)
	dificultades := make(map[string]int)
	for _, ejercicio := range ejercicios {
		totalTiempoEnSegundos += ejercicio.TiempoEnSegundos
		totalCalorias += ejercicio.Calorias
		totalPuntos += ejercicio.Puntos
		tiposSet[ejercicio.Tipo] = true
		dificultades[ejercicio.Dificultad]++
	}
	r.TiempoEnMinutos = totalTiempoEnSegundos / 60
	r.Calorias = totalCalorias
	r.Tipos = unirKeys(mapKeysAStringSlice(tiposSet), ", ")
	r.Puntos = totalPuntos
	if r.Dificultad == "" {
		r.Dificultad = keyMaxima(dificultades)
	}
	return nil
}

// unirKeys une una lista de cadenas usando un separador dado.
//
// Parámetros:
//   - 'elementos' será un slice de string que se unirán.
//   - 'separador' como un string que se usará como separador entre los elementos.
// Retorna:
//   - Un string que contiene todos los elementos del slice unidos por el separador especificado.
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

// mapKeysAStringSlice convierte un mapa de bool a una lista de String.
//
// Parámetros:
//   - 'mapa' será un mapa de claves de tipo String y valores de tipo bool.
//
// Retorna:
//   - Un slice de String que contiene todas las claves del mapa.
func mapKeysAStringSlice(mapa map[string]bool) []string {
	claves := make([]string, 0, len(mapa))
	for clave := range mapa {
		claves = append(claves, clave)
	}
	return claves
}

// keyMaxima devuelve la clave con el valor máximo en un mapa de enteros.
//
// Parámetros:
//   - 'mapa' será un mapa de claves tipo String y valores de tipo int.
//
// Retorna:
//   - Un String que es la clave con el valor máximo en el mapa.
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