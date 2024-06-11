package ejercicio

import (
    "testing"
)

func TestAgregarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
        Calorias:          100,
    }
    if err := gestor.AgregarEjercicio(ejercicio); err != nil {
        t.Errorf("Error al agregar ejercicio: %v", err)
    }
    if err := gestor.AgregarEjercicio(ejercicio); err == nil {
        t.Error("Se esperaba un error al intentar agregar un ejercicio duplicado, pero no se recibió error")
    }
}

func TestListarEjercicios(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio1 := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
        Calorias:          100,
        Tipo:              "Fuerza",
        Puntos:            10,
        Dificultad:        "Media",
    }
    ejercicio2 := &Ejercicio{
        Nombre:            "Flexiones",
        TiempoEnSegundos: 45,
        Calorias:          80,
        Tipo:              "Fuerza",
        Puntos:            8,
        Dificultad:        "Alta",
    }
    gestor.AgregarEjercicio(ejercicio1)
    gestor.AgregarEjercicio(ejercicio2)
    ejercicios := gestor.ListarEjercicios()
    if len(ejercicios) != 2 {
        t.Errorf("Se esperaban 2 ejercicios, pero se obtuvieron %d", len(ejercicios))
    }
    if ejercicios[0].Nombre != "Sentadillas" {
        t.Errorf("Se esperaba el ejercicio 'Sentadillas', pero se obtuvo '%s'", ejercicios[0].Nombre)
    }
    if ejercicios[1].Nombre != "Flexiones" {
        t.Errorf("Se esperaba el ejercicio 'Flexiones', pero se obtuvo '%s'", ejercicios[1].Nombre)
    }
}

func TestEliminarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio{
        Nombre: "Burpees",
    }
    gestor.AgregarEjercicio(ejercicio)
    if err := gestor.EliminarEjercicio("Burpees"); err != nil {
        t.Errorf("Error al eliminar el ejercicio: %v", err)
    }
    if _, err := gestor.ConsultarEjercicio("Burpees"); err == nil {
        t.Error("Se esperaba un error al intentar consultar un ejercicio eliminado, pero no se recibió error")
    }
}

func TestConsultarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio{
        Nombre:            "Plancha",
        TiempoEnSegundos: 30,
    }
    gestor.AgregarEjercicio(ejercicio)
    resultado, err := gestor.ConsultarEjercicio("Plancha")
    if err != nil {
        t.Errorf("Error al consultar el ejercicio: %v", err)
    }
    if resultado.Nombre != "Plancha" {
        t.Errorf("El nombre del ejercicio esperado era 'Plancha', se obtuvo: %s", resultado.Nombre)
    }
}

func TestModificarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
        Calorias:          100,
    }
    gestor.AgregarEjercicio(ejercicio)
    nuevoEjercicio := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 45,
        Calorias:          150,
    }
    if err := gestor.ModificarEjercicio("Sentadillas", nuevoEjercicio); err != nil {
        t.Errorf("Error al modificar el ejercicio: %v", err)
    }
    ejercicios := gestor.ListarEjercicios()
    if len(ejercicios) != 1 {
        t.Errorf("Se esperaba 1 ejercicio después de la modificación, pero se obtuvieron %d", len(ejercicios))
    }
    if ejercicios[0].Nombre != "Sentadillas" || ejercicios[0].TiempoEnSegundos != 45 || ejercicios[0].Calorias != 150 {
        t.Errorf("El ejercicio no se modificó correctamente. Se esperaba Nombre='Sentadillas', TiempoEnSegundos=45, Calorias=150, pero se obtuvo Nombre='%s', TiempoEnSegundos=%d, Calorias=%d", ejercicios[0].Nombre, ejercicios[0].TiempoEnSegundos, ejercicios[0].Calorias)
    }
}

func TestObtenerEjercicioPorNombre(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio1 := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
    }
    ejercicio2 := &Ejercicio{
        Nombre:            "Flexiones",
        TiempoEnSegundos: 45,
    }
    gestor.AgregarEjercicio(ejercicio1)
    gestor.AgregarEjercicio(ejercicio2)
    nombres := []string{"Sentadillas", "Flexiones", "Plancha"}
    ejercicios := gestor.ObtenerEjercicioPorNombre(nombres)
    if len(ejercicios) != 2 {
        t.Errorf("Se esperaban 2 ejercicios, pero se obtuvieron %d", len(ejercicios))
    }
}

