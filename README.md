<h1 align="center">CONCLUSIÓN</h1>

<p align="center">
  <img src="Extras\duck.gif" alt="Texto alternativo">
</p>

<!-- <p align="center">
  <img src="https://i.kym-cdn.com/photos/images/newsfeed/002/805/512/39d.jpg" alt="Texto alternativo">
</p>*/ -->

<h3 align="center">Resumen</h2>

El proyecto consiste en la implementación de un sistema de gestión de ejercicios y rutinas que permite a los usuarios el manejo de información relacionado con distintos tipos de ejercicios y así compilar estos mismos en rutinas personalizadas. Usa estructuras de datos dinámicas como listas doblemente enlazadas para así poder gestionar los datos de manera más eficiente. Este programa implementándose en Go, utilizando testeos que verifican la funcionalidad del sistema tanto para los métodos de los ejercicios como para los de las rutinas.

<h4>Conclusiones de la implementación</h4>

- Estructuras de datos: El uso de las listas doblemente enlazadas permite manipulaciones eficientes de los datos, como inserciones y eliminaciones, que son frecuentes ne la gestión de los ejercicios y rutinas.
- Modularidad y escalabilidad: El diseño modular del sistema, que tiene gestores separados para los ejercicios y las rutinas, facilita así la escalabilidad y el mantenimiento del sistema. Esta separación asegura que las modificaciones en una parte del sistema afecten mínimamente sobre otras partes.
- Pruebas: La implementación de los testeos para cada funcionalidad principal asegura la robustez y confiabilidad en el sistema, así como también los casos de borde, como gestión de duplicados y validación de entradas.
- Integración y extensibilidad: El código tiene una estructura que permite integrarlo fácilmente con otros sistemas y así también extensibilidad para añadir más características, como otros ejercicios.

<h3 align="center">Explicación del Código Fuente</h2>

El código fuente se divide en 4 carpetas principales:

<h4>- /extras</h4>

Contiene 1 archivo completamente esencial para el funcionamiento del proyecto.

<h4>- /informacion</h4>

Contiene archivos con la información de los ejercicios y rutinas en formato CSV tanto para guardar la información sobre estos archivos o para cargarla sobre los mismos.

  - `ejercicios.csv` - Archivo para almacenar los datos de ejercicios.
  - `rutinas.csv` - Archivo para almacenar los datos de rutinas.

<h4>- /main</h4>

Contiene el archivo principal que manejará todas las funciones y permitirá la utilización por parte del usuario en forma de menú interactivo.

  - `main.go` - Punto de entrada principal de la aplicación con la interfaz de línea de comandos.

<h4>- /paquete</h4>

Contiene 3 sub carpetas que manejarán 3 aspectos importantes, los cuales serán el almacenamiento, los ejercicios y las rutinas.

- <h4>- /almacenamiento</h4>
    Contiene el archivo con el manejo de archivos CSV.

    - `almacenamiento.go` - Maneja la persistencia de datos, como cargar y guardar en archivos CSV.
- <h4>- /ejercicio</h4>
    Contiene los archivos encargados de los ejercicios.

    - `ejercicio.go` - Define las estructuras y métodos ABMCL para la gestión de ejercicios.
    - `ejercicio_test.go` - Contiene pruebas unitarias para el módulo de ejercicios.
- <h4>- /rutina</h4>
    Contiene los archivos encargados de las rutinas.

    - `rutina.go` - Define las estructuras y métodos ABMCL para la gestión de rutinas.
    - `rutina_test.go` - Contiene pruebas unitarias para el módulo de rutinas.

<h1 align="center">REFERENCIAS</h1>

## DOCUMENTACIÓN UTILIZADA
- [bufio](https://pkg.go.dev/bufio)
- [fmt](https://pkg.go.dev/fmt)
- [os](https://pkg.go.dev/os)
- [strconv](https://pkg.go.dev/strconv)
- [strings](https://pkg.go.dev/strings)
- [gocsv](https://pkg.go.dev/github.com/gocarina/gocsv)
- [errors](https://pkg.go.dev/errors)
- [sort](https://pkg.go.dev/sort)
- [list](https://github.com/untref-ayp2/data-structures/tree/main/list)
- [testing](https://pkg.go.dev/testing)
- [assert](https://pkg.go.dev/github.com/stretchr/testify/assert)

<h1 align="center">CONTENIDO UTILIZADO PARA LA PROGRAMACIÓN DEL CÓDIGO</h1>

- [what is ligma](https://www.youtube.com/watch?v=R6e1_IlvmQs)
- [meow](https://www.youtube.com/watch?v=cwyTleTL06Y)
- [spinetta jede - las wachas son del viento](https://www.youtube.com/watch?v=jZB6io7G-no)
- [never gonna give you up](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
- [hace rato me picó un mosquito](https://www.youtube.com/watch?v=wwBW8_yRJLM)
- [can't start](https://www.youtube.com/watch?v=_Nwn9ybsCRk)
- [música para estudiar como un filósofo del siglo XVIII y alcanzar la verdad absoluta](https://www.youtube.com/watch?v=1_0OOWfZZxw)
- [métodos piqueteros](https://www.youtube.com/watch?v=kbXQaP3Xoik)
- [punteros](https://youtu.be/qclZUQYZTzg)
- [java](https://youtu.be/XQtilPmhgUs)
- [el chavo del 8 RKT](https://youtu.be/sgyqafoCVcA)
- [el profesor freezer](https://youtu.be/MkW-T-AN70Q)
- [hello world in COW language](https://youtu.be/fLBMVl1zL8Q)
- [lunfardo](https://github.com/WhiteHeadbanger/Lunfardo)
- [paint](https://seoi.net/penint/)