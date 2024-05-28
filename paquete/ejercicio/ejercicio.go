// Package ejercicio gestiona una colección de ejercicios físicos utilizando una lista doblemente enlazada. Ofrece funcionalidades para agregar, eliminar, consultar y listar ejercicios, ideal para aplicaciones de fitness o entrenamiento.
package ejercicio

import (
    "errors"
    lista "github.com/untref-ayp2/data-structures/list"
)

//Estructura de tipo Ejercicio que tiene varios campos de tipo string/int los cuales tendrán csv anclados.
type Ejercicio struct {
    Nombre        string    `csv:"Nombre"`
    Descripcion   string    `csv:"Descripcion"`
    Tiempo        int       `csv:"Tiempo estimado"`
    Calorias      int       `csv:"Calorias quemadas"`
    Tipos         string    `csv:"Tipo de ejercicio"`
    Puntos        int       `csv:"Puntos por tipo de ejercicio"`
    Dificultad    string    `csv:"Nivel de dificultad"`
}

//Estructura de tipo GestorEjercicios que recibe como campo 'ejercicios' de tipo puntero a una lista enlazada doble que recibe como parámetros un puntero a estructura de tipo Ejercicio.
type GestorEjercicios struct {
    ejercicios *lista.DoubleLinkedList[*Ejercicio]
}

//Función 'NuevoGestorEjercicios' crea un nuevo objeto 'GestorEjercicios', el cual inicializa el campo 'ejercicios' recién creado con una nueva instancia de 'DoubleLinkedList' que almacena punteros a objetos de tipo 'Ejercicio' y devuelve un puntero a 'GestorEjercicios'.
func NuevoGestorEjercicios() *GestorEjercicios {
    return &GestorEjercicios{
        ejercicios: lista.NewDoubleLinkedList[*Ejercicio](),
    }
}

// AgregarEjercicio añade un nuevo ejercicio a la lista si este no existe ya. Devuelve un error si el ejercicio ya se encuentra en la lista.
func (gestor *GestorEjercicios) AgregarEjercicio(ejercicio *Ejercicio) error {
    nodo := gestor.ejercicios.Find(ejercicio)
    if nodo != nil {
        return errors.New("el ejercicio ya existe")
    }
    gestor.ejercicios.Append(ejercicio)
    return nil
}

// ObtenerTodosLosEjercicios devuelve una lista de todos los ejercicios almacenados en la gestión.
func (g *GestorEjercicios) ObtenerTodosLosEjercicios() []*Ejercicio {
    var ejercicios []*Ejercicio
    for node := g.ejercicios.Head(); node != nil; node = node.Next() {
        ejercicios = append(ejercicios, node.Data())
    }
    return ejercicios
}

// EliminarEjercicio busca y elimina un ejercicio por su nombre. Devuelve un error si el ejercicio no se encuentra en la lista.
func (g *GestorEjercicios) EliminarEjercicio(nombre string) error {
    for node := g.ejercicios.Head(); node != nil; node = node.Next() {
        if node.Data().Nombre == nombre {
            g.ejercicios.Remove(node.Data())
            return nil
        }
    }
    return errors.New("ejercicio no encontrado")
}

// ConsultarEjercicio busca un ejercicio por su nombre y lo devuelve si lo encuentra. Devuelve un error si el ejercicio no se encuentra en la lista.
func (g *GestorEjercicios) ConsultarEjercicio(nombre string) (*Ejercicio, error) {
    for node := g.ejercicios.Head(); node != nil; node = node.Next() {
        if node.Data().Nombre == nombre {
            return node.Data(), nil
        }
    }
    return nil, errors.New("ejercicio no encontrado")
}