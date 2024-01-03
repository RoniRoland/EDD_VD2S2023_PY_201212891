UNIVESIDAD DE SAN CARLOS DE GUATEMALA

FACULTAD DE INGENIERIA

ESCUELA DE CIENCAS Y SISTEMAS

ESTRUCTURAS DE DATOS

SECCIÓN A

VACACIONES SEGUNDO SEMESTRE 2023

AUX. CRISTIAN SUY

EDGAR ROLANDO RAMIREZ LOPEZ

201212891

Guatemala, Diciembre del 2023



# <p align="center"> **MANUAL TECNICO** </p>

# Introduccion

Manual Técnico de la Plataforma de Tutorías de la Escuela de Ciencias y Sistemas de la Universidad de San Carlos de Guatemala. Este manual proporciona una guía detallada sobre la implementación y la eficiencia de las estructuras de datos utilizadas en la aplicación, diseñada para mejorar la gestión de estudiantes tutores y sus tutorías.

La aplicación, desarrollada en el lenguaje de programación Go, integra diversas estructuras avanzadas, como Árboles B para la asignación de cursos, Tablas Hash para la administración de usuarios, Grafos para gestionar relaciones entre cursos, Árboles de Merkle para el control de libros certificados, y técnicas de codificación y encriptación para garantizar la seguridad de la información.

Este manual se presenta en dos enfoques: la creación de un cliente-servidor utilizando un framework con una interfaz gráfica en React, y la implementación de una interfaz estática con un servidor en Go, que utiliza HTML, CSS y Javascript para interactuar con la aplicación. Cada sección proporciona información detallada sobre la estructura correspondiente y cómo ha sido optimizada para mejorar la eficiencia y la experiencia del usuario.

# Estructura del Proyecto

## Organización de Carpetas

El proyecto sigue una estructura organizada con carpetas específicas para cada componente, dividido en las carpetas "backend" y "frontend" para gestionar las lógicas de servidor y cliente respectivamente.


En la carpeta "backend," se implementan las estructuras de datos esenciales que forman el núcleo de la aplicación. Cada subcarpeta cumple una función específica:

- ArbolB: Almacena la lógica para la asignación eficiente de cursos a tutores.
- TablaHash: Gestiona la información de usuarios, tanto tutores como estudiantes.
- Grafos: Controla las relaciones entre cursos, utilizando la estructura de grafos.
- MerkleTree: Administra la certificación de libros mediante la estructura de árbol de Merkle.
- CodificacionDecodificacion: Ofrece métodos para manipular archivos PDF de tutores.
- Encriptacion: Implementa funciones para garantizar la seguridad de las contraseñas de los usuarios.


La carpeta "frontend" contiene la interfaz gráfica de la aplicación, desarrollada con React y Vite. Aquí, cada componente se organiza de manera intuitiva.

## Árbol B

El paquete arbolB contiene la implementación del árbol B, una estructura de búsqueda y organización de datos eficiente. Este árbol se utiliza para gestionar la asignación de cursos a tutores, manteniendo un equilibrio adecuado y permitiendo búsquedas eficientes en el sistema. 

## Arbol Merkle

La carpeta arbolMerkle contiene la implementación de un árbol Merkle y sus nodos asociados, diseñado para verificar la integridad de los bloques de datos en una cadena de bloques. A continuación, se presenta un resumen de las funciones y estructuras implementadas:

## Grafos

Un grafo es una estructura matemática que representa relaciones entre un conjunto de elementos. Está compuesto por nodos (también llamados vértices) y arcos (también llamados bordes o enlaces), que conectan los nodos. Los grafos se utilizan para modelar situaciones donde hay conexiones o interacciones entre diferentes entidades.

La carpeta grafo contiene la implementación de un grafo dirigido mediante listas de adyacencia.

## Peticiones

En esta sección, se presentan funciones y estructuras de datos relacionadas con la gestión de peticiones en el sistema. Las peticiones son acciones o solicitudes que los usuarios pueden realizar, y el código proporcionado muestra cómo estas peticiones están estructuradas y gestionadas.

## Tabla Hash

La carpeta tablaHash contiene la implementación de una tabla hash en el backend del sistema. Aquí se almacenan y gestionan las personas, utilizando su carné como clave de búsqueda. 

# Descripcion del codigo

## Árbol B

A continuación, se describe brevemente cada componente del código:

- ### Función insertar_rama
    La función insertar_rama se encarga de realizar la inserción de un nodo en una rama del árbol B. La lógica garantiza que, al insertar un nuevo nodo, se verifique si es necesario realizar una división de la rama, manteniendo así el equilibrio del árbol.

- ### Función dividir
    La función dividir se ejecuta cuando una rama alcanza su capacidad máxima, dividiéndola y generando un nuevo nodo que servirá como punto de división.

