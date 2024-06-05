// Package rutina gestiona rutinas de ejercicios con funcionalidades para crear, modificar, eliminar y consultar rutinas, integrando cálculos basados en un gestor de ejercicios.
package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"errors"
	"os"

	"github.com/gocarina/gocsv"
	list "github.com/untref-ayp2/data-structures/list"
)

// Rutina es una estructura con las características que tendrá su campo.
//
// Funcionamiento:
//
//	Se define el nombre de la rutina
//	Se define el conjunto de ejercicios que contiene la rutina
//	Se define el tiempo en segundos
//	Se define la cantidad total de calorías quemadas
//	Se define la dificultad de la rutina
//	Se define el tipo de rutina
//	Se define la cantidad de puntos totales
type Rutina struct {
	Nombre        string `csv:"Nombre"`        //Nombre de la rutina
	Ejercicios    string `csv:"Ejercicios"`    //Ejercicios que contiene la rutina
	Tiempo        int    `csv:"Tiempo"`        //Tiempo total en segundo
	Calorias      int    `csv:"Calorias"`      //Calorías quemadas en total
	Dificultad    string `csv:"Dificultad"`    //Dificultad de la rutina
	Tipos         string `csv:"Tipos"`         //Tipo de rutina
	PuntosPorTipo int    `csv:"PuntosPorTipo"` //Puntos totales por tipo
}

