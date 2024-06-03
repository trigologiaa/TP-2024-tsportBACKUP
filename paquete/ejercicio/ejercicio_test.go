package ejercicio

import (
    "testing"
)

//  TestAgregarEjercicio verifica la funcionalidad de AgregarEjercicio. Comprueba que se puede agregar un ejercicio correctamente y que no se pueden agregar duplicados.
//
//  Funcionamiento:
//      Se declara la variable 'gestor' que es el llamado a la función 'NuevoGestorEjercicios' que crea una nueva instancia de 'GestorEjercicios'
//      Se declara la variable 'ejercicio' que es una nueva instancia de 'Ejercicio' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Tiempo
//          Se inicializa el campo Calorias
//      }
//      Si se produce un problema al agregar el ejercicio {
//          Se genera un error
//      }
//      Si se produce un problema al agregar el ejercicio por estar duplicado {
//          Se genera un error
//      }
func TestAgregarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio {
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

//  TestObtenerTodosLosEjercicios verifica la funcionalidad de ObtenerTodosLosEjercicios. 
//  Comprueba que se pueden obtener todos los ejercicios agregados correctamente.
//
//  Funcionamiento:
//      Se declara la variable 'gestor' que es el llamado a la función 'NuevoGestorEjercicios' que crea una nueva instancia de 'GestorEjercicios'
//      Se declara la variable 'ejercicio1' que es una nueva instancia de 'Ejercicio' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Tiempo
//          Se inicializa el campo Calorias
//          Se inicializa el campo Tipo
//          Se inicializa el campo Puntos
//          Se inicializa el campo Dificultad
//      }
//      Se declara la variable 'ejercicio2' que es una nueva instancia de 'Ejercicio' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Tiempo
//          Se inicializa el campo Calorias
//          Se inicializa el campo Tipo
//          Se inicializa el campo Puntos
//          Se inicializa el campo Dificultad
//      }
//      Se agrega 'ejercicio1' a 'gestor' llamando al método 'AgregarEjercicio'
//      Se agrega 'ejercicio2' a 'gestor' llamando al método 'AgregarEjercicio'
//      Se llama al método 'ObtenerTodosLosEjercicios' que devuelve un slice de punteros a Ejercicio, asignado a 'ejercicios'
//      Si la cantidad de ejercicios no es la esperada {
//          Se genera un error
//      }
//      Si el nombre del primer ejercicio no es el esperado {
//          Se genera un error
//      }
//      Si el nombre del segundo ejercicio no es el esperado {
//          Se genera un error
//      }
func TestObtenerTodosLosEjercicios(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio1 := &Ejercicio{
        Nombre: "Sentadillas",
        Tiempo: 60,
        Calorias: 100,
        Tipo: "Fuerza",
        Puntos: 10,
        Dificultad: "Media",
    }
    ejercicio2 := &Ejercicio{
        Nombre: "Flexiones",
        Tiempo: 45,
        Calorias: 80,
        Tipo: "Fuerza",
        Puntos: 8,
        Dificultad: "Alta",
    }
    gestor.AgregarEjercicio(ejercicio1)
    gestor.AgregarEjercicio(ejercicio2)
    ejercicios := gestor.ObtenerTodosLosEjercicios()
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

//  TestEliminarEjercicio verifica la funcionalidad de EliminarEjercicio. Comprueba que se puede eliminar un ejercicio correctamente y que un ejercicio eliminado no puede ser consultado.
//
//  Funcionamiento:
//      Se declara la variable 'gestor' que es el llamado a la función 'NuevoGestorEjercicios' que crea una nueva instancia de 'GestorEjercicios'
//      Se declara la variable 'ejercicio' que es una nueva instancia de 'Ejercicio' {
//          Se inicializa el campo Nombre
//      }
//      Se agrega el ejercicio a 'gestor' llamando al método 'AgregarEjercicio'
//      Si se produce un problema al eliminar el ejercicio {
//          Se genera un error
//      }
//      Si se produce un problema al consultar un ejercicio eliminado {
//          Se genera un error
//      }
func TestEliminarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio {
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

//  TestConsultarEjercicio verifica la funcionalidad de ConsultarEjercicio. Asegura que se puede consultar un ejercicio correctamente y que los datos devueltos son correctos.
//
//  Funcionamiento:
//      Se declara la variable 'gestor' que es el llamado a la función 'NuevoGestorEjercicios' que crea una nueva instancia de 'GestorEjercicios'
//      Se declara la variable 'ejercicio' que es una nueva instancia de 'Ejercicio' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Tiempo
//      }
//      Se agrega el ejercicio a 'gestor' llamando al método 'AgregarEjercicio'
//      Se llama al método 'ConsultarEjercicio' con el campo Nombre como parámetro, devolviendo un puntero a Ejercicio asignado a 'resultado' y un error asignado a 'err'
//      Si se produce un problema al consultar el ejercicio {
//          Se genera un error
//      }
//      Si se produce un problema al verificar el Nombre del ejercicio {
//          Se genera un error
//      }
func TestConsultarEjercicio(t *testing.T) {
    gestor := NuevoGestorEjercicios()
    ejercicio := &Ejercicio {
        Nombre: "Plancha", 
        Tiempo: 30,
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