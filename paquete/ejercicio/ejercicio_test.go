package ejercicio

import (
    "testing"
)

func TestAgregarEjercicio(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
        CaloriasDeEjercicio:            100,
    }
    if err := gestorDeEjercicios.AgregarEjercicio(ejercicio); err != nil {
        tester.Errorf("Error al agregar ejercicio: %v", err)
    }
    if err := gestorDeEjercicios.AgregarEjercicio(ejercicio); err == nil {
        tester.Error("Se esperaba un error al intentar agregar un ejercicio duplicado, pero no se recibió error")
    }
}

func TestListarEjercicios(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio1 := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
        CaloriasDeEjercicio:            100,
        TipoDeEjercicio:                "Fuerza",
        PuntosPorTipoDeEjercicio:       10,
        DificultadDeEjercicio:          "Media",
    }
    ejercicio2 := &Ejercicio {
        NombreDeEjercicio:              "Flexiones",
        TiempoEnSegundosDeEjercicio:    45,
        CaloriasDeEjercicio:            80,
        TipoDeEjercicio:                "Fuerza",
        PuntosPorTipoDeEjercicio:       8,
        DificultadDeEjercicio:          "Alta",
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio1)
    gestorDeEjercicios.AgregarEjercicio(ejercicio2)
    ejercicios := gestorDeEjercicios.ListarEjercicios()
    if len(ejercicios) != 2 {
        tester.Errorf("Se esperaban 2 ejercicios, pero se obtuvieron %d", len(ejercicios))
    }
    if ejercicios[0].NombreDeEjercicio != "Sentadillas" {
        tester.Errorf("Se esperaba el ejercicio 'Sentadillas', pero se obtuvo '%s'", ejercicios[0].NombreDeEjercicio)
    }
    if ejercicios[1].NombreDeEjercicio != "Flexiones" {
        tester.Errorf("Se esperaba el ejercicio 'Flexiones', pero se obtuvo '%s'", ejercicios[1].NombreDeEjercicio)
    }
}

func TestEliminarEjercicio(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio := &Ejercicio {
        NombreDeEjercicio:  "Burpees",
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio)
    if err := gestorDeEjercicios.EliminarEjercicio("Burpees"); err != nil {
        tester.Errorf("Error al eliminar el ejercicio: %v", err)
    }
    if _, err := gestorDeEjercicios.ConsultarEjercicio("Burpees"); err == nil {
        tester.Error("Se esperaba un error al intentar consultar un ejercicio eliminado, pero no se recibió error")
    }
}

func TestConsultarEjercicio(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio := &Ejercicio {
        NombreDeEjercicio:              "Plancha",
        TiempoEnSegundosDeEjercicio:    30,
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio)
    resultado, err := gestorDeEjercicios.ConsultarEjercicio("Plancha")
    if err != nil {
        tester.Errorf("Error al consultar el ejercicio: %v", err)
    }
    if resultado.NombreDeEjercicio != "Plancha" {
        tester.Errorf("El nombre del ejercicio esperado era 'Plancha', se obtuvo: %s", resultado.NombreDeEjercicio)
    }
}

func TestModificarEjercicio(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
        CaloriasDeEjercicio:            100,
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio)
    nuevoEjercicio := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    45,
        CaloriasDeEjercicio:            150,
    }
    if err := gestorDeEjercicios.ModificarEjercicio("Sentadillas", nuevoEjercicio); err != nil {
        tester.Errorf("Error al modificar el ejercicio: %v", err)
    }
    ejercicios := gestorDeEjercicios.ListarEjercicios()
    if len(ejercicios) != 1 {
        tester.Errorf("Se esperaba 1 ejercicio después de la modificación, pero se obtuvieron %d", len(ejercicios))
    }
    if ejercicios[0].NombreDeEjercicio != "Sentadillas" || ejercicios[0].TiempoEnSegundosDeEjercicio != 45 || ejercicios[0].CaloriasDeEjercicio != 150 {
        tester.Errorf("El ejercicio no se modificó correctamente. Se esperaba Nombre='Sentadillas', TiempoEnSegundos=45, Calorias=150, pero se obtuvo Nombre='%s', TiempoEnSegundos=%d, Calorias=%d", ejercicios[0].NombreDeEjercicio, ejercicios[0].TiempoEnSegundosDeEjercicio, ejercicios[0].CaloriasDeEjercicio)
    }
}

func TestObtenerEjercicioPorNombre(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio1 := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
    }
    ejercicio2 := &Ejercicio {
        NombreDeEjercicio:              "Flexiones",
        TiempoEnSegundosDeEjercicio:    45,
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio1)
    gestorDeEjercicios.AgregarEjercicio(ejercicio2)
    nombres := []string{"Sentadillas", "Flexiones", "Plancha"}
    ejercicios := gestorDeEjercicios.ObtenerEjercicioPorNombre(nombres)
    if len(ejercicios) != 2 {
        tester.Errorf("Se esperaban 2 ejercicios, pero se obtuvieron %d", len(ejercicios))
    }
}