func TestFiltrarPorTiposYDificultad(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio1 := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
        Tipo:              "Fuerza",
        Dificultad:        "Media",
    }
    ejercicio2 := &Ejercicio{
        Nombre:            "Flexiones",
        TiempoEnSegundos: 45,
        Tipo:              "Fuerza",
        Dificultad:        "Alta",
    }
    ejercicio3 := &Ejercicio{
        Nombre:            "Carrera",
        TiempoEnSegundos: 300,
        Tipo:              "Cardio",
        Dificultad:        "Alta",
    }
    gestor.AgregarEjercicio(ejercicio1)
    gestor.AgregarEjercicio(ejercicio2)
    gestor.AgregarEjercicio(ejercicio3)
    ejerciciosFiltrados := gestor.FiltrarPorTiposYDificultad([]string{"Fuerza"}, "Alta")
    if len(ejerciciosFiltrados) != 1 {
        t.Errorf("Se esperaba 1 ejercicio filtrado, pero se obtuvieron %d", len(ejerciciosFiltrados))
    }
}

func TestOrdenarTiempoMenorAMayor(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio1 := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
    }
    ejercicio2 := &Ejercicio{
        Nombre:            "Flexiones",
        TiempoEnSegundos: 45,
    }
    ejercicio3 := &Ejercicio{
        Nombre:            "Plancha",
        TiempoEnSegundos: 30,
    }
    gestor.AgregarEjercicio(ejercicio1)
    gestor.AgregarEjercicio(ejercicio2)
    gestor.AgregarEjercicio(ejercicio3)
    ejercicios := gestor.ListarEjercicios()
    ejerciciosOrdenados := gestor.OrdenarTiempoMenorAMayor(ejercicios)
    if ejerciciosOrdenados[0].Nombre != "Plancha" {
        t.Errorf("Se esperaba el ejercicio 'Plancha' como el primero en la lista ordenada, pero se obtuvo '%s'", ejerciciosOrdenados[0].Nombre)
    }
}

func TestFiltrarPorTipoPuntosYDuracionMaxima(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio1 := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
        Tipo:              "Fuerza",
        Puntos:            10,
    }
    ejercicio2 := &Ejercicio{
        Nombre:            "Flexiones",
        TiempoEnSegundos: 45,
        Tipo:              "Fuerza",
        Puntos:            8,
    }
    ejercicio3 := &Ejercicio{
        Nombre:            "Carrera",
        TiempoEnSegundos: 300,
        Tipo:              "Cardio",
        Puntos:            12,
    }
    gestor.AgregarEjercicio(ejercicio1)
    gestor.AgregarEjercicio(ejercicio2)
    gestor.AgregarEjercicio(ejercicio3)
    ejerciciosFiltrados := gestor.FiltrarPorTipoPuntosYDuracionMaxima("Fuerza", 1)
    if len(ejerciciosFiltrados) != 2 {
        t.Errorf("Se esperaban 2 ejercicios filtrados, pero se obtuvieron %d", len(ejerciciosFiltrados))
    }
}

func TestOrdenarPorPuntajeDescendente(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio1 := &Ejercicio{
        Nombre:            "Sentadillas",
        TiempoEnSegundos: 60,
        Puntos:            10,
    }
    ejercicio2 := &Ejercicio{
        Nombre:            "Flexiones",
        TiempoEnSegundos: 45,
        Puntos:            8,
    }
    ejercicio3 := &Ejercicio{
        Nombre:            "Carrera",
        TiempoEnSegundos: 300,
        Puntos:            12,
    }
    gestor.AgregarEjercicio(ejercicio1)
    gestor.AgregarEjercicio(ejercicio2)
    gestor.AgregarEjercicio(ejercicio3)
    ejercicios := gestor.ListarEjercicios()
    ejerciciosOrdenados := gestor.OrdenarPorPuntajeDescendente(ejercicios)
    if ejerciciosOrdenados[0].Nombre != "Carrera" {
        t.Errorf("Se esperaba el ejercicio 'Carrera' como el primero en la lista ordenada, pero se obtuvo '%s'", ejerciciosOrdenados[0].Nombre)
    }
}