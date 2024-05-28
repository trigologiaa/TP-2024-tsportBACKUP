package rutina

import (
    "testing"
    "TP-2024-TSPORT/paquete/ejercicio"
)

// setup inicializa y configura los gestores de ejercicios y rutinas para las pruebas. Agrega ejercicios iniciales al gestor de ejercicios y retorna ambos gestores.
func setup() (*GestorRutinas, *ejercicio.GestorEjercicios) {
    gestorEjercicios := ejercicio.NuevoGestorEjercicios()
    gestorEjercicios.AgregarEjercicio(&ejercicio.Ejercicio{Nombre: "Sentadillas", Tiempo: 60, Calorias: 100, Dificultad: "Media"})
    gestorEjercicios.AgregarEjercicio(&ejercicio.Ejercicio{Nombre: "Flexiones", Tiempo: 30, Calorias: 50, Dificultad: "Alta"})
    gestorRutinas := NuevoGestorRutinas(gestorEjercicios)
    return gestorRutinas, gestorEjercicios
}

// TestAgregarRutina verifica la capacidad del gestor para agregar rutinas y manejar duplicados.
func TestAgregarRutina(t *testing.T) {
    gestorRutinas, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    rutina := &Rutina{
        Nombre: "Rutina Intensa",
        Ejercicios: ejercicios[0].Nombre + ", " + ejercicios[1].Nombre,
    }
    if err := gestorRutinas.AgregarRutina(rutina); err != nil {
        t.Errorf("AgregarRutina() error = %v, wantErr %v", err, false)
    }
    if err := gestorRutinas.AgregarRutina(rutina); err == nil {
        t.Errorf("Se esperaba un error al intentar agregar una rutina duplicada, pero no se obtuvo error")
    }
}

// TestEliminarRutina verifica que las rutinas se puedan eliminar correctamente y que no puedan ser consultadas después de su eliminación.
func TestEliminarRutina(t *testing.T) {
    gestorRutinas, _ := setup()
    rutina := &Rutina{Nombre: "Rutina Temporal"}
    gestorRutinas.AgregarRutina(rutina)
    if err := gestorRutinas.EliminarRutina("Rutina Temporal"); err != nil {
        t.Errorf("EliminarRutina() error = %v, wantErr %v", err, false)
    }
    if _, err := gestorRutinas.ConsultarRutina("Rutina Temporal"); err == nil {
        t.Error("Se esperaba un error al buscar una rutina eliminada, pero no se obtuvo error")
    }
}

// TestModificarRutina verifica que se puedan modificar las rutinas existentes correctamente.
func TestModificarRutina(t *testing.T) {
    gestorRutinas, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    original := &Rutina{
        Nombre:     "Rutina Mañana",
        Ejercicios: ejercicios[0].Nombre,
    }
    gestorRutinas.AgregarRutina(original)
    modificada := &Rutina{
        Nombre:     "Rutina Mañana",
        Ejercicios: ejercicios[1].Nombre,
    }
    if err := gestorRutinas.ModificarRutina("Rutina Mañana", modificada); err != nil {
        t.Errorf("ModificarRutina() error = %v", err)
    }
    resultado, err := gestorRutinas.ConsultarRutina("Rutina Mañana")
    if err != nil {
        t.Errorf("ConsultarRutina() error = %v", err)
    }
    if resultado.Ejercicios != modificada.Ejercicios {
        t.Errorf("Esperado Ejercicios = %s, Obtenido Ejercicios = %s", modificada.Ejercicios, resultado.Ejercicios)
    }
}

// TestListarRutinas verifica que la función ListarRutinas filtre las rutinas por dificultad correctamente.
func TestListarRutinas(t *testing.T) {
    gestorRutinas, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    gestorRutinas.AgregarRutina(&Rutina{
        Nombre:     "Rutina Ligera",
        Dificultad: "Baja",
        Ejercicios: ejercicios[0].Nombre,
    })
    gestorRutinas.AgregarRutina(&Rutina{
        Nombre:     "Rutina Intensa",
        Dificultad: "Alta",
        Ejercicios: ejercicios[1].Nombre,
    })
    resultados := gestorRutinas.ListarRutinas("Alta")
    if len(resultados) != 1 || resultados[0].Nombre != "Rutina Intensa" {
        t.Errorf("ListarRutinas() failed, expected 1 result named 'Rutina Intensa', got %d results", len(resultados))
    }
}