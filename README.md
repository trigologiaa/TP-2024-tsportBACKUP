# Rutinas de Entrenamiento

## Consigna General
Diseñar una aplicación para la gestión de rutinas de entrenamiento que ofrezca una experiencia completa y personalizada para usuarios de todos los niveles de condición física. La aplicación permitirá a los usuarios elegir entre seguir una rutina predefinida o crear una nueva de manera dinámica. Para ello, la aplicación utilizará algoritmos de búsqueda avanzados y estructuras de datos eficientes para almacenar y organizar los ejercicios, tiempos y tipos de entrenamiento.

Además, se implementará un sistema de etiquetado mediante mapas para categorizar los ejercicios según su tipo. 

Una característica destacada será la capacidad de la aplicación para armar automáticamente una rutina nueva, teniendo en cuenta una entrada en calor, ejercicios específicos según los parámetros establecidos por el usuario y un estiramiento final, todo ello dentro de los límites de tiempo definidos.

Para optimizar la selección de ejercicios, se empleará el concepto de "máxima cantidad" utilizando técnicas de programación dinámica, garantizando así una rutina efectiva y adaptada a las necesidades individuales de cada usuario, maximizando la cantidad (u otro aspecto) de ejercicios a realizar en un período dado.

La aplicación que se realizará tendrá una interfaz de línea de comandos, mediante la cual se podrá interactuar con el sistema.

## Definiciones
Un ejercicio tendrá:
- nombre (por ejemplo, "sentadillas")
- descripción (por ejemplo, "flexión de rodillas hasta formar un ángulo de 90 grados")
- tiempo estimado (en segundos)
- calorias quemadas (por ejemplo, 100)
- tipo de ejercicio (por ejemplo, cardio, fuerza, flexibilidad, etc). Puede tener más de una categoría
- puntos por tipo de ejercicio (puntos cardio, puntos fuerza, puntos flexibilidad, etc). Puede tener puntaje en más de una categoría
- nivel de dificultad (por ejemplo, principiante, intermedio, avanzado)

### ABMCL de Ejercicios
Se deberá implementar un ABMCL (Alta, Baja, Modificación, Consulta y Listado) de ejercicios, permitiendo al usuario agregar, eliminar, modificar y consultar ejercicios. Además, se deberá permitir listar todos los ejercicios disponibles.
El listado de ejercicios permitirá filtrar por tipo de ejercicio y nivel de dificultad. Asimismo, se podrá buscar por nombre de ejercicio (total o parcial), o por cantidad mínima de calorías quemadas.

Debe considerarse que los ejercicios tendrán etiquetas para definir su tipo de ejercicio y dificultad. Las etiquetas serán un conjunto de palabras clave que permitirán clasificar los ejercicios. Por ejemplo, un ejercicio de sentadillas podría tener las etiquetas "fuerza" y "piernas". Sin embargo, las dificultades son únicas; por ejemplo, las sentadillas serán de dificultad "media" sólamente.

### ABMCL de Rutinas
Se deberá implementar un ABMCL de rutinas, permitiendo al usuario agregar, eliminar, modificar y consultar rutinas. Además, se deberá permitir listar todas las rutinas disponibles.

Para armar una rutina, se debe colocar el nombre de la rutina. La duración, tipo de ejercicios, calorías quemadas totales y dificultad deben calcularse automáticamente. La duración de la rutina será la suma de los tiempos estimados de los ejercicios que la componen. El tipo de ejercicios y la dificultad de la rutina serán los más frecuentes entre los ejercicios que la componen. Las calorías quemadas totales serán la suma de las calorías quemadas de los ejercicios que la componen. Lo mismo sucederá con los puntos en cada una de las dimensiones (cardio, fuerza, etc)

### Generación Automágica de Rutinas, parte 1
Se deberá implementar un algoritmo que permita generar automáticamente una rutina nueva, teniendo en cuenta los siguientes parámetros:
- Nombre de la rutina
- Duración total de la rutina (en minutos)
- Tipo de ejercicios a incluir (por ejemplo, cardio, fuerza, flexibilidad, etc.)
- Nivel de dificultad de los ejercicios a incluir (por ejemplo, principiante, intermedio, avanzado)

El algoritmo deberá seleccionar los ejercicios que cumplan con los parámetros establecidos y que maximicen la cantidad de ejercicios a realizar en el tiempo definido. Para ello, se empleará el concepto de "máxima cantidad". No pueden repetirse ejercicios en la rutina.

### Generación Automágica de Rutinas, parte 2
Se deberá implementar un algoritmo que permita generar automáticamente una rutina nueva, teniendo en cuenta los siguientes parámetros:
- Nombre de la rutina
- Calorías totales a quemar

El algoritmo deberá seleccionar los ejercicios que cumplan con los parámetros establecidos y que minimicen la duración de la rutina. Para ello, se empleará el concepto de "mínima duración". No pueden repetirse ejercicios en la rutina.

### Generación Automágica de Rutinas, parte 3
Se deberá implementar un algoritmo que permita generar automáticamente una rutina nueva, teniendo en cuenta los siguientes parámetros:
- Nombre de la rutina
- Tipo de puntos a maximizar (cardio, fuerza, flexibilidad)
- Duración máxima de la rutina

El algoritmo deberá seleccionar los ejercicios que cumplan con los parámetros establecidos y que maximicen el puntaje de la rutina en la dimensión solicitada. Para ello, se empleará el concepto de "duración fija, máximos puntos". No pueden repetirse ejercicios en la rutina.

### Gestión de Datos
Se deberá implementar un sistema de almacenamiento de datos que permita guardar y recuperar los ejercicios y rutinas creados por los usuarios. Se deberá utilizar un archivo de texto para almacenar los datos de manera persistente.
Se sugiere una estructura de archios tipo CSV para el almacenamiento de los datos. Deberán definir la estructura de los archivos y los campos que contendrán, adecuados para guardar tanto los ejercicios como las rutinas.

**Tip:** [https://pkg.go.dev/github.com/gocarina/gocsv](https://pkg.go.dev/github.com/gocarina/gocsv)

## Formato de Entrega
Se deberá realizar el proyecto en este mismo repositorio. La entrega consistirá en:
1. Código fuente completo del proyecto.
2. Pruebas sobre las funcionalidades presentadas.
3. Archivos de datos con ejemplos de ejercicios y rutinas.
4. Informe en formato markdown (INFORME.md) que incluya:
    - Conclusiones sobre la implementación del proyecto
    - Referencias

### Primera Entrega
Para la primera entrega se hará foco en el desarrollo de las estructuras de datos y mecanismos para la creación y consulta/búsqueda de ejercicios y rutinas.
Deberá probarse cada aspecto no trivial desarrollado (validaciones, restricciones de dominio, formato de salidas por pantalla).

### Segunda Entrega
La segunda entrega incorporará los algoritmos de generación automágica de rutinas, y el menú interactivo.
Asimismo se deberán mejorar los algoritmos de búsqueda, bajo la luz de los conocimientos nuevos adquiridos en la segunda mitad del cuatrimestre.
