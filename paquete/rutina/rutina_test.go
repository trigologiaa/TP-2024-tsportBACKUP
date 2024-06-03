package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"testing"
	"github.com/stretchr/testify/assert"
)

//  setup inicializa y configura los gestores de ejercicios y rutinas para las pruebas. Agrega ejercicios iniciales al gestor de ejercicios y retorna ambos gestores.
//
//  Retorna:
//      - Un puntero a GestorRutinas por ser una nueva instancia de este.
//      - Un puntero a GestorEjercicios, creado y configurado con los ejercicios.
//  Funcionamiento:
//      Se declara la variable 'gestorEjercicios' que es el llamado a la función 'NuevoGestorEjercicios' que crea una nueva instancia de 'GestorEjercicios'
//      Se declara la variable 'ejercicio' que es una nueva instancia de 'Ejercicio' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Tiempo
//          Se inicializa el campo Calorias
//          Se inicializa el campo Dificultad
//      }
//      Se declara la variable 'ejercicio' que es una nueva instancia de 'Ejercicio' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Tiempo
//          Se inicializa el campo Calorias
//          Se inicializa el campo Dificultad
//      }
//      Se declara la variable 'gestorRutinas' que es el llamado a la función 'NuevoGestorRutinas', que crea una nueva instancia de 'GestorRutinas' con el gestor de ejercicios previamente creado
//      Se retornan 'gestorRutinas' y 'gestorEjercicios'
func setup() (*GestorRutinas, *ejercicio.GestorEjercicios) {
    gestorEjercicios := ejercicio.NuevoGestorEjercicios()
    gestorEjercicios.AgregarEjercicio(&ejercicio.Ejercicio {
        Nombre: "Sentadillas", 
        Tiempo: 60, 
        Calorias: 100, 
        Dificultad: "Media",
    })
    gestorEjercicios.AgregarEjercicio(&ejercicio.Ejercicio {
        Nombre: "Flexiones", 
        Tiempo: 30, 
        Calorias: 50, 
        Dificultad: "Alta",
    })
    gestorRutinas := NuevoGestorRutinas(gestorEjercicios)
    return gestorRutinas, gestorEjercicios
}

