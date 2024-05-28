package ejercicio

import (
    "testing"
)

// TestAgregarEjercicio verifica la funcionalidad de AgregarEjercicio. Comprueba que se puede agregar un ejercicio correctamente y que no se pueden agregar duplicados.
func TestAgregarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio{
        Nombre: "Sentadillas",
        Tiempo: 60,
        Calorias: 100,
    }
    if err := gestor.AgregarEjercicio(ejercicio); err != nil {
        t.Errorf("Error al agregar ejercicio: %v", err)
    }
    if err := gestor.AgregarEjercicio(ejercicio); err == nil {
        t.Error("Se esperaba un error al intentar agregar un ejercicio duplicado, pero no se recibió error")
    }
}

// TestEliminarEjercicio verifica la funcionalidad de EliminarEjercicio. Comprueba que se puede eliminar un ejercicio correctamente y que un ejercicio eliminado no puede ser consultado.
func TestEliminarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio{Nombre: "Burpees"}
    gestor.AgregarEjercicio(ejercicio)
    if err := gestor.EliminarEjercicio("Burpees"); err != nil {
        t.Errorf("Error al eliminar el ejercicio: %v", err)
    }
    if _, err := gestor.ConsultarEjercicio("Burpees"); err == nil {
        t.Error("Se esperaba un error al intentar consultar un ejercicio eliminado, pero no se recibió error")
    }
}

// TestConsultarEjercicio verifica la funcionalidad de ConsultarEjercicio. Asegura que se puede consultar un ejercicio correctamente y que los datos devueltos son correctos.
func TestConsultarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio{Nombre: "Plancha", Tiempo: 30}
    gestor.AgregarEjercicio(ejercicio)
    resultado, err := gestor.ConsultarEjercicio("Plancha")
    if err != nil {
        t.Errorf("Error al consultar el ejercicio: %v", err)
    }
    if resultado.Nombre != "Plancha" {
        t.Errorf("El nombre del ejercicio esperado era 'Plancha', se obtuvo: %s", resultado.Nombre)
    }
}