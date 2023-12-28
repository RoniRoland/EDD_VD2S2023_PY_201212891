package main

import (
	"EDD_VD2S2023_PY_201212891/estructuras/AVLTree"
	"EDD_VD2S2023_PY_201212891/estructuras/Lists"
	"EDD_VD2S2023_PY_201212891/estructuras/MatrizDispersa"
	"EDD_VD2S2023_PY_201212891/estructuras/PriorityQueue"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var circularList *Lists.CircularList = &Lists.CircularList{Start: nil, Lenght: 0}
var doubleList *Lists.DoubleList = &Lists.DoubleList{Start: nil, Lenght: 0}
var priorityQueue *PriorityQueue.Queue = &PriorityQueue.Queue{First: nil, Lenght: 0}
var matrizDispersa *MatrizDispersa.Matriz = &MatrizDispersa.Matriz{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}
var arbolCursos *AVLTree.AVLTree = &AVLTree.AVLTree{Root: nil}
var loggeado_estudiante string = ""

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("No se puede limpiar la pantalla en este sistema operativo.")
	}
}

func pressEnterToContinue() {
	fmt.Print("\nPresiona Enter para continuar...")
	fmt.Scanln()
}

func main() {
	opt := 0
	exit := false
	for !exit {
		clearScreen()
		fmt.Println("=====================================================")
		fmt.Println("||                   LOGIN                         ||")
		fmt.Println("||                                                 ||")
		fmt.Println("||              1.- Inicio de Sesion               ||")
		fmt.Println("||              2.- Salir                          ||")
		fmt.Println("||                                                 ||")
		fmt.Println("=====================================================")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			clearScreen()
			Login()
		case 2:
			exit = true
		}
	}
}

func Login() {
	usuario := ""
	password := ""
	fmt.Println("=====================================================")
	fmt.Println("||                   LOGIN                         ||")
	fmt.Println("=====================================================")
	fmt.Print("               Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("               Password: ")
	fmt.Scanln(&password)

	if usuario == "ADMIN_201212891" && password == "Admin" {
		fmt.Println("\n*****Sesion Iniciada Correctamente******")
		pressEnterToContinue()
		AdminMenu()
	} else if doubleList.Find(usuario, password) {
		fmt.Println("\n*****Sesion Iniciada Correctamente Alumno ", usuario, "*******")
		loggeado_estudiante = usuario
		pressEnterToContinue()
		StudentMenu()
	} else {
		fmt.Println("\n Contraseña y usuario INCORRECTOS")
		pressEnterToContinue()
		clearScreen()
	}
}

func AdminMenu() {
	opt := 0
	exit := false
	for !exit {
		clearScreen()
		fmt.Println("═══════════════════════════════════════════════════════════════════")
		fmt.Println("||                   	MENU ADMINISTRADOR                       ||")
		fmt.Println("||                                                               ||")
		fmt.Println("||              1.- Carga de Estudiantes Tutores                 ||")
		fmt.Println("||              2.- Carga de Estudiantes                         ||")
		fmt.Println("||              3.- Cargar de Cursos                             ||")
		fmt.Println("||              4.- Control de Estudiantes                       ||")
		fmt.Println("||              5.- Reportes                                     ||")
		fmt.Println("||              6.- Salir                                        ||")
		fmt.Println("||                                                               ||")
		fmt.Println("===================================================================")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			TutorLoadCSV()
		case 2:
			StudentsLoadCSV()
		case 3:
			ClassLoadCSV()
		case 4:
			ControlEstudiantes()
		case 5:
			Reports()
		case 6:
			exit = true
		}

	}
}

func Reports() {
	opt := 0
	exit := false
	for !exit {
		clearScreen()
		fmt.Println("═══════════════════════════════════════════════════════════════════")
		fmt.Println("||                   	MENU DE REPORTES                          ||")
		fmt.Println("||                                                               ||")
		fmt.Println("||              1.- Reporte de Alumnos                           ||")
		fmt.Println("||              2.- Reporte de Tutores                           ||")
		fmt.Println("||              3.- Reporte de Asignacion                        ||")
		fmt.Println("||              4.- Reporte de Cursos                            ||")
		fmt.Println("||              5.- Salir                                        ||")
		fmt.Println("||                                                               ||")
		fmt.Println("===================================================================")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			doubleList.Reporte()
			pressEnterToContinue()
			exit = true
		case 2:
			circularList.Reportev2()
			pressEnterToContinue()
			exit = true
		case 3:
			matrizDispersa.Reporte("Matriz.jpg")
			pressEnterToContinue()
			exit = true
		case 4:
			arbolCursos.Graficar()
			pressEnterToContinue()
			exit = true
		case 5:
			exit = true
		}

	}
}

