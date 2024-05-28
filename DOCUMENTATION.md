<h1 align="center">ARCHIVOS</h1>

# ejercicio.go

## type Ejercicio struct
Es una estructura que representa un ejercicio. Contiene varios campos con información relevante sobre el ejercicio:
- Nombre: El nombre del ejercicio.
- Descripcion: Una descripción breve del ejercicio.
- Tiempo: La duración del ejercicio en segundos.
- Calorias: Cantidad de calorías que se estima se queman realizando el ejercicio.
- Tipos: Tipo de ejercicio, por ejemplo, cardiovascular, fuerza, etc.
- Puntos: Puntos ganados por realizar el ejercicio, que podrían estar relacionados con algún sistema de recompensas o puntos de experiencia.
- Dificultad: Nivel de dificultad del ejercicio.

Cada campo tiene asociado un tag csv que indica cómo debería ser etiquetado al escribir o leer este tipo de estructura en un formato CSV.

## type GestorEjercicios struct
Es una estructura que actúa como un gestor para un conjunto de ejercicios. Usa una lista doblemente enlazada para almacenar y gestionar objetos del tipo Ejercicio:
- ejercicios: Un puntero a una DoubleLinkedList que almacena objetos de tipo Ejercicio.

## func NuevoGestorEjercicios() *GestorEjercicios
Esta función crea y retorna una nueva instancia de GestorEjercicios. Inicializa ejercicios con una nueva lista doblemente enlazada vacía de tipo Ejercicio.

## func (g *GestorEjercicios) AgregarEjercicio(ejercicio *Ejercicio) error
Este método intenta agregar un nuevo ejercicio a la lista doblemente enlazada. Primero verifica si el ejercicio ya existe en la lista (usando Find). Si el ejercicio ya existe, retorna un error. Si no, lo añade al final de la lista (usando Append) y retorna nil para indicar que no hubo errores.

## func (g *GestorEjercicios) ObtenerTodosLosEjercicios() []*Ejercicio
Este método retorna un slice de todos los ejercicios almacenados en la lista doblemente enlazada. Recorre todos los nodos de la lista desde el cabeza hasta el final, agregando cada ejercicio a un slice que luego retorna.

## func (g *GestorEjercicios) EliminarEjercicio(nombre string) error
Este método busca y elimina un ejercicio por su nombre. Recorre la lista y si encuentra un ejercicio con el nombre especificado, lo elimina de la lista (usando Remove) y retorna nil. Si no encuentra el ejercicio, retorna un error.

## func (g *GestorEjercicios) ConsultarEjercicio(nombre string) (*Ejercicio, error)
Este método busca un ejercicio por su nombre. Similar a EliminarEjercicio, recorre la lista buscando un ejercicio que coincida con el nombre proporcionado. Si lo encuentra, retorna el ejercicio y nil. Si no lo encuentra, retorna nil y un error.

# ejercicio_test.go

## func TestAgregarEjercicio(t *testing.T)
Esta función prueba el método AgregarEjercicio del gestor de ejercicios. Realiza las siguientes comprobaciones:
- Creación de un gestor de ejercicios: Inicializa un nuevo gestor de ejercicios llamando a NuevoGestorEjercicios().
- Agregar un ejercicio nuevo: Crea un nuevo ejercicio llamado "Sentadillas" y trata de agregarlo al gestor. Si falla (retornando un error), la prueba falla y se registra un mensaje de error.
- Agregar un ejercicio duplicado: Intenta agregar el mismo ejercicio una segunda vez. Si no falla (es decir, si no retorna un error), la prueba falla porque se esperaba un error al intentar agregar un duplicado.

## func TestEliminarEjercicio(t *testing.T)
Esta función prueba el método EliminarEjercicio del gestor de ejercicios:
- Creación de un gestor y agregar un ejercicio: Inicializa un nuevo gestor de ejercicios y agrega un ejercicio llamado "Burpees".
- Eliminar el ejercicio: Intenta eliminar el ejercicio "Burpees". Si hay un error al eliminarlo, la prueba falla y registra un mensaje de error.
- Consultar el ejercicio eliminado: Intenta consultar el ejercicio después de haber sido eliminado. Si no falla (es decir, si retorna un ejercicio o no da un error), la prueba falla porque se esperaba que el ejercicio ya no estuviera disponible.

## func TestConsultarEjercicio(t *testing.T)
Esta función prueba el método ConsultarEjercicio del gestor de ejercicios:
- Creación de un gestor y agregar un ejercicio: Inicializa un nuevo gestor y agrega un ejercicio llamado "Plancha".
- Consultar el ejercicio: Intenta recuperar el ejercicio "Plancha" del gestor. Si falla al recuperar el ejercicio (es decir, si retorna un error), la prueba falla y se registra un mensaje de error.
- Verificación del resultado: Comprueba que el nombre del ejercicio recuperado sea correcto. Si el nombre no coincide con "Plancha", la prueba falla indicando que se recibió un nombre incorrecto.

