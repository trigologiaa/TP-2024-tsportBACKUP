package almacenamiento

import (
    "io"
    "os"
    "github.com/gocarina/gocsv"
    "TP-2024-TSPORT/paquete/ejercicio"
    "TP-2024-TSPORT/paquete/rutina"
)

// Inicialización con configuraciones específicas para la escritura de CSV
func init() {
    gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
        return gocsv.DefaultCSVWriter(out)
    })
    gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
        return gocsv.DefaultCSVReader(in)
    })
}

// GuardarEjercicios guarda una lista de ejercicios en un archivo CSV
func GuardarEjercicios(ejercicios []*ejercicio.Ejercicio, filepath string) error {
    file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    return gocsv.MarshalFile(&ejercicios, file)
}

// CargarEjercicios carga una lista de ejercicios desde un archivo CSV
func CargarEjercicios(filepath string) ([]*ejercicio.Ejercicio, error) {
    ejercicios := []*ejercicio.Ejercicio{}
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    if err := gocsv.UnmarshalFile(file, &ejercicios); err != nil {
        return nil, err
    }
    return ejercicios, nil
}

// GuardarRutinas guarda una lista de rutinas en un archivo CSV
func GuardarRutinas(rutinas []*rutina.Rutina, gestorEj *ejercicio.GestorEjercicios, filepath string) error {
    file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    for _, r := range rutinas {
        r.CalcularPropiedades(gestorEj)
    }
    return gocsv.MarshalFile(&rutinas, file)
}

// CargarRutinas carga una lista de rutinas desde un archivo CSV
func CargarRutinas(filepath string) ([]*rutina.Rutina, error) {
    rutinas := []*rutina.Rutina{}
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    if err := gocsv.UnmarshalFile(file, &rutinas); err != nil {
        return nil, err
    }
    return rutinas, nil
}