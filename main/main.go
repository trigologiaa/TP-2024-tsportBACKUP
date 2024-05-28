package main

import (
	/*"TP-2024-TSPORT/paquete/ejercicio"/*
	  "TP-2024-TSPORT/paquete/rutina"
	  "TP-2024-TSPORT/paquete/almacenamiento"*/
	//"TP-2024-TSPORT/extras/interaccion"
	"TP-2024-TSPORT/extras/interaccion"
)

func main() {
    interaccion.Portada()
    interaccion.Inicio()
    /*e1 := &ejercicio.Ejercicio{
        Nombre:      "Burpees",
        Descripcion: "Ejercicio de cuerpo completo",
        Tiempo:      30,
        Calorias:    10,
        Tipos:       "Cardio, Fuerza",
        Puntos:      2,
        Dificultad:  "Intermedio",
    }
    e2 := &ejercicio.Ejercicio{
        Nombre:      "Plancha",
        Descripcion: "Ejercicio de fuerza para abdomen",
        Tiempo:      60,
        Calorias:    20,
        Tipos:       "Fuerza",
        Puntos:      1,
        Dificultad:  "Principiante",
    }
    gestorEjercicios := ejercicio.NuevoGestorEjercicios()
    if err := gestorEjercicios.AgregarEjercicio(e1); err != nil {
        log.Fatal(err)
    }
    if err := gestorEjercicios.AgregarEjercicio(e2); err != nil {
        log.Fatal(err)
    }
    r := &rutina.Rutina{
        Nombre:     "Rutina de calentamiento",
        Ejercicios: e1.Nombre + ", " + e2.Nombre,
    }
    gestorRutinas := rutina.NuevoGestorRutinas(gestorEjercicios)
    if err := gestorRutinas.AgregarRutina(r); err != nil {
        log.Fatal(err)
    }
    if err := almacenamiento.GuardarEjercicios([]*ejercicio.Ejercicio{e1, e2}, "informacion/ejercicios.csv"); err != nil {
        log.Fatal(err)
    }
    if err := almacenamiento.GuardarRutinas([]*rutina.Rutina{r}, gestorEjercicios, "informacion/rutinas.csv"); err != nil {
        log.Fatal(err)
    }
    ejerciciosCargados, err := almacenamiento.CargarEjercicios("informacion/ejercicios.csv")
    if err != nil {
        log.Fatal(err)
    }
    rutinasCargadas, err := almacenamiento.CargarRutinas("informacion/rutinas.csv")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Ejercicios cargados:")
    for _, e := range ejerciciosCargados {
        fmt.Printf("%s: %s\n", e.Nombre, e.Descripcion)
    }
    fmt.Println("\nRutinas cargadas:")
    for _, r := range rutinasCargadas {
        fmt.Printf("%s: %s\n", r.Nombre, r.Dificultad)
    }*/
}