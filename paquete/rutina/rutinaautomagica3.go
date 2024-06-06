package rutina
//no está completo
import (
    "errors"
    "math/rand"
    "time"
    "TP-2024-TSPORT/paquete/ejercicio"
)

// GenerarRutinaAutomagica3 genera automáticamente una rutina nueva que cumple con los parámetros especificados.
//
// Parámetros:
//   - nombre: Nombre de la rutina.
//   - tipoPuntos: Tipo de puntos a maximizar (cardio, fuerza, flexibilidad).
//   - duracionMaxima: Duración máxima de la rutina en minutos.
//   - gestor: GestorEjercicios que proporciona acceso a los ejercicios disponibles.
//
// Retorna:
//   - Un puntero a Rutina que representa la rutina generada.
//   - Un error en caso de que no se puedan encontrar ejercicios que cumplan con los requisitos.
func (g *GestorRutinas) GenerarRutinaAutomagica3(nombre string, tipoPuntos string, duracionMaxima int) (*Rutina, error) {
    ejercicios := g.gestorEjercicios.ObtenerTodosLosEjercicios()
    ejerciciosFiltrados := filtrarEjerciciosPorTipoPuntosRutinaAutomagica3(ejercicios, tipoPuntos)
    if len(ejerciciosFiltrados) == 0 {
        return nil, errors.New("no hay ejercicios disponibles para el tipo de puntos especificado")
    }
    //ejerciciosSeleccionados := seleccionarEjerciciosAleatoriosRutinaAutomagica3(ejerciciosFiltrados, duracionMaxima)
    rutina := &Rutina{
        
    }
    return rutina, nil
}

// Filtra los ejercicios por el tipo de puntos a maximizar.
func filtrarEjerciciosPorTipoPuntosRutinaAutomagica3(ejercicios []*ejercicio.Ejercicio, tipoPuntos string) []*ejercicio.Ejercicio {
    var ejerciciosFiltrados []*ejercicio.Ejercicio
    for _, ejercicio := range ejercicios {
        if ejercicio.Tipo == tipoPuntos {
            ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
        }
    }
    return ejerciciosFiltrados
}

// Selecciona ejercicios aleatorios hasta alcanzar la duración máxima de la rutina.
func seleccionarEjerciciosAleatoriosRutinaAutomagica3(ejercicios []*ejercicio.Ejercicio, duracionMaxima int) []*ejercicio.Ejercicio {
    rand.Seed(time.Now().UnixNano())
    var ejerciciosSeleccionados []*ejercicio.Ejercicio
    duracionTotal := 0
    for duracionTotal < duracionMaxima && len(ejercicios) > 0 {
        indice := rand.Intn(len(ejercicios))
        ejercicio := ejercicios[indice]
        if duracionTotal+ejercicio.Tiempo <= duracionMaxima {
            ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejercicio)
            duracionTotal += ejercicio.Tiempo
        }
        // Eliminar el ejercicio seleccionado para evitar repeticiones
        ejercicios = append(ejercicios[:indice], ejercicios[indice+1:]...)
    }
    return ejerciciosSeleccionados
}

// Calcula la duración total de los ejercicios.
func calcularDuracionTotalRutinaAutomagica3(ejercicios []*ejercicio.Ejercicio) int {
    var duracionTotal int
    for _, ejercicio := range ejercicios {
        duracionTotal += ejercicio.Tiempo
    }
    return duracionTotal
}

// Calcula los puntos totales de la rutina para un tipo de puntos específico.
func calcularPuntosTotalesRutinaAutomagica3(ejercicios []*ejercicio.Ejercicio, tipoPuntos string) int {
    var puntosTotales int
    for _, ejercicio := range ejercicios {
        if ejercicio.Tipo == tipoPuntos {
            puntosTotales += ejercicio.Puntos
        }
    }
    return puntosTotales
}