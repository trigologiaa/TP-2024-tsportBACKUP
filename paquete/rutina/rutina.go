// Package rutina gestiona rutinas de ejercicios con funcionalidades para crear, modificar, eliminar y consultar rutinas, integrando cálculos basados en un gestor de ejercicios.
package rutina

import (
    "errors"
    list "github.com/untref-ayp2/data-structures/list"
    "TP-2024-TSPORT/paquete/ejercicio"
)

// Rutina define las características de una rutina de ejercicios.
type Rutina struct {
    Nombre          string  `csv:"Nombre"`
    Ejercicios      string  `csv:"Ejercicios"`
    Tiempo          int     `csv:"Tiempo estimado"`
    Calorias        int     `csv:"Calorias"`
    Dificultad      string  `csv:"Dificultad"`
    Tipos           string  `csv:"Tipos"`
    PuntosPorTipo   int     `csv:"PuntosPorTipo"`
}

// GestorRutinas maneja una lista doblemente enlazada de rutinas de ejercicios, proporcionando métodos para agregar, eliminar, modificar y consultar rutinas.
type GestorRutinas struct {
    rutinas          *list.DoubleLinkedList[*Rutina]
    gestorEjercicios *ejercicio.GestorEjercicios
}

// NuevoGestorRutinas crea e inicializa un nuevo gestor de rutinas con acceso a un gestor de ejercicios existente.
func NuevoGestorRutinas(gestorEj *ejercicio.GestorEjercicios) *GestorRutinas {
    return &GestorRutinas{
        rutinas:          list.NewDoubleLinkedList[*Rutina](),
        gestorEjercicios: gestorEj,
    }
}

// AgregarRutina añade una nueva rutina a la lista, evitando duplicados.
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

// ConsultarRutina devuelve los datos de una rutina buscando por su nombre.
func (g *GestorRutinas) ConsultarRutina(nombre string) (*Rutina, error) {
    for node := g.rutinas.Head(); node != nil; node = node.Next() {
        if node.Data().Nombre == nombre {
            return node.Data(), nil
        }
    }
    return nil, errors.New("rutina no encontrada")
}

// ListarRutinas devuelve una lista de rutinas que coinciden con una dificultad específica.
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
        tiposSet[ejercicio.Tipos] = true
        dificultades[ejercicio.Dificultad]++
    }
    r.Tiempo = totalTiempo
    r.Calorias = totalCalorias
    r.Ejercicios = joinKeys(nombresEjercicios, ", ")
    r.Tipos = joinKeys(mapKeysToStringSlice(tiposSet), ", ")
    r.PuntosPorTipo = totalPuntos
    if r.Dificultad == "" {
        r.Dificultad = maxKey(dificultades)
    }
}

// joinKeys une una lista de cadenas usando un separador dado.
func joinKeys(items []string, separator string) string {
    result := ""
    for i, item := range items {
        if i > 0 {
            result += separator
        }
        result += item
    }
    return result
}

// mapKeysToStringSlice convierte un mapa de booleanos a una lista de cadenas.
func mapKeysToStringSlice(m map[string]bool) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    return keys
}

// maxKey devuelve la clave con el valor máximo en un mapa de enteros.
func maxKey(m map[string]int) string {
    maxCount := -1
    maxKey := ""
    for key, count := range m {
        if count > maxCount {
            maxCount = count
            maxKey = key
        }
    }
    return maxKey
}