# rutina.go

## type Rutina struct
Esta estructura representa una rutina de ejercicios, incluyendo varios detalles sobre la misma:
- Nombre: El nombre de la rutina.
- Ejercicios: Una lista de los nombres de los ejercicios incluidos, almacenados como una cadena.
- Duracion: La duración total de la rutina en segundos.
- Calorias: Total de calorías estimadas que se queman realizando la rutina.
- Dificultad: Nivel de dificultad de la rutina (e.g., fácil, medio, difícil).
- Tipos: Tipos de ejercicios incluidos en la rutina, representados como una cadena.
- PuntosPorTipo: Puntos obtenidos por la realización de la rutina, posiblemente sumando puntos de todos los ejercicios.

## type GestorRutinas struct
Este tipo estructurado actúa como un gestor o administrador para un conjunto de rutinas:
- rutinas: Un puntero a una lista doblemente enlazada que almacena objetos del tipo Rutina.
- gestorEjercicios: Un puntero a GestorEjercicios, que es utilizado para calcular propiedades de las rutinas basadas en los ejercicios disponibles.

## func NuevoGestorRutinas(gestorEj *ejercicio.GestorEjercicios) *GestorRutinas
Constructor para GestorRutinas. Inicializa y retorna una nueva instancia de GestorRutinas con una lista doblemente enlazada vacía para almacenar rutinas y un enlace al gestor de ejercicios proporcionado.

## func (g *GestorRutinas) AgregarRutina(rutina *Rutina) error
Método para agregar una nueva rutina al gestor. Verifica si una rutina con el mismo nombre ya existe; si es así, retorna un error. Si no, calcula las propiedades de la rutina con CalcularPropiedades y la añade a la lista de rutinas.

## func (g *GestorRutinas) EliminarRutina(nombre string) error
Método para eliminar una rutina por su nombre. Busca la rutina en la lista y si la encuentra, la elimina y retorna nil. Si no la encuentra, retorna un error.

## func (g *GestorRutinas) ModificarRutina(nombre string, nuevaRutina *Rutina) error
Método para modificar una rutina existente. Busca la rutina por nombre y, si la encuentra, reemplaza sus datos con los de nuevaRutina después de calcular sus propiedades. Si no encuentra la rutina, retorna un error.

## func (g *GestorRutinas) ConsultarRutina(nombre string) (*Rutina, error)
Método para consultar una rutina por nombre. Si la encuentra, retorna la rutina y nil. Si no, retorna nil y un error.

## func (g *GestorRutinas) ListarRutinas(dificultad string) []*Rutina
Método para obtener una lista de todas las rutinas que coinciden con una dificultad específica. Si dificultad está vacío, retorna todas las rutinas.

## func (r *Rutina) CalcularPropiedades(gestor *ejercicio.GestorEjercicios)
Calcula y actualiza propiedades de la Rutina basándose en los ejercicios disponibles en el gestorEjercicios.

## func joinKeys(items []string, separator string) string
Función auxiliar para unir elementos de un slice en una cadena, separados por separator.

## func mapKeysToStringSlice(m map[string]bool) []string
Convierte las claves de un mapa a un slice de cadenas.

## func maxKey(m map[string]int) string
Determina la clave con el valor máximo en un mapa de enteros.

# rutina_test.go

## func setup() (*GestorRutinas, *ejercicio.GestorEjercicios)
Esta función inicializa el entorno para las pruebas. Retorna dos objetos: un gestor de rutinas y un gestor de ejercicios, ambos inicializados y listos para ser usados en las pruebas. El gestor de ejercicios es cargado con dos ejercicios predefinidos ("Sentadillas" y "Flexiones") para garantizar que hay datos con los cuales trabajar durante las pruebas.

## func TestAgregarRutina(t *testing.T)
Esta función prueba la capacidad de agregar rutinas al gestor de rutinas. Se verifica que:
- Se pueda agregar una rutina sin errores.
- No se pueda agregar la misma rutina dos veces (verifica la gestión de duplicados).

## func TestEliminarRutina(t *testing.T)
Prueba la funcionalidad de eliminar una rutina del gestor. Se verifica que:
- Una rutina se pueda eliminar correctamente.
- Una vez eliminada, intentar consultarla debería resultar en un error, confirmando que fue eliminada.

## func TestModificarRutina(t *testing.T)
Verifica la funcionalidad de modificar una rutina existente. Se asegura que:
- La rutina se pueda modificar sin problemas.
- Después de la modificación, los cambios en los detalles de la rutina (en este caso, los ejercicios asociados) se reflejen correctamente al consultar la rutina.

## func TestListarRutinas(t *testing.T)
Prueba la capacidad del gestor de rutinas para filtrar y listar rutinas basadas en su dificultad. Se verifica que:
- Solo las rutinas que coinciden con el criterio de búsqueda (en este caso, la dificultad "Alta") sean retornadas.
- La cantidad de rutinas y los nombres de las rutinas retornadas sean correctos según los datos esperados.