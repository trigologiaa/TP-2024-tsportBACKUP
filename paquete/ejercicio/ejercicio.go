// Package ejercicio gestiona una colección de ejercicios físicos utilizando una lista doblemente enlazada.
// Ofrece funcionalidades para agregar, eliminar, consultar y listar ejercicios, ideal para aplicaciones de fitness o entrenamiento.
package ejercicio

import (
	"errors"
	"os"

	"github.com/gocarina/gocsv"
	lista "github.com/untref-ayp2/data-structures/list"
)

// Ejercicio es una estructura con las características que tendrá su campo.
//
// Funcionamiento:
//
//	Se define el nombre del ejercicio
//	Se define la descripción
//	Se define el tiempo en segundos
//	Se define la cantidad de calorías quemadas
//	Se define el tipo de ejercicio
//	Se define el puntaje por tipo de ejercicio
//	Se define el nivel de fidicultad
type Ejercicio struct {
	Nombre      string `csv:"Nombre"`      // Nombre del ejercicio
	Descripcion string `csv:"Descripcion"` // Descripción del ejercicio
	Tiempo      int    `csv:"Tiempo"`      // Tiempo estimado en segundos
	Calorias    int    `csv:"Calorias"`    // Calorías quemadas
	Tipo        string `csv:"Tipo"`        // Tipo de ejercicio
	Puntos      int    `csv:"Puntos"`      // Puntos por tipo de ejercicio
	Dificultad  string `csv:"Dificultad"`  // Nivel de dificultad
}

// GestorEjercicios es una estructura para gestionar los ejercicios.
//
// Funcionamiento:
//
//	Se añade el campo 'ejercicios' que es un puntero a una lista doblemente enlazada de punteros a 'Ejercicio'
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
//
// Funcionamiento:
//
//	Se retorna una nueva instancia 'GestorEjercicios' con su dirección de memoria, creando un puntero {
//	    Inicializa 'ejercicios' con una nueva lista doblemente enlazada
//	}
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
//
// Funcionamiento:
//
//	Se declara la variable 'nodo' que verá si existe el ejercicio en la lista
//	Si el ejercicio existe en la lista {
//	    Se retorna un error
//	}
//	Si no, se agrega el ejercicio a la lista
//	Se retorna nil
func (gestor *GestorEjercicios) AgregarEjercicio(ejercicio *Ejercicio) error {
	nodo := gestor.ejercicios.Find(ejercicio)
	if nodo != nil {
		return errors.New("el ejercicio ya existe")
	}
	gestor.ejercicios.Append(ejercicio)
	return nil
}

// ObtenerTodosLosEjercicios devuelve una lista de todos los ejercicios almacenados en la gestión.
//
// Parámetros:
//   - Sin parámetros.
//
// Retorna:
//   - Un slice con todos los ejercicios.
//
// Funcionamiento:
//
//	Se declara la variable 'ejercicios' de tipo slice de punteros a Ejercicio (vacío)
//	Se recorre la lista de ejercicios {
//	    Se agregan los punteros desde la data de 'Ejercicio' a 'ejercicios'
//	}
//	Se retorna el slice 'ejercicios'
func (g *GestorEjercicios) ObtenerTodosLosEjercicios() []*Ejercicio {
	var ejercicios []*Ejercicio
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		ejercicios = append(ejercicios, nodo.Data())
	}
	return ejercicios
}

// EliminarEjercicio busca un ejercicio por su nombre y lo elimina de la lista.
// Parámetros:
//   - 'nombre' será el nombre del ejercicio a buscar.
//
// Retorna:
//   - Un nil en caso de que se haya eliminado el ejercicio.
//   - Un error en caso de que no se encuentre en la lista.
//
// Funcionamiento:
//
//	Se recorre la lista de ejercicios {
//	    Si el nombre del ejercicio es el nombre buscado {
//	        Se elimina el nodo que contiene el ejercicio de la lista
//	        Se retorna nil
//	    }
//	}
//	Si no se encuentra el ejercicio, se retorna un error
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
//
// Funcionamiento:
//
//	Se recorre la lista de ejercicios {
//	    Si el nombre del ejercicio es el nombre buscado {
//	        Se retornan los datos del ejercicio encontrado
//	    }
//	    Si no, se retorna un error
//	}
func (g *GestorEjercicios) ConsultarEjercicio(nombre string) (*Ejercicio, error) {
	for nodo := g.ejercicios.Head(); nodo != nil; nodo = nodo.Next() {
		if nodo.Data().Nombre == nombre {
			return nodo.Data(), nil
		}
	}
	return nil, errors.New("ejercicio no encontrado")
}

// GuardarEjercicios guarda una lista de ejercicios en un archivo CSV.
//
// Parámetros:
//   - 'ejercicios' será un slice de punteros a la estructura 'Ejercicio'.
//   - 'nombreDeArchivo' será un String con la ruta del archivo donde se guardarán los ejercicios.
//
// Retorna:
//   - Un error en caso de que ocurra un problema al abrir o escribir en el archivo.
//   - El método 'MarshalFile' en caso de que se haya ejecutado correctamente.
//
// Funcionamiento:
//
//	Se declaran las variables 'archivo' y 'error' que funcionarán como receptores del retorno del método 'OpenFile' de 'os', que recibe los parámetros 'nombreDeArchivo' como string, 'os.O_RDWR (Permite la lectura y escritura en el archivo) | os.O_CREATE (Crea el archivo si no existe) | os.O_TRUNC (Vacía el archivo si ya existe)' como int y '0644' como 'fs.FileMode'
//	Si hubo un error al abrir el archivo {
//	    Se retorna un error
//	}
//	Se difiere el método 'Close' hasta que el método 'GuardarEjercicios' termine
//	Se retorna el método 'MarshalFile' que serializa los datos a formato CSV y los escribe en el archivo especificado, recibiendo como argumentos '&ejercicios' que es la dirección de memoria y lo deja como interface y 'archivo' como puntero a 'os.File'
func GuardarEjercicios(ejercicios []*Ejercicio, nombreDeArchivo string) error {
	archivo, err := os.OpenFile(nombreDeArchivo, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer archivo.Close()
	return gocsv.MarshalFile(&ejercicios, archivo)
}

// CargarEjercicios carga una lista de ejercicios desde un archivo CSV.
//
// Parámetros:
//   - 'nombreDeArchivo' será un String con la ruta del archivo desde donde se cargarán los ejercicios.
//
// Retorna:
//   - Un slice de punteros a la estructura 'Ejercicio'.
//   - Un error en caso de que ocurra un problema al abrir o leer el archivo.
//
// Funcionamiento:
//
//	Se declara la variable 'ejercicios' que es un slice vacío de punteros a la estructura Ejercicio
//	Se declaran las variables 'archivo' y 'err' que funcionarán como receptores del retorno del método 'Open' de 'os', que recibe como argumento a 'nombreDeArchivo' de tipo String
//	Si hubo un error al abrir el archivo {
//	    Se retorna nil y un error
//	}
//	Se difiere al método 'Close' hasta que el método 'CargarEjercicios' termine
//	Si hubo un error durante la deserialización {
//	    Se retorna nil y un error
//	}
//	Se retorna 'ejercicios' y nil
func CargarEjercicios(nombreDeArchivo string) ([]*Ejercicio, error) {
	ejercicios := []*Ejercicio{}
	archivo, err := os.Open(nombreDeArchivo)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	if err := gocsv.UnmarshalFile(archivo, &ejercicios); err != nil {
		return nil, err
	}
	return ejercicios, nil
}
