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

El programa es una aplicación de gestión educativa que utiliza diversas estructuras de datos para organizar información académica. El árbol AVL almacena datos de los cursos, manteniendo un equilibrio eficiente. Las listas dobles y circulares gestionan tutores con sus cursos y a los alumnos. La matriz dispersa registra la relación entre estudiantes y tutores en distintos cursos. La cola de prioridad organiza tutores según su rendimiento académico. Estas estructuras permiten una administración completa y eficaz de alumnos y tutores, facilitando la consulta y modificación de información educativa. El sistema se complementa con visualizaciones gráficas generadas a partir de los datos almacenados.

# Estructura del Proyecto

## Organización de Carpetas

El proyecto sigue una estructura organizada con carpetas específicas para cada componente.

- Lists: Contiene implementaciones de listas dobles y listas dobles circulares.
- MatrizDispersa: Contiene la implementación de la matriz dispersa.
- PriorityQueue: Contiene la implementación de la cola de prioridad.
- AVLTree: Contiene la implementacion del arbol avl.
- GeneradorArchivos: Contiene la implementacion para la creacion de los reportes.

## Listas 

Las listas dobles y circulares gestionan información relevante sobre tutores y cursos. La lista doble permite la inserción ordenada de alumnos, facilitando su recuperación y manipulación. Por otro lado, la lista circular organiza a los tutores asignados a los cursos. Ambas estructuras proporcionan funciones para agregar, buscar, reemplazar y visualizar información, contribuyendo al manejo integral de los datos académicos.



## Árbol AVL

El Árbol AVL desempeña un papel crucial en la organización y gestión de los datos académicos del sistema. Este árbol de búsqueda binario autoequilibrado almacena información tanto de estudiantes como de cursos, permitiendo búsquedas eficientes y operaciones de inserción y eliminación optimizadas. La función de inserción garantiza el equilibrio del árbol, manteniendo alturas adecuadas de los subárboles. Además de gestionar los estudiantes, el Árbol AVL almacena datos cruciales de los cursos, proporcionando una estructura jerárquica que facilita la recuperación de información detallada sobre el rendimiento académico y la distribución de estudiantes en diferentes asignaturas.

## Matriz Dispersa 

La matriz dispersa se especializa en almacenar relaciones entre estudiantes y tutores en diferentes cursos. Optimiza la gestión de información mediante nodos que enlazan columnas y filas. La inserción de elementos se realiza de manera eficiente, ajustando dinámicamente la matriz. Sus funciones clave, como `Insertar_Elemento` y `Reporte`, permiten registrar y visualizar las relaciones de manera clara y estructurada. En esta matriz los encabezados son los tutores y filas los alumnos.

## Cola de Prioridad

La cola de prioridad clasifica a los tutores según su rendimiento académico. Basándose en las calificaciones, asigna prioridades que determinan la posición en la cola. Las funciones `AddQue` y `Dequeue` facilitan la gestión dinámica de la cola, permitiendo la adición ordenada y la eliminación eficiente de tutores. La lectura de datos desde un archivo CSV agrega flexibilidad al sistema, mientras que la función `First_Queue` proporciona una vista detallada del tutor en la cima de la cola, mejorando la toma de decisiones académicas.

## Generar Archivos

El paquete GenerarArchivos desempeña un papel crucial en el sistema, facilitando la creación, escritura y ejecución de archivos de reporte visual. La función `CrearArchivo` verifica la existencia del archivo especificado y lo crea si aún no existe. Esto es fundamental para generar informes de salida. La función `EscribirArchivo` permite escribir contenido en el archivo, proporcionando una interfaz para la actualización de datos en informes ya existentes. La función `Ejecutar` utiliza la herramienta `dot` para transformar archivos de descripción de gráficos en lenguaje DOT en imágenes, permitiendo la visualización de estructuras de datos como listas y matrices dispersas. Este paquete ofrece una integración esencial para generar y manipular informes visuales en el sistema.

# Descripcion del codigo

## Listas 
### CircularList.go

- `Add(studentID int, name string, class string, score int)`
Añade un nuevo tutor a la lista doble circular. Si la lista está vacía, crea un nuevo nodo. En caso contrario, busca la posición adecuada según el StudentID y lo inserta.

- `Show()`
Muestra en consola la información de todos los tutores en la lista.

- `Find(class string) bool`
Busca si existe algún tutor asociado a un curso específico.

- `FindTutor(class string) *CircularNode`
Devuelve el nodo tutor asociado a un curso específico.

- `TutorExistsForClass(class string) (*Tutor, bool)`
Verifica si hay algún tutor asociado a un curso específico y devuelve su información.

- `ReplaceTutor(existingTutor *Tutor, newTutor *Tutor)`
Reemplaza un tutor existente con uno nuevo.

- `Reportev2()`
Genera un archivo de visualización en formato DOT y crea una imagen de la lista doble circular.

### DoubleList.go
- `Add(studentID int, name string)`
Añade un nuevo alumno a la lista doble. Si la lista está vacía, crea un nuevo nodo. En caso contrario, lo añade al final.

- `Find(studentID string, password string) bool`
Busca si existe un alumno con el StudentID y contraseña proporcionados.

- `ReadCSV(ruta string)`
Lee datos desde un archivo CSV y añade nodos a la lista doble.

- `Reporte()`
Genera un archivo de visualización en formato DOT y crea una imagen de la lista doble.

### node.go

- **Alumn:** Representa la información de un estudiante con campos como StudentID para el identificador del estudiante y Name para el nombre del estudiante.

- **Tutor:** Representa la información de un tutor con campos como StudentID para el identificador del tutor, Name para el nombre del tutor, Class para la clase asignada y Score para la puntuación académica del tutor.

- **DoubleListNode:** Define un nodo para una lista doblemente enlazada que contiene un puntero a la estructura Alumn, así como punteros al siguiente y al anterior nodo en la lista.

- **CircularNode:** Define un nodo para una lista circular que contiene un puntero a la estructura Tutor, así como punteros al siguiente y al anterior nodo en la lista.

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