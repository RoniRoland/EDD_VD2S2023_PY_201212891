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

## Matriz Dispersa
### Matriz.go
- `Matriz`
Contiene la estructura general de la matriz dispersa, con funciones para buscar columnas y filas, insertar columnas y filas, e insertar elementos en la matriz.

- `buscarColumna(carnet_tutor int, curso string) *NodoMatriz`
Busca una columna específica en la matriz dispersa.

- `buscarFila(carnet_estudiante int) *NodoMatriz`
Busca una fila específica en la matriz dispersa.

- `insertarColumna(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz`
Inserta una nueva columna en la matriz dispersa.

- `nuevaColumna(x int, carnet_Tutor int, curso string) *NodoMatriz`
Crea un nuevo nodo de columna y lo inserta en la matriz dispersa.

- `insertarFila(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz`
Inserta una nueva fila en la matriz dispersa.

- `nuevaFila(y int, carnet_estudiante int, curso string) *NodoMatriz`
Crea un nuevo nodo de fila y lo inserta en la matriz dispersa.

- `Insertar_Elemento(carnet_estudiante int, carnet_tutor int, curso string)`
Inserta un elemento en la matriz dispersa.

- `Reporte(nombre string)`
Genera un archivo de visualización en formato DOT y crea una imagen de la matriz dispersa.

### nodoMatriz.go
- **Dato:** Representa la información almacenada en cada celda de la matriz. Contiene campos como Carnet_Tutor para el identificador del tutor, Carnet_Estudiante para el identificador del estudiante y Curso para el curso asociado.

- **NodoMatriz:** Define un nodo para la matriz dispersa. Contiene punteros a nodos adyacentes (siguiente, anterior, arriba y abajo), así como coordenadas (posición X e Y) en la matriz. Además, tiene un puntero a la estructura Dato que almacena la información específica de la celda.

## Cola Prioridad
### PriorityQueue.go
- `AddQue(studentID int, name string, class string, score int)`
Añade un nuevo tutor a la cola de prioridad según su puntaje.

- `Dequeue()`
Elimina al primer tutor de la cola.

- `LeerCSV(ruta string)`
Lee datos desde un archivo CSV y añade tutores a la cola de prioridad.

- `First_Queue()`
Muestra en consola la información del tutor en la cima de la cola.

- `Queue`
Estructura principal que representa la cola de prioridad, con un puntero al primer nodo y un contador de longitud.

- `PriorityQueueNode`
Estructura que representa cada nodo en la cola de prioridad, conteniendo la información del tutor, un puntero al siguiente nodo, y la prioridad asignada según el puntaje.

- `First_Queue()`
Muestra en consola la información del tutor en la cima de la cola.

### node.go
- **Tutor:** Representa la información asociada a un tutor, incluyendo campos como StudentID para el identificador del estudiante, Name para el nombre del tutor, Class para la clase asociada y Score para la puntuación del tutor.

- **QueueNode:** Define un nodo para la cola de prioridad. Contiene un puntero a la estructura Tutor que almacena la información del tutor, así como un puntero al siguiente nodo en la cola (Next). Además, tiene un campo Priority que indica la prioridad del tutor en la cola.

## Árbol AVL
### treeNode.go
- Alumno
Estructura que representa la información de un alumno en el contexto del árbol AVL, con campos para el ID del estudiante y el nombre.

- AVLNode
Estructura que representa un nodo en el árbol AVL, con un puntero al alumno, a los nodos hijo izquierdo y derecho, y la altura del nodo.

### tree.go
- `AVLTree`
Estructura principal que representa el árbol AVL, con un puntero a la raíz del árbol.

- `getHeight(node *AVLNode) int`
Función que devuelve la altura de un nodo dado. Utilizada para equilibrar el árbol.

- `calculateBalanceFactor(node *AVLNode) int`
Función que calcula el factor de equilibrio de un nodo, determinado por la diferencia entre las alturas de los nodos hijo izquierdo y derecho.

- `rotateRight(node *AVLNode) *AVLNode`
Función que realiza una rotación a la derecha en un nodo del árbol AVL para mantener su equilibrio.

- `rotateLeft(node *AVLNode) *AVLNode`
Función que realiza una rotación a la izquierda en un nodo del árbol AVL para mantener su equilibrio.

- `insert(node *AVLNode, studentID int, name string) *AVLNode`
Función que inserta un nuevo nodo en el árbol AVL, manteniendo su equilibrio mediante rotaciones.

- `Insert(studentID int, name string)`
Función externa que facilita la inserción de un nuevo alumno en el árbol AVL.

- `inOrderTraversal(node *AVLNode)`
Función que realiza un recorrido en orden del árbol AVL e imprime los alumnos en consola.

- `Insertar_Elemento(carnet_estudiante int, nombre_estudiante string)`
Función que proporciona una interfaz externa para la inserción de estudiantes en el árbol AVL.

## Main
### main.go
- `main:` Función principal que presenta un menú de inicio de sesión y permite acceder al sistema como administrador o estudiante.

- `Login:` Gestiona el proceso de inicio de sesión solicitando usuario y contraseña. Permite iniciar sesión como administrador o estudiante.

- `AdminMenu:` Presenta el menú de opciones para el usuario administrador, que incluye la carga de datos, control de estudiantes y generación de reportes.

- `Reports:` Ofrece opciones para generar diferentes tipos de informes sobre alumnos, tutores, asignaciones y cursos.

- `StudentMenu:` Muestra las opciones disponibles para un estudiante, como ver tutores disponibles y asignarse a un tutor.

- `AsignarCurso:` Permite a un estudiante asignarse a un curso y tutor específico.

- `TutorLoadCSV:` Carga tutores desde un archivo CSV a la cola de prioridad.

- `StudentsLoadCSV:` Carga estudiantes desde un archivo CSV a la lista doblemente enlazada.

- `ClassLoadCSV:` Carga cursos desde un archivo JSON al árbol AVL de cursos.

- `ControlEstudiantes:` Muestra un menú de control de estudiantes para aceptar o rechazar tutores basándose en su puntaje.

- `clearScreen:` Limpia la pantalla de la consola para mejorar la presentación de los menús.

- `pressEnterToContinue:` Muestra un mensaje y espera que el usuario presione Enter para continuar.

# Conclusiones

El programa es una implementación de diversas estructuras de datos para gestionar información relacionada con estudiantes, tutores y cursos. Se utilizan listas doblemente enlazadas, listas circulares, matrices dispersas y colas de prioridad. Cada estructura se adapta a un propósito específico: las listas para almacenar información de tutores, las matrices para representar relaciones entre tutores y estudiantes, y la cola de prioridad para gestionar la evaluación de tutores según su rendimiento. La modularidad y eficiencia de estas estructuras permiten una gestión efectiva de los datos. La interfaz con archivos CSV facilita la entrada y salida de información. En conjunto, el programa proporciona una herramienta versátil para el manejo y evaluación de tutores y estudiantes.