func TestFiltrarPorTiposYDificultad(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio1 := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
        TipoDeEjercicio:                "Fuerza",
        DificultadDeEjercicio:          "Media",
    }
    ejercicio2 := &Ejercicio {
        NombreDeEjercicio:              "Flexiones",
        TiempoEnSegundosDeEjercicio:    45,
        TipoDeEjercicio:                "Fuerza",
        DificultadDeEjercicio:          "Alta",
    }
    ejercicio3 := &Ejercicio {
        NombreDeEjercicio:              "Carrera",
        TiempoEnSegundosDeEjercicio:    300,
        TipoDeEjercicio:                "Cardio",
        DificultadDeEjercicio:          "Alta",
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio1)
    gestorDeEjercicios.AgregarEjercicio(ejercicio2)
    gestorDeEjercicios.AgregarEjercicio(ejercicio3)
    ejerciciosFiltrados := gestorDeEjercicios.FiltrarPorTiposYDificultad([]string{"Fuerza"}, "Alta")
    if len(ejerciciosFiltrados) != 1 {
        tester.Errorf("Se esperaba 1 ejercicio filtrado, pero se obtuvieron %d", len(ejerciciosFiltrados))
    }
}

func TestOrdenarTiempoMenorAMayor(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio1 := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
    }
    ejercicio2 := &Ejercicio {
        NombreDeEjercicio:              "Flexiones",
        TiempoEnSegundosDeEjercicio:    45,
    }
    ejercicio3 := &Ejercicio {
        NombreDeEjercicio:              "Plancha",
        TiempoEnSegundosDeEjercicio:    30,
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio1)
    gestorDeEjercicios.AgregarEjercicio(ejercicio2)
    gestorDeEjercicios.AgregarEjercicio(ejercicio3)
    ejercicios := gestorDeEjercicios.ListarEjercicios()
    ejerciciosOrdenados := gestorDeEjercicios.OrdenarTiempoMenorAMayor(ejercicios)
    if ejerciciosOrdenados[0].NombreDeEjercicio != "Plancha" {
        tester.Errorf("Se esperaba el ejercicio 'Plancha' como el primero en la lista ordenada, pero se obtuvo '%s'", ejerciciosOrdenados[0].NombreDeEjercicio)
    }
}

func TestFiltrarPorTipoPuntosYDuracionMaxima(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio1 := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
        TipoDeEjercicio:                "Fuerza",
        PuntosPorTipoDeEjercicio:       10,
    }
    ejercicio2 := &Ejercicio {
        NombreDeEjercicio:              "Flexiones",
        TiempoEnSegundosDeEjercicio:    45,
        TipoDeEjercicio:                "Fuerza",
        PuntosPorTipoDeEjercicio:       8,
    }
    ejercicio3 := &Ejercicio {
        NombreDeEjercicio:              "Carrera",
        TiempoEnSegundosDeEjercicio:    300,
        TipoDeEjercicio:                "Cardio",
        PuntosPorTipoDeEjercicio:       12,
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio1)
    gestorDeEjercicios.AgregarEjercicio(ejercicio2)
    gestorDeEjercicios.AgregarEjercicio(ejercicio3)
    ejerciciosFiltrados := gestorDeEjercicios.FiltrarPorTipoPuntosYDuracionMaxima("Fuerza", 1)
    if len(ejerciciosFiltrados) != 2 {
        tester.Errorf("Se esperaban 2 ejercicios filtrados, pero se obtuvieron %d", len(ejerciciosFiltrados))
    }
}

func TestOrdenarPorPuntajeDescendente(tester *testing.T) {
    gestorDeEjercicios := NuevoGestorDeEjercicios()
    ejercicio1 := &Ejercicio {
        NombreDeEjercicio:              "Sentadillas",
        TiempoEnSegundosDeEjercicio:    60,
        PuntosPorTipoDeEjercicio:       10,
    }
    ejercicio2 := &Ejercicio {
        NombreDeEjercicio:              "Flexiones",
        TiempoEnSegundosDeEjercicio:    45,
        PuntosPorTipoDeEjercicio:       8,
    }
    ejercicio3 := &Ejercicio {
        NombreDeEjercicio:              "Carrera",
        TiempoEnSegundosDeEjercicio:    300,
        PuntosPorTipoDeEjercicio:       12,
    }
    gestorDeEjercicios.AgregarEjercicio(ejercicio1)
    gestorDeEjercicios.AgregarEjercicio(ejercicio2)
    gestorDeEjercicios.AgregarEjercicio(ejercicio3)
    ejercicios := gestorDeEjercicios.ListarEjercicios()
    ejerciciosOrdenados := gestorDeEjercicios.OrdenarPorPuntajeDescendente(ejercicios)
    if ejerciciosOrdenados[0].NombreDeEjercicio != "Carrera" {
        tester.Errorf("Se esperaba el ejercicio 'Carrera' como el primero en la lista ordenada, pero se obtuvo '%s'", ejerciciosOrdenados[0].NombreDeEjercicio)
    }
}