- ### Función Insertar
    La función Insertar permite la inserción de un nuevo tutor en el árbol B. Si la raíz aún no está inicializada, se crea una nueva rama. En caso contrario, se invoca la función insertar_rama para llevar a cabo la inserción.

- ### Funciones de Graficación
    Las funciones Graficar, grafo, grafoRamas, y conexionRamas se encargan de generar un archivo DOT para visualizar el árbol B mediante Graphviz. Estas funciones facilitan la creación de representaciones visuales del árbol, proporcionando una herramienta útil para entender su estructura.

- ### Funciones de Búsqueda
    Las funciones Buscar y buscarArbol se utilizan para buscar un tutor específico en el árbol B. La búsqueda se realiza de manera recursiva, y los resultados se almacenan en una lista simple para su posterior manipulación.

- ### Funciones para Gestión de Libros y Publicaciones
    Las funciones GuardarLibro y GuardarPublicacion permiten la asociación de libros y publicaciones a un tutor específico en el árbol B, enriqueciendo la información almacenada en la estructura.

- ### Lista Enlazada
    En el siguiente archivo listaEnlazada.go, se encuentra la implementación de una lista simple, utilizada para almacenar resultados de búsquedas en el árbol B. La lista simple ofrece operaciones básicas como inserción y recorrido, contribuyendo a la manipulación eficiente de los datos recuperados durante las búsquedas.

- ### NodoB
    El archivo nodoB.go define la estructura NodoB, que representa un nodo del árbol B. Este nodo almacena información sobre un tutor y enlaces a otros nodos.

    -  #### Estructura NodoB
        - Tutor: Un puntero a la estructura Tutores, que contiene información sobre un tutor.
        - Siguiente: Un puntero al siguiente nodo en la rama del árbol B.
        - Anterior: Un puntero al nodo anterior en la rama del árbol B.
        - Izquierdo: Un puntero al nodo hijo izquierdo en la rama del árbol B.
        - Derecho: Un puntero al nodo hijo derecho en la rama del árbol B.

    - #### Función Insertar
        - Esta función se utiliza para insertar un nuevo nodo en una rama del árbol B.
        - Implementa la lógica necesaria para mantener el orden y equilibrio del árbol B.

- ### RamaB
    El archivo ramaB.go define la estructura RamaB, que representa una rama del árbol B.

    - #### Estructura RamaB
        - Primero: Un puntero al primer nodo en la rama.
        - Hoja: Un booleano que indica si la rama es una hoja.
        - Contador: Un entero que lleva el conteo de nodos en la rama.

   - #### Función Insertar
        - Esta función se encarga de insertar un nuevo nodo en la rama del árbol B.
        - Implementa la lógica para manejar la inserción en diferentes casos según el estado de la rama y el valor del nuevo nodo.

## Árbol Merkle

A continuación, se presenta un resumen de las funciones y estructuras implementadas:

### ArbolMerkle.go
- Estructura ArbolMerkle
    - RaizMerkle: Puntero a la raíz del árbol Merkle.
    - BloqueDeDatos: Puntero al primer bloque de datos enlazado.
    - CantidadBloques: Número de bloques de datos en la cadena.

- Función AgregarBloque
    - Agrega un nuevo bloque de datos a la cadena, indicando el estado, nombre del libro y carnet del tutor.

- Función GenerarArbol
    - Genera el árbol Merkle a partir de los bloques de datos existentes.

- Función generarHash
    - Crea una lista de nodos hoja del árbol Merkle y luego construye el árbol recursivamente.

- Función encriptarSha3
    - Aplica la función hash SHA-3 a una cadena y devuelve la representación hexadecimal del resultado.

- Función Graficar
    - Genera un archivo DOT para representar gráficamente el árbol Merkle y ejecuta el comando para generar la imagen.

- Función retornarValoresArbol
    - Retorna una cadena que representa las conexiones entre nodos del árbol Merkle para su visualización gráfica.

### NodoMerkle.go
- Estructura NodoMerkle
    - Izquierda: Puntero al hijo izquierdo del nodo.
    - Derecha: Puntero al hijo derecho del nodo.
    - Bloque: Puntero al bloque de datos asociado al nodo.
    - Valor: Valor hash del nodo.

- Estructura InformacionBloque
    - Fecha: Fecha y hora del bloque.
    - Accion: Acción realizada en el bloque (p. ej., "nulo" o "modificación").
    - Nombre: Nombre del libro asociado al bloque.
    - Tutor: Carnet del tutor asociado al bloque.

- Estructura NodoBloqueDatos
    - Siguiente: Puntero al siguiente bloque de datos.
    - Anterior: Puntero al bloque de datos anterior.
    - Valor: Puntero a la información del bloque.

