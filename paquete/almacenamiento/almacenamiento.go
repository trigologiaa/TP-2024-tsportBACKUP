package almacenamiento

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"TP-2024-TSPORT/paquete/rutina"
	"io"
	"os"
	"path/filepath"
	"github.com/gocarina/gocsv"
)


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
//   - ejercicios: slice de punteros a la estructura Ejercicio que se guardará en el archivo.
//   - nombreDeArchivo: ruta del archivo donde se guardarán los ejercicios.
//
// Retorna:
//   - Un error en caso de problemas al abrir o escribir en el archivo.
//   - El método MarshalFile en caso de ejecución correcta.
func GuardarEjercicios(ejercicios []*ejercicio.Ejercicio, nombreDeArchivo string) error {
	rutaCompleta := filepath.Join("/home/lauty/Documents/AYP2/TP-2024-tsportBACKUP/informacion", nombreDeArchivo)
	archivo, err := os.OpenFile(rutaCompleta, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer archivo.Close()
	return gocsv.MarshalFile(&ejercicios, archivo)
}

// CargarEjercicios carga una lista de ejercicios desde un archivo CSV.
//
// Parámetros:
//   - nombreDeArchivo: ruta del archivo desde donde se cargarán los ejercicios.
//
// Retorna:
//   - Un slice de punteros a Ejercicio.
//   - Un error si hay problemas al abrir o leer el archivo.
func CargarEjercicios(nombreDeArchivo string) ([]*ejercicio.Ejercicio, error) {
    archivo, err := os.Open(filepath.Join("D:/UNTREF/AlgoritmosyProgramaciónII/TP-2024-tsport/informacion", nombreDeArchivo))
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

// CargarEjerciciosDeUna carga una lista de ejercicios desde un archivo CSV.
//
// Retorna:
//   - Un slice de punteros a Ejercicio.
//   - Un error si hay problemas al abrir o leer el archivo.
func CargarEjerciciosDeUna() ([]*ejercicio.Ejercicio, error) {
	nombreDeArchivo := "ejercicios.csv"
	archivo, err := os.Open(filepath.Join("D:/UNTREF/AlgoritmosyProgramaciónII/TP-2024-tsport/informacion", nombreDeArchivo))
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
//   - rutinas: slice de punteros a la estructura Rutina que se guardará en el archivo.
//   - gestorDeEjercicios: puntero a GestorDeEjercicios usado para calcular las propiedades de las rutinas.
//   - nombreDeArchivo: ruta del archivo donde se guardarán las rutinas.
//
// Retorna:
//   - Un error en caso de problemas al abrir o escribir en el archivo.
func GuardarRutinas(rutinas []*rutina.Rutina, gestorDeEjercicios *ejercicio.GestorDeEjercicios, nombreDeArchivo string) error {
    rutaCompleta := filepath.Join("D:/UNTREF/AlgoritmosyProgramaciónII/TP-2024-tsport/informacion", nombreDeArchivo)
    archivo, err := os.OpenFile(rutaCompleta, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0644)
    if err != nil {
        return err
    }
    defer archivo.Close()
    for _, rutina := range rutinas {
        rutina.CalcularPropiedades(gestorDeEjercicios)
    }
    return gocsv.MarshalFile(&rutinas, archivo)
}

// CargarRutinas carga una lista de rutinas desde un archivo CSV.
//
// Parámetros:
//   - nombreDeArchivo: ruta del archivo desde donde se cargarán las rutinas.
//
// Retorna:
//   - Un slice de punteros a Rutina.
//   - Un error si hay problemas al abrir o leer el archivo.
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

// CargarRutinasDeUna carga una lista de rutinas desde un archivo CSV.
//
// Retorna:
//   - Un slice de punteros a Rutina.
//   - Un error si hay problemas al abrir o leer el archivo.
func CargarRutinasDeUna() ([]*rutina.Rutina, error) {
    nombreDeArchivo := "rutinas.csv"
    archivo, err := os.Open(filepath.Join("D:/UNTREF/AlgoritmosyProgramaciónII/TP-2024-tsport/informacion", nombreDeArchivo))
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