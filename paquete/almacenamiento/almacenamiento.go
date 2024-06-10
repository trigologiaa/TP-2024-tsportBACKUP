package almacenamiento

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"TP-2024-TSPORT/paquete/rutina"
	"io"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

// init genera una inicialización con configuraciones específicas para la escritura de CSV.
//
// Funcionamiento:
//
//	Se establece un escritor de CSV seguro por defecto llamando al método 'SetCSWriter' de 'gocsv' usando como argumento una función anónima que usa 'out' de 'io.Writer' como argumento y retorna un puntero a 'gocsv.SafeCSVWriter' {
//	    Se retorna un llamado al método 'DefaultCSVWriter' de 'gocsv' recibiendo como argumento a 'out'.
//	}
//	Se establece un lector de CSV por defecto llamando al método 'SetCSVReader' de 'gocsv' usando como argumento una función anónima que usa 'in' de 'io.Reader' como argumento y retorna un 'CSVReader' de 'gocsv' {
//	    Se retorna un llamado al método 'DefaultCSVReader' de 'gocsv' recibiendo como argumento a 'in
//	}
func init() {
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		return gocsv.DefaultCSVWriter(out)
	})
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		return gocsv.DefaultCSVReader(in)
	})
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
func GuardarEjercicios(ejercicios []*ejercicio.Ejercicio, nombreDeArchivo string) error {
	// Especifica la ruta completa del archivo
	rutaCompleta := filepath.Join("/home/lauty/Documents/AYP2/TP-2024-tsportBACKUP/informacion", nombreDeArchivo)
	archivo, err := os.OpenFile(rutaCompleta, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
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
func CargarEjercicios(nombreDeArchivo string) ([]*ejercicio.Ejercicio, error) {
	archivo, err := os.Open(filepath.Join("/home/lauty/Documents/AYP2/TP-2024-tsportBACKUP/informacion", nombreDeArchivo))
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	ejercicios := []*ejercicio.Ejercicio{}
	if err := gocsv.UnmarshalFile(archivo, &ejercicios); err != nil {
		return nil, err
	}
	return ejercicios, nil
}

func CargarEjerciciosDeUna() ([]*ejercicio.Ejercicio, error) {
	nombreDeArchivo := "ejercicios.csv"
	archivo, err := os.Open(filepath.Join("/home/lauty/Documents/AYP2/TP-2024-tsportBACKUP/informacion", nombreDeArchivo))
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	ejercicios := []*ejercicio.Ejercicio{}
	if err := gocsv.UnmarshalFile(archivo, &ejercicios); err != nil {
		return nil, err
	}
	return ejercicios, nil
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
func GuardarRutinas(rutinas []*rutina.Rutina, gestor *ejercicio.GestorEjercicios, nombreDeArchivo string) error {
	// Especifica la ruta completa del archivo
	rutaCompleta := filepath.Join("/home/lauty/Documents/AYP2/TP-2024-tsportBACKUP/informacion", nombreDeArchivo)
	archivo, err := os.OpenFile(rutaCompleta, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
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
func CargarRutinas(nombreDeArchivo string) ([]*rutina.Rutina, error) {
	archivo, err := os.Open(filepath.Join("/home/lauty/Documents/AYP2/TP-2024-tsportBACKUP/informacion", nombreDeArchivo))
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	rutinas := []*rutina.Rutina{}
	if err := gocsv.UnmarshalFile(archivo, &rutinas); err != nil {
		return nil, err
	}
	return rutinas, nil
}