## Cola Prioridad
Aquí se presenta un resumen de las funciones y estructuras implementadas:
### nodoGrafo.go

- Estructura NodoListaAdyacencia
    - Siguiente: Puntero al siguiente nodo en la misma columna.
    - Abajo: Puntero al siguiente nodo en la misma fila.
    - Valor: Valor asociado al nodo.

- Estructura PeticionGrafo
    - NombreArchivo: Nombre del archivo asociado a la petición.

### grafo.go

- Estructura Grafo
    - Principal: Puntero al nodo principal (cabeza) del grafo.

- Función insertarColumna
    - Inserta un nuevo nodo en la columna especificada.

- Función insertarFila
    - Inserta un nuevo nodo en la fila especificada.

- Función InsertarValores
    - Inserta valores en el grafo, creando nodos en la fila y columna correspondientes.

- Función Reporte
    - Genera un archivo DOT para representar gráficamente el grafo dirigido y ejecuta el comando para generar la imagen.

- Función retornarValoresMatriz
    - Retorna una cadena que representa las conexiones entre nodos del grafo para su visualización gráfica.


## Peticiones
### generar.go
- CrearArchivo(nombre_archivo string): Esta función crea un archivo con el nombre proporcionado si no existe previamente.

- EscribirArchivo(contenido string, nombre_archivo string): Escribe el contenido proporcionado en el archivo especificado.

- Ejecutar(nombre_imagen string, archivo string): Ejecuta el comando dot para generar una imagen (formato jpg) a partir de un archivo de definición de grafo en formato DOT. Utiliza el programa Graphviz para realizar esta conversión.

### peticion.go
**Estructuras de Peticiones**

- PeticionLogin: Representa la petición de inicio de sesión con nombre de usuario, contraseña y si el usuario es un tutor.
- PeticionRegistroTutor: Representa la petición de registro de un tutor con carné, nombre, curso y contraseña.
- PeticionRegistroAlumno: Representa la petición de registro de un alumno con carné, nombre, contraseña y cursos.
- PeticionCursos: Representa la petición de obtener la lista de cursos.
- Cursos: Estructura que contiene el código de un curso y una lista de posts asociados.
- PeticionLibro: Representa la petición de un usuario para agregar un libro con carné, nombre y contenido.
- PeticionPublicacion: Representa la petición de un usuario para realizar una publicación con carné y contenido.

## Tabla Hash
A continuación, se presenta una descripción del código:
### nodoHash.go
- Persona: Estructura que representa a una persona con un carné, nombre, contraseña y cursos.

- NodoHash: Estructura de un nodo de la tabla hash que contiene una llave y una referencia a una persona.

### tablaHash.go
- TablaHash: Estructura principal que representa la tabla hash. Contiene un mapa que asigna enteros a nodos hash, la capacidad de la tabla y la cantidad de elementos utilizados.

- calculoIndice(carnet int) int: Calcula el índice en la tabla hash utilizando el carné.

- capacidadTabla(): Verifica y ajusta la capacidad de la tabla si se supera un umbral de utilización.

- nuevaCapacidad() int: Calcula una nueva capacidad para la tabla en caso de ajuste.

- reInsertar(capacidadAnterior int): Reinserta los nodos en la tabla hash con la nueva capacidad después de un ajuste.

- reCalculoIndice(carnet int, contador int) int: Realiza el reajuste del índice en caso de colisión.

- nuevoIndice(nuevoIndice int) int: Calcula un nuevo índice ajustado a la capacidad de la tabla.

- Insertar(carnet int, nombre string, password string): Inserta una persona en la tabla hash.

- Buscar(carnet string, password string) bool: Busca una persona en la tabla hash por carné y contraseña.

- ConvertirArreglo() []NodoHash: Convierte la tabla hash a un arreglo de nodos.

## Main
A continuación, se presenta una descripción del código main.
### main.go
- Persona: Estructura que representa a una persona con un carné, nombre, contraseña y cursos.

- NodoHash: Estructura de un nodo de la tabla hash que contiene una llave y una referencia a una persona.

# Conclusiones

El programa es una implementación de diversas estructuras de datos para gestionar información relacionada con estudiantes, tutores y cursos. Se utilizan listas doblemente enlazadas, listas circulares, matrices dispersas y colas de prioridad. Cada estructura se adapta a un propósito específico: las listas para almacenar información de tutores, las matrices para representar relaciones entre tutores y estudiantes, y la cola de prioridad para gestionar la evaluación de tutores según su rendimiento. La modularidad y eficiencia de estas estructuras permiten una gestión efectiva de los datos. La interfaz con archivos CSV facilita la entrada y salida de información. En conjunto, el programa proporciona una herramienta versátil para el manejo y evaluación de tutores y estudiantes.