func StudentMenu() {
	opt := 0
	exit := false
	for !exit {
		clearScreen()
		fmt.Println("═══════════════════════════════════════════════════════════════════")
		fmt.Println("||                   	MENU ESTUDIANTE                          ||")
		fmt.Println("||                                                               ||")
		fmt.Println("||              1.- Ver Tutores disponibles                      ||")
		fmt.Println("||              2.- Asignarse a tutores                          ||")
		fmt.Println("||              3.- Salir                                        ||")
		fmt.Println("||                                                               ||")
		fmt.Println("===================================================================")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			circularList.Show()
			pressEnterToContinue()
		case 2:
			AsignarCurso()
		case 3:
			exit = true
		}

	}
}

func AsignarCurso() {
	opcion := ""
	salir := false
	for !salir {
		fmt.Print("\nIngresar el codigo del curso a asignarse: ")
		fmt.Scanln(&opcion)
		if arbolCursos.Busqueda(opcion) {
			if circularList.Find(opcion) {
				TutorBuscado := circularList.FindTutor(opcion)
				estudiante, err := strconv.Atoi(loggeado_estudiante)
				if err != nil {
					break
				}
				matrizDispersa.Insertar_Elemento(estudiante, TutorBuscado.Tutor.StudentID, opcion)
				fmt.Println("\n***CURSO asignado CORRECTAMENTE***")
				pressEnterToContinue()
				break
			} else {
				fmt.Println("\n No existen tutores para este curso!!!")
				pressEnterToContinue()
				break
			}
		} else {
			fmt.Println("\n ERROR! El curso NO EXISTE en el sistema!!!")
			pressEnterToContinue()
			break
		}

	}
}

func TutorLoadCSV() {
	ruta := ""
	clearScreen()
	fmt.Println("=====================================================")
	fmt.Println("||              CARGA DE TUTORES                    ||")
	fmt.Println("=====================================================")
	fmt.Print("             Nombre del Archivo CSV: ")
	fmt.Scanln(&ruta)
	rutaCompleta := ruta + ".csv"
	priorityQueue.LeerCSV(rutaCompleta)
	pressEnterToContinue()
}

func StudentsLoadCSV() {
	ruta := ""
	clearScreen()
	fmt.Println("=====================================================")
	fmt.Println("||              CARGA DE ESTUDIANTES               ||")
	fmt.Println("=====================================================")
	fmt.Print("          Nombre del Archivo CSV: ")
	fmt.Scanln(&ruta)
	rutaCompleta := ruta + ".csv"
	doubleList.ReadCSV(rutaCompleta)
	pressEnterToContinue()
}

func ClassLoadCSV() {
	ruta := ""
	clearScreen()
	fmt.Println("=====================================================")
	fmt.Println("||              CARGA DE CURSOS               ||")
	fmt.Println("=====================================================")
	fmt.Print("          Nombre del Archivo JSON: ")
	fmt.Scanln(&ruta)
	rutaCompleta := ruta + ".json"
	arbolCursos.LeerJson(rutaCompleta)
	pressEnterToContinue()
}

func ControlEstudiantes() {
	opcion := 0
	salir := false

	for !salir {
		clearScreen()
		priorityQueue.First_Queue()
		fmt.Println("═══════════════════════════════════════════════════════════════")
		fmt.Println("||                  1.- Aceptar                              ||")
		fmt.Println("||                  2.- Rechazar                             ||")
		fmt.Println("||                  3.- Salir                                ||")
		fmt.Println("═══════════════════════════════════════════════════════════════")
		fmt.Print("               Ingrese opcion: ")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			if tutor, exists := circularList.TutorExistsForClass(priorityQueue.First.Tutor.Class); exists {
				// Si ya hay un tutor para el curso, comparar notas y retener al tutor con la nota más alta
				if priorityQueue.First.Tutor.Score > tutor.Score {
					// Reemplazar al tutor existente con el nuevo tutor
					circularList.ReplaceTutor(tutor, (*Lists.Tutor)(priorityQueue.First.Tutor))
				}
				priorityQueue.Dequeue()
			} else {
				// Si no hay tutor para el curso, agregar al nuevo tutor
				circularList.Add(priorityQueue.First.Tutor.StudentID, priorityQueue.First.Tutor.Name, priorityQueue.First.Tutor.Class, priorityQueue.First.Tutor.Score)
				priorityQueue.Dequeue()
			}
		} else if opcion == 2 {
			priorityQueue.Dequeue()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}