// GestorRutinas es una estructura para gestionar las rutinas.
//
// Funcionamiento:
//
//	Se define el campo 'rutinas' como puntero a una lista doblemente enlazada de punteros a 'Rutina'
//	Se define el campo 'gestorEjercicios' como puntero a una instancia de 'GestorEjercicios' del paquete 'ejercicio'
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
//
// Funcionamiento:
//
//	Se retorna una nueva instancia de 'GestorRutinas' con su dirección de memoria, creando un nuevo puntero {
//	    Inicializa 'rutinas' con una lista doblemente enlazada
//	    Inicializa 'gestorEjercicios' con el valor del parámetro 'gestorEj'
//	}
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
//
// Funcionamiento:
//
//	Se recorre la lista de rutinas {
//	    Si el nombre de la rutina ya existe {
//	        Se retorna un error
//	    }
//	}
//	Si no se encuentra una rutina, se calculan las propiedades de 'gestorEjercicios'
//	Se agrega la rutina al final de la lista
//	Se retorna nil
func (g *GestorRutinas) AgregarRutina(rutina *Rutina) error {
	for node := g.rutinas.Head(); node != nil; node = node.Next() {
		if node.Data().Nombre == rutina.Nombre {
			return errors.New("la rutina ya existe")
		}
	}
	rutina.CalcularPropiedades(g.gestorEjercicios)
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
//
// Funcionamiento:
//
//	Se recorre la lista de rutinas {
//	    Si el nombre de la rutina es el nombre buscado {
//	        Se elimina el nodo que contiene la rutina de la lista
//	        Se retorna nil
//	    }
//	}
//	Si no se encuentra la rutina, se retorna un error
func (g *GestorRutinas) EliminarRutina(nombre string) error {
	for node := g.rutinas.Head(); node != nil; node = node.Next() {
		if node.Data().Nombre == nombre {
			g.rutinas.Remove(node.Data())
			return nil
		}
	}
	return errors.New("rutina no encontrada")
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
//
// Funcionamiento:
//
//	Se recorre la lista de rutinas {
//	    Si el nombre de la rutina es el nombre buscado {
//	        Se calculan las propiedades de la nueva rutina pasando el gestorEjercicios como argumento
//	        Se actualizan los datos del nodo con nuevaRutina
//	        Se retorna nil
//	    }
//	}
//	Si no se encuentra la rutina, se retorna un error
func (g *GestorRutinas) ModificarRutina(nombre string, nuevaRutina *Rutina) error {
	for node := g.rutinas.Head(); node != nil; node = node.Next() {
		if node.Data().Nombre == nombre {
			nuevaRutina.CalcularPropiedades(g.gestorEjercicios)
			node.SetData(nuevaRutina)
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
//
// Funcionamiento:
//
//	Se recorre la lista de rutinas {
//	    Si el nombre de la rutina es el nombre buscado {
//	        Se retornan los datos del nodo (la rutina) y un nil
//	    }
//	}
//	Si no se encuentra la rutina, se retorna un nil y un error
func (g *GestorRutinas) ConsultarRutina(nombre string) (*Rutina, error) {
	for node := g.rutinas.Head(); node != nil; node = node.Next() {
		if node.Data().Nombre == nombre {
			return node.Data(), nil
		}
	}
	return nil, errors.New("rutina no encontrada")
}

// ListarRutinas devuelve una lista de rutinas que coinciden con una dificultad específica
//
// Parámetros:
//   - 'dificultad' será la dificultad de la rutina a buscar.
//
// Retorna:
//   - Un slice de punteros a Rutina que coinciden con la dificultad especificada.
//
// Funcionamiento:
//
//	Se declara la variable 'resultado' de tipo slice de punteros a Rutina (vacío)
//	Se recorre la lista de rutinas {
//	    Si la dificultad de la rutina es la dificultad buscada {
//	        Se agrega la rutina del nodo actual al slice 'resultado'
//	    }
//	}
//	Se retorna 'resultado'
func (g *GestorRutinas) ListarRutinas(dificultad string) []*Rutina {
	resultado := []*Rutina{}
	for node := g.rutinas.Head(); node != nil; node = node.Next() {
		if node.Data().Dificultad == dificultad {
			resultado = append(resultado, node.Data())
		}
	}
	return resultado
}

// CalcularPropiedades calcula y actualiza las propiedades de una rutina basadas en los ejercicios disponibles.
//
// Parámetros:
//   - 'gestor' como puntero al gestorEjercicios que proporciona acceso a los ejercicios disponibles.
//
// Funcionamiento:
//
//	Se crea la variable 'ejercicios' de tipo slice de punteros de Ejercicio, que contiene el resultado del método 'ObtenerTodosLosEjercicios' de gestor
//	Se declaran las variables 'totalTiempo', 'totalCalorias' y 'totalPuntos' de tipo int para acumular tiempo, calorías y puntos totales de los ejercicios
//	Se declara la variable 'nombresEjercicios' de tipo slice de string (vacío)
//	Se declara la variable 'tiposSet' de tipo map para almacenar los tipos únicos de ejercicios
//	Se declara la variable 'dificultades' de tipo mal para contar la frecuencia de cada nivel de dificultad
//	Se recorre la lista de ejercicios {
//	    Se suma el tiempo del ejercicio actual a 'totalTiempo'
//	    Se suma las calorías del ejercicio actual a 'totalCalorias'
//	    Se suma los puntos del ejercicio actual a 'totalPuntos'
//	    Se añade el nombre del ejercicio actual al slice 'nombresEjercicios'
//	    Se marca el tipo del ejercicio actual como presente en el map 'tiposSet'
//	    Se incrementa el conteo de la dificultad del ejercicio actual en el map 'dificultades'
//	}
//	Se asigna el tiempo acumulado de 'totalTiempo' a la rutina
//	Se asigna las calorías acumuladas de 'totalCalorias' a la rutina
//	Se combina los nombres de los ejercicios
//	Se combina los tipos de los ejercicios
//	Se asigna los puntos totales de 'totalPuntos' a la rutina
//	Si la dificultad de la rutina es una cadena vacía {
//	    Se asigna el nivel de dificultad más frecuente del map 'dificultades'
//	}
func (r *Rutina) CalcularPropiedades(gestor *ejercicio.GestorEjercicios) {
	ejercicios := gestor.ObtenerTodosLosEjercicios()
	var totalTiempo, totalCalorias, totalPuntos int
	nombresEjercicios := make([]string, 0)
	tiposSet := make(map[string]bool)
	dificultades := make(map[string]int)
	for _, ejercicio := range ejercicios {
		totalTiempo += ejercicio.Tiempo
		totalCalorias += ejercicio.Calorias
		totalPuntos += ejercicio.Puntos
		nombresEjercicios = append(nombresEjercicios, ejercicio.Nombre)
		tiposSet[ejercicio.Tipo] = true
		dificultades[ejercicio.Dificultad]++
	}
	r.Tiempo = totalTiempo
	r.Calorias = totalCalorias
	r.Ejercicios = unirKeys(nombresEjercicios, ", ")
	r.Tipos = unirKeys(mapKeysAStringSlice(tiposSet), ", ")
	r.PuntosPorTipo = totalPuntos
	if r.Dificultad == "" {
		r.Dificultad = keyMaxima(dificultades)
	}
}

// unirKeys une una lista de cadenas usando un separador dado.
//
// Parámetros:
//   - 'elementos' será un slice de string que se unirán.
//   - 'separador' como un string que se usará como separador entre los elementos.
//
// Retorna:
//   - Un string que contiene todos los elementos del slice unidos por el separador especificado.
//
// Funcionamiento:
//
//	Se declara la variable 'resultado' de tipo string vacío
//	Se recorren los elementos de 'elemento' con 'indice' como índice del elemento actual {
//	    Si 'indice' no es el primer elemento  de 'elementos' {
//	        Se agrega el separador al resultado
//	    }
//	    Si no, se agrega el elemento actual de 'elemento' a 'resultado'
//	}
//	Se retorna 'resultado'
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
//
// Funcionamiento:
//
//	Se declara la variable 'claves' de tipo slice de String con longitud inicial de 0 y capacidad igual a la longitud de 'mapa'
//	Se recorre cada clave del map 'mapa' {
//	    Se agrega la clave actual al slice 'claves'
//	}
//	Se retorna 'claves'
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
//
// Funcionamiento:
//
//	Se declara la variable 'recuentoMaximo' de tipo int con un valor de -1
//	Se declara la variable 'claveMaxima' de tipo String vacío
//	Se recorre cada par clave-valor del map 'mapa' {
//	    Si el valor 'cuenta' es mayor al valor 'recuentoMaximo' {
//	        Se actualiza 'recuentoMaximo' al valor de 'cuenta'
//	        Se actualiza 'claveMaxima' al valor de 'clave'
//	    }
//	}
//	Se retorna 'claveMaxima'
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

// GuardarRutinas guarda una lista de rutinas en un archivo CSV.
//
// Parámetros:
//   - 'rutinas' será un slice de punteros a la estructura 'Rutina'.
//   - 'gestor' será un puntero a 'GestorEjercicios', usado para calcular las propiedades de las rutinas.
//   - 'nombreDeArchivo' será un String con la ruta del archivo donde se guardarán las rutinas.
//
// Retorna:
//   - Un error en caso de que ocurra un problema al abrir o escribir en el archivo.
//
// Funcionamiento:
//
//	Se declaran las variables 'archivo' y 'error' que funcionarán como receptores del retorno del método 'OpenFile' de 'os', que recibe los parámetros 'nombreDeArchivo' como string, 'os.O_RDWR (Permite la lectura y escritura en el archivo) | os.O_CREATE (Crea el archivo si no existe) | os.O_TRUNC (Vacía el archivo si ya existe)' como int y '0644' como 'fs.FileMode'
//	Si hubo un error al abrir el archivo {
//	    Se retorna un error
//	}
//	Se difiere al método 'Close' hasta que el método 'GuardarRutinas' termine
//	Se recorre cada ´rutina´ de 'rutinas' {
//	    Se llama al método 'CalcularPropiedades' pasándole como argumento 'gestor'
//	}
//	Se retorna el método 'MarshalFile' que serializa los datos a formato CSV y los escribe en el archivo especificado, recibiendo como argumentos '&rutinas' que es la dirección de memoria y lo deja como interface y 'archivo' como puntero a 'os.File'
func GuardarRutinas(rutinas []*Rutina, gestor *ejercicio.GestorEjercicios, nombreDeArchivo string) error {
	archivo, err := os.OpenFile(nombreDeArchivo, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer archivo.Close()
	for _, rutina := range rutinas {
		rutina.CalcularPropiedades(gestor)
	}
	return gocsv.MarshalFile(&rutinas, archivo)
}

// CargarRutinas carga una lista de rutinas desde un archivo CSV.
//
// Parámetros:
//   - 'nombreDeArchivo' será una cadena con la ruta del archivo desde donde se cargarán las rutinas.
//
// Retorna:
//   - Un slice de punteros a la estructura 'Rutina'.
//   - Un error en caso de que ocurra un problema al abrir o leer el archivo.
//
// Funcionamiento:
//
//	Se declara la variable 'rutinas' que es un slice vacío de punteros a la estructura 'Rutina'
//	Se declaran las variables 'archivo' y 'err' que funcionarán como receptores del retorno del método 'Open' de 'os', que recibe como argumento a 'nombreDeArchivo' de tipo String
//	Si hubo un error al abrir el archivo {
//	    Se retorna nil y un error
//	}
//	Se difiere el método 'Close' hasta que el método 'CargarRutinas' termine
//	Si hubo un error durante la deserealización {
//	    Se retorna nil y un error
//	}
//	Se retorna 'rutinas' y nil
func CargarRutinas(nombreDeArchivo string) ([]*Rutina, error) {
	rutinas := []*Rutina{}
	archivo, err := os.Open(nombreDeArchivo)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	if err := gocsv.UnmarshalFile(archivo, &rutinas); err != nil {
		return nil, err
	}
	return rutinas, nil
}