// TestAgregarRutina verifica la capacidad del gestor para agregar rutinas y manejar duplicados.
//
// Funcionamiento:
//      Se declaran las variables 'gestorRutinas' y 'gestorEjercicios' llamando a la función 'setup' que inicializa ambos gestores
//      Se declara la variable ejercicios que es el llamado a la función 'ObtenerTodosLosEjercicios' de 'gestorEjercicios', dando una lista de todos los ejercicios
//      Se declara la variable 'rutina' que es una nueva instancia de 'Rutina' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Ejercicios
//      }
//      Si se produce un problema al agregar la rutina {
//          Se genera un error
//      }
//      Si se produce un problema al agregar la rutina por estar duplicada {
//          Se genera un error
//      }
func TestAgregarRutina(t *testing.T) {
    gestorRutinas, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    rutina := &Rutina {
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
//
// Funcionamiento:
//      Se declara la variable 'gestorRutinas' llamando a la función 'setup' que inicializa el gestor de rutinas
//      Se declara la variable 'rutina' que es una nueva instancia de 'Rutina' {
//          Se inicializa el campo Nombre
//      }
//      Se agrega la rutina a 'gestorRutinas' llamando al método 'AgregarRutina'
//      Si se produce un problema al eliminar la rutina {
//          Se genera un error
//      }
//      Si se produce un problema al consultar una rutina eliminada {
//          Se genera un error
//      }
func TestEliminarRutina(t *testing.T) {
    gestorRutinas, _ := setup()
    rutina := &Rutina {
        Nombre: "Rutina Temporal",
    }
    gestorRutinas.AgregarRutina(rutina)
    if err := gestorRutinas.EliminarRutina("Rutina Temporal"); err != nil {
        t.Errorf("EliminarRutina() error = %v, wantErr %v", err, false)
    }
    if _, err := gestorRutinas.ConsultarRutina("Rutina Temporal"); err == nil {
        t.Error("Se esperaba un error al buscar una rutina eliminada, pero no se obtuvo error")
    }
}

// TestModificarRutina verifica que se puedan modificar las rutinas existentes correctamente.
//
// Funcionamiento:
//      Se declaran las variables 'gestorRutinas' y 'gestorEjercicios' llamando a la función 'setup' que inicializa ambos gestores
//      Se declara la variable ejercicios que es el llamado a la función 'ObtenerTodosLosEjercicios' de 'gestorEjercicios', dando una lista de todos los ejercicios
//      Se declara la variable 'original' que es una nueva instancia de 'Rutina' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Ejercicios
//      }
//      Se agrega la rutina original a 'gestorRutinas' llamando al método 'AgregarRutina'
//      Se declara la variable 'modificada' que es una nueva instancia de 'Rutina' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Ejercicios
//      }
//      Si se produce un problema al modificar la rutina {
//          Se genera un error
//      }
//      Se consulta la rutina modificada llamando al método 'ConsultarRutina'
//      Si se produce un problema al consultar la rutina modificada {
//          Se genera un error
//      }
//      Si los ejercicios de la rutina modificada no coinciden con los esperados {
//          Se genera un error
//      }
func TestModificarRutina(t *testing.T) {
    gestorRutinas, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    original := &Rutina {
        Nombre:     "Rutina Mañana",
        Ejercicios: ejercicios[0].Nombre,
    }
    gestorRutinas.AgregarRutina(original)
    modificada := &Rutina {
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

//  TestConsultarRutina verifica la capacidad del gestor para consultar rutinas por su nombre.
//
//  Funcionamiento:
//      Se declaran las variables 'gestorRutinas' y 'gestorEjercicios' llamando a la función 'setup' que inicializa ambos gestores
//      Se declara la variable ejercicios que es el llamado a la función 'ObtenerTodosLosEjercicios' de 'gestorEjercicios', dando una lista de todos los ejercicios
//      Se declara la variable 'rutina' que es una nueva instancia de 'Rutina' {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Ejercicios
//      }
//      Si se produce un problema al agregar la rutina {
//          Se genera un error
//      }
//      Se consulta la rutina agregada llamando al método 'ConsultarRutina'
//      Si se produce un problema al consultar la rutina {
//          Se genera un error
//      }
//      Si el nombre de la rutina consultada no coincide con el esperado {
//          Se genera un error
//      }
func TestConsultarRutina(t *testing.T) {
    gestorRutinas, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    rutina := &Rutina {
        Nombre: "Rutina Consultada",
        Ejercicios: ejercicios[0].Nombre + ", " + ejercicios[1].Nombre,
    }
    if err := gestorRutinas.AgregarRutina(rutina); err != nil {
        t.Errorf("AgregarRutina() error = %v, wantErr %v", err, false)
    }
    resultado, err := gestorRutinas.ConsultarRutina("Rutina Consultada")
    if err != nil {
        t.Errorf("ConsultarRutina() error = %v", err)
    }
    if resultado.Nombre != rutina.Nombre {
        t.Errorf("Esperado Nombre = %s, Obtenido Nombre = %s", rutina.Nombre, resultado.Nombre)
    }
}

// TestListarRutinas verifica que la función ListarRutinas filtre las rutinas por dificultad correctamente.
//
// Funcionamiento:
//      Se declaran las variables 'gestorRutinas' y 'gestorEjercicios' llamando a la función 'setup' que inicializa ambos gestores
//      Se declara la variable ejercicios que es el llamado a la función 'ObtenerTodosLosEjercicios' de 'gestorEjercicios', dando una lista de todos los ejercicios
//      Se llama al método 'AgregarRutina' de 'gestorRutinas' pasándole como parámetro una nueva instancia de 'Rutina' con dificultad baja {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Dificultad
//          Se inicializa el campo Ejercicios
//      }
//      Se llama al método 'AgregarRutina' de 'gestorRutinas' pasándole como parámetro una nueva instancia de 'Rutina' con dificultad alta {
//          Se inicializa el campo Nombre
//          Se inicializa el campo Dificultad
//          Se inicializa el campo Ejercicios
//      }
//      Se declara la variable resultados que es el llamado al método 'ListarRutinas' pasándole como parámetro de dificultad 'Alta'
//      Si la cantidad de rutinas obtenidas no es la esperada o el nombre de la rutina no es el esperado {
//          Se genera un error
//      }
func TestListarRutinas(t *testing.T) {
    gestorRutinas, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    gestorRutinas.AgregarRutina(&Rutina {
        Nombre:     "Rutina Ligera",
        Dificultad: "Baja",
        Ejercicios: ejercicios[0].Nombre,
    })
    gestorRutinas.AgregarRutina(&Rutina {
        Nombre:     "Rutina Intensa",
        Dificultad: "Alta",
        Ejercicios: ejercicios[1].Nombre,
    })
    resultados := gestorRutinas.ListarRutinas("Alta")
    if len(resultados) != 1 || resultados[0].Nombre != "Rutina Intensa" {
        t.Errorf("ListarRutinas() failed, expected 1 result named 'Rutina Intensa', got %d results", len(resultados))
    }
}

//  TestCalcularPropiedades verifica que se calculen y actualicen correctamente las propiedades de una rutina.
//
//  Funcionamiento:
//      Se declara la variable 'gestorEjercicios' llamando a la función 'setup', ignorando el primer retorno
//      Se declara la variable ejercicios que es el llamado a la función 'ObtenerTodosLosEjercicios' de 'gestorEjercicios', dando una lista de todos los ejercicios
//      Se declara la variable 'rutina' que es una nueva instancia de 'Rutina' {
//          Se inicializa el campo Nombre
//      }
//      Se llama al método 'CalcularPropiedades' de 'rutina' pasando 'gestorEjercicios' como argumento
//      Si el tiempo total de la rutina no coincide con el esperado {
//          Se genera un error
//      }
//      Si las calorías totales de la rutina no coinciden con las esperadas {
//          Se genera un error
//      }
//      Si la dificultad de la rutina no coincide con la esperada {
//          Se genera un error
//      }
func TestCalcularPropiedades(t *testing.T) {
    _, gestorEjercicios := setup()
    ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
    rutina := &Rutina {
        Nombre: "Rutina Propiedades",
    }
    rutina.CalcularPropiedades(gestorEjercicios)
    assert.Equal(t, 90, rutina.Tiempo, "El tiempo total de la rutina no es el esperado")
    assert.Equal(t, 150, rutina.Calorias, "Las calorías totales de la rutina no son las esperadas")
    assert.Equal(t, "Media", rutina.Dificultad, "La dificultad de la rutina no es la esperada")
    assert.Equal(t, ejercicios[0].Nombre + ", " + ejercicios[1].Nombre, rutina.Ejercicios, "Los ejercicios de la rutina no son los esperados